package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	// axc prefix address:  axc1v8vkkymvhe2sf7gd2092ujc6hweta38xadu2pj
	// taxc prefix address: taxc1v8vkkymvhe2sf7gd2092ujc6hweta38xnc4wpr
	PegAccount = sdk.AccAddress(crypto.AddressHash([]byte("AximchainPegAccount")))
)
