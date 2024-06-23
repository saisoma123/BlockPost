package blockpost

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/saisoma123/BlockPost/x/blockpost/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

// Defined Handler type due to sdk.Handler deprecation

type Handler func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error)

// Handles incoming messages and calls appropriate handler with keeper

func NewHandler(keeper keeper.Keeper) Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgBlockPostMessage:
			return handleMsgBlockPostMessageStore(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized blockpost Msg type: %v", reflect.TypeOf(msg).String())
			return nil, sdkerrors.ErrUnknownRequest.Wrap(errMsg)
		}
	}
}

// Handles message storing and checks for any errors during store process

func handleMsgBlockPostMessageStore(ctx sdk.Context, keeper keeper.Keeper, msg sdk.Msg) (*sdk.Result, error) {
	return nil, nil
}
