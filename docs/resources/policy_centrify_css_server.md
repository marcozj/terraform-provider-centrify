---
subcategory: "Policy Configuration"
---

# centrify_css_server attribute

**centrify_css_server** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        centrify_css_server {
            authentication_enabled = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            pass_through_mode = 2
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_centrify_css_server.tf)

## Argument Reference

Optional:

- `authentication_enabled` - (Boolean) Enable authentication policy controls.
- `challenge_rule` (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched)
- `pass_through_mode` - (Number) Apply pass-through duration: Never (Default) `0`, If Same Source and Target `1`, If Same Source `2`, If Same Target `3`.
