---
subcategory: "Policy Configuration"
---

# centrify_client attribute

**centrify_client** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        centrify_client {
            authentication_enabled = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            allow_no_mfa_mech = true
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_centrify_client.tf)

## Argument Reference

Optional:

- `authentication_enabled` - (Boolean) Enable authentication policy controls.
- `challenge_rule` (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched)
- `allow_no_mfa_mech` - (Boolean) Allow users without a valid authentication factor to log in.
