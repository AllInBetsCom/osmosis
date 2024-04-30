package apptesting

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmoutils/cosmwasm"
	"github.com/osmosis-labs/osmosis/v25/x/cosmwasmpool/cosmwasm/msg"
	"github.com/osmosis-labs/osmosis/v25/x/cosmwasmpool/cosmwasm/msg/transmuter"
	"github.com/osmosis-labs/osmosis/v25/x/cosmwasmpool/model"

	cosmwasmpooltypes "github.com/osmosis-labs/osmosis/v25/x/cosmwasmpool/types"
)

const (
	DefaultTransmuterDenomA       = "axlusdc"
	DefaultTransmuterDenomB       = "gravusdc"
	TransmuterContractName        = "transmuter"
	TransmuterMigrateContractName = "transmuter_migrate"
	DefaultCodeId                 = 1

	osmosisRepository         = "osmosis"
	osmosisRepoTransmuterPath = "x/cosmwasmpool/bytecode"
)

// PrepareCosmWasmPool sets up a cosmwasm pool with the default parameters.
func (s *KeeperTestHelper) PrepareCosmWasmPool() cosmwasmpooltypes.CosmWasmExtension {
	return s.PrepareCustomTransmuterPool(s.TestAccs[0], []string{DefaultTransmuterDenomA, DefaultTransmuterDenomB})
}

// PrepareCustomTransmuterPool sets up a transmuter pool with the default parameters assuming that
// the transmuter contract is stored under x/cosmwasmpool/bytecode in the Osmosis repository.
func (s *KeeperTestHelper) PrepareCustomTransmuterPool(owner sdk.AccAddress, denoms []string) cosmwasmpooltypes.CosmWasmExtension {
	return s.PrepareCustomTransmuterPoolCustomProject(owner, denoms, osmosisRepository, osmosisRepoTransmuterPath)
}

// PrepareCustomTransmuterPoolCustomProject sets up a transmuter pool with the custom parameters.
// Gives flexibility for the helper to be reused outside of the Osmosis repository by providing custom
// project name and bytecode path.
func (s *KeeperTestHelper) PrepareCustomTransmuterPoolCustomProject(owner sdk.AccAddress, denoms []string, projectName, byteCodePath string) cosmwasmpooltypes.CosmWasmExtension {
	// Mint some assets to the account.
	s.FundAcc(s.TestAccs[0], DefaultAcctFunds)

	// Upload contract code and get the code id.
	codeId := s.StoreCosmWasmPoolContractCodeCustomProject(TransmuterContractName, projectName, byteCodePath)

	// Add code id to the whitelist.
	s.App.CosmwasmPoolKeeper.WhitelistCodeId(s.Ctx, codeId)

	// Generate instantiate message bytes.
	instantiateMsgBz := s.GetTransmuterInstantiateMsgBytes(denoms)

	// Generate msg create pool.
	validCWPoolMsg := model.NewMsgCreateCosmWasmPool(codeId, owner, instantiateMsgBz)

	// Create pool.
	poolId, err := s.App.PoolManagerKeeper.CreatePool(s.Ctx, validCWPoolMsg)
	s.Require().NoError(err)

	// Get and return the pool.
	pool, err := s.App.CosmwasmPoolKeeper.GetPoolById(s.Ctx, poolId)
	s.Require().NoError(err)

	return pool
}

// GetDefaultTransmuterInstantiateMsgBytes returns the default instantiate message for the transmuter contract
// with DefaultTransmuterDenomA and DefaultTransmuterDenomB as the pool asset denoms.
func (s *KeeperTestHelper) GetDefaultTransmuterInstantiateMsgBytes() []byte {
	return s.GetTransmuterInstantiateMsgBytes([]string{DefaultTransmuterDenomA, DefaultTransmuterDenomB})
}

