---
subcategory: "Resources"
---

# centrifyvault_vaultdomain (Data Source)

This data source gets information of domain.

## Example Usage

```terraform
data "centrifyvault_vaultdomain" "example.com" {
    name = "example.com"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomain)

## Search Attributes

### Required

- `name` - (String) Name of the domain.

## Attributes Reference

- `id` - (String) ID of the domain.
- `name` - (String) Name of the domain.
- `description` - (String) Description of the domain.
- `verify_domain` - (Boolean) Whether to verify the Domain upon creation.
- `forest_id` - (String) Object ID of the forest.
- `parent_id` - (String) Object ID of parent domain.
- Policy
  - `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid.
- Advanced -> Account Reconciliation and Zone Role workflow settings
  - `administrative_account_id` - (String) ID of the domain administrative account. This is either a vaulted domain account id or directory object id.
  - `administrative_account_name` - (String) Name of administrative account.
  - `administrative_account_password` - (String, Sensitive) Password of administrative account. This is required if `administrative_account_id` is a directory object id. To use non-vaulted account, password must be provided.
  - `auto_domain_account_maintenance` - (Boolean) Enable Automatic Domain Account Maintenance.
  - `manual_domain_account_unlock` - (Boolean) Enable Manual Domain Account Unlock.
  - `auto_local_account_maintenance` - (Boolean) Enable Automatic Local Account Maintenance.
  - `manual_local_account_unlock` - (Boolean) Enable Manual Local Account Unlock.
  - `provisioning_admin_id` - (String) Provisioning Administrative Account (must be managed). This is a vaulted domain account id.
  - `reconciliation_account_name` - (Sring) Reconciliation account name.
  - `enable_zonerole_workflow` - (Boolean) Enable zone role requests for systems in the domain.
  - `assigned_zonerole` - (Block Set) List of assignable Zone Roles. Refer to [assigned_zonerole](./attribute_assigned_zonerole.md) attribute for details.
  - `assigned_zonerole_approver` - (Block List) List of approvers for Zone Role request. Refer to [workflow_approver](./attribute_workflow_approver.md) and [assigned_zonerole_approver](./attribute_assigned_zonerole.md) for details.
- Advanced -> Security Settings
  - `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts per AD account added for this domain.
  - `enable_password_rotation` - (Boolean) Enable periodic password rotation.
  - `password_rotate_interval` - (Number) Password rotation interval (days).
  - `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin.
  - `minimum_password_age` - (Number) Minimum Password Age (days).
  - `password_profile_id` - (String) Password complexity profile id.
- Advanced -> Maintenance Settings
  - `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup.
  - `password_historycleanup_duration` - (Number) Password history cleanup (days).
- Advanced -> Domain/Zone Tasks
  - `enable_zone_joined_check` - (Boolean) Enable periodic domain/zone joined check
  - `zone_joined_check_interval` - (Number) Domain/zone joined check interval (minutes).
  - `enable_zonerole_cleanup` - (Boolean) Enable periodic removal of expired zone role assignments.
  - `zonerole_cleanup_interval` - (Number) Expired zone role assignment removal interval (hours).
- Zone Role Workflow
  - `enable_zonerole_workflow` - (Boolean) Enable zone role requests for systems in this domain.
  - `assigned_zonerole` - (Block Set) List of zone Role. Windows and Unix/Linux only.
  - `assigned_zonerole_approver` - (Block List) List of approvers for Zone Role request.
- `connector_list` (Set of String) List of Connector IDs.
