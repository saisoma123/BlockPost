package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
	}
}

// Adds message to KVStore with unique id
func (k Keeper) addMessage(ctx sdk.Context, creator string, message string) (*sdk.Result, error) {
	return nil, nil
}

// Generates a unique ID per string stored
func (k Keeper) genMessageID(ctx sdk.Context) string {
	return ""
}

// Retrieves message from store with id
func (k Keeper) getMessage(ctx sdk.Context, id string) (types.MsgBlockPostMessage, error) {
	return types.MsgBlockPostMessage{}, nil
}

// Retrieves all messages from store
func (k Keeper) getAllMessages(ctx sdk.Context) ([]types.MsgBlockPostMessage, error) {
	return []types.MsgBlockPostMessage{}, nil
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
