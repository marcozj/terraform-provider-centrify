---
subcategory: "Policy Configuration"
---

# radius attribute

**radius** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        radius {
            allow_radius = true
            require_challenges = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            send_vendor_attributes = true
            allow_external_radius = true
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_radius.tf)

## Argument Reference

Optional:

- `allow_radius` - (Boolean) Allow RADIUS client connections. This enables you to extend MFA from Centrify Privileged Access Service to thick clients (e.g. VPNs) that support RADIUS.
- `require_challenges` - (Boolean) Require authentication challenge. Only applicable if `allow_radius` is `true`.
- `default_profile_id` - (String) Default authentication profile. Only applicable if `require_challenges` is `true`.
- `send_vendor_attributes` - (Boolean) Send vendor specific attributes. Only applicable if `allow_radius` is `true`.
- `allow_external_radius` - (Boolean) Allow 3rd Party RADIUS Authentication. This enables you to add 3rd party authentication solutions (e.g. RSA SecurID) to the list of supported authentication mechanisms for your users.
