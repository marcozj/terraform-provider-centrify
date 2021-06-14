// Used by creation of set
data "centrify_user" "admin" {
    username = "admin@centrify.com.207"
}

// Used by creation of set
data "centrify_role" "system_admin" {
    name = "System Administrator"
}

// Existing System set
// Only works for user created set
data "centrify_manualset" "aws_systems" {
    type = "Server"
    name = "AWS Systems"
}

output "id" {
    value = data.centrify_manualset.aws_systems.id
}
output "name" {
    value = data.centrify_manualset.aws_systems.name
}
output "type" {
    value = data.centrify_manualset.aws_systems.type
}
output "subtype" {
    value = data.centrify_manualset.aws_systems.subtype
}
output "description" {
    value = data.centrify_manualset.aws_systems.description
}