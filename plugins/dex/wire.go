package dex

import (
	"github.com/aximchain/beacon-node/plugins/dex/order"
	"github.com/aximchain/beacon-node/plugins/dex/store"
	"github.com/aximchain/beacon-node/plugins/dex/types"
	"github.com/aximchain/beacon-node/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(Genesis{}, "dex/Genesis", nil)

	cdc.RegisterConcrete(order.NewOrderMsg{}, "dex/NewOrder", nil)
	cdc.RegisterConcrete(order.CancelOrderMsg{}, "dex/CancelOrder", nil)

	cdc.RegisterConcrete(types.ListMsg{}, "dex/ListMsg", nil)
	cdc.RegisterConcrete(types.TradingPair{}, "dex/TradingPair", nil)

	cdc.RegisterConcrete(types.ListMiniMsg{}, "dex/ListMiniMsg", nil)

	cdc.RegisterConcrete(order.FeeConfig{}, "dex/FeeConfig", nil)
	cdc.RegisterConcrete(order.OrderBookSnapshot{}, "dex/OrderBookSnapshot", nil)
	cdc.RegisterConcrete(order.ActiveOrders{}, "dex/ActiveOrders", nil)
	cdc.RegisterConcrete(store.RecentPrice{}, "dex/RecentPrice", nil)
}
