---
subcategory: "Resources"
---

# centrify_secretfolder (Resource)

This resource allows you to create/update/delete secret folder.

## Example Usage

```terraform
resource "centrify_secretfolder" "level1_folder" {
    name = "Level 1 Folder"
    description = "Level 1 Folder"
}

resource "centrify_secretfolder" "level2_folder" {
    name = "Level 2 Folder"
    description = "Level 2 Folder"
    parent_id = centrify_secretfolder.level1_folder.id
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_secret)

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

## Import

Secret Folder can be imported using the resource `id`, e.g.

```shell
terraform import centrify_secretfolder.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` and `member_permission` aren't supported in import process.
