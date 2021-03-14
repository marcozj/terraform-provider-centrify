---
page_title: "centrifyvault_role Data Source - terraform-provider-centrifyvault"
description: |-
  This data source gets information of role.
---

# centrifyvault_role (Data Source)

This data source gets information of role.

## Example Usage

```terraform
data "centrifyvault_role" "system_admin" {
    name = "System Admnistrator"
}
```

More examples can be found [here](../../examples/centrifyvault_role/)

## Search Attributes

### Required

- `name` - (String) Name of the role.

## Attributes Reference

- `id` - id of the role.
- `name` - Name property.
