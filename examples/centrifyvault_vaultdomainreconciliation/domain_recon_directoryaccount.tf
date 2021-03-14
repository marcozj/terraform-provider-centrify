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
    name = "ad_admin@example.com"
    object_type = "User"
}

resource "centrifyvault_vaultdomainreconciliation" "example_com_recon" {
    domain_id = data.centrifyvault_vaultdomain.example_com.id // For Terraform managed domain, this can be directly from resource rather than data source
    administrative_account_name = data.centrifyvault_directoryobject.ad_admin.system_name
    administrative_account_id = data.centrifyvault_directoryobject.ad_admin.id
    administrative_account_password = "xxxxxxx" // Actual password of "ad_admin@example.com"

    auto_domain_account_maintenance = true
    manual_domain_account_unlock = true
    auto_local_account_maintenance = true
    manual_local_account_unlock = true
}
