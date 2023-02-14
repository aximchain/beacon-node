package tx

import (
	"github.com/aximchain/axc-cosmos-sdk/x/auth"

	"github.com/aximchain/flash-node/wire"
)

func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(&auth.StdTx{}, "auth/StdTx", nil)
}
