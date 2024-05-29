package clientchains

import (
	"math"

	errorsmod "cosmossdk.io/errors"
	assetstypes "github.com/ExocoreNetwork/exocore/x/assets/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func (p Precompile) GetClientChains(
	ctx sdk.Context,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	if len(args) > 0 {
		ctx.Logger().Error(
			"GetClientChains",
			"err", errorsmod.Wrapf(
				assetstypes.ErrInvalidInputParameter, "no input is required",
			),
		)
		return method.Outputs.Pack(false, nil)
	}
	infos, err := p.assetsKeeper.GetAllClientChainInfo(ctx)
	if err != nil {
		ctx.Logger().Error(
			"GetClientChains",
			"err", err,
		)
		return method.Outputs.Pack(false, nil)
	}
	ids := make([]uint16, 0, len(infos))
	for id := range infos {
		// technically LZ supports uint32, but unfortunately all the precompiles
		// based it on uint16, so we have to stick with it.
		// TODO: change it to uint32 here and in other precompiles.
		if id > math.MaxUint16 {
			ctx.Logger().Error(
				"GetClientChains",
				"err", errorsmod.Wrapf(
					assetstypes.ErrInvalidInputParameter, "client chain id is too large",
				),
			)
			return method.Outputs.Pack(false, nil)
		}
		// #nosec G701 // already checked
		convID := uint16(id)
		ids = append(ids, convID)
	}
	return method.Outputs.Pack(true, ids)
}
