---
subcategory: "Policy Configuration"
---

# secret_set attribute

**secret_set** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("DataVault|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        secret_set {
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_secret_set.tf)

## Argument Reference

Optional:

- `challenge_rule` (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched)
