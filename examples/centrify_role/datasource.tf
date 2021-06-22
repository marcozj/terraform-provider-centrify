data "centrify_role" "system_admin" {
    name = "System Administrator"
}

output "id" {
    value = data.centrify_role.system_admin.id
}
output "name" {
    value = data.centrify_role.system_admin.name
}
output "description" {
    value = data.centrify_role.system_admin.description
}
output "adminrights" {
    value = data.centrify_role.system_admin.adminrights
}
output "member" {
    value = data.centrify_role.system_admin.member
}