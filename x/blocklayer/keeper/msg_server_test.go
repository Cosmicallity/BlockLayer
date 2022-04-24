package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Cosmicallity/BlockLayer/x/blocklayer/types"
    "github.com/Cosmicallity/BlockLayer/x/blocklayer/keeper"
    keepertest "github.com/Cosmicallity/BlockLayer/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlocklayerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
