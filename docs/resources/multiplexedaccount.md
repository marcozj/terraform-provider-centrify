---
subcategory: "Resources"
---

# centrifyvault_multiplexedaccount (Resource)

This resource allows you to create/update/delete multiplexed account.

## Example Usage

```terraform
resource "centrifyvault_multiplexedaccount" "testmultiplex" {
    name = "Account for TestWindowsService"
    description = "Multiplexed account for TestWindowsService"
    accounts = [
        centrifyvault_vaultaccount.test_svc1.id,
        centrifyvault_vaultaccount.test_svc2.id,
    ]
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_service)

## Argument Reference

### Required

- `name` - (String) The name of the multiplexed account.
- `accounts` (Set of String) IDs of the 2 accounts. There must be 2 accounts.

### Optional

- `description` - (String) Description of the multiplexed account.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
