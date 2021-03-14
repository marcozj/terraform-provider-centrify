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

data "centrifyvault_vaultdatabase" "sql-centrifysuite" {
    name = "SQL-CENTRIFYSUITE"
}

data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}