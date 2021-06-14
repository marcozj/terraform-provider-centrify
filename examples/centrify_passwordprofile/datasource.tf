data "centrify_passwordprofile" "win_profile" {
    name = "Windows Profile"
}

data "centrify_passwordprofile" "custom_profile" {
    name = "Generic Password Profile"
    //profile_type = "UserDefined"
}
output "id" {
    value = data.centrify_passwordprofile.custom_profile.id
}
output "name" {
    value = data.centrify_passwordprofile.custom_profile.name
}
output "profile_type" {
    value = data.centrify_passwordprofile.custom_profile.profile_type
}
output "description" {
    value = data.centrify_passwordprofile.custom_profile.description
}
output "minimum_password_length" {
    value = data.centrify_passwordprofile.custom_profile.minimum_password_length
}
output "maximum_password_length" {
    value = data.centrify_passwordprofile.custom_profile.maximum_password_length
}
output "at_least_one_lowercase" {
    value = data.centrify_passwordprofile.custom_profile.at_least_one_lowercase
}
output "at_least_one_uppercase" {
    value = data.centrify_passwordprofile.custom_profile.at_least_one_uppercase
}
output "at_least_one_digit" {
    value = data.centrify_passwordprofile.custom_profile.at_least_one_digit
}
output "no_consecutive_repeated_char" {
    value = data.centrify_passwordprofile.custom_profile.no_consecutive_repeated_char
}
output "at_least_one_special_char" {
    value = data.centrify_passwordprofile.custom_profile.at_least_one_special_char
}
output "maximum_char_occurrence_count" {
    value = data.centrify_passwordprofile.custom_profile.maximum_char_occurrence_count
}
output "special_charset" {
    value = data.centrify_passwordprofile.custom_profile.special_charset
}
output "first_character_type" {
    value = data.centrify_passwordprofile.custom_profile.first_character_type
}
output "last_character_type" {
    value = data.centrify_passwordprofile.custom_profile.last_character_type
}
output "minimum_alphabetic_character_count" {
    value = data.centrify_passwordprofile.custom_profile.minimum_alphabetic_character_count
}
output "minimum_non_alphabetic_character_count" {
    value = data.centrify_passwordprofile.custom_profile.minimum_non_alphabetic_character_count
}
