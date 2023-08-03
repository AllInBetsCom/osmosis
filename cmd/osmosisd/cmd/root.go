package cmd

import (
	// "fmt"

	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/osmosis-labs/osmosis/v17/app/params"

	tmcmds "github.com/cometbft/cometbft/cmd/tendermint/commands"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/joho/godotenv"

	"github.com/cosmos/cosmos-sdk/client/config"

	osmosis "github.com/osmosis-labs/osmosis/v17/app"
)

type AssetList struct {
	Assets []Asset `json:"assets"`
}
type Asset struct {
	DenomUnits []DenomUnit `json:"denom_units"`
	Symbol     string      `json:"symbol"`
	Base       string      `json:"base"`
	Traces     []Trace     `json:"traces"`
}

type DenomUnit struct {
	Denom    string `json:"denom"`
	Exponent uint64 `json:"exponent"`
}

type Trace struct {
	Type         string `json:"type"`
	Counterparty struct {
		BaseDenom string `json:"base_denom"`
	} `json:"counterparty"`
}

type DenomUnitMap struct {
	Base       string
	DenomUnits []DenomUnit `json:"denom_units"`
}

var (
	//go:embed "osmosis-1-assetlist.json" "osmo-test-5-assetlist.json"
	assetFS   embed.FS
	mainnetId = "osmosis-1"
	testnetId = "osmo-test-5"
)

func loadAssetList(initClientCtx client.Context, cmd *cobra.Command, basedenomToIBC, IBCtoBasedenom bool) (map[string]DenomUnitMap, map[string]string) {
	var assetList AssetList

	chainId := GetChainId(initClientCtx, cmd)

	fileName := ""
	if chainId == mainnetId || chainId == "" {
		fileName = "cmd/osmosisd/cmd/osmosis-1-assetlist-manual.json"
	} else if chainId == testnetId {
		fileName = "cmd/osmosisd/cmd/osmo-test-5-assetlist-manual.json"
	} else {
		return nil, nil
	}

	// The order of precedence for asset list is:
	//  - If the manually generated asset list (generated via the `update-asset-list` cli cmd) exists (noted by -manual.json ending), use it.
	//  - If the manually generated asset list does not exist, fall back to the embedded asset list.
	localFile, err := os.Open(fileName)
	if err != nil {
		// If we can't open the local file, fall back to the embedded file.
		if chainId == mainnetId || chainId == "" {
			fileName = "osmosis-1-assetlist.json"
		} else if chainId == testnetId {
			fileName = "osmo-test-5-assetlist.json"
		} else {
			return nil, nil
		}
		embeddedFile, err := assetFS.Open(fileName)
		if err != nil {
			return nil, nil
		}
		defer embeddedFile.Close()

		byteValue, _ := io.ReadAll(embeddedFile)
		err = json.Unmarshal(byteValue, &assetList)
		if err != nil {
			return nil, nil
		}
	} else {
		// If the local file opens successfully, use it instead of the embedded file.
		defer localFile.Close()
		byteValue, _ := io.ReadAll(localFile)
		err = json.Unmarshal(byteValue, &assetList)
		if err != nil {
			return nil, nil
		}
	}

	baseMap := make(map[string]DenomUnitMap)
	baseMapRev := make(map[string]string)

	if basedenomToIBC {
		for _, asset := range assetList.Assets {
			DenomUnitMap := DenomUnitMap{
				Base:       asset.Base,
				DenomUnits: asset.DenomUnits,
			}
			baseMap["base"+strings.ToLower(asset.Symbol)] = DenomUnitMap
		}
	}
	if IBCtoBasedenom {
		for _, asset := range assetList.Assets {
			baseMapRev[asset.Base] = "base" + strings.ToLower(asset.Symbol)
		}
	}
	return baseMap, baseMapRev
}

type customWriter struct {
	originalOut io.Writer
	baseMap     map[string]string
}

