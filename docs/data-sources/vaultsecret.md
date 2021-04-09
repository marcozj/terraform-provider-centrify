---
subcategory: "Resources"
---

# centrifyvault_vaultsecret (Data Source)

This data source gets information of secret.

## Example Usage

```terraform
data "centrifyvault_vaultsecret" "test_secret" {
    secret_name = "testsecret"
    checkout = true
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsecret)

## Search Attributes

### Required

- `secret_name` - (String) Name of the secret.

### Optional

- `parent_path` - (String) Path of parent folder.
- `checkout` - (Boolean) Whether to retrieve secret content. Default is `false`. If `true`, `secret_text` will be populated.

## Attributes Reference

- `id` - (String) ID of the secret.
- `secret_name` - (String) Name of the secret.
- `type` - (String) Type of the secret.
- `description` - (String) Description of the secret.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `folder_id` - (String) ID of the folder where the secret is located.
- `parent_path` - (String) Path of parent folder.
- `secret_text` - (String, Sensitive) Content of the secret.
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
