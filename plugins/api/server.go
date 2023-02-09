package api

import (
	"github.com/gorilla/mux"

	"github.com/aximchain/axc-cosmos-sdk/client/context"
	keyscli "github.com/aximchain/axc-cosmos-sdk/client/keys"
	"github.com/aximchain/axc-cosmos-sdk/crypto/keys"

	"github.com/aximchain/flash-node/common"
	"github.com/aximchain/flash-node/plugins/tokens"
	"github.com/aximchain/flash-node/wire"
)

// config consts
const maxPostSize int64 = 1024 * 1024 * 0.5 // ~500KB

type server struct {
	router *mux.Router

	// settings
	maxPostSize int64

	// handler dependencies
	ctx context.CLIContext
	cdc *wire.Codec

	// stores for handlers
	keyBase keys.Keybase
	tokens  tokens.Mapper

	accStoreName string
}

// NewServer provides a new server structure.
func newServer(ctx context.CLIContext, cdc *wire.Codec) *server {
	kb, err := keyscli.GetKeyBase()
	if err != nil {
		panic(err)
	}

	return &server{
		router:       mux.NewRouter(),
		maxPostSize:  maxPostSize,
		ctx:          ctx,
		cdc:          cdc,
		keyBase:      kb,
		tokens:       tokens.NewMapper(cdc, common.TokenStoreKey),
		accStoreName: common.AccountStoreName,
	}
}
