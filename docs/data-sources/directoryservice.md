---
subcategory: "Settings"
---

# centrifyvault_directoryservice (Data Source)

This data source gets information of directory service.

## Example Usage

```terraform
data "centrifyvault_directoryservice" "demo_lab" {
    // name is the actual Active Directory doman name
    name = "demo.lab"
    type = "Active Directory"
}

data "centrifyvault_directoryservice" "federated_dir" {
    // name must be "Federated Directory Service"
    name = "Federated Directory Service"
    type = "Federated Directory"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_role/role_member_with_federatedgroup.tf) and [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_role/role_member_with_adgroup.tf)

## Search Attributes

### Required

- `name` - (String) Name of the Directory Service. When type is `Federated Directory`, name must be `Federated Directory Service`.
- `type` - (String) Type of the Directory Service. Can be set to `Centrify Directory`, `Active Directory`, `Federated Directory`, `Google Directory` or `LDAP Directory`.

## Attributes Reference

- `id` - (String) The ID of this resource.
- `status` - (String) Status of the Directory Service.
