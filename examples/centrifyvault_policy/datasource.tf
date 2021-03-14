data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}