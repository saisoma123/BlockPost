package keeper

import (
	"fmt"
	loggy "log"

	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	gonanoid "github.com/matoous/go-nanoid/v2"

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
func (k Keeper) AddMessage(ctx sdk.Context, creator string, message string) (string, error) {
	// Opens the KVStore
	store := k.storeService.OpenKVStore(ctx)

	messageID := genMessageID()

	// Instantiates MsgBlockPostMessage with sender address and message
	msg := types.MsgBlockPostMessage{
		Creator: creator,
		Message: message,
	}

	// Converts message to binary format for storing in KVStore
	bz, err := k.cdc.Marshal(&msg)

	if err != nil {
		return "", errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, "error marshalling message: %v", err)
	}

	// Stores marshalled message with ID converted to bytes array as key
	set_err := store.Set([]byte(messageID), bz)
	if set_err != nil {
		return "", errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "error storing message with ID %s: %v", messageID, set_err)
	}

	return messageID, nil
}

// Generates a unique ID per string stored
func genMessageID() string {

	// Generates a new, compact NanoID
	id, err := gonanoid.New()
	if err != nil {
		loggy.Fatalf("Failed to generate NanoID: %v", err)
	}
	return id
}

// Retrieves message from store with id
func (k Keeper) GetMessage(ctx sdk.Context, id string) (string, error) {
	//Opens store
	store := k.storeService.OpenKVStore(ctx)

	// Retrieves the marshalled message
	bz, err := store.Get([]byte(id))
	if bz == nil {
		return "", errorsmod.Wrapf(sdkerrors.ErrNotFound, "message with ID %s not found", id)
	}
	if err != nil {
		return "", errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "error retrieving message with ID %s: %v", id, err)
	}

	// Unmarshalls the message into msg and returns the Message field
	var msg types.MsgBlockPostMessage
	unmarshal_error := k.cdc.Unmarshal(bz, &msg)
	if unmarshal_error != nil {
		return "", errorsmod.Wrapf(sdkerrors.ErrJSONUnmarshal, "error unmarshalling message with ID %s: %v", id, unmarshal_error)
	}
	return msg.Message, nil
}

// Retrieves all messages from store
func (k Keeper) GetAllMessages(ctx sdk.Context) ([]string, error) {
	// Opens the store
	store := k.storeService.OpenKVStore(ctx)

	var messages []string

	iterator, err := store.Iterator(nil, nil)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	// Simply iterates throguh the store and unmarshals each message and appends
	// to an array, and the array is returned

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgBlockPostMessage

		err := k.cdc.Unmarshal(iterator.Value(), &msg)
		if err != nil {
			return nil, nil
		}

		if msg.Message == "" {
			continue
		}

		messages = append(messages, msg.Message)
	}

	return messages, nil
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
