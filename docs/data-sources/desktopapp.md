---
subcategory: "Resources"
---

# centrifyvault_desktopapp (Data Source)

This data source gets information of desktop app.

## Example Usage

```terraform
data "centrifyvault_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_desktopapp)

## Search Attributes

### Required

- `name` - (String) Name of the Desktop App.

## Attributes Reference

- `id` - (String) ID of the desktop app.
- `template_name` - (String) Template of the Desktop App.
- `application_alias` - (String) The alias name of the published RemoteApp program.
- `application_host_id` - (String) ID of application host. The Windows system that is hosting the desktop application.
- `login_credential_type` - (String) Host login credential type. The credentials used for the RDP connection to the application host.
- `description` - (String) Description of the Desktop App.
- `application_account_id` - (String) ID of the shared account.
- `command_line` - (String) Command line. The command line used to initiate launching the desktop application. Values can be entered directly or substituted at run-time using the command line arguments table.
- `command_parameter` - (Block Set) Command Line Arguments.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched).
- `policy_script` - (String) Use script to specify authentication rules (configured rules are ignored).
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
