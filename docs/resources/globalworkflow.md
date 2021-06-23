---
subcategory: "Settings"
---

# centrify_globalworkflow (Resource)

This resource allows you to configure global workflows including account, agent auth, secrets and privilege elevation workflows.

## Example Usage

```terraform
resource "centrify_globalworkflow" "account_wf" {
    type = "wf"
    settings {
        enabled = true
        approver {
            guid = data.centrify_role.system_admin.id
            name = data.centrify_role.system_admin.name
            type = "Role"
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrify/tree/main/examples/centrify_globalworkflow)

## Argument Reference

### Required

- `type` - (String) Type of workflow. Can be set to `wf`, `agentAuthWorkflow`, `secretsWorkflow`, or `privilegeElevationWorkflow`.

### Optional

- `settings` - (Block List, Max: 1) Workflow approver settings. (see below [reference for settings](#reference-for-settings))

## [Reference for `settings`]

Optional:

- `enabled` - (Boolean) Enable workflow for all accounts/systems/secrets.
- `approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
