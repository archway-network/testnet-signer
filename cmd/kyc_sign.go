package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/archway-network/augusta-testnet-signer/types"
	cosmosFlag "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"net/mail"
	"os"
	"strings"
)

var signIDCmd = &cobra.Command{
	Use:   "sign_id [key_name]",
	Short: "Sign an id message with your key to prove that you own the key",
	Long: `This command asks you some personal details like legal name, email address and github handle, 
           and uses the key associated with [key_name] to sign the message containing that data and print it`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyringBackend, err := cmd.Flags().GetString(cosmosFlag.FlagKeyringBackend)
		if err != nil {
			return err
		}

		backendDir, err := cmd.Flags().GetString(cosmosFlag.FlagKeyringDir)
		if err != nil {
			return err
		}

		kr, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, backendDir, os.Stdin)
		if err != nil {
			return err
		}

		keyName := args[0]
		_, err = kr.Key(keyName)
		if err != nil {
			return err
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please Enter your Full Legal Name:")
		fullLegalName, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		fullLegalName = strings.TrimSpace(fullLegalName)
		if len(fullLegalName) < 3 || len(fullLegalName) > 512 {
			return fmt.Errorf("full legal name need to be between 3 and 256 letter")
		}

		fmt.Printf("Please enter your github handle:")
		githubHandle, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		githubHandle = strings.TrimSpace(githubHandle)
		if strings.ContainsAny(githubHandle, " \t\n\r") {
			return fmt.Errorf("github handle cannot contain whitespace character")
		}
		if len(githubHandle) < 1 || len(githubHandle) > 39 {
			return fmt.Errorf("github handle need to be between 1 and 38 letter")
		}

		fmt.Printf("Please enter your email address (Use same email address, you would use in kyc form):")
		emailAddress, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		emailAddress = strings.TrimSpace(emailAddress)
		_, err = mail.ParseAddress(emailAddress)
		if err != nil {
			return err
		}

		container, err := types.CreateContainer(fullLegalName, githubHandle, emailAddress, keyName, kr)
		if err != nil {
			return err
		}

		verified, err := container.VerifySubmission()
		if err != nil {
			return fmt.Errorf("error while verifying id submission, err: %v", err)
		}
		if !verified {
			return fmt.Errorf("signature verification failed. This may be due to bug in the program")
		}

		fmt.Println("Your Augusta incentivized testnet address is: ", container.ID.AccountAddress)
		fmt.Println("Amino encoded Public key is:", container.ID.PubKey)

		marshalledContainer, err := json.MarshalIndent(container, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println("Please submit json below the line in the form")
		fmt.Println("-----------------------------")
		fmt.Println(string(marshalledContainer))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(signIDCmd)
}
