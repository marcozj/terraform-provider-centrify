---
subcategory: "Access"
---

# centrifyvault_user (Data Source)

This data source gets information of Centrify Directory User.

## Example Usage

```terraform
data "centrifyvault_user" "admin" {
    username = "admin@example.com"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_user)

## Search Attributes

### Required

- `username` - (String) The username in loginid@suffix format.

## Attributes Reference

- `id` - id of user.
- `username` - (String) The username in loginid@suffix format.
- `email` - (String) Email address.
- `displayname` - (String) Display name.
- `password` - (String, Sensitive) Password of the user.
- `confirm_password` - (String, Sensitive) Password of the user.
- `password_never_expire` - (Boolean) Password never expires.
- `force_password_change_next` - (Boolean) Require password change at next login.
- `auth_client` - (Boolean) Is OAuth confidential client.
- `send_email_invite` - (Boolean) Send email invite for user profile setup.
- `description` - (String) Description of the user.
- `office_number` - (String) Office phone number.
- `home_number` - (String) Home phone number.
- `mobile_number` - (String) Mobile phone number.
- `redirect_mfa_user_id` - (String) Redirect multi factor authentication to a different user account. This is the ID of another user.
- `manager_username` - (String) Username of the manager.
