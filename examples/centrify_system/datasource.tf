data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrify_domain" "example_com" {
    name = "demo.lab"
}

data "centrify_manualset" "test_set" {
  type = "Server"
  name = "Test Set"
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

data "centrify_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}

data "centrify_connector" "connector1" {
    name = "dc01" // Connector name registered in Centrify
}

// Data source for existing system
data "centrify_system" "system" {
    name = "centos1"
    fqdn = "centos1.demo.lab"
    computer_class = "Unix"
}

output "id" {
    value = data.centrify_system.system.id
}
output "name" {
    value = data.centrify_system.system.name
}
output "fqdn" {
    value = data.centrify_system.system.fqdn
}
output "computer_class" {
    value = data.centrify_system.system.computer_class
}
output "session_type" {
    value = data.centrify_system.system.session_type
}
output "description" {
    value = data.centrify_system.system.description
}
output "port" {
    value = data.centrify_system.system.port
}
output "use_my_account" {
    value = data.centrify_system.system.use_my_account
}
output "management_mode" {
    value = data.centrify_system.system.management_mode
}
output "management_port" {
    value = data.centrify_system.system.management_port
}
output "system_timezone" {
    value = data.centrify_system.system.system_timezone
}
output "status" {
    value = data.centrify_system.system.status
}
output "proxyuser" {
    value = data.centrify_system.system.proxyuser
}
output "proxyuser_managed" {
    value = data.centrify_system.system.proxyuser_managed
}
output "checkout_lifetime" {
    value = data.centrify_system.system.checkout_lifetime
}
output "allow_remote_access" {
    value = data.centrify_system.system.allow_remote_access
}
output "allow_rdp_clipboard" {
    value = data.centrify_system.system.allow_rdp_clipboard
}
output "default_profile_id" {
    value = data.centrify_system.system.default_profile_id
}
output "privilege_elevation_default_profile_id" {
    value = data.centrify_system.system.privilege_elevation_default_profile_id
}
output "local_account_automatic_maintenance" {
    value = data.centrify_system.system.local_account_automatic_maintenance
}
output "local_account_manual_unlock" {
    value = data.centrify_system.system.local_account_manual_unlock
}
output "domain_id" {
    value = data.centrify_system.system.domain_id
}
output "remove_user_on_session_end" {
    value = data.centrify_system.system.remove_user_on_session_end
}
output "allow_multiple_checkouts" {
    value = data.centrify_system.system.allow_multiple_checkouts
}
output "enable_password_rotation" {
    value = data.centrify_system.system.enable_password_rotation
}
output "password_rotate_interval" {
    value = data.centrify_system.system.password_rotate_interval
}
output "enable_password_rotation_after_checkin" {
    value = data.centrify_system.system.enable_password_rotation_after_checkin
}
output "minimum_password_age" {
    value = data.centrify_system.system.minimum_password_age
}
output "password_profile_id" {
    value = data.centrify_system.system.password_profile_id
}
output "enable_password_history_cleanup" {
    value = data.centrify_system.system.enable_password_history_cleanup
}
output "password_historycleanup_duration" {
    value = data.centrify_system.system.password_historycleanup_duration
}
output "enable_sshkey_rotation" {
    value = data.centrify_system.system.enable_sshkey_rotation
}
output "sshkey_rotate_interval" {
    value = data.centrify_system.system.sshkey_rotate_interval
}
output "minimum_sshkey_age" {
    value = data.centrify_system.system.minimum_sshkey_age
}
output "sshkey_algorithm" {
    value = data.centrify_system.system.sshkey_algorithm
}
output "enable_sshkey_history_cleanup" {
    value = data.centrify_system.system.enable_sshkey_history_cleanup
}
output "sshkey_historycleanup_duration" {
    value = data.centrify_system.system.sshkey_historycleanup_duration
}
output "agent_auth_workflow_enabled" {
    value = data.centrify_system.system.agent_auth_workflow_enabled
}
output "agent_auth_workflow_approver" {
    value = data.centrify_system.system.agent_auth_workflow_approver
}
output "privilege_elevation_workflow_approver" {
    value = data.centrify_system.system.privilege_elevation_workflow_approver
}
output "use_domainadmin_for_zonerole_workflow" {
    value = data.centrify_system.system.use_domainadmin_for_zonerole_workflow
}
output "enable_zonerole_workflow" {
    value = data.centrify_system.system.enable_zonerole_workflow
}
output "use_domain_assignment_for_zoneroles" {
    value = data.centrify_system.system.use_domain_assignment_for_zoneroles
}
output "assigned_zonerole" {
    value = data.centrify_system.system.assigned_zonerole
}
output "use_domain_assignment_for_zonerole_approvers" {
    value = data.centrify_system.system.use_domain_assignment_for_zonerole_approvers
}
output "assigned_zonerole_approver" {
    value = data.centrify_system.system.assigned_zonerole_approver
}
output "connector_list" {
    value = data.centrify_system.system.connector_list
}
output "challenge_rule" {
    value = data.centrify_system.system.challenge_rule
}
output "privilege_elevation_rule" {
    value = data.centrify_system.system.privilege_elevation_rule
}
