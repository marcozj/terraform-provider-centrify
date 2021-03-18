---
subcategory: "Resources"
---

# centrifyvault_sshkey (Resource)

This resource allows you to create/update/delete ssh key.

## Example Usage

```terraform
resource "centrifyvault_sshkey" "test_key" {
  name = "Test Key"
    description = "Test RSA key"
    private_key = file("rsa.key")
    passphrase = ""
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_sshkey)

## Argument Reference

### Required

- `name` - (String) Name of the SSH Key.

### Optional

- `description` - (String) Description of the SSH Key
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default SSH Key Challenge Profile ID (used if no conditions matched).
- `private_key` - (String, Sensitive) SSH private key.
- `passphrase` - (String, Sensitive) Passphrase to use for encrypting the PrivateKey.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.
