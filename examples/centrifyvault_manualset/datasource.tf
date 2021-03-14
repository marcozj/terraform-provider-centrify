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
