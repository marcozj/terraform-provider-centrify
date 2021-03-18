---
subcategory: "Settings"
---

# centrifyvault_globalgroupmappings (Resource)

This resource allows you to create/update/delete global federated group mapping.

## Example Usage

```terraform
resource "centrifyvault_globalgroupmappings" "group_mappings" {
    mapping {
        attribute_value = "Idp Group 1"
        group_name = "Okta PAS Admin"
    }
    mapping {
        attribute_value = "Idp Group 2"
        group_name = "Azure PAS Users"
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_globalgroupmappings)

## Argument Reference

### Optional

- `mapping` - (Block Set) (see [below reference for mapping](#reference-for-mapping))

## [Reference for `mapping`]

Required:

- `attribute_value` - (String) Group attribute value. This is group or role name from IdP side.
- `group_name` - (String) Group name. This is the virtual group in Centrify side. If the group doesn't exist, it will be created.
