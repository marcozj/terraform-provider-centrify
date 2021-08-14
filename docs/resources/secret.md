---
subcategory: "Resources"
---

# centrify_secret (Resource)

This resource allows you to create/update/delete secret.

## Example Usage

```terraform
resource "centrify_secret" "test_secret" {
    secret_name = "Test Secret"
    description = "Test Secret"
    secret_text = "xxxxxxxxxxxxx"
    type = "Text"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_secret)

## Argument Reference

### Required

- `secret_name` - (String) Name of the secret.
- `type` - (String) Type of the secret. Can be set to `Text` only at the moment. (TODO: `File`)

### Optional

- `description` - (String) Description of the secret.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `folder_id` - (String) ID of the folder where the secret is located.
- `parent_path` - (String) Path of parent folder.
- `secret_text` - (String, Sensitive) Content of the secret.
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Import

Secret can be imported using the resource `id`, e.g.

```shell
terraform import centrify_secret.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` and `set` aren't supported in import process.
