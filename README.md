# Augusta testnet signer utility
This utility enables anybody who wishes to participate in challenges involving sending tx on the testnet, generate a id json with signature
that proves ownership of a key used to submit tx on blockchain.

# Installation

Go `1.16` is required.

If you haven't already, download and install Go. See the official [go.dev documentation](https://golang.org/doc/install). Make sure your `GOBIN` and `GOPATH` are setup.

## Get the source code & install

```bash
git clone git@github.com:archway-network/augusta-testnet-signer.git
cd augusta-testnet-signer
make install
```

# How it works? 

## Generate a key (Optional)
**Note:** Only do this if you want to generate new key for testnet, if you already have a key in keystore you want to use
skip this step.

### Description
Adds a new key to the keystore. This command is compatible with the cosmos-sdk key management functionality.

### Command
```shell
augusta-testnet-signer keys add <key-name>
```

## Generate a signed id json

### Description
Generates a signed id message that we will need to submit in the form, after asking series of questions. Please make sure to fill the details correctly.

```shell
% augusta-testnet-signer sign_id my-key
Please enter information as accurate as possible, information entered here must match your KYC
Please Enter your Full Legal Name:FirstName MiddleName LastName
Please enter your github handle:mygithub
Please enter your email address:myemail@domain.com
Your Augusta incentivized testnet address is:  archway1lf26gv87sxvkj59e3f9q2fh6q8phqwgje6g3xg
Amino encoded Public key is: 61rphyEDtd8YCbk465UwocPsEcaSNn3IHx7zUa7tUdoAOuy/iyw=
Please submit json below the line in the form
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


