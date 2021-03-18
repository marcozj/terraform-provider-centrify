---
subcategory: "Resources"
---

# centrifyvault_vaultaccount (Data Source)

This data source gets information of account.

## Example Usage

```terraform
data "centrifyvault_vaultaccount" "centos1_local_account" {
    name = "local_account"
    host_id = centrifyvault_vaultsystem.centos1.id
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultaccount)

## Search Attributes

### Required

- `name` - (String) Name of the account.

### Optional

- `host_id` - (String) ID of the system it belongs to.
- `domain_id` - (String) ID of the domain it belongs to.
- `database_id` - (String) ID of the database it belongs to.
- `cloudprovider_id` - (String) ID of the cloud provider it belongs to.
- `access_key_id` - (String) AWS access key id. Only applicable if this is cloud provider IAM account and `cloudprovider_id` is set.
- `checkout` - (Boolean) Whether to checkout the password, sshkey or AWS secret.
- `checkin` - (Boolean) Whether to checkin the password immediately after checkout. Only applicable if the account's credential type is password.
- `key_pair_type` - (String) SSH Key type. Can be set to `PublicKey`, `PrivateKey`, or `PPK`. Only appliable if the account's credential type is SSH key.
- `passphrase` - (String, Sensitive) Passphrase to use for encrypting the PrivateKey.

## Attributes Reference

- `id` - id of the account.
- `name` - name property.
- `password` - Account's clear text password.
- `private_key` - Account's SSh private key.
- `secret_access_key` - Cloud provider IAM account's secret key.
