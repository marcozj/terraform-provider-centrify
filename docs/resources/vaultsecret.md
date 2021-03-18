---
subcategory: "Resources"
---

# centrifyvault_vaultsecret (Resource)

This resource allows you to create/update/delete secret.

## Example Usage

```terraform
resource "centrifyvault_vaultsecret" "test_secret" {
    secret_name = "Test Secret"
    description = "Test Secret"
    secret_text = "xxxxxxxxxxxxx"
    type = "Text"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsecret)

## Argument Reference

### Required

- `secret_name` - (String) Name of the secret.
- `type` - (String) Type of the secret. Can be set to `Text` only at the moment. (TODO: `File`)

### Optional

- `description` - (String) Description of the secret.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `folder_id` - (String) ID of the folder where the secret is located.
- `parent_path` - (String) Path of parent folder
- `secret_text` - (String, Sensitive) Content of the secret.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.
