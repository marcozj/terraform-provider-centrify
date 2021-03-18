---
subcategory: "Policy Configuration"
---

# system_set attribute

**system_set** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("Server|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        system_set {
            // Account Policy
            checkout_lifetime = 60
            // System Policy
            allow_remote_access = true
            allow_rdp_clipboard = true
            local_account_automatic_maintenance = true
            local_account_manual_unlock = true
            // Security Settings
            remove_user_on_session_end = true
            allow_multiple_checkouts = true
            enable_password_rotation = true
            password_rotate_interval = 80
            enable_password_rotation_after_checkin = true
            minimum_password_age = 30
            minimum_sshkey_age = 30
            enable_sshkey_rotation = true
            sshkey_rotate_interval = 90
            sshkey_algorithm = "RSA_2048"
            // Maintenance Settings
            enable_password_history_cleanup = true
            password_historycleanup_duration = 120
            enable_sshkey_history_cleanup = true
            sshkey_historycleanup_duration = 120
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_system_set.tf)

## Argument Reference

Optional:

- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid. Enter the maximum number of minutes users are allowed to have a password checked out. After the number of minutes specified, the Centrify Privileged Access Service automatically checks the password back in. The minimum checkout lifetime is 15 minutes. If the policy is not defined, the default checkout lifetime is 60 minutes. You can set this policy globally or on an individual system. Policies defined globally apply to all systems except where you have explicitly defined a system-specific policy. Range between `15` to `2147483647`.
- `allow_remote_access` - (Boolean) Allow access from a public network (web client only). Enable it if you want to allow connections from outside of the firewall to access the selected system. When disabled, users will be denied access if they attempt to log on to the selected system from a connection outside of the firewall. You can set this policy globally or on an individual system. Policies defined globally apply to all systems except where you have explicitly defined a system-specific policy. **Note:** Access to systems from outside of the firewall is not supported by local clients.
- `allow_rdp_clipboard` - (Boolean) Allow RDP client to sync local clipboard with remote session. When enabled, allows users to copy texts or images from the local machine and paste them to the remote session, or vice versa. Applies to RDP native client and web client on supported browsers only.
- `local_account_automatic_maintenance` - (Boolean) Enable local account automatic maintenance.
- `local_account_manual_unlock` - (Boolean) Enable local account manual unlock.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `privilege_elevation_rule` - (Block List) Privilege Elevation Challenge Rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `privilege_elevation_default_profile_id` - (String) Default Privilege Elevation Profile (used if no conditions matched).
- `remove_user_on_session_end` - (Bollean) Remove local accounts upon session termination (Windows only). When enabled, the client removes local accounts created when a session is started and their local system profiles and settings after the session terminates. This setting affects Windows systems only.
- `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts for this system. Specifies whether multiple users can have the same account password checked out at the same time for a selected system. Enable it if only one user is allowed check out the password for a selected system at any given time. If disabled, the user must check the password in and have a new password generated before another user can access the system with the updated password. Enable it if you want to allow multiple users to have the account password checked out at the same time for a selected system. If you select Yes, multiple users can access the system without waiting for the password to be checked in.
- `enable_password_rotation` - (Boolean) Enable periodic password rotation
  - `password_rotate_interval` - (Number) Password rotation interval (days). Enter the maximum number of days to allow between automated password changes for managed accounts. You can set this policy to comply with your organization's password expiration policies. For example, your organization might require passwords to be changed every 90 days. You can use this policy to automatically update managed passwords at a maximum of every 90 days. If the policy is not defined, password are not rotated. Range between `1` to `2147483647`.
- `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin.
- `minimum_password_age` - (Number) Minimum Password Age (days). Minimum amount of days old a password must be before it is rotated. Range between `0` to `2147483647`.
- `minimum_sshkey_age` - (Number) Minimum SSH Key Age (days). Minimum amount of days old an SSH key must be before it is rotated. Range between `0` to `2147483647`.
- `enable_sshkey_rotation` - (Boolean) Enable periodic SSH key rotation.
  - `sshkey_rotate_interval` - (Number) SSH key rotation interval (days). Enter the maximum number of days to allow between automated SSH key changes for managed accounts. Range between `1` to `2147483647`.
- `sshkey_algorithm` - (String) SSH Key Generation Algorithm. Can be set to `RSA_1024`, `RSA_2048`, `ECDSA_P256`, `ECDSA_P384`, `ECDSA_P521`, `EdDSA_Ed448` or `EdDSA_Ed25519`.
- `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup.
  - `password_historycleanup_duration` - (Number) Password history cleanup (days). Deletes retired passwords automatically that were last modified either equal to or greater than the number of days specified here. Enter the number of days after which retired passwords matching the duration are deleted. The minimum value is 90 days. The default value is 365 days unless otherwise set in the Security settings under the Settings tab. Range between `90` to `2147483647`.
- `enable_sshkey_history_cleanup` - (Boolean) Enable periodic SSH key history cleanup.
  - `sshkey_historycleanup_duration` - (Number) SSH key history cleanup (days). Deletes retired SSH keys automatically that were last modified either equal to or greater than the number of days specified here. Enter the number of days after which retired SSH key matching the duration are deleted. The minimum value is 90 days. The default value is 365 days unless otherwise set in the Security settings under the Settings tab. Range between `1` to `2147483647`.
