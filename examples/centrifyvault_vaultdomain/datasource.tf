data "centrifyvault_passwordprofile" "domain_pw_pf" {
    name = "Domain Profile"
}

data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

data "centrifyvault_manualset" "test_set" {
    type = "VaultDomain"
    name = "Test Set"
}

data "centrifyvault_connector" "connector1" {
    name = "connector_host1" // Connector name registered in Centrify
}

data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}

// Data source for existing domain
data "centrifyvault_vaultdomain" "example_com" {
    name = "example.com"
}

output "id" {
    value = data.centrifyvault_vaultdomain.example_com.id
}
output "name" {
    value = data.centrifyvault_vaultdomain.example_com.name
}
output "description" {
    value = data.centrifyvault_vaultdomain.example_com.description
}
output "forest_id" {
    value = data.centrifyvault_vaultdomain.example_com.forest_id
}
output "parent_id" {
    value = data.centrifyvault_vaultdomain.example_com.parent_id
}
output "checkout_lifetime" {
    value = data.centrifyvault_vaultdomain.example_com.checkout_lifetime
}
output "allow_multiple_checkouts" {
    value = data.centrifyvault_vaultdomain.example_com.allow_multiple_checkouts
}
output "enable_password_rotation" {
    value = data.centrifyvault_vaultdomain.example_com.enable_password_rotation
}
output "password_rotate_interval" {
    value = data.centrifyvault_vaultdomain.example_com.password_rotate_interval
}
output "enable_password_rotation_after_checkin" {
    value = data.centrifyvault_vaultdomain.example_com.enable_password_rotation_after_checkin
}
output "minimum_password_age" {
    value = data.centrifyvault_vaultdomain.example_com.minimum_password_age
}
output "password_profile_id" {
    value = data.centrifyvault_vaultdomain.example_com.password_profile_id
}
output "enable_password_history_cleanup" {
    value = data.centrifyvault_vaultdomain.example_com.enable_password_history_cleanup
}
output "password_historycleanup_duration" {
    value = data.centrifyvault_vaultdomain.example_com.password_historycleanup_duration
}
output "enable_zone_joined_check" {
    value = data.centrifyvault_vaultdomain.example_com.enable_zone_joined_check
}
output "zone_joined_check_interval" {
    value = data.centrifyvault_vaultdomain.example_com.zone_joined_check_interval
}
output "enable_zonerole_cleanup" {
    value = data.centrifyvault_vaultdomain.example_com.enable_zonerole_cleanup
}
output "zonerole_cleanup_interval" {
    value = data.centrifyvault_vaultdomain.example_com.zonerole_cleanup_interval
}
output "connector_list" {
    value = data.centrifyvault_vaultdomain.example_com.connector_list
}
output "administrative_account_id" {
    value = data.centrifyvault_vaultdomain.example_com.administrative_account_id
}
output "administrator_display_name" {
    value = data.centrifyvault_vaultdomain.example_com.administrator_display_name
}
output "administrative_account_name" {
    value = data.centrifyvault_vaultdomain.example_com.administrative_account_name
}
output "administrative_account_password" {
    value = data.centrifyvault_vaultdomain.example_com.administrative_account_password
}
output "auto_domain_account_maintenance" {
    value = data.centrifyvault_vaultdomain.example_com.auto_domain_account_maintenance
}
output "auto_local_account_maintenance" {
    value = data.centrifyvault_vaultdomain.example_com.auto_local_account_maintenance
}
output "manual_domain_account_unlock" {
    value = data.centrifyvault_vaultdomain.example_com.manual_domain_account_unlock
}
output "manual_local_account_unlock" {
    value = data.centrifyvault_vaultdomain.example_com.manual_local_account_unlock
}
output "provisioning_admin_id" {
    value = data.centrifyvault_vaultdomain.example_com.provisioning_admin_id
}
output "reconciliation_account_name" {
    value = data.centrifyvault_vaultdomain.example_com.reconciliation_account_name
}
output "enable_zonerole_workflow" {
    value = data.centrifyvault_vaultdomain.example_com.enable_zonerole_workflow
}
output "assigned_zonerole" {
    value = data.centrifyvault_vaultdomain.example_com.assigned_zonerole
}
output "assigned_zonerole_approver" {
    value = data.centrifyvault_vaultdomain.example_com.assigned_zonerole_approver
}