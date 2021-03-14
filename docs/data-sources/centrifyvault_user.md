---
page_title: "centrifyvault_user Data Source - terraform-provider-centrifyvault"
description: |-
  This data source gets information of Centrify Directory User.
---

# centrifyvault_user (Data Source)

This data source gets information of Centrify Directory User.

## Example Usage

```terraform
data "centrifyvault_user" "admin" {
    username = "admin@example.com"
}
```

More examples can be found [here](../../examples/centrifyvault_user/)

## Search Attributes

### Required

- `username` - (String) The username in loginid@suffix format.

## Attributes Reference

- `id` - id of user.
- `username` - Username property.
- `email` - Email address property.
- `displayname` - Display name property.
