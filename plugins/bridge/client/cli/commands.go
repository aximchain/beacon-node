package cli

import (
	"github.com/aximchain/axc-cosmos-sdk/client"
	"github.com/aximchain/axc-cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, cdc *codec.Codec) {
	bridgeCmd := &cobra.Command{
		Use:   "bridge",
		Short: "bridge commands",
	}

	bridgeCmd.AddCommand(
		client.PostCommands(
			BindCmd(cdc),
			TransferOutCmd(cdc),
			UnbindCmd(cdc),
		)...,
	)

	bridgeCmd.AddCommand(client.LineBreak)

	bridgeCmd.AddCommand(
		client.GetCommands(
			QueryProphecy(cdc))...,
	)
	cmd.AddCommand(bridgeCmd)
}
