package cmd

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	cosmosFlag "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"os"
)

var rootCmd = &cobra.Command{Use: "archway-testnet-signer"}

func init() {
	rootCmd.PersistentFlags().String(cosmosFlag.FlagKeyringBackend, "os", "Keyring backend to use, default value is: os")
	rootCmd.PersistentFlags().String(cosmosFlag.FlagKeyringDir, os.ExpandEnv("$HOME/")+".augusta", "Keyring backend directory")
	rootCmd.PersistentFlags().String(cli.OutputFlag, "text", "Output format (text|json)")
}

func Execute() {
	initClientCtx := client.Context{}.WithInput(os.Stdin)
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &initClientCtx)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
