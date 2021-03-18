---
subcategory: "Resources"
---

# centrifyvault_vaultsecret (Data Source)

This data source gets information of secret.

## Example Usage

```terraform
data "centrifyvault_vaultsecret" "test_secret" {
    secret_name = "testsecret"
    checkout = true
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsecret)

## Search Attributes

### Required

- `secret_name` - (String) Name of the secret.

### Optional

- `folder_id` - (String) ID of the folder where the secret is located.
- `checkout` - (Boolean) Whether to retrieve secret content. Default is `false`. If `true`, `secret_text` will be populated.

## Attributes Reference

- `id` - id of the secret.
- `description` - description property.
- `parent_path` - parent_path property.
- `secret_text` - secret_text property.
