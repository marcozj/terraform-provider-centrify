// Existing domain
data "centrify_domain" "example_com" {
    //name = "example.com"
    name = "demo.lab"
}

// Existing vaulted domain account used as administrative account
data "centrify_account" "ad_admin_vaulted" {
    name = "ad_admin"
    domain_id = data.centrify_domain.example_com.id
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

// This example sets an vaulted account as domain administrative account
resource "centrify_domainconfiguration" "domain_config" {
    domain_id = data.centrify_domain.example_com.id // For Terraform managed domain, this can be directly from resource rather than data source
    administrative_account_id = data.centrify_account.ad_admin_vaulted.id // For Terraform managed vaulted account, this can be directly from resource rather than data source
    auto_domain_account_maintenance = true
    manual_domain_account_unlock = true
    auto_local_account_maintenance = true
    manual_local_account_unlock = true

    provisioning_admin_id = data.centrify_account.ad_admin_vaulted.id // This can be different from administrative_account_id
    reconciliation_account_name = "centrify_lapr"

    // Zone Role Workflow
    enable_zonerole_workflow = true
    assigned_zonerole {
      name = "Windows Login/Global" // name is in format of "<zone role name>/<zone name>"
    }
    assigned_zonerole_approver {
        guid = data.centrify_role.system_admin.id
        name = data.centrify_role.system_admin.name
        type = "Role"
    }
}