func (cw *customWriter) Write(p []byte) (n int, err error) {
	// Convert byte slice to string.
	s := string(p)

	// Buffer to hold the new string.
	var buf strings.Builder

	// Index where the current denom starts. -1 if we're not currently in a denom.
	denomStart := -1

	// Counter for slashes encountered
	slashCounter := 0

	re, err := regexp.Compile("[^a-zA-Z0-9/-]")
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(s); i++ {
		if denomStart == -1 {
			// If we're not currently in a denom, check if this character starts a new denom.
			// Check for "ibc/" or "factory/" prefix.
			if strings.HasPrefix(s[i:], "ibc/") || strings.HasPrefix(s[i:], "factory/") {
				slashCounter = 0
				denomStart = i
				continue
			}
			// Write the character to the buffer.
			buf.WriteByte(s[i])
		} else {
			// For factory denoms, we keep track of slashes to find the second slash.
			if s[i] == '/' {
				slashCounter++
			}

			// For "ibc/" we find the end by length, for "factory/" we find the end by second slash and regex.
			if (strings.HasPrefix(s[denomStart:], "ibc/") && i-denomStart == 68) || (strings.HasPrefix(s[denomStart:], "factory/") && slashCounter == 2 && re.MatchString(string(s[i]))) {
				// We've reached the end of the line containing the denom.
				denom := s[denomStart:i]
				if replacement, ok := cw.baseMap[denom]; ok {
					// If the denom is in the map, write the replacement to the buffer.
					buf.WriteString(replacement)
				} else {
					// If the denom is not in the map, write the original denom to the buffer.
					buf.WriteString(denom)
				}
				// Write the new line character to the buffer.
				buf.WriteByte(s[i])

				// We're no longer in a denom.
				denomStart = -1
				slashCounter = 0
			}
		}
	}

	// If we're still in a denom at the end of the string, write the rest of the denom to the buffer.
	if denomStart != -1 {
		denom := s[denomStart:]
		if replacement, ok := cw.baseMap[denom]; ok {
			buf.WriteString(replacement)
		} else {
			buf.WriteString(denom)
		}
	}

	// Write the new string to the original output.
	return cw.originalOut.Write([]byte(buf.String()))
}

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := osmosis.MakeEncodingConfig()
	homeEnvironment := getHomeEnvironment()
	homeDir, err := environmentNameToPath(homeEnvironment)
	if err != nil {
		// Failed to convert home environment to home path, using default home
		homeDir = osmosis.DefaultNodeHome
	}

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(homeDir).
		WithViper("OSMOSIS")

	// Allows you to add extra params to your client.toml
	// gas, gas-price, gas-adjustment, and human-readable-denoms
	SetCustomEnvVariablesFromClientToml(initClientCtx)
	humanReadableDenomsInput, humanReadableDenomsOutput := GetHumanReadableDenomEnvVariables()

	rootCmd := &cobra.Command{
		Use:   "osmosisd",
		Short: "Start osmosis app",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			// Only loads asset list into a map if human readable denoms are enabled.
			assetMap, assetMapRev := map[string]DenomUnitMap{}, map[string]string{}
			if humanReadableDenomsInput || humanReadableDenomsOutput {
				assetMap, assetMapRev = loadAssetList(initClientCtx, cmd, humanReadableDenomsInput, humanReadableDenomsOutput)
			}

			// If enabled, CLI output will be parsed and human readable denominations will be used in place of ibc denoms.
			if humanReadableDenomsOutput {
				initClientCtx.Output = &customWriter{originalOut: os.Stdout, baseMap: assetMapRev}
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			customAppTemplate, customAppConfig := initAppConfig()

			// If enabled, CLI input will be parsed and human readable denominations will be automatically converted to ibc denoms.
			if humanReadableDenomsInput {
				// Parse and replace denoms in args
				for i, arg := range args {
					lowerCaseArg := strings.ToLower(arg)
					lowerCaseArgArray := strings.Split(lowerCaseArg, ",")
					re := regexp.MustCompile(`^(\d+)(.+)`)

					for i, lowerCaseArg := range lowerCaseArgArray {
						match := re.FindStringSubmatch(lowerCaseArg)
						if len(match) == 3 {
							// If the index has a length of 3 then it has a number and a denom (this is a coin object)
							// Note, index 0 is the entire string, index 1 is the number, and index 2 is the denom
							if _, ok := assetMap[match[2]]; ok {
								// In this case, we just need to replace the denom with the base denom and retain the number
								lowerCaseArgArray[i] = match[1] + assetMap[match[2]].Base
							}
						} else {
							if _, ok := assetMap[lowerCaseArg]; ok {
								// In this case, we just need to replace the denom with the base denom
								lowerCaseArgArray[i] = assetMap[lowerCaseArg].Base
							}
						}
					}
					args[i] = strings.Join(lowerCaseArgArray, ",")
				}

				// Parse and replace denoms in flags
				cmd.Flags().VisitAll(func(flag *pflag.Flag) {
					lowerCaseFlagValue := strings.ToLower(flag.Value.String())
					lowerCaseFlagValueArray := strings.Split(lowerCaseFlagValue, ",")

					re := regexp.MustCompile(`^(\d+)(.+)`)

					for i, lowerCaseFlagValue := range lowerCaseFlagValueArray {
						match := re.FindStringSubmatch(lowerCaseFlagValue)
						if len(match) == 3 {
							// If the index has a length of 3 then it has a number and a denom (this is a coin object)
							// Note, index 0 is the entire string, index 1 is the number, and index 2 is the denom
							if _, ok := assetMap[match[2]]; ok {
								// In this case, we just need to replace the denom with the base denom and retain the number
								lowerCaseFlagValueArray[i] = strings.Replace(lowerCaseFlagValue, match[2], assetMap[match[2]].Base, -1)
							}
						} else {
							if _, ok := assetMap[lowerCaseFlagValue]; ok {
								// Otherwise, we just need to replace the denom with the base denom
								lowerCaseFlagValueArray[i] = assetMap[lowerCaseFlagValue].Base
							}
						}
					}
					newLowerCaseFlagValue := strings.Join(lowerCaseFlagValueArray, ",")
					if lowerCaseFlagValue != newLowerCaseFlagValue {
						if err := cmd.Flags().Set(flag.Name, newLowerCaseFlagValue); err != nil {
							fmt.Println("Failed to set flag:", err)
						}
					}
				})
			}

			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig)
		},
		SilenceUsage: true,
	}

	genAutoCompleteCmd(rootCmd)

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func getHomeEnvironment() string {
	envPath := filepath.Join(osmosis.DefaultNodeHome, ".env")

	// Use default node home if can't get environment
	err := godotenv.Load(envPath)
	if err != nil {
		// Failed to load, using default home directory
		return EnvMainnet
	}
	val := os.Getenv(EnvVariable)
	return val
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	type OsmosisMempoolConfig struct {
		ArbitrageMinGasPrice string `mapstructure:"arbitrage-min-gas-fee"`
	}

	type CustomAppConfig struct {
		serverconfig.Config

		OsmosisMempoolConfig OsmosisMempoolConfig `mapstructure:"osmosis-mempool"`
	}

	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	srvCfg := serverconfig.DefaultConfig()
	srvCfg.API.Enable = true
	srvCfg.StateSync.SnapshotInterval = 1500
	srvCfg.StateSync.SnapshotKeepRecent = 2
	srvCfg.MinGasPrices = "0uosmo"

	// 128MB IAVL cache
	srvCfg.IAVLCacheSize = 781250

	memCfg := OsmosisMempoolConfig{ArbitrageMinGasPrice: "0.01"}

	OsmosisAppCfg := CustomAppConfig{Config: *srvCfg, OsmosisMempoolConfig: memCfg}

	OsmosisAppTemplate := serverconfig.DefaultConfigTemplate + `
###############################################################################
###                      Osmosis Mempool Configuration                      ###
###############################################################################

[osmosis-mempool]
# This is the max allowed gas any tx. 
# This is only for local mempool purposes, and thus	is only ran on check tx.
max-gas-wanted-per-tx = "25000000"

# This is the minimum gas fee any arbitrage tx should have, denominated in uosmo per gas
# Default value of ".005" then means that a tx with 1 million gas costs (.005 uosmo/gas) * 1_000_000 gas = .005 osmo
arbitrage-min-gas-fee = ".005"

# This is the minimum gas fee any tx with high gas demand should have, denominated in uosmo per gas
# Default value of ".0025" then means that a tx with 1 million gas costs (.0025 uosmo/gas) * 1_000_000 gas = .0025 osmo
min-gas-price-for-high-gas-tx = ".0025"
`

	return OsmosisAppTemplate, OsmosisAppCfg
}

