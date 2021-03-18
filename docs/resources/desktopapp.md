---
subcategory: "Applications"
---

# centrifyvault_desktopapp (Resource)

This resource allows you to create/update/delete DesktopApp.

## Example Usage

```terraform
resource "centrifyvault_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
    template_name = "GenericDesktopApplication"
    description = "Test Desktop Application"
    application_host_id = data.centrifyvault_vaultsystem.apphost.id
    login_credential_type = "SharedAccount"
    application_account_id = data.centrifyvault_vaultaccount.shared_account.id
    application_alias = "pas_desktopapp"
    
    command_line = "--ini=ini\\web_myapp.ini --username={user.User} --password={user.Password}"
    command_parameter {
        name = "system"
        type = "Server"
        target_object_id = data.centrifyvault_vaultsystem.my_app.id
    }
    command_parameter {
        name = "user"
        type = "VaultAccount"
        target_object_id = data.centrifyvault_vaultaccount.admin.id
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_desktopapp)

## Argument Reference

### Required

- `name` - (String) Name of the Desktop App.
- `template_name` - (String) Template of the Desktop App. Can be set to `GenericDesktopApplication`, `Ssms`, `Toad` or `VpxClient`.
- `application_alias` - (String) The alias name of the published RemoteApp program.
- `application_host_id` - (String) ID of application host. The Windows system that is hosting the desktop application.
- `login_credential_type` - (String) Host login credential type. The credentials used for the RDP connection to the application host. Can be set to `ADCredential`, `SetByUser`, `AlternativeAccount`, or `SharedAccount`.

### Optional

- `description` - (String) Description of the Desktop App.
- `application_account_id` - (String) ID of the shared account. Required if `login_credential_type` is `SharedAccount`.
- `command_line` - (String) Command line. The command line used to initiate launching the desktop application. Values can be entered directly or substituted at run-time using the command line arguments table.
- `command_parameter` - (Block Set) Command Line Arguments. (see [reference for `command_parameter`](#reference-for-command_parameter)) Run-time argument substitutions for the command line.

- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched).
- `policy_script` - (String) Use script to specify authentication rules (configured rules are ignored)
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Reference for `command_parameter`

Required:

- `name` - (String) Name of the parameter.
- `type` - (String) Type of the parameter. Can be set to `int`, `date`, `string`, `VaultAccount`, `CloudProviders`, `VaultDatabase`, `Device`, `VaultDomain`, `ResourceProfile`, `Role`, `DataVault`, `Subscriptions`, `SshKeys`, `Server`, `User`.

Optional:

- `target_object_id` - (String) ID of selected parameter value
