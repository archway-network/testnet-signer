# Augusta testnet signer utility
This utility enables anybody who wishes to participate in challenges involving sending tx on the testnet, generate a id json with signature
that proves ownership of a key used to submit tx on blockchain.

# How it works? 

## Generate a key (Optional)
*Note:* Only do this if you want to generate new key for testnet, if you already have a key in keystore you want to use
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
% augusta-testnet-signer sign_id my-key --keyring-backend test
Your Augusta incentivized testnet address is:  augusta1lf26gv87sxvkj59e3f9q2fh6q8phqwgje6g3xg
Please Enter your Full Legal Name:1
Please enter your github handle:2
Please enter your email address (Use same email address, you would use in kyc form):3
Please submit json below the line in the form
-----------------------------
{
  "id": {
    "full_legal_name": "1",
    "github_handle": "2",
    "email_address": "3",
    "account_address": "augusta1lf26gv87sxvkj59e3f9q2fh6q8phqwgje6g3xg",
    "pub_key": "PubKeySecp256k1{03B5DF1809B938EB9530A1C3EC11C692367DC81F1EF351AEED51DA003AECBF8B2C}"
  },
  "signature": "4DEq8ZnxvAj8nrsvVRE/WgW3hoozAnDe1k64E1ZU+dcZXU6U8a3tilkzHkhmgNLTyWd5FI2g+p8hRwSMb0xCTg=="
}
```


