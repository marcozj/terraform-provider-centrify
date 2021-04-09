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
- `password` - (String, Sensitive) Password of the account.
- `private_key` - String, Sensitive) Account's SSh private key.
- `secret_access_key` - String, Sensitive) Cloud provider IAM account's secret key.

- `use_proxy_account` - (Boolean) Use proxy account to manage this account.
- `managed` - (Boolean) If this account is managed. By enabling this option the credential will be automatically changed and become unknown to other applications or users.
- `description` - (String) Description of the account.
- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid.
- `challenge_rule` - (Block List) Password checkout challenge rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default password checkout profile (used if no conditions matched).
- `access_secret_checkout_default_profile_id` - (String) "Default secret access key checkout challenge rule ID. Only applicable to AWS IAM user.
- `access_secret_checkout_rule` - (Block List) Secret Access Key Checkout Challenge Rules. Only applicable to AWS IAM user. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `sshkey_id` - (String) ID of the SSH key.
- `access_key` - (Block Set) AWS Access Keys (see [reference for `access_key`](#reference-for-access_key))
- `is_admin_account` - (Boolean) Whether this is an administrative account.
- `is_root_account` - (Boolean) Whether this is an root account for cloud provider. Only applicable if `credential_type` is `AwsAccessKey`.
- `host_id` - (String) ID of the system it belongs to.
- `domain_id` - (String) ID of the domain it belongs to.
- `database_id` - (String) ID of the database it belongs to.
- `cloudprovider_id` - (String) ID of the cloud provider it belongs to.
- `workflow_enabled` - (Boolean) Enable account workflow.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
- `credential_type` - (String) Type of account credential.
- `credential_name` - (String) Name of SSH Key.
