package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

type queryServer struct {
	k Keeper
}

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

func (q queryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params := q.k.GetParams(sdk.UnwrapSDKContext(ctx))
	return &types.QueryParamsResponse{Params: params}, nil
}

// Calls the keeper GetMessage function and returns the message as a response
func (q queryServer) Message(ctx context.Context, req *types.QueryMessageRequest) (*types.QueryMessageResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	message, err := q.k.GetMessage(sdkCtx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryMessageResponse{Message: message}, nil
}

// Calls the keeper GetAllMessages function and returns the messages as a response
func (q queryServer) Messages(ctx context.Context, req *types.QueryAllMessagesRequest) (*types.QueryAllMessagesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	messages, err := q.k.GetAllMessages(sdkCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryAllMessagesResponse{Messages: messages}, nil
}
