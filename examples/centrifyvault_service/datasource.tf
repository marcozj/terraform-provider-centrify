data "centrifyvault_vaultdomain" "example_com" {
    name = "demo.lab"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

data "centrifyvault_multiplexedaccount" "testmultiplex" {
  name = "Test Multiplex"
}
