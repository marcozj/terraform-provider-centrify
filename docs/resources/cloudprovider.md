---
subcategory: "Resources"
---

# centrifyvault_cloudprovider (Resource)

This resource allows you to create/update/delete cloud provider.

## Example Usage

```terraform
resource "centrifyvault_cloudprovider" "demo_aws" {
    type = "Aws"
    name = "Demo AWS"
    cloud_account_id = "xxxxxxxxxx"
    description = "Demo AWS"
    enable_interactive_password_rotation = true
    prompt_change_root_password = true
    enable_password_rotation_reminders = true
    password_rotation_reminder_duration = 20
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_cloudprovider)

## Argument Reference

### Required

- `cloud_account_id` - (String) Account ID of the cloud provider.
- `name` - (String) Name of the cloud provider.
- `type` - (String) Type of the cloud provider. Can be set to `Aws`.

### Optional

- `description` - (String) Description of the cloud provider.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Root Account Login Profile (used if no conditions matched).
- `enable_interactive_password_rotation` - (Boolean) Enable interactive password rotation. When enabled, allows on demand rotation of your root account password. Requires the Centrify Browser Extension.
- `prompt_change_root_password` - (Boolean) Prompt to change root password every login and password checkin. Displays a prompt with an option to rotate the root account password after every root account login attempt or password checkin.
- `enable_password_rotation_reminders` - (Boolean) Enable password rotation reminders. Displays a banner message with an option to rotate the root account password after the specified minimum number of days since last rotation has expired.
- `password_rotation_reminder_duration` - (Number) Minimum number of days since last rotation to trigger a reminder. Range between `1` to `2147483647`.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_attribute/sets.md) attribute for details.
