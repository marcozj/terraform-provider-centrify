data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_vaultdomain" "example_com" {
    name = "demo.lab"
}

data "centrifyvault_manualset" "test_set" {
  type = "Server"
  name = "Test Set"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

data "centrifyvault_connector" "connector1" {
    name = "dc01" // Connector name registered in Centrify
}

// Existing system
data "centrifyvault_vaultsystem" "centos1" {
    name = "centos1"
    fqdn = "centos1.demo.lab"
    computer_class = "Unix"
}

data "centrifyvault_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}