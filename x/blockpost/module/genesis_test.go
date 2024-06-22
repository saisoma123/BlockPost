package blockpost_test

import (
	"testing"

	keepertest "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/saisoma123/BlockPost/testutil/nullify"
	blockpost "github.com/saisoma123/BlockPost/x/blockpost/module"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlockpostKeeper(t)
	blockpost.InitGenesis(ctx, k, genesisState)
	got := blockpost.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
