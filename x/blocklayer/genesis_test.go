package blocklayer_test

import (
	"testing"

	keepertest "github.com/Cosmicallity/BlockLayer/testutil/keeper"
	"github.com/Cosmicallity/BlockLayer/testutil/nullify"
	"github.com/Cosmicallity/BlockLayer/x/blocklayer"
	"github.com/Cosmicallity/BlockLayer/x/blocklayer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlocklayerKeeper(t)
	blocklayer.InitGenesis(ctx, *k, genesisState)
	got := blocklayer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