// initRootCmd initializes root commands when creating a new root command for simd.
func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	cfg := sdk.GetConfig()
	cfg.Seal()

	debugCmd := debug.Cmd()
	debugCmd.AddCommand(ConvertBech32Cmd())
	debugCmd.AddCommand(DebugProtoMarshalledBytes())

	rootCmd.AddCommand(
		// genutilcli.InitCmd(osmosis.ModuleBasics, osmosis.DefaultNodeHome),
		forceprune(),
		InitCmd(osmosis.ModuleBasics, osmosis.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, osmosis.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		ExportDeriveBalancesCmd(),
		StakedToCSVCmd(),
		AddGenesisAccountCmd(osmosis.DefaultNodeHome),
		genutilcli.GenTxCmd(osmosis.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, osmosis.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(osmosis.ModuleBasics),
		PrepareGenesisCmd(osmosis.DefaultNodeHome, osmosis.ModuleBasics),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(osmosis.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		tmcmds.RollbackStateCmd,
		debugCmd,
		ConfigCmd(),
		ChangeEnvironmentCmd(),
		PrintEnvironmentCmd(),
		UpdateAssetListCmd(osmosis.DefaultNodeHome, osmosis.ModuleBasics),
	)

	server.AddCommands(rootCmd, osmosis.DefaultNodeHome, newApp, createOsmosisAppAndExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(osmosis.DefaultNodeHome),
	)
	// add rosetta
	rootCmd.AddCommand(server.RosettaCommand(encodingConfig.InterfaceRegistry, encodingConfig.Marshaler))
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	wasm.AddModuleInitFlags(startCmd)
}

// queryCommand adds transaction and account querying commands.
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	osmosis.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand adds transaction signing, encoding / decoding, and broadcasting commands.
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
	)

	osmosis.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// newApp initializes and returns a new Osmosis app.
func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return osmosis.NewOsmosisApp(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		appOpts,
		wasmOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshot(snapshotStore, snapshottypes.NewSnapshotOptions(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval)), cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent)))),
	)
}

