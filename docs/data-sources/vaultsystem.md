---
subcategory: "Resources"
---

# centrifyvault_vaultsystem (Data Source)

This data source gets information of system.

## Example Usage

```terraform
data "centrifyvault_vaultsystem" "demo_system" {
    name = "demosystem"
    fqdn = "demosystem.example.com"
    computer_class = "Unix"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsystem)

## Search Attributes

### Required

- `name` - (String) The name of the system.
- `fqdn` - (String) Hostname or IP address of the system.

### Optional

- `computer_class` - (String) Type of the system. Can be set to `Windows`, `Unix`, `CiscoIOS`, `CiscoNXOS`, `JuniperJunos`, `HpNonStopOS`, `IBMi`, `CheckPointGaia`, `PaloAltoNetworksPANOS`, `F5NetworksBIGIP`, `CiscoAsyncOS`, `VMwareVMkernel`, `GenericSsh` or `CustomSsh`.

## Attributes Reference

- `id` - (String) ID of the system.
- `name` - (String) The name of the system.
- `fqdn` - (String) Hostname or IP address of the system.
- `computer_class` - (String) Type of the system.
- `description` - (String) Description of the system.
- `port` - (Number) Port that used to connect to the system.
- `system_timezone` - (String) System time zone.
- `use_my_account` (Boolean) Enable Use My Account - Unix/Linux only. Check this box once you have made the required changes to OpenSSH on this system.
- `proxyuser` - (String) - Proxy user name.
- `proxyuser_password` - (String, Sensitive) Proxy user password.
- `proxyuser_managed` - (Boolean) Manage proxy user credential. By selecting this option the credential will be automatically changed and become unknown to other applications or users.
- `management_mode` - (String) Management mode of the system. For Windows only.
- `management_port` - (Number) Management port for account management. For Windows, F5, PAN-OS and VMKernel only. For Windows, it is used when `management_mode` is set to either `WinRMOverHttp` or `WinRMOverHttps`.
- `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid.
- `allow_remote_access` - (Boolean) Allow access from a public network (web client only). Specifies whether remote connections are allowed from a public network for a selected system.
- `allow_rdp_clipboard` - (Boolean) Allow RDP client to sync local clipboard with remote session. When enabled, allows users to copy texts or images from the local machine and paste them to the remote session, or vice versa. Applies to RDP native client and web client on supported browsers only.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default System Login Profile (used if no conditions matched).
- `privilege_elevation_default_profile_id` - (String) Default Privilege Elevation Profile (used if no conditions matched).
- `privilege_elevation_rule` - (Block List) Privilege Elevation Challenge Rules. Refer to [privilege_elevation_rule](./attribute_challengerule.md) attribute for details.
- `local_account_automatic_maintenance` - (Boolean) Local Account Automatic Maintenance. Specifies whether local account automatic maintenance is enabled.
- `local_account_manual_unlock` - (Boolean) Local Account Manual Unlock - Windows only. Specifies whether local account manual unlock is enabled.
- `administrative_account_id` - (String) Local administrative account id. Set a local administrative account or join a domain with a configured provisioning administrative account for account reconciliation operations.
- `domain_id` - (String) AD domain that this system belongs to. When this is set, "Domain Administrative Account" is automatically populated.
- `remove_user_on_session_end` - (Boolean) Remove local accounts upon session termination - Windows only. When enabled, the client removes local accounts created when a session is started and their local system profiles and settings after the session terminates. This setting affects Windows systems only.
- `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts for this system. Specifies whether multiple users can have the same account password checked out at the same time for a selected system.
- `enable_password_rotation` - (Boolean) Enable periodic password rotation. Specifies whether managed password should be rotated periodically.
- `password_rotate_interval` - (Number) Password rotation interval (days). Rotates managed passwords automatically at the interval you specify.
- `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin. Specifies whether managed password should be rotated after it's checked in.
- `minimum_password_age` - (Number) Minimum Password Age (days). Minimum amount of days old a password must be before it is rotated.
- `password_profile_id` - (String) Password complexity profile id.
- `enable_sshkey_rotation` - (Boolean) Enable periodic SSH key rotation - SSH system only. Specifies whether managed SSH key should be rotated periodically.
- `sshkey_rotate_interval` - (Number) SSH key rotation interval (days) - SSH system only. Rotates managed SSH key automatically at the interval you specify.
- `minimum_sshkey_age` - (Number) Minimum SSH Key Age (days). Minimum amount of days old an SSH key must be before it is rotated.
- `sshkey_algorithm` - (String) SSH Key Generation Algorithm - SSH system only. Specifies the algorithm to use when generating SSH keys during manual or automatic SSH key rotation.
- `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup. Specifies whether retired passwords should be deleted periodically.
- `password_historycleanup_duration` - (Number) Password history cleanup (days). Deletes retired passwords automatically that were last modified either equal to or greater than the number of days specified here.
- `enable_sshkey_history_cleanup` - (Boolean) Enable periodic SSH key history cleanup - SSH system only. Specifies whether retired passwords should be deleted periodically.
- `sshkey_historycleanup_duration` - (Number) SSH key history cleanup (days) - SSH system only. Deletes retired SSH keys automatically that were last modified either equal to or greater than the number of days specified here.
- `agent_auth_workflow_enabled` - (Boolean) Enable Agent Auth Workflow. Windows and Unix/Linux only.
- `agent_auth_workflow_approver` - (Block List) List of approvers. Windows and Unix/Linux only. Refer to [agent_auth_workflow_approver](./attribute_workflow_approver.md) attribute for details.
- `privilege_elevation_workflow_enabled` - (Boolean) Enable Privilege Elevation Request Workflow. Windows and Unix/Linux only.
- `privilege_elevation_workflow_approver` - (Block List) List of approvers. Windows and Unix/Linux only. Refer to [privilege_elevation_workflow_approver](./attribute_workflow_approver.md) attribute for details.
- `use_domainadmin_for_zonerole_workflow` - (Boolean) Use Domain Administrator Account for Zone Role Workflow operations - Windows and Unix/Linux only.
- `enable_zonerole_workflow` - (Boolean) Enable zone role requests for this system - Windows and Unix/Linux only.
- `use_domain_assignment_for_zoneroles` - (Boolean) Assignable Zone Roles - Use domain assignments.
- `assigned_zonerole` - (Block Set) List of zone Role. Windows and Unix/Linux only. Refer to [assigned_zonerole](./attribute_assigned_zonerole.md) attribute for details.
- `use_domain_assignment_for_zonerole_approvers` - (Boolean) Approver list - Use domain assignments
- `assigned_zonerole_approver` - (Block List) List of approvers for Zone Role request. Windows and Unix/Linux only. Refer to [workflow_approver](./attribute_workflow_approver.md) and [assigned_zonerole_approver](./attribute_assigned_zonerole.md) for details.
- `connector_list` (Set of String) List of Connector IDs. Refer to [connector_list](./attribute_connector_list.md) attribute for details.
