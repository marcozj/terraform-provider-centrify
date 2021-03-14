---
page_title: "oath_otp attribute in policy - terraform-provider-centrify"
description: |-
  oath_otp attribute in centrifyvault_policy Resource.
---

# oath_otp attribute

**oath_otp** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

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
        oath_otp {
            allow_otp = true
        }
    }
}
```

More examples can be found [here](../../examples/centrifyvault_policy/oath_otp.tf)

## Argument Reference

Optional:

- `allow_otp` - (Boolean) Allow OATH OTP integration.