// createOsmosisAppAndExport creates and exports the new Osmosis app, returns the state of the new Osmosis app for a genesis file.
func createOsmosisAppAndExport(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
	appOpts servertypes.AppOptions, modulesToExport []string,
) (servertypes.ExportedApp, error) {
	encCfg := osmosis.MakeEncodingConfig() // Ideally, we would reuse the one created by NewRootCmd.
	encCfg.Marshaler = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	loadLatest := height == -1
	homeDir := cast.ToString(appOpts.Get(flags.FlagHome))
	app := osmosis.NewOsmosisApp(logger, db, traceStore, loadLatest, map[int64]bool{}, homeDir, 0, appOpts, osmosis.EmptyWasmOpts)

	if !loadLatest {
		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailWhiteList, modulesToExport)
}

func UpdateAssetListCmd(defaultNodeHome string, mbm module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-asset-list [chain-id]",
		Short: "Updates asset list used by the CLI to replace ibc denoms with human readable names",
		Long: `Updates asset list used by the CLI to replace ibc denoms with human readable names.
Outputs:
	- cmd/osmosisd/cmd/osmosis-1-assetlist-manual.json for osmosis-1
	- cmd/osmosisd/cmd/osmo-test-5-assetlist-manual.json for osmo-test-5
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			assetListURL := ""
			fileName := ""

			if args[0] == mainnetId || args[0] == "" {
				assetListURL = "https://raw.githubusercontent.com/osmosis-labs/assetlists/main/osmosis-1/osmosis-1.assetlist.json"
				fileName = "cmd/osmosisd/cmd/osmosis-1-assetlist-manual.json"
			} else if args[0] == testnetId {
				assetListURL = "https://raw.githubusercontent.com/osmosis-labs/assetlists/main/osmo-test-5/osmo-test-5.assetlist.json"
				fileName = "cmd/osmosisd/cmd/osmo-test-5-assetlist-manual.json"
			} else {
				return nil
			}

			// Try to fetch the asset list from the URL
			resp, err := http.Get(assetListURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Create a new file
			out, err := os.Create(fileName)
			if err != nil {
				return err
			}
			defer out.Close()

			// Copy the response body to the new file
			_, err = io.Copy(out, resp.Body)
			return err
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

func genAutoCompleteCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "enable-cli-autocomplete [bash|zsh|fish|powershell]",
		Short: "Generates cli completion scripts",
		Long: `To configure your shell to load completions for each session, add to your profile:

# bash example
echo '. <(osmosisd enable-cli-autocomplete bash)' >> ~/.profile
source ~/.profile

# zsh example
echo '. <(osmosisd enable-cli-autocomplete zsh)' >> ~/.zshrc
source ~/.zshrc
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				_ = cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				_ = cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				_ = cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	})
}
