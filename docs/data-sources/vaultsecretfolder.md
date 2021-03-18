---
subcategory: "Resources"
---

# centrifyvault_vaultsecretfolder (Data Source)

This data source gets information of secret folder.

## Example Usage

```terraform
data "centrifyvault_vaultsecretfolder" "level1_folder" {
    name = "Level 1 Folder"
}

data "centrifyvault_vaultsecretfolder" "level2_folder" {
    name = "Level 2 Folder"
    parent_path = "Level 1 Folder"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsecret)

## Search Attributes

### Required

- `name` - (String) The name of the secret folder.

### Optional

- `parent_path` - (String) Parent folder path of an secret folder.

## Attributes Reference

- `id` - id of the secret folder.
- `description` - description property.
- `parent_path` - parent_path property.
