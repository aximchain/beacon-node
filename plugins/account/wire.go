package account

import (
	"github.com/aximchain/beacon-node/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(SetAccountFlagsMsg{}, "scripts/SetAccountFlagsMsg", nil)
}
