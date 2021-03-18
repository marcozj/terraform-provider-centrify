---
subcategory: "Resources"
---

# centrifyvault_vaultdomainreconciliation (Resource)

This resource allows you to create/update/delete domain reconciliation settings.

## Example Usage

```terraform
resource "centrifyvault_vaultdomainreconciliation" "domain_recon" {
    domain_id = centrifyvault_vaultdomain.example_com.id
    administrative_account_id = centrifyvault_vaultaccount.ad_admin_vaulted.id
    auto_domain_account_maintenance = true
    manual_domain_account_unlock = true
    auto_local_account_maintenance = true
    manual_local_account_unlock = true

    provisioning_admin_id = centrifyvault_vaultaccount.ad_admin_vaulted.id
    reconciliation_account_name = "centrify_lapr"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdomainreconciliation)

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
