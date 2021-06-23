---
subcategory: "Policy Configuration"
---

# oath_otp attribute

**oath_otp** is a sub attribute in settings attribute within **centrify_policy** Resource.

## Example Usage

```terraform
resource "centrify_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrify_role.system_admin.id,
    ]
    
    settings {
        oath_otp {
            allow_otp = true
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/blob/main/examples/centrify_policy/policy_oath_otp.tf)

## Argument Reference

Optional:

- `allow_otp` - (Boolean) Allow OATH OTP integration.
