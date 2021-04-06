---
subcategory: "Resources"
---

# centrifyvault_vaultdomainconfiguration (Resource)

Certain domain configurations can only be done when domain administrative account is available. But domain administrative account can only be created after domain creation. Therefore, we need a resource type take care of post domain creation configuration.
This resource allows you to create/update/delete domain reconciliation and Zone Role workflow configurations.

## Example Usage

```terraform
resource "centrifyvault_vaultdomainconfiguration" "domain_config" {
    domain_id = centrifyvault_vaultdomain.example_com.id
    administrative_account_id = centrifyvault_vaultaccount.ad_admin_vaulted.id
    auto_domain_account_maintenance = true
    manual_domain_account_unlock = true
    auto_local_account_maintenance = true
    manual_local_account_unlock = true

    provisioning_admin_id = centrifyvault_vaultaccount.ad_admin_vaulted.id
    reconciliation_account_name = "centrify_lapr"

    // Zone Role Workflow
    enable_zonerole_workflow = true
    assigned_zonerole {
      name = "Windows Login/Global" // name is in format of "<zone role name>/<zone name>"
    }
    assigned_zonerole_approver {
        guid = data.centrifyvault_role.system_admin.id
        name = data.centrifyvault_role.system_admin.name
        type = "Role"
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomainconfiguratoin)

## Argument Reference

### Required

- `domain_id` - (String) ID of the domain whose reconciliation settings are to be configured.
- `administrative_account_id` - (String) ID of the domain administrative account. This is either a vaulted domain account id or directory object id.

### Optional

- `administrative_account_name` - (String) Name of administrative account. This is required if `administrative_account_id` is a directory object id.
- `administrative_account_password` - (String, Sensitive) Password of administrative account. This is required if `administrative_account_id` is a directory object id. To use non-vaulted account, password must be provided.
- `auto_domain_account_maintenance` - (Boolean) Enable Automatic Domain Account Maintenance. Default is `false`.
- `manual_domain_account_unlock` - (Boolean) Enable Manual Domain Account Unlock. Default is `false`.
- `auto_local_account_maintenance` - (Boolean) Enable Automatic Local Account Maintenance. Default is `false`.
- `manual_local_account_unlock` - (Boolean) Enable Manual Local Account Unlock. Default is `false`.
- `provisioning_admin_id` - (String) Provisioning Administrative Account (must be managed). This is a vaulted domain account id.
- `reconciliation_account_name` - (Sring) Reconciliation account name.
- `enable_zonerole_workflow` - (Boolean) Enable zone role requests for systems in the domain.
- `assigned_zonerole` - (Block Set) List of assignable Zone Roles. Refer to [assigned_zonerole](./attribute_assigned_zonerole.md) attribute for details.
- `assigned_zonerole_approver` - (Block List) List of approvers for Zone Role request. Refer to [workflow_approver](./attribute_workflow_approver.md) and [assigned_zonerole_approver](./attribute_assigned_zonerole.md) for details.
