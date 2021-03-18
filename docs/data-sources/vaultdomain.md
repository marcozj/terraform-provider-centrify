---
subcategory: "Resources"
---

# centrifyvault_vaultdomain (Data Source)

This data source gets information of domain.

## Example Usage

```terraform
data "centrifyvault_vaultdomain" "example.com" {
    name = "example.com"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomain)

## Search Attributes

### Required

- `name` - (String) Name of the domain.

## Attributes Reference

- `id` - id of the domain.
- `name` - name property.
