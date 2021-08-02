---
subcategory: "Settings"
---

# centrify_federatedgroup (Data Resource)

This data source gets information of federated group.

## Example Usage

```terraform
data "centrify_federatedgroup" "fedgroup1" {
  name = "Okta Infra Admins"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_role/role_member_with_federatedgroup.tf)

## Argument Reference

### Required

- `name` - (String) Name of the fedreated group.