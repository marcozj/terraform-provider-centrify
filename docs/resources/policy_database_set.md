---
subcategory: "Policy Configuration"
---

# database_set attribute

**database_set** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("VaultDatabase|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        database_set {
            // Account Security
            checkout_lifetime = 60
            // Security Settings
            allow_multiple_checkouts = true
            enable_password_rotation = true
            password_rotate_interval = 90
            enable_password_rotation_after_checkin = true
            minimum_password_age = 70
            // Maintenance Settings
            enable_password_history_cleanup = true
            password_historycleanup_duration = 120
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_database_set.tf)

## Argument Reference

Optional:

- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid. Enter the maximum number of minutes users are allowed to have a password checked out. After the number of minutes specified, the Centrify Privileged Access Service automatically checks the password back in. The minimum checkout lifetime is 15 minutes. If the policy is not defined, the default checkout lifetime is 60 minutes. You can set this policy globally or on an individual database. Policies defined globally apply to all database except where you have explicitly defined a database-specific policy. Range between `15` to `2147483647`.
- `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts for related accounts. Specifies whether multiple users can have the same database account password checked out at the same time. Enable it if only one user is allowed to check out the password at any given time. If disabled, the user must check the password in and have a new password generated before another user can check out the updated password. Enable it if you want to allow multiple users to have the account password checked out at the same time without waiting for the password to be checked in.
- `enable_password_rotation` - (Boolean) Enable periodic password rotation.
  - `password_rotate_interval` - (Number) Password rotation interval (days). Rotates managed passwords automatically at the interval you specify. Enter the maximum number of days to allow between automated password changes for managed accounts. You can set this policy to comply with your organization's password expiration policies. For example, your organization might require passwords to be changed every 90 days. You can use this policy to automatically update managed passwords at a maximum of every 90 days. If the policy is not defined, password are not rotated. Range between `1` to `2147483647`.
- `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin. Specifies whether managed password should be rotated after it's checked in.
- `minimum_password_age` - (Number) Minimum Password Age (days). Range between `0` to `2147483647`.
- `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup. Specifies whether retired passwords should be deleted periodically.
  - `password_historycleanup_duration` - (Number) Password history cleanup (days). Deletes retired passwords automatically that were last modified either equal to or greater than the number of days specified here. Enter the number of days after which retired passwords matching the duration are deleted. The minimum value is 90 days. The default value is 365 days unless otherwise set in the Security settings under the Settings tab. Range between `90` to `2147483647`.
