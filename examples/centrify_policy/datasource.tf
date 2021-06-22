data "centrify_role" "system_admin" {
    name = "System Administrator"
}

data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

// Data source for existing policy
data "centrify_policy" "mypolicy" {
    name = "Default Policy"
}
output "id" {
    value = data.centrify_policy.mypolicy.id
}
output "name" {
    value = data.centrify_policy.mypolicy.name
}
output "description" {
    value = data.centrify_policy.mypolicy.description
}
output "link_type" {
    value = data.centrify_policy.mypolicy.link_type
}
output "policy_assignment" {
    value = data.centrify_policy.mypolicy.policy_assignment
}
output "settings" {
    value = data.centrify_policy.mypolicy.settings
}