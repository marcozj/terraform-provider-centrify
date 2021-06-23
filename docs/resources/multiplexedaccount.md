---
subcategory: "Resources"
---

# centrify_multiplexedaccount (Resource)

This resource allows you to create/update/delete multiplexed account.

## Example Usage

```terraform
resource "centrify_multiplexedaccount" "testmultiplex" {
    name = "Account for TestWindowsService"
    description = "Multiplexed account for TestWindowsService"
    accounts = [
        centrify_account.test_svc1.id,
        centrify_account.test_svc2.id,
    ]
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_service)

## Argument Reference

### Required

- `name` - (String) The name of the multiplexed account.
- `accounts` (Set of String) IDs of the 2 accounts. There must be 2 accounts.

### Optional

- `description` - (String) Description of the multiplexed account.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.

## Import

Multiplexed Account can be imported using the resource `id`, e.g.

```shell
terraform import centrify_multiplexedaccount.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` isn't support in import process.
