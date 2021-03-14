data "centrifyvault_passwordprofile" "win_profile" {
    name = "Windows Profile"
}

data "centrifyvault_passwordprofile" "custom_profile" {
    name = "Generic Password Profile"
    //profile_type = "UserDefined"
}