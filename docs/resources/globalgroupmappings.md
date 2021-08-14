---
subcategory: "Settings"
---

# centrify_globalgroupmappings (Resource)

This resource allows you to create/update/delete global federated group mapping.

~> **WARNING:** Multiple `centrify_globalgroupmappings` resources will produce inconsistent behavior! Do NOT use more than once!

## Example Usage

```terraform
resource "centrify_globalgroupmappings" "group_mappings" {
    bulkupdate = true
    mapping = {
        "Idp Group 1" = "Okta PAS Admin"
        "Idp Group 2" = "Azure PAS Users"
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_globalgroupmappings)

## Argument Reference

### Optional

- `bulkupdate` - (Bollean) When this is set to true, one API call is issued to perform create/update/delete for multiple mappings instead of one API call per mapping. This improves performance when there are large number of mappings. Default is `true`. **NOTE:** When this is true, existing mappings not managed by Terraform will be removed when create/update/delete actions are performed.
- `mapping` - (Map) Map entries in <group attribute value> = <group name> format. Group attribute value is group or role name from IdP side. Group name is the virtual group in Centrify side. If the group doesn't exist, it will be created.
