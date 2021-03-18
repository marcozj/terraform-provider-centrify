---
subcategory: "Resources"
---

# centrifyvault_multiplexedaccount (Data Source)

This data source gets information of multiplexed account.

## Example Usage

```terraform
data "centrifyvault_multiplexedaccount" "testmultiplex" {
  name = "Test Multiplex"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_service)

## Search Attributes

### Required

- `name` - (String) The name of the multiplexed account.

## Attributes Reference

- `id` - id of the multiplexed account.
- `name` - name property.
- `description` - description property.
- `active_account` - active_account property.
