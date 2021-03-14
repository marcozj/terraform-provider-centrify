data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_vaultsystem" "apphost" {
    name = "apphost"
    fqdn = "apphost.demo.lab"
    computer_class = "Windows"
}

data "centrifyvault_vaultsystem" "my_app" {
    name = "My App"
    fqdn = "192.168.18.15"
    computer_class = "CustomSsh"
}

data "centrifyvault_vaultdomain" "demo_lab" {
    name = "demo.lab"
}

data "centrifyvault_vaultaccount" "shared_account" {
    name = "shared_account"
    domain_id = data.centrifyvault_vaultdomain.demo_lab.id
}

data "centrifyvault_vaultaccount" "admin" {
    name = "admin"
    host_id = data.centrifyvault_vaultsystem.my_app.id
}

data "centrifyvault_manualset" "test_set" {
  type = "Application"
  name = "Test Set"
  subtype = "Desktop"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}