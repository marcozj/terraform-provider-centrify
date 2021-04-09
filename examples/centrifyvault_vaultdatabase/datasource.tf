data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

data "centrifyvault_manualset" "test_set" {
  type = "VaultDatabase"
  name = "Test Set"
}

data "centrifyvault_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}

data "centrifyvault_connector" "connector1" {
  name = "dc01" // Connector name registered in Centrify
}

data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

// Data source for existing database
data "centrifyvault_vaultdatabase" "database" {
    name = "SQL-CENTRIFYSUITE"
}

output "id" {
    value = data.centrifyvault_vaultdatabase.database.id
}
output "name" {
    value = data.centrifyvault_vaultdatabase.database.name
}
output "hostname" {
    value = data.centrifyvault_vaultdatabase.database.hostname
}
output "database_class" {
    value = data.centrifyvault_vaultdatabase.database.database_class
}
output "instance_name" {
    value = data.centrifyvault_vaultdatabase.database.instance_name
}
output "service_name" {
    value = data.centrifyvault_vaultdatabase.database.service_name
}
output "description" {
    value = data.centrifyvault_vaultdatabase.database.description
}
output "port" {
    value = data.centrifyvault_vaultdatabase.database.port
}
output "checkout_lifetime" {
    value = data.centrifyvault_vaultdatabase.database.checkout_lifetime
}
output "allow_multiple_checkouts" {
    value = data.centrifyvault_vaultdatabase.database.allow_multiple_checkouts
}
output "enable_password_rotation" {
    value = data.centrifyvault_vaultdatabase.database.enable_password_rotation
}
output "password_rotate_interval" {
    value = data.centrifyvault_vaultdatabase.database.password_rotate_interval
}
output "enable_password_rotation_after_checkin" {
    value = data.centrifyvault_vaultdatabase.database.enable_password_rotation_after_checkin
}
output "minimum_password_age" {
    value = data.centrifyvault_vaultdatabase.database.minimum_password_age
}
output "password_profile_id" {
    value = data.centrifyvault_vaultdatabase.database.password_profile_id
}
output "enable_password_history_cleanup" {
    value = data.centrifyvault_vaultdatabase.database.enable_password_history_cleanup
}
output "password_historycleanup_duration" {
    value = data.centrifyvault_vaultdatabase.database.password_historycleanup_duration
}
output "connector_list" {
    value = data.centrifyvault_vaultdatabase.database.connector_list
}
