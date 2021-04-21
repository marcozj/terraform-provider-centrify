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

data "centrifyvault_user" "approver" {
  username = "approver@example.com"
}

// Data source for desktop app

data "centrifyvault_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
}

output "id" {
    value = data.centrifyvault_desktopapp.test_desktopapp.id
}
output "name" {
    value = data.centrifyvault_desktopapp.test_desktopapp.name
}
output "template_name" {
    value = data.centrifyvault_desktopapp.test_desktopapp.template_name
}
output "description" {
    value = data.centrifyvault_desktopapp.test_desktopapp.description
}
output "application_host_id" {
    value = data.centrifyvault_desktopapp.test_desktopapp.application_host_id
}
output "login_credential_type" {
    value = data.centrifyvault_desktopapp.test_desktopapp.login_credential_type
}
output "application_account_id" {
    value = data.centrifyvault_desktopapp.test_desktopapp.application_account_id
}
output "application_alias" {
    value = data.centrifyvault_desktopapp.test_desktopapp.application_alias
}
output "command_line" {
    value = data.centrifyvault_desktopapp.test_desktopapp.command_line
}
output "command_parameter" {
    value = data.centrifyvault_desktopapp.test_desktopapp.command_parameter
}
output "default_profile_id" {
    value = data.centrifyvault_desktopapp.test_desktopapp.default_profile_id
}
output "workflow_enabled" {
    value = data.centrifyvault_desktopapp.test_desktopapp.workflow_enabled
}
output "workflow_approver" {
    value = data.centrifyvault_desktopapp.test_desktopapp.workflow_approver
}
output "challenge_rule" {
    value = data.centrifyvault_desktopapp.test_desktopapp.challenge_rule
}
output "policy_script" {
    value = data.centrifyvault_desktopapp.test_desktopapp.policy_script
}