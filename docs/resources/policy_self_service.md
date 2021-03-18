---
subcategory: "Policy Configuration"
---

# self_service attribute

**self_service** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        self_service {
            // Password Reset
            account_selfservice_enabled = true
            password_reset_enabled = true
            pwreset_allow_for_aduser = true
            pwreset_with_cookie_only = true
            login_after_reset = true
            pwreset_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            max_reset_attempts = 5
            // Account Unlock
            account_unlock_enabled = true
            unlock_allow_for_aduser = true
            unlock_with_cookie_only = true
            show_locked_message = true
            unlock_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            // Active Directory Self Service Settings
            use_ad_admin = false
            // Additional Policy Parameters
            max_reset_allowed = 6
            max_time_allowed = 50
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_self_service.tf)

## Argument Reference

Optional:

- `account_selfservice_enabled` - (Boolean) Enable account self service controls.
- `password_reset_enabled` - (Boolean) Enable password reset.
- `pwreset_allow_for_aduser` - (Boolean) Allow for Active Directory users.
- `pwreset_with_cookie_only` - (Boolean) Only allow from browsers with identity cookie.
- `login_after_reset` - (Boolean) User must log in after successful password reset.
- `pwreset_auth_profile_id` - (String) Password reset authentication profile.
- `max_reset_attempts` - (Number) Maximum consecutive password reset attempts per session. Range between `0` to `10`.
- `account_unlock_enabled` - (Boolean) Enable account unlock.
- `unlock_allow_for_aduser` - (Boolean) Allow for Active Directory users.
- `unlock_with_cookie_only` - (Boolean) Only allow from browsers with identity cookie.
- `show_locked_message` - (Boolean) Show a message to end users in desktop login that account is locked (default no)
- `unlock_auth_profile_id` - (String) Account unlock authentication profile.
- `use_ad_admin` - (Boolean) Use AD admin for AD self-service. when it is `false`, it will use connector running on privileged account.
- `ad_admin_user` - (String) Admin user name. Applicable if `use_ad_admin` is `true`.
- `admin_user_password` - (Block List, Max: 1) Admin user password attributes.
  - `type` - (String) Password type. Must be `SafeString`.
  - `value` - (String, Sensitive) Actual password.
- `max_reset_allowed` - (Number) Maximum forgotten password resets allowed within window (default 10). Range between `0` to `10`.
- `max_time_allowed` - (Number) Capture window for forgotten password resets (default 60 minutes). Valid values are `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90` or `100`.
