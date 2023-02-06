package bridge

import (
	"github.com/aximchain/beacon-node/plugins/bridge/keeper"
	"github.com/aximchain/beacon-node/plugins/bridge/types"
)

var (
	NewKeeper = keeper.NewKeeper
)

type (
	Keeper = keeper.Keeper

	TransferOutMsg = types.TransferOutMsg
	BindMsg        = types.BindMsg
	UnbindMsg      = types.UnbindMsg
)
