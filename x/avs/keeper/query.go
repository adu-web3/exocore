package keeper

import (
	"context"

	"github.com/ExocoreNetwork/exocore/x/avs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryAVSInfo(ctx context.Context, req *types.QueryAVSInfoReq) (*types.QueryAVSInfoResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return k.GetAVSInfo(c, req.AVSAddres)
}
