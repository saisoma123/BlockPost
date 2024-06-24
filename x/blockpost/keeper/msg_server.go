package keeper

import (
	"context"

	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) BlockPostMessage(goCtx context.Context, msg *types.MsgBlockPostMessage) (*types.MsgBlockPostMessageResponse, error) {
	return &types.MsgBlockPostMessageResponse{}, nil
}

var _ types.MsgServer = msgServer{}
