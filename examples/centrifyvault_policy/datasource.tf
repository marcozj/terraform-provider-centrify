data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

// Data source for existing policy
data "centrifyvault_policy" "mypolicy" {
    name = "Default Policy"
}
output "id" {
    value = data.centrifyvault_policy.mypolicy.id
}
output "name" {
    value = data.centrifyvault_policy.mypolicy.name
}
output "description" {
    value = data.centrifyvault_policy.mypolicy.description
}
output "link_type" {
    value = data.centrifyvault_policy.mypolicy.link_type
}
output "policy_assignment" {
    value = data.centrifyvault_policy.mypolicy.policy_assignment
}
output "settings" {
    value = data.centrifyvault_policy.mypolicy.settings
}