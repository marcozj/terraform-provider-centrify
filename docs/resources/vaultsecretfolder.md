---
subcategory: "Resources"
---

# centrifyvault_vaultsecretfolder (Resource)

This resource allows you to create/update/delete secret folder.

## Example Usage

```terraform
resource "centrifyvault_vaultsecretfolder" "level1_folder" {
    name = "Level 1 Folder"
    description = "Level 1 Folder"
}

resource "centrifyvault_vaultsecretfolder" "level2_folder" {
    name = "Level 2 Folder"
    description = "Level 2 Folder"
    parent_id = centrifyvault_vaultsecretfolder.level1_folder.id
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsecret)

## Argument Reference

### Required

- `name` - (String) The name of the secret folder.

### Optional

- `description` - (String) Description of the secret folder.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `parent_id` - (String) Parent folder ID of an secret folder.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `member_permission` - (Block Set) Set member permissions. Refer to [member_permission attribute](./attribute_permission.md) for details.
