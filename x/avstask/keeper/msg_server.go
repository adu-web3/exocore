package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	"github.com/ExocoreNetwork/exocore/x/avstask/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = &Keeper{}

func (k *Keeper) RegisterAVSTask(ctx context.Context, req *types.RegisterAVSTaskReq) (*types.RegisterAVSTaskResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	avs, err := k.avsKeeper.GetAVSInfo(c, req.FromAddress)
	if err != nil {
		return nil, err
	}
	if avs.GetInfo() == nil {
		return nil, errorsmod.Wrap(types.ErrNotYetRegistered, fmt.Sprintf("RegisterAVSTask: avs address is %s", req.FromAddress))
	}
	err = k.SetAVSTaskInfo(c, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
