---
subcategory: "Policy Configuration"
---

# centrify_services attribute

**centrify_services** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        centrify_services {
            authentication_enabled = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            // Session Parameters
            session_lifespan = 23
            allow_session_persist = true
            default_session_persist = true
            persist_session_lifespan = 30
            // Other Settings
            allow_iwa = true
            iwa_set_cookie = true
            iwa_satisfies_all = true
            use_certauth = true
            certauth_skip_challenge = true
            certauth_set_cookie = true
            certauth_satisfies_all = true
            allow_no_mfa_mech = true
            auth_rule_federated = false
            federated_satisfies_all = true
            block_auth_from_same_device = false
            continue_failed_sessions = true
            stop_auth_on_prev_failed = true
            remember_last_factor = true
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_centrify_services.tf)

## Argument Reference

Optional:

- `authentication_enabled` - (Boolean) Enable authentication policy controls.
- `challenge_rule` (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched)
- `session_lifespan` (Number) Hours until session expires (default 12). Range between `1` to `9999`.
- `allow_session_persist` - (Boolean) Allow 'Keep me signed in' checkbox option at login (session spans browser sessions).
- `default_session_persist` - (Boolean) Default 'Keep me signed in' checkbox option to enabled.
- `persist_session_lifespan` - (Number) Hours until session expires when 'Keep me signed in' option enabled (default 2 weeks). Range between `1` to `9999`.
- `allow_iwa` - (Boolean) Allow IWA connections (bypasses authentication rules and default profile).
- `iwa_set_cookie` - (Boolean) Set identity cookie for IWA connections.
- `iwa_satisfies_all` - (Boolean) IWA connections satisfy all MFA mechanisms.
- `use_certauth` - (Boolean) Use certificates for authentication.
- `certauth_skip_challenge` - (Boolean) Certificate authentication bypasses authentication rules and default profile.
- `certauth_set_cookie` - (Boolean) Set identity cookie for connections using certificate authentication.
- `certauth_satisfies_all` - (Boolean) Connections using certificate authentication satisfy all MFA mechanisms.
- `allow_no_mfa_mech` - (Boolean) Allow users without a valid authentication factor to log in.
- `auth_rule_federated` - (Boolean) Apply additional authentication rules to federated users.
- `federated_satisfies_all` - (Boolean) Connections via Federation satisfy all MFA mechanisms.
- `block_auth_from_same_device` - (Boolean) Allow additional authentication from same device.
- `continue_failed_sessions` - (Boolean) Continue with additional challenges after failed challenge.
- `stop_auth_on_prev_failed` - (Boolean) Do not send challenge request when previous challenge response failed.
- `remember_last_factor` - (Boolean) Remember and suggest last used authentication factor.
