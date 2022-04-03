# Augusta testnet signer utility

This utility generates a signed JSON-formatted ID to prove ownership of a key used to submit tx on the blockchain. This testnet signer utility is required for all users who participate in testnet challenges that involve sending tx on the testnet. The challenge submission requires the ID message be signed by the primary key used to deploy the smart contract.

# Installation

Go `1.16` is required.

If you haven't already, download and install Go. See the official [go.dev documentation](https://golang.org/doc/install). Make sure your `GOBIN` and `GOPATH` are set correctly.

## Get the source code and install

```bash
git clone git@github.com:archway-network/testnet-signer.git
cd testnet-signer
make install
```

# How the testnet signer utility works 

## Generate a new key 

**Note:** Required only to generate a new key for the testnet. If you want to use an existing key in the keystore, you can skip this step.

### Description
Adds a new key to the keystore. This command is compatible with the Cosmos SDK key management functionality.

### Command
```shell
testnet-signer keys add <key-name>
```

## Generate a signed JSON-formatted ID

### Description
Generates the required signed id message to submit in the testnet challenge form. Be sure to provide accurate details.

```shell
% testnet-signer sign_id my-key
Enter information as accurately as possible. Information entered here must match your KYC.
Your full legal name:FirstName MiddleName LastName
Your GitHub handle:mygithub
Your email address:myemail@domain.com
Your Augusta incentivized testnet address is:  archway1lf26gv87sxvkj59e3f9q2fh6q8phqwgje6g3xg
Amino encoded Public key is: 61rphyEDtd8YCbk465UwocPsEcaSNn3IHx7zUa7tUdoAOuy/iyw=
Submit JSON below the line in the form.
-----------------------------
{
  "id": {
    "full_legal_name": "FirstName MiddleName LastName",
    "github_handle": "mygithub",
    "email_address": "myemail@domain.com",
    "account_address": "archway1lf26gv87sxvkj59e3f9q2fh6q8phqwgje6g3xg",
    "pub_key": "61rphyEDtd8YCbk465UwocPsEcaSNn3IHx7zUa7tUdoAOuy/iyw="
  },
  "signature": "Fnsuzh71v9FJtaz6hdRWsKstGeE1mexEClq67OPuzaZdBKmurXo8P6Himu69mmEsCcz+YGtQV/204XSX0lmnMQ=="
}

```


