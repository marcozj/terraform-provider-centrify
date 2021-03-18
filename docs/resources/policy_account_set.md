---
subcategory: "Policy Configuration"
---

# account_set attribute

**account_set** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("VaultAccount|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        account_set {
            checkout_lifetime = 60
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            access_secret_checkout_dfault_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_account_set.tf)

## Argument Reference

Optional:

- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid. Enter the maximum number of minutes users are allowed to have a password checked out. After the number of minutes specified, the Centrify Privileged Access Service automatically checks the password back in. The minimum checkout lifetime is 15 minutes. If the policy is not defined, the default checkout lifetime is 60 minutes. You can set this policy globally or on an individual account. Policies defined globally apply to all accounts except where you have explicitly defined a account-specific policy. Range between `15` to `2147483647`.
- `challenge_rule` - (Block List) Password Checkout Challenge Rule. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Password Checkout Profile (used if no conditions matched).
- `access_secret_checkout_rule` - (Block List) Secret Access Key Checkout Challenge Rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `access_secret_checkout_dfault_profile_id` - (String) Default Secret Access Key Checkout Profile (used if no conditions matched).
