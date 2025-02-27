package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrNoClientChainKey = errorsmod.Register(
		ModuleName, 0,
		"there is no stored key for the input chain index",
	)

	ErrNoClientChainAssetKey = errorsmod.Register(
		ModuleName, 1,
		"there is no stored key for the input assetID",
	)

	ErrNoStakerAssetKey = errorsmod.Register(
		ModuleName, 2,
		"there is no stored key for the input staker and assetID",
	)

	ErrSubAmountIsMoreThanOrigin = errorsmod.Register(
		ModuleName, 3,
		"the amount that want to decrease is more than the original state amount",
	)

	ErrNoOperatorAssetKey = errorsmod.Register(
		ModuleName, 4,
		"there is no stored key for the input operator address and assetID",
	)

	ErrParseAssetsStateKey = errorsmod.Register(
		ModuleName, 5,
		"assets state key can't be parsed",
	)

	ErrInvalidCliCmdArg = errorsmod.Register(
		ModuleName, 6,
		"the input client command arguments are invalid",
	)

	ErrInputPointerIsNil = errorsmod.Register(
		ModuleName, 7,
		"the input pointer is nil",
	)

	ErrInvalidOperatorAddr = errorsmod.Register(
		ModuleName,
		8,
		"the operator address isn't a valid account address",
	)

	ErrUnknownAppChainID = errorsmod.Register(
		ModuleName, 9,
		"the app chain id is unknown or invalid",
	)

	ErrInvalidEvmAddressFormat = errorsmod.Register(
		ModuleName, 10,
		"the evm address format is error",
	)

	ErrInvalidLzUaTopicIDLength = errorsmod.Register(
		ModuleName, 11,
		"the LZUaTopicID length isn't equal to HashLength",
	)

	ErrNoParamsKey = errorsmod.Register(
		ModuleName, 12,
		"there is no stored key for deposit module params",
	)

	ErrNotEqualToLzAppAddr = errorsmod.Register(
		ModuleName, 13,
		"the address isn't equal to the layerZero gateway address",
	)

	ErrInvalidGenesisData = errorsmod.Register(
		ModuleName, 14,
		"the genesis data supplied is invalid",
	)

	ErrInvalidInputParameter = errorsmod.Register(
		ModuleName, 15,
		"the input parameter is invalid",
	)
)
