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

data "centrifyvault_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}

data "centrifyvault_connector" "connector1" {
    name = "dc01" // Connector name registered in Centrify
}

// Data source for existing system
data "centrifyvault_vaultsystem" "system" {
    name = "centos1"
    fqdn = "centos1.demo.lab"
    computer_class = "Unix"
}

output "id" {
    value = data.centrifyvault_vaultsystem.system.id
}
output "name" {
    value = data.centrifyvault_vaultsystem.system.name
}
output "fqdn" {
    value = data.centrifyvault_vaultsystem.system.fqdn
}
output "computer_class" {
    value = data.centrifyvault_vaultsystem.system.computer_class
}
output "session_type" {
    value = data.centrifyvault_vaultsystem.system.session_type
}
output "description" {
    value = data.centrifyvault_vaultsystem.system.description
}
output "port" {
    value = data.centrifyvault_vaultsystem.system.port
}
output "use_my_account" {
    value = data.centrifyvault_vaultsystem.system.use_my_account
}
output "management_mode" {
    value = data.centrifyvault_vaultsystem.system.management_mode
}
output "management_port" {
    value = data.centrifyvault_vaultsystem.system.management_port
}
output "system_timezone" {
    value = data.centrifyvault_vaultsystem.system.system_timezone
}
output "status" {
    value = data.centrifyvault_vaultsystem.system.status
}
output "proxyuser" {
    value = data.centrifyvault_vaultsystem.system.proxyuser
}
output "proxyuser_managed" {
    value = data.centrifyvault_vaultsystem.system.proxyuser_managed
}
output "checkout_lifetime" {
    value = data.centrifyvault_vaultsystem.system.checkout_lifetime
}
output "allow_remote_access" {
    value = data.centrifyvault_vaultsystem.system.allow_remote_access
}
output "allow_rdp_clipboard" {
    value = data.centrifyvault_vaultsystem.system.allow_rdp_clipboard
}
output "default_profile_id" {
    value = data.centrifyvault_vaultsystem.system.default_profile_id
}
output "privilege_elevation_default_profile_id" {
    value = data.centrifyvault_vaultsystem.system.privilege_elevation_default_profile_id
}
output "local_account_automatic_maintenance" {
    value = data.centrifyvault_vaultsystem.system.local_account_automatic_maintenance
}
output "local_account_manual_unlock" {
    value = data.centrifyvault_vaultsystem.system.local_account_manual_unlock
}
output "domain_id" {
    value = data.centrifyvault_vaultsystem.system.domain_id
}
output "remove_user_on_session_end" {
    value = data.centrifyvault_vaultsystem.system.remove_user_on_session_end
}
output "allow_multiple_checkouts" {
    value = data.centrifyvault_vaultsystem.system.allow_multiple_checkouts
}
output "enable_password_rotation" {
    value = data.centrifyvault_vaultsystem.system.enable_password_rotation
}
output "password_rotate_interval" {
    value = data.centrifyvault_vaultsystem.system.password_rotate_interval
}
output "enable_password_rotation_after_checkin" {
    value = data.centrifyvault_vaultsystem.system.enable_password_rotation_after_checkin
}
output "minimum_password_age" {
    value = data.centrifyvault_vaultsystem.system.minimum_password_age
}
output "password_profile_id" {
    value = data.centrifyvault_vaultsystem.system.password_profile_id
}
output "enable_password_history_cleanup" {
    value = data.centrifyvault_vaultsystem.system.enable_password_history_cleanup
}
output "password_historycleanup_duration" {
    value = data.centrifyvault_vaultsystem.system.password_historycleanup_duration
}
output "enable_sshkey_rotation" {
    value = data.centrifyvault_vaultsystem.system.enable_sshkey_rotation
}
output "sshkey_rotate_interval" {
    value = data.centrifyvault_vaultsystem.system.sshkey_rotate_interval
}
output "minimum_sshkey_age" {
    value = data.centrifyvault_vaultsystem.system.minimum_sshkey_age
}
output "sshkey_algorithm" {
    value = data.centrifyvault_vaultsystem.system.sshkey_algorithm
}
output "enable_sshkey_history_cleanup" {
    value = data.centrifyvault_vaultsystem.system.enable_sshkey_history_cleanup
}
output "sshkey_historycleanup_duration" {
    value = data.centrifyvault_vaultsystem.system.sshkey_historycleanup_duration
}
output "agent_auth_workflow_enabled" {
    value = data.centrifyvault_vaultsystem.system.agent_auth_workflow_enabled
}
output "agent_auth_workflow_approver" {
    value = data.centrifyvault_vaultsystem.system.agent_auth_workflow_approver
}
output "privilege_elevation_workflow_approver" {
    value = data.centrifyvault_vaultsystem.system.privilege_elevation_workflow_approver
}
output "use_domainadmin_for_zonerole_workflow" {
    value = data.centrifyvault_vaultsystem.system.use_domainadmin_for_zonerole_workflow
}
output "enable_zonerole_workflow" {
    value = data.centrifyvault_vaultsystem.system.enable_zonerole_workflow
}
output "use_domain_assignment_for_zoneroles" {
    value = data.centrifyvault_vaultsystem.system.use_domain_assignment_for_zoneroles
}
output "assigned_zonerole" {
    value = data.centrifyvault_vaultsystem.system.assigned_zonerole
}
output "use_domain_assignment_for_zonerole_approvers" {
    value = data.centrifyvault_vaultsystem.system.use_domain_assignment_for_zonerole_approvers
}
output "assigned_zonerole_approver" {
    value = data.centrifyvault_vaultsystem.system.assigned_zonerole_approver
}
output "connector_list" {
    value = data.centrifyvault_vaultsystem.system.connector_list
}
output "challenge_rule" {
    value = data.centrifyvault_vaultsystem.system.challenge_rule
}
output "privilege_elevation_rule" {
    value = data.centrifyvault_vaultsystem.system.privilege_elevation_rule
}
