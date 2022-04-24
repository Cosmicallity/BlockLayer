package keeper

import (
	"github.com/Cosmicallity/BlockLayer/x/blocklayer/types"
)

var _ types.QueryServer = Keeper{}
