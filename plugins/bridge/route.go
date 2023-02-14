package bridge

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"

	"github.com/aximchain/flash-node/plugins/bridge/types"
)

func Routes(keeper Keeper) map[string]sdk.Handler {
	routes := make(map[string]sdk.Handler)
	routes[types.RouteBridge] = NewHandler(keeper)
	return routes
}
