---
page_title: "centrifyvault_multiplexedaccount Data Source - terraform-provider-centrifyvault"
description: |-
  This data source gets information of multiplexed account.
---

# centrifyvault_multiplexedaccount (Data Source)

This data source gets information of multiplexed account.

## Example Usage

```terraform
data "centrifyvault_multiplexedaccount" "testmultiplex" {
  name = "Test Multiplex"
}
```

More examples can be found [here](../../examples/centrifyvault_service/)

## Search Attributes

### Required

- `name` - (String) The name of the multiplexed account.

## Attributes Reference

- `id` - id of the multiplexed account.
- `name` - name property.
- `description` - description property.
- `active_account` - active_account property.
