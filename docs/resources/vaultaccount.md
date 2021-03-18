---
subcategory: "Resources"
---

# centrifyvault_vaultaccount (Resource)

This resource allows you to create/update/delete account.

## Example Usage (Resource)

```terraform
resource "centrifyvault_vaultaccount" "unix_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrifyvault_vaultsystem.unix1.id
    description = "Test Account for Unix"
    use_proxy_account = false
    checkout_lifetime = 70
    managed = false
}
```

Examples of system account can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsystem)

Examples of database account can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdatabase)

Examples of domain account can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomain)

Examples of cloud provider account can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_cloudprovider)

## Argument Reference

### Required

- `name` - (String) Name of the account.
- `credential_type` - (String) Credential type of the account. Can be set to `Password`, `SshKey` or `AwsAccessKey`.

### Optional

- `use_proxy_account` - (Boolean) Use proxy account to manage this account. Only applicable if `credential_type` is `Password` and the system that this account belongs to has `proxyuser` configured.
- `managed` - (Boolean) If this account is managed. By enabling this option the credential will be automatically changed and become unknown to other applications or users.
- `description` - (String) Description of the account.
- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid. Range between `15` to `2147483647`. **Note:** Do NOT set this if it is IAM user.
- `challenge_rule` - (Block List) Password checkout challenge rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default password checkout profile (used if no conditions matched).
- `access_secret_checkout_default_profile_id` - (String) "Default secret access key checkout challenge rule ID. Only applicable to AWS IAM user.
- `access_secret_checkout_rule` - (Block List) Secret Access Key Checkout Challenge Rules. Only applicable to AWS IAM user. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `password` - (String, Sensitive) Password of the account. Only applicable if `credential_type` is `Password`.
- `sshkey_id` - (String) ID of the SSH key. Only applicable if `credential_type` is `SshKey`.
- `access_key` - (Block Set) AWS Access Keys (see [reference for `access_key`](#reference-for-access_key))
- `is_admin_account` - (Boolean) Whether this is an administrative account.
- `is_root_account` - (Boolean) Whether this is an root account for cloud provider. Only applicable if `credential_type` is `AwsAccessKey`.
- `host_id` - (String) ID of the system it belongs to.
- `domain_id` - (String) ID of the domain it belongs to.
- `database_id` - (String) ID of the database it belongs to.
- `cloudprovider_id` - (String) ID of the cloud provider it belongs to.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the account belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Reference for `access_key`

Required:

- `access_key_id` - (String) AWS access key id.
- `secret_access_key` - (String, Sensitive) AWS secret access key.
