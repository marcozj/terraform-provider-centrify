---
page_title: "centrifyvault_passwordprofile Data Source - terraform-provider-centrifyvault"
description: |-
  This data source gets information of password profile.
---

# centrifyvault_passwordprofile (Data Source)

This data source gets information of password profile.

## Example Usage

```terraform
data "centrifyvault_passwordprofile" "win_profile" {
    name = "Windows Profile"
}
```

More examples can be found [here](../../examples/centrifyvault_passwordprofile/)

## Search Attributes

### Required

- `name` - (String) The name of password profile.

### Optional

- `profile_type` - (String) The type of password profile.

## Attributes Reference

- `id` - id of the password profile.
- `name` - name property.
- `profile_type` - profile_type property.
- `description` - (String) description property.
