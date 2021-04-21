---
subcategory: "Resources"
---

# centrifyvault_vaultdomain (Resource)

This resource allows you to create/update/delete domain.

## Example Usage

```terraform
resource "centrifyvault_vaultdomain" "example_lab" {
    name = "example.lab"
    description = "example.lab domain"
    // Policy menu
    checkout_lifetime = 90

    // Advanced -> Security Settings
    allow_multiple_checkouts = true
    enable_password_rotation = true
    password_rotate_interval = 90
    enable_password_rotation_after_checkin = true
    minimum_password_age = 120
    // Advanced -> Maintenance Settings
    enable_password_history_cleanup = true
    password_historycleanup_duration = 100
    // Advanced -> Domain/Zone Tasks
    enable_zone_joined_check = true
    zone_joined_check_interval = 90
    enable_zonerole_cleanup = true
    zonerole_cleanup_interval = 6
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomain)

## Argument Reference

### Required

- `name` - (String) The domain name.

### Optional

- Settings
  - `description` - (String) Description of the domain.
  - `verify_domain` - (Boolean) Whether to verify the Domain upon creation. Default is `true`.
  - `forest_id` - (String) Object ID of the forest. (experiment)
  - `parent_id` - (String) Object ID of parent domain. (experiment)
- Policy
  - `checkout_lifetime` - (Number) Checkout lifetime (minutes). Specifies the number of minutes that a checked out password is valid. Range between `15` to `2147483647`.
- Advanced -> Account Reconciliation and Zone Role workflow settings
  Domain reconciliation and Zone Role workflow configurations require vaulted domain account or domain account directly from AD. Vaulted domain account can only be added after the domain creation therefore reconciliation configurations can't be set during domain creation. It is handled separately by [`centrifyvault_vaultdomainconfiguration`](./centrifyvault_vaultdomainconfiguration.md) resource.
- Advanced -> Security Settings
  - `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts per AD account added for this domain.
  - `enable_password_rotation` - (Boolean) Enable periodic password rotation.
  - `password_rotate_interval` - (Number) Password rotation interval (days).
  - `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin.
  - `minimum_password_age` - (Number) Minimum Password Age (days). Range between `0` to `2147483647`.
  - `password_profile_id` - (String) Password complexity profile id.
- Advanced -> Maintenance Settings
  - `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup.
  - `password_historycleanup_duration` - (Number) Password history cleanup (days). Range between `90` to `2147483647`.
- Advanced -> Domain/Zone Tasks
  - `enable_zone_joined_check` - (Boolean) Enable periodic domain/zone joined check
  - `zone_joined_check_interval` - (Number) Domain/zone joined check interval (minutes). Range between `1` to `2147483647`. Default is `1440`.
  - `enable_zonerole_cleanup` - (Boolean) Enable periodic removal of expired zone role assignments.
  - `zonerole_cleanup_interval` - (Number) Expired zone role assignment removal interval (hours). Range between `1` to `2147483647`. Default is `6`.
- Zone Role Workflow
  - `enable_zonerole_workflow` - (Boolean) Enable zone role requests for systems in this domain.
  - `assigned_zonerole` - (Block Set) List of zone Role. Windows and Unix/Linux only. Applicable only if `use_domainadmin_for_zonerole_workflow` is `true` and`enable_zonerole_workflow` is `true` and `use_domain_assignment_for_zoneroles` is `false`. Refer to [assigned_zonerole](./attribute_assigned_zonerole.md) attribute for details.
  - `assigned_zonerole_approver` - (Block List) List of approvers for Zone Role request. Windows and Unix/Linux only. Applicable only if `use_domainadmin_for_zonerole_workflow` is `true` and`enable_zonerole_workflow` is `true` and `use_domain_assignment_for_zonerole_approvers` is `false`. Refer to [workflow_approver](./attribute_workflow_approver.md) and [assigned_zonerole_approver](./attribute_assigned_zonerole.md) for details.
- `connector_list` (Set of String) List of Connector IDs. Refer to [connector_list](./attribute_connector_list.md) attribute for details.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Import

Domain can be imported using the resource `id`, e.g.

```shell
terraform import centrifyvault_vaultdomain.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` and `set` aren't support in import process.
