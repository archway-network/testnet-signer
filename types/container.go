package types

import (
	"encoding/base64"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

type Container struct {
	ID        ID     `json:"id"`
	Signature string `json:"signature"`
}

func (c Container) VerifySubmission(kr keyring.Keyring, keyName string) (bool, error) {
	marshalledId, err := json.Marshal(c.ID)
	if err != nil {
		return false, err
	}

	testKey, err := kr.Key(keyName)
	if err != nil {
		return false, err
	}

	sigBytes, err := base64.StdEncoding.DecodeString(c.Signature)
	if err != nil {
		return false, err
	}

	return testKey.GetPubKey().VerifySignature(marshalledId, sigBytes), nil
}
