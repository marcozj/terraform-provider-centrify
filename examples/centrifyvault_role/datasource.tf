data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

output "id" {
    value = data.centrifyvault_role.system_admin.id
}
output "name" {
    value = data.centrifyvault_role.system_admin.name
}
output "description" {
    value = data.centrifyvault_role.system_admin.description
}
output "adminrights" {
    value = data.centrifyvault_role.system_admin.adminrights
}
output "member" {
    value = data.centrifyvault_role.system_admin.member
}