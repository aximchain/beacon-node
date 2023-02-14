package testutils

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/bank"

	"github.com/aximchain/flash-node/common/types"
	"github.com/aximchain/flash-node/plugins/dex/order"
	"github.com/aximchain/flash-node/plugins/tokens"
	"github.com/aximchain/flash-node/wire"
)

func MakeCodec() *wire.Codec {
	var cdc = wire.NewCodec()

	wire.RegisterCrypto(cdc) // Register crypto.
	bank.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc) // Register Msgs
	tokens.RegisterWire(cdc)
	types.RegisterWire(cdc)
	cdc.RegisterConcrete(order.NewOrderMsg{}, "dex/NewOrder", nil)
	cdc.RegisterConcrete(order.CancelOrderMsg{}, "dex/CancelOrder", nil)

	cdc.RegisterConcrete(order.OrderBookSnapshot{}, "dex/OrderBookSnapshot", nil)
	cdc.RegisterConcrete(order.ActiveOrders{}, "dex/ActiveOrders", nil)

	return cdc
}
