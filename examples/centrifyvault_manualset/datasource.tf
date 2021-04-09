// Used by creation of set
data "centrifyvault_user" "admin" {
    username = "admin@centrify.com.207"
}

// Used by creation of set
data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

// Existing System set
// Only works for user created set
data "centrifyvault_manualset" "aws_systems" {
    type = "Server"
    name = "AWS Systems"
}

output "id" {
    value = data.centrifyvault_manualset.aws_systems.id
}
output "name" {
    value = data.centrifyvault_manualset.aws_systems.name
}
output "type" {
    value = data.centrifyvault_manualset.aws_systems.type
}
output "subtype" {
    value = data.centrifyvault_manualset.aws_systems.subtype
}
output "description" {
    value = data.centrifyvault_manualset.aws_systems.description
}