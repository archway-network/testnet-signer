package cmd

import (
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{Use: "keys"}

func init() {
	keysCmd.AddCommand(keys.AddKeyCommand())
	keysCmd.AddCommand(keys.ShowKeysCmd())
	keysCmd.AddCommand(keys.ListKeysCmd())
	rootCmd.AddCommand(keysCmd)
}
