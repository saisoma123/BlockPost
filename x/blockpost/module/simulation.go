package blockpost

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/saisoma123/BlockPost/testutil/sample"
	blockpostsimulation "github.com/saisoma123/BlockPost/x/blockpost/simulation"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

// avoid unused import issue
var (
	_ = blockpostsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgBlockPostMessage = "op_weight_msg_block_post_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBlockPostMessage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blockpostGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blockpostGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgBlockPostMessage int
	simState.AppParams.GetOrGenerate(opWeightMsgBlockPostMessage, &weightMsgBlockPostMessage, nil,
		func(_ *rand.Rand) {
			weightMsgBlockPostMessage = defaultWeightMsgBlockPostMessage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBlockPostMessage,
		blockpostsimulation.SimulateMsgBlockPostMessage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgBlockPostMessage,
			defaultWeightMsgBlockPostMessage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blockpostsimulation.SimulateMsgBlockPostMessage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
