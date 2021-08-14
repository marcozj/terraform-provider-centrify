---
subcategory: "Access"
---

# centrify_role_membership (Resource)

This resource allows you to create/update/delete role membership for either existing or new role.

~> **WARNING:** `centrify_role_membership` will conflict with itself if used more than once with the same role.

~> **NOTE:** Do NOT use both `centrify_role` and `centrify_role_membership` to manage role membership for the same role.

## Example Usage

```terraform
// Existing federated (virtual) group
data "centrify_federatedgroup" "fedgroup1" {
  name = "Okta Infra Admins"
}

// Existing role whose membership to be managed
data "centrify_role" "testrole" {
    name = "Test Role"
}

// Role membership for exsting role
resource "centrify_role_membership" "testrolemembers" {
  role_id = data.centrify_role.testrole.id

  // Existing federated (virtual) group
  member {
    id = data.centrify_federatedgroup.fedgroup1.id
    type = "Group"
  }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_role_membership)

## Argument Reference

### Required

- `role_id` - (String) ID of the role.

### Optional

- `member` - (Block Set) (see [below reference for member](#reference-for-member))

## [Reference for `member`]

Required:

- `id` - (String) ID of the member.
- `type` - (String) Type of the member. Can be set to `User`, `Group` or `Role`.

## Import

Role membership can be imported using the resource `id`, e.g.

```shell
terraform import centrify_role_membership.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```
