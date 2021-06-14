data "centrify_role" "system_admin" {
  name = "System Administrator"
}

data "centrify_manualset" "test_set" {
  type = "VaultDatabase"
  name = "Test Set"
}

data "centrify_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}

data "centrify_connector" "connector1" {
  name = "dc01" // Connector name registered in Centrify
}

data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

// Data source for existing database
data "centrify_database" "database" {
    name = "SQL-CENTRIFYSUITE"
}

output "id" {
    value = data.centrify_database.database.id
}
output "name" {
    value = data.centrify_database.database.name
}
output "hostname" {
    value = data.centrify_database.database.hostname
}
output "database_class" {
    value = data.centrify_database.database.database_class
}
output "instance_name" {
    value = data.centrify_database.database.instance_name
}
output "service_name" {
    value = data.centrify_database.database.service_name
}
output "description" {
    value = data.centrify_database.database.description
}
output "port" {
    value = data.centrify_database.database.port
}
output "checkout_lifetime" {
    value = data.centrify_database.database.checkout_lifetime
}
output "allow_multiple_checkouts" {
    value = data.centrify_database.database.allow_multiple_checkouts
}
output "enable_password_rotation" {
    value = data.centrify_database.database.enable_password_rotation
}
output "password_rotate_interval" {
    value = data.centrify_database.database.password_rotate_interval
}
output "enable_password_rotation_after_checkin" {
    value = data.centrify_database.database.enable_password_rotation_after_checkin
}
output "minimum_password_age" {
    value = data.centrify_database.database.minimum_password_age
}
output "password_profile_id" {
    value = data.centrify_database.database.password_profile_id
}
output "enable_password_history_cleanup" {
    value = data.centrify_database.database.enable_password_history_cleanup
}
output "password_historycleanup_duration" {
    value = data.centrify_database.database.password_historycleanup_duration
}
output "connector_list" {
    value = data.centrify_database.database.connector_list
}
