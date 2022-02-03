package types

import (
	"encoding/base64"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Container struct {
	ID        ID     `json:"id"`
	Signature string `json:"signature"`
}

func CreateContainer(fullLegalName, githubHandle, emailAddress, pubKeyId string, kr keyring.Keyring) (Container, error) {
	var container Container

	keyInfo, err := kr.Key(pubKeyId)
	if err != nil {
		return container, err
	}

	pubKey := keyInfo.GetPubKey()

	pubKeyBytes, err := legacy.Cdc.Marshal(pubKey)
	if err != nil {
		return container, err
	}

	pubKeyStr := base64.StdEncoding.EncodeToString(pubKeyBytes)
	container.ID.PubKey = pubKeyStr
	container.ID.AccountAddress = sdk.AccAddress(keyInfo.GetPubKey().Address()).String()
	container.ID.EmailAddress = emailAddress
	container.ID.FullLegalName = fullLegalName
	container.ID.GithubHandle = githubHandle

	marshalledContainer, err := json.Marshal(container.ID)
	if err != nil {
		return container, err
	}

	signature, _, err := kr.Sign(pubKeyId, marshalledContainer)
	if err != nil {
		return container, err
	}

	container.Signature = base64.StdEncoding.EncodeToString(signature)

	return container, nil
}

func (c Container) VerifySubmission() (bool, error) {
	marshalledId, err := json.Marshal(c.ID)
	if err != nil {
		return false, err
	}

	pubkeyBytes, err := base64.StdEncoding.DecodeString(c.ID.PubKey)
	if err != nil {
		return false, err
	}

	var pubkey cryptotypes.PubKey
	err = legacy.Cdc.Unmarshal(pubkeyBytes, &pubkey)
	if err != nil {
		return false, err
	}

	sigBytes, err := base64.StdEncoding.DecodeString(c.Signature)
	if err != nil {
		return false, err
	}

	return pubkey.VerifySignature(marshalledId, sigBytes), nil
}
