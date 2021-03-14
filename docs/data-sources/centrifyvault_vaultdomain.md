---
page_title: "centrifyvault_vaultdomain Data Source - terraform-provider-centrifyvault"
description: |-
  This data source gets information of domain.
---

# centrifyvault_vaultdomain (Data Source)

This data source gets information of domain.

## Example Usage

```terraform
data "centrifyvault_vaultdomain" "example.com" {
    name = "example.com"
}
```

More examples can be found [here](../../examples/centrifyvault_vaultdomain/)

## Search Attributes

### Required

- `name` - (String) Name of the domain.

## Attributes Reference

- `id` - id of the domain.
- `name` - name property.
