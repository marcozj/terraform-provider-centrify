---
subcategory: "Settings"
---

# centrifyvault_directoryobject (Data Source)

This data source gets information of directory object such as account and group.

## Example Usage

```terraform
// data source for Active Directory domain demo.lab
data "centrifyvault_directoryservice" "demo_lab" {
    name = "demo.lab"
    type = "Active Directory"
}

// data source for AD user ad.user@demo.lab
data "centrifyvault_directoryobject" "ad_user" {
    directory_services = [
        data.centrifyvault_directoryservice.demo_lab.id
    ]
    name = "ad.user"
    object_type = "User"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_role)

## Search Attributes

### Required

- `directory_services` - (Set of String) List of ID of directory services.
- `name` - (String) Name of the directory object.
- `object_type` - (String) Type of the directory object. Can be set to `User` or `Group`.

## Attributes Reference

- `id` - (String) The ID of this resource.
- `system_name` - (String) UPN of the directory object. This is directory user or group UPN.
- `display_name` - (String) Display name of the directory object.
- `distinguished_name` - (String) Distinguished name of the directory object.
- `forest` - (String) Forest/Domain name of the directory object.
