package v15

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	"github.com/osmosis-labs/osmosis/v15/app/upgrades"
	cltypes "github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
	cosmwasmpooltypes "github.com/osmosis-labs/osmosis/v15/x/cosmwasmpool/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"
	protorevtypes "github.com/osmosis-labs/osmosis/v15/x/protorev/types"
	valsetpreftypes "github.com/osmosis-labs/osmosis/v15/x/valset-pref/types"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v15 upgrade.
const UpgradeName = "v15"

// pool ids to migrate
const (
	stOSMO_OSMOPoolId   = 833
	stJUNO_JUNOPoolId   = 817
	stSTARS_STARSPoolId = 810
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		// TODO: move cosmwasmpooltypes.StoreKet to v16 upgrade handler.
		Added:   []string{poolmanagertypes.StoreKey, cltypes.StoreKey, valsetpreftypes.StoreKey, protorevtypes.StoreKey, icqtypes.StoreKey, packetforwardtypes.StoreKey, cosmwasmpooltypes.StoreKey},
		Deleted: []string{},
	},
}
