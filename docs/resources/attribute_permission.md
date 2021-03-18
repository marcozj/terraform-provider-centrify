---
subcategory: "Common Attribute"
---

# Permission and member_permission attribute

Permission & member_permission are a common attribute in various resources.

## Example Usage

```terraform
resource "centrifyvault_manualset" "my_accounts" {
    name = "My Accounts"
    type = "VaultAccount"
    description = "This Set contains my accounts."

    permission {
        principal_id = data.centrifyvault_user.admin.id
        principal_name = data.centrifyvault_user.admin.id
        principal_type = "User"
        rights = ["Grant","View","Edit","Delete"]
    }

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View"]
    }

    member_permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Checkout","Login","FileTransfer","Edit","Delete","UpdatePassword","WorkspaceLogin","RotatePassword"]
    }
}
```

## Argument Reference

### Required

- `principal_id` - (String) ID of the principal.
- `principal_name` - (String) User/role/group name.
- `principal_type` - (String) Principal type. Can be set to `User`, `Role` or `Group`.
- `rights` - (Set of String) Permissions: Grant,View,Edit,Delete.
