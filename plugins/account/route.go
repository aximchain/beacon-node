package account

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/auth"
)

func routes(accKeeper auth.AccountKeeper) map[string]sdk.Handler {
	routes := make(map[string]sdk.Handler)
	routes[AccountFlagsRoute] = NewHandler(accKeeper)
	return routes
}
