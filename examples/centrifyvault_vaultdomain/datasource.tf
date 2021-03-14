data "centrifyvault_vaultdomain" "example_com" {
    name = "example.com"
}

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