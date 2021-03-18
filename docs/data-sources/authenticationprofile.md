---
subcategory: "Settings"
---

# centrifyvault_authenticationprofile (Data Source)

This data source gets information of authentication profile.

## Example Usage

```terraform
data "centrifyvault_authenticationprofile" "new_device" {
    name = "Default New Device Login Profile"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_authenticationprofile)

## Search Attributes

### Required

- `name` (String) The name of the authentication profile.

## Attributes Reference

- `id` - (String) The ID of this resource.
- `pass_through_duration` - (Number) Pass through duration of the authentication profile.
- `uuid` - (String) UUID of the authentication profile.
