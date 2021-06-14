data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrify_system" "apphost" {
    name = "apphost"
    fqdn = "apphost.demo.lab"
    computer_class = "Windows"
}

data "centrify_system" "my_app" {
    name = "My App"
    fqdn = "192.168.18.15"
    computer_class = "CustomSsh"
}

data "centrify_domain" "demo_lab" {
    name = "demo.lab"
}

data "centrify_account" "shared_account" {
    name = "shared_account"
    domain_id = data.centrify_domain.demo_lab.id
}

data "centrify_account" "admin" {
    name = "admin"
    host_id = data.centrify_system.my_app.id
}

data "centrify_manualset" "test_set" {
  type = "Application"
  name = "Test Set"
  subtype = "Desktop"
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

data "centrify_user" "approver" {
  username = "approver@example.com"
}

// Data source for desktop app

data "centrify_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
}

output "id" {
    value = data.centrify_desktopapp.test_desktopapp.id
}
output "name" {
    value = data.centrify_desktopapp.test_desktopapp.name
}
output "template_name" {
    value = data.centrify_desktopapp.test_desktopapp.template_name
}
output "description" {
    value = data.centrify_desktopapp.test_desktopapp.description
}
output "application_host_id" {
    value = data.centrify_desktopapp.test_desktopapp.application_host_id
}
output "login_credential_type" {
    value = data.centrify_desktopapp.test_desktopapp.login_credential_type
}
output "application_account_id" {
    value = data.centrify_desktopapp.test_desktopapp.application_account_id
}
output "application_alias" {
    value = data.centrify_desktopapp.test_desktopapp.application_alias
}
output "command_line" {
    value = data.centrify_desktopapp.test_desktopapp.command_line
}
output "command_parameter" {
    value = data.centrify_desktopapp.test_desktopapp.command_parameter
}
output "default_profile_id" {
    value = data.centrify_desktopapp.test_desktopapp.default_profile_id
}
output "workflow_enabled" {
    value = data.centrify_desktopapp.test_desktopapp.workflow_enabled
}
output "workflow_approver" {
    value = data.centrify_desktopapp.test_desktopapp.workflow_approver
}
output "challenge_rule" {
    value = data.centrify_desktopapp.test_desktopapp.challenge_rule
}
output "policy_script" {
    value = data.centrify_desktopapp.test_desktopapp.policy_script
}