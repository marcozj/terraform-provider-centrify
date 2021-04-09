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
- `name` - (String) The name of the multiplexed account.
- `description` - (String) Description of the multiplexed account.
- `accounts` - (Block Set) List of assigned account IDs.
- `account1_id` - (String) ID of assigned account1.
- `account2_id` - (String) ID of assigned account2.
- `account1` - (String) Name of assigned account1.
- `account2` - (String) Name of assigned account2.
- `active_account` - (String) Name of current active account.
