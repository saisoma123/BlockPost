package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

func (k msgServer) BlockPostMessage(goCtx context.Context, msg *types.MsgBlockPostMessage) (*types.MsgBlockPostMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBlockPostMessageResponse{}, nil
}
