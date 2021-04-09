---
subcategory: "Access"
---

# centrifyvault_role (Data Source)

This data source gets information of role.

## Example Usage

```terraform
data "centrifyvault_role" "system_admin" {
    name = "System Admnistrator"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_role)

## Search Attributes

### Required

- `name` - (String) Name of the role.

## Attributes Reference

- `id` - (String) ID of the member.
- `type` - (String) Type of the member.
- `description` - (String) Description of an role.
- `adminrights` - (Set of String) List of administrative rights.
- `member` - (Block Set) (see [below reference for member](#reference-for-member))
