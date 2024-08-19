package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSuperfluidDelegate{}, "osmosis/superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegate{}, "osmosis/superfluid-undelegate", nil)
	cdc.RegisterConcrete(&MsgLockAndSuperfluidDelegate{}, "osmosis/lock-and-superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUnbondLock{}, "osmosis/superfluid-unbond-lock", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegateAndUnbondLock{}, "osmosis/sf-undelegate-and-unbond-lock", nil)
	// TODO: Remove in v27 once comfortable with new gov message
	cdc.RegisterConcrete(&SetSuperfluidAssetsProposal{}, "osmosis/set-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&UpdateUnpoolWhiteListProposal{}, "osmosis/update-unpool-whitelist", nil)
	// TODO: Remove in v27 once comfortable with new gov message
	cdc.RegisterConcrete(&RemoveSuperfluidAssetsProposal{}, "osmosis/del-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&MsgUnPoolWhitelistedPool{}, "osmosis/unpool-whitelisted-pool", nil)
	cdc.RegisterConcrete(&MsgUnlockAndMigrateSharesToFullRangeConcentratedPosition{}, "osmosis/unlock-and-migrate", nil)
	cdc.RegisterConcrete(&MsgCreateFullRangePositionAndSuperfluidDelegate{}, "osmosis/full-range-and-sf-delegate", nil)
	cdc.RegisterConcrete(&MsgAddToConcentratedLiquiditySuperfluidPosition{}, "osmosis/add-to-cl-superfluid-position", nil)
	cdc.RegisterConcrete(&MsgUnbondConvertAndStake{}, "osmosis/unbond-convert-and-stake", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSuperfluidDelegate{},
		&MsgSuperfluidUndelegate{},
		&MsgLockAndSuperfluidDelegate{},
		&MsgSuperfluidUnbondLock{},
		&MsgSuperfluidUndelegateAndUnbondLock{},
		&MsgUnPoolWhitelistedPool{},
		&MsgUnlockAndMigrateSharesToFullRangeConcentratedPosition{},
		&MsgCreateFullRangePositionAndSuperfluidDelegate{},
		&MsgAddToConcentratedLiquiditySuperfluidPosition{},
		&MsgUnbondConvertAndStake{},
	)

	registry.RegisterImplementations(
		(*govtypesv1.Content)(nil),
		&SetSuperfluidAssetsProposal{},
		&RemoveSuperfluidAssetsProposal{},
		&UpdateUnpoolWhiteListProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
