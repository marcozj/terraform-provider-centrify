---
subcategory: "Settings"
---

# centrify_federatedgroup (Resource)

This resource allows you to create federated group by leveraging on `centrify_globalgroupmappings` resource. It doesn't actuall create federated group but search the group internally created by `centrify_globalgroupmappings` instead.

## Example Usage

```terraform
resource "centrify_globalgroupmappings" "group_mappings" {
    bulkupdate = true
    mapping = {
        "Idp Group 1" = "Okta Infra Admins"
        "Idp Group 2" = "Azure PAS Users" // Assuming "Azure PAS Users" doesn't exist yet and will be created by this resource
    }
}

// New federated group created by centrify_globalgroupmappings
resource "centrify_federatedgroup" "fedgroup2" {
    name = centrify_globalgroupmappings.group_mappings.mapping["Idp Group 2"] // Reference to "Idp Group 2" map which returns "Azure PAS Users"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_role/role_member_with_federatedgroup.tf)

## Argument Reference

### Required

- `name` - (String) Name of the fedreated group. Do NOT set string of federated group directly, make sure to reference to the map entry from `centrify_globalgroupmappings` resource so that the federated group is created.