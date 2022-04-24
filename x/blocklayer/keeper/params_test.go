package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/Cosmicallity/BlockLayer/testutil/keeper"
	"github.com/Cosmicallity/BlockLayer/x/blocklayer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlocklayerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
