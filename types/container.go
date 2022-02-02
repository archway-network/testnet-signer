package types

import (
	"encoding/base64"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type Container struct {
	ID        ID     `json:"id"`
	Signature string `json:"signature"`
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
