data "centrifyvault_vaultdomain" "example_com" {
    name = "demo.lab"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

// Multiplex account data source
data "centrifyvault_multiplexedaccount" "testmultiplex" {
  name = "Test Multiplex"
}
output "id" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.id
}
output "name" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.name
}
output "description" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.description
}
output "account1_id" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.account1_id
}
output "account2_id" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.account2_id
}
output "account1" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.account1
}
output "account2" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.account2
}
output "accounts" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.accounts
}
output "active_account" {
    value = data.centrifyvault_multiplexedaccount.testmultiplex.active_account
}


// Service data source
data "centrifyvault_service" "testservice" {
    service_name = "TestWindowsService"
}

output "id" {
    value = data.centrifyvault_service.testservice.id
}
output "service_name" {
    value = data.centrifyvault_service.testservice.service_name
}
output "system_id" {
    value = data.centrifyvault_service.testservice.system_id
}
output "description" {
    value = data.centrifyvault_service.testservice.description
}
output "service_type" {
    value = data.centrifyvault_service.testservice.service_type
}
output "enable_management" {
    value = data.centrifyvault_service.testservice.enable_management
}
output "admin_account_id" {
    value = data.centrifyvault_service.testservice.admin_account_id
}
output "multiplexed_account_id" {
    value = data.centrifyvault_service.testservice.multiplexed_account_id
}
output "restart_service" {
    value = data.centrifyvault_service.testservice.restart_service
}
output "restart_time_restriction" {
    value = data.centrifyvault_service.testservice.restart_time_restriction
}
output "days_of_week" {
    value = data.centrifyvault_service.testservice.days_of_week
}
output "restart_start_time" {
    value = data.centrifyvault_service.testservice.restart_start_time
}
output "restart_end_time" {
    value = data.centrifyvault_service.testservice.restart_end_time
}
output "use_utc_time" {
    value = data.centrifyvault_service.testservice.use_utc_time
}
