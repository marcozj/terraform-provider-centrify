---
subcategory: "Policy Configuration"
---

# password_settings attribute

**password_settings** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        password_settings {
            // Password Requirements
            min_length = 12
            max_length = 24
            require_digit = true
            require_mix_case = true
            require_symbol = true
            // Display Requirements
            show_password_complexity = true
            complexity_hint = "Whatever ......."
            // Additional Requirements
            no_of_repeated_char_allowed = 2
            check_weak_password = true
            allow_include_username = true
            allow_include_displayname = true
            require_unicode = true
            // Password Age
            min_age_in_days = 10
            max_age_in_days = 90
            password_history = 10
            expire_soft_notification = 35
            expire_hard_notification = 72
            expire_notification_mobile = true
            // Capture Settings
            bad_attempt_threshold = 5
            capture_window = 20
            lockout_duration = 30
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_password_settings.tf)

## Argument Reference

Optional:

- `min_length` - (Number) Minimum password length (default 8). Range between `4` to `16`.
- `max_length` - (Number) Maximum password length (default 64). Range between `4` to `64`.
- `require_digit` - (Boolean) Require at least one digit (default yes).
- `require_mix_case` - (Boolean) Require at least one upper case and one lower case letter (default yes).
- `require_symbol` - (Boolean) Require at least one symbol (default no).
- `show_password_complexity` - (Boolean) Show password complexity requirements when entering a new password (default no). Enabling this policy displays the password complexity requirements when updating a user account password.
- `complexity_hint` - (String) Password complexity requirements for directory services other than Centrify Directory. Password complexity requirements for Centrify Directory users are automatically discovered but all other directory services require manually entering a complexity requirement string. Only applicable when `show_password_complexity` is `true`.
- `no_of_repeated_char_allowed` - (Number) Limit the number of consecutive repeated characters. Password cannot contain consecutive repeated characters equal to or more than the set value. The default is to allow consecutive repeated characters. Range between `2` to `64`.
- `check_weak_password` - (Boolean) Check against weak password.
- `allow_include_username` - (Boolean) Allow username as part of password.
- `allow_include_displayname` - (Boolean) Allow display name as part of password.
- `require_unicode` - (Boolean) Require at least one Unicode characters. Require at least one unicode character which categorized as an alphabetic character but is not uppercase or lowercase.
- `min_age_in_days` - (Number) Minimum password age before change is allowed. The default is 0 days. Users will not be allowed to change or reset their password until the current password is at least this old. Range between `0` to `998`.
- `max_age_in_days` - (Number) Maximum password age. The default is 365 days. After the password expires, users are prompted to enter their current password and then enter a new one. Enter 0 (zero) if you do not want to specify a password expiration period. Range between `0` to `3650`.
- `password_history` - (Number) Password history (default 3). Range between `0` to `25`.
- `expire_soft_notification` - (Number) Password Expiration Notification (default 14 days). Select the number of days before a user's password expires to begin posting a notification of expiration through a portal banner and daily emails. This policy applies to Centrify Directory users and Active Directory accounts. Valid values are `7`, `14`, `21`, `28`, `35`, `42` or `49`.
- `expire_hard_notification` - (Number) Escalated Password Expiration Notification (default 48 hours). Select the number of hours before a user's password expires to present a change password dialog. The dialog is automatically displayed when the user logs in. This policy applies to Centrify Directory users and Active Directory accounts. **NOTE:** This policy is not supported on mobile clients. Valid values are `24`, `48`, `72`, `96` or `120`.
- `expire_notification_mobile` - (Boolean) Enable password expiration notifications on enrolled mobile devices. When enabled, password expiration notifications are sent to registered mobile devices
- `bad_attempt_threshold` - (Number) Maximum consecutive bad password attempts allowed within window (default Off). Enter `Off` to allow the user an unlimited number of failed attempts. Users are locked out for the time period you specify in the "Lockout duration before password re-attempt allowed" policy when they fail in the attempt after the number you select.
- `capture_window` - (Number) Capture window for consecutive bad password attempts (default 30 minutes). Enter the number of minutes to define the time period before the number of failed password attempts is reset. This time period is only applicable when the "Maximum consecutive bad password attempts allowed within window" policy defines the number of failed attempts allowed and is not set to Off. The user is locked out for the time period you set in the "Lockout duration before password re-attempt allowed" policy. After that, the user can attempt to log in again. Range between `1` to `2147483647`.
- `lockout_duration` - (Number) Lockout duration before password re-attempt allowed (default 30 minutes). Enter the number of minutes users must wait before they can attempt to log in again after lockout. Range between `1` to `2147483647`.
