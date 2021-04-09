---
subcategory: "Resources"
---

# centrifyvault_cloudprovider (Data Source)

This data source gets information of authentication cloud provider.

## Example Usage

```terraform
data "centrifyvault_cloudprovider" "my_aws" {
    name = "My AWS"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_cloudprovider)

## Search Attributes

### Required

- `name` - (String) Name of the cloud provider.
- `cloud_account_id` - (String) Account ID of the cloud provider.

## Attributes Reference

- `id` - (String) The ID of this resource.
- `name` - (String) Name of the cloud provider.
- `cloud_account_id` - (String) Account ID of the cloud provider.
- `description` - (String) Description of the cloud provider.
- `challenge_rule` - (Block List) Authentication rules.
- `default_profile_id` - (String) Default Root Account Login Profile (used if no conditions matched).
- `enable_interactive_password_rotation` - (Boolean) Enable interactive password rotation. When enabled, allows on demand rotation of your root account password. Requires the Centrify Browser Extension.
- `prompt_change_root_password` - (Boolean) Prompt to change root password every login and password checkin. Displays a prompt with an option to rotate the root account password after every root account login attempt or password checkin.
- `enable_password_rotation_reminders` - (Boolean) Enable password rotation reminders. Displays a banner message with an option to rotate the root account password after the specified minimum number of days since last rotation has expired.
- `password_rotation_reminder_duration` - (Number) Minimum number of days since last rotation to trigger a reminder.
