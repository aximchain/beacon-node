package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/node/wire"
)

func RegisterWire(cdc *wire.Codec) {
	// Register AppAccount
	cdc.RegisterInterface((*sdk.Account)(nil), nil)
	cdc.RegisterInterface((*NamedAccount)(nil), nil)
	cdc.RegisterInterface((*IToken)(nil), nil)

	cdc.RegisterConcrete(&AppAccount{}, "axcchain/Account", nil)

	cdc.RegisterConcrete(&Token{}, "axcchain/Token", nil)
	cdc.RegisterConcrete(&MiniToken{}, "axcchain/MiniToken", nil)
}
