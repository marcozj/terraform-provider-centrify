---
subcategory: "Policy Configuration"
---

# cloudproviders_set attribute

**cloudproviders_set** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("CloudProviders|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        cloudproviders_set {
            enable_interactive_password_rotation = true
            prompt_change_root_password = true
            enable_password_rotation_reminders = true
            password_rotation_reminder_duration = 20
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_cloudproviders_set.tf)

## Argument Reference

Optional:

- `challenge_rule` - (Block List) Root Account Login Challenge Rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Root Account Login Profile (used if no conditions matched).
- `enable_interactive_password_rotation` - (Boolean) Enable interactive password rotation. When enabled, allows on demand rotation of your root account password. Requires the Centrify Browser Extension.
- `prompt_change_root_password` - (Boolean) Prompt to change root password every login and password checkin. Displays a prompt with an option to rotate the root account password after every root account login attempt or password checkin.
- `enable_password_rotation_reminders` - (Boolean) Enable password rotation reminders. Displays a banner message with an option to rotate the root account password after the specified minimum number of days since last rotation has expired.
- `password_rotation_reminder_duration` - (Number) Minimum number of days since last rotation to trigger a reminder. Range between `1` to `2147483647`.
