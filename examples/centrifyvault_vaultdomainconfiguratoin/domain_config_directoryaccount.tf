// Existing domain in Centrify Vault
data "centrifyvault_vaultdomain" "example_com" {
    name = "example.com"
}

// data source for example.com domain
data "centrifyvault_directoryservice" "example_com_dir" {
    name = "example.com"
    type = "Active Directory"
}

// data source for AD accounnt ad_admin@example.com
data "centrifyvault_directoryobject" "ad_admin" {
    directory_services = [
        data.centrifyvault_directoryservice.example_com_dir.id
    ]
    //name = "ad_admin@example.com"
    name = "ad_admin@demo.lab"
    object_type = "User"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

// This example sets an non-vaulted Active Directory account as domain administrative account
resource "centrifyvault_vaultdomainconfiguration" "example_com_config" {
    domain_id = data.centrifyvault_vaultdomain.example_com.id // For Terraform managed domain, this can be directly from resource rather than data source
    administrative_account_name = data.centrifyvault_directoryobject.ad_admin.system_name
    administrative_account_id = data.centrifyvault_directoryobject.ad_admin.id
    administrative_account_password = "xxxxxxx" // Actual password of "ad_admin@example.com"

    auto_domain_account_maintenance = true
    manual_domain_account_unlock = true
    auto_local_account_maintenance = true
    manual_local_account_unlock = true

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
