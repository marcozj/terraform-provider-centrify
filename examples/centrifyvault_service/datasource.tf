data "centrifyvault_vaultdomain" "example_com" {
    name = "demo.lab"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

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