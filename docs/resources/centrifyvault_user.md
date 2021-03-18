---
subcategory: "Access"
---

# centrifyvault_user (Resource)

This resource allows you to create/update/delete Centrify Directory User.

## Example Usage

```terraform
resource "centrifyvault_user" "testuser" {
    username = "testuser@example.com"
    email = "testuser@example.com"
    displayname = "Test User"
    description = "Test user"
    password = "xxxxxxxxxxxx"
    confirm_password = "xxxxxxxxxxxx"
    password_never_expire = true
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_user)

## Argument Reference

### Required

- `username` - (String) The username in loginid@suffix format.

### Optional

- `email` - (String) Email address.
- `displayname` - (String) Display name.
- `password` - (String, Sensitive) Password of the user.
- `confirm_password` - (String, Sensitive) Password of the user.
- `password_never_expire` - (Boolean) Password never expires. Default is `false`. When this is set to `true`, `force_password_change_next` should not be set to `true`.
- `force_password_change_next` - (Boolean) Require password change at next login. Default is `true`. When this is set to `true`, `password_never_expire` should not be set to `true`.
- `auth_client` - (Boolean) Is OAuth confidential client. Default is `false`.
- `send_email_invite` - (Boolean) Send email invite for user profile setup.
- `description` - (String) Description of the user.
- `office_number` - (String) Office phone number.
- `home_number` - (String) Home phone number.
- `mobile_number` - (String) Mobile phone number.
- `redirect_mfa_user_id` - (String) Redirect multi factor authentication to a different user account. This is the ID of another user.
- `manager_username` - (String) Username of the manager.
- `roles` -  (Set of String) Add to list of Roles.
