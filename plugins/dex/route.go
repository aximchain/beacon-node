package dex

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/gov"
	"github.com/aximchain/flash-node/plugins/dex/types"

	"github.com/aximchain/flash-node/plugins/dex/list"
	"github.com/aximchain/flash-node/plugins/dex/order"
	"github.com/aximchain/flash-node/plugins/tokens"
)

// Routes exports dex message routes
func Routes(dexKeeper *DexKeeper, tokenMapper tokens.Mapper, govKeeper gov.Keeper) map[string]sdk.Handler {
	routes := make(map[string]sdk.Handler)
	orderHandler := order.NewHandler(dexKeeper)
	routes[order.RouteNewOrder] = orderHandler
	routes[order.RouteCancelOrder] = orderHandler
	routes[types.ListRoute] = list.NewHandler(dexKeeper, tokenMapper, govKeeper)
	return routes
}