// GetTransmuterInstantiateMsgBytes returns the instantiate message for the transmuter contract with the
// given pool asset denoms.
func (s *KeeperTestHelper) GetTransmuterInstantiateMsgBytes(poolAssetDenoms []string) []byte {
	instantiateMsg := msg.InstantiateMsg{
		PoolAssetDenoms: poolAssetDenoms,
	}

	instantiateMsgBz, err := json.Marshal(instantiateMsg)
	s.Require().NoError(err)

	return instantiateMsgBz
}

// StoreCosmWasmPoolContractCode stores the cosmwasm pool contract code in the wasm keeper and returns the code id.
// contractName is the name of the contract file in the x/cosmwasmpool/bytecode directory without the .wasm extension.
func (s *KeeperTestHelper) StoreCosmWasmPoolContractCode(contractName string) uint64 {
	return s.StoreCosmWasmPoolContractCodeCustomProject(contractName, osmosisRepository, osmosisRepoTransmuterPath)
}

// StoreCosmWasmPoolContractCodeCustomProject stores the cosmwasm pool contract code in the wasm keeper and returns the code id.
// contractName is the name of the contract file in the x/cosmwasmpool/bytecode directory without the .wasm extension.
// It has the flexibility of being used from outside the Osmosis repository by providing custom project name and bytecode path.
func (s *KeeperTestHelper) StoreCosmWasmPoolContractCodeCustomProject(contractName, projectName, byteCodePath string) uint64 {
	cosmwasmpoolModuleAddr := s.App.AccountKeeper.GetModuleAddress(cosmwasmpooltypes.ModuleName)
	s.Require().NotNil(cosmwasmpoolModuleAddr)

	// Allow the cosmwasm pool module to upload code.
	params := s.App.WasmKeeper.GetParams(s.Ctx)
	err := s.App.WasmKeeper.SetParams(s.Ctx, wasmtypes.Params{
		CodeUploadAccess: wasmtypes.AccessConfig{
			Permission: wasmtypes.AccessTypeAnyOfAddresses,
			Addresses:  []string{cosmwasmpoolModuleAddr.String()},
		},
		InstantiateDefaultPermission: params.InstantiateDefaultPermission,
	})
	s.Require().NoError(err)

	code := s.GetContractCodeCustomProject(contractName, projectName, byteCodePath)

	instantiateConfig := wasmtypes.AccessConfig{Permission: wasmtypes.AccessTypeAnyOfAddresses, Addresses: []string{cosmwasmpoolModuleAddr.String()}}
	codeID, _, err := s.App.ContractKeeper.Create(s.Ctx, cosmwasmpoolModuleAddr, code, &instantiateConfig)
	s.Require().NoError(err)

	return codeID
}

func (s *KeeperTestHelper) GetContractCode(contractName string) []byte {
	return s.GetContractCodeCustomProject(contractName, "osmosis", "x/cosmwasmpool/bytecode")
}

// GetContractCode returns the contract code for the given contract name.
// Assumes that the contract code is stored under x/cosmwasmpool/bytecode.
func (s *KeeperTestHelper) GetContractCodeCustomProject(contractName string, projectName string, path string) []byte {
	workingDir, err := os.Getwd()
	s.Require().NoError(err)

	projectRootPath := fmt.Sprintf("/%s/", projectName)
	projectRootIndex := strings.LastIndex(workingDir, projectRootPath) + len(projectRootPath)
	workingDir = workingDir[:projectRootIndex]
	code, err := os.ReadFile(workingDir + path + "/" + contractName + ".wasm")
	s.Require().NoError(err)

	return code
}

// JoinTransmuterPool joins the given pool with the given coins from the given address.
func (s *KeeperTestHelper) JoinTransmuterPool(lpAddress sdk.AccAddress, poolId uint64, coins sdk.Coins) {
	pool, err := s.App.CosmwasmPoolKeeper.GetPoolById(s.Ctx, poolId)
	s.Require().NoError(err)
	// add liquidity by joining the pool
	request := transmuter.JoinPoolExecuteMsgRequest{}
	cosmwasm.MustExecute[transmuter.JoinPoolExecuteMsgRequest, msg.EmptyStruct](s.Ctx, s.App.ContractKeeper, pool.GetContractAddress(), lpAddress, coins, request)
}
