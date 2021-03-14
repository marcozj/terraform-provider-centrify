
resource "centrifyvault_passwordprofile" "test_pw_profile" {
    name = "Test Password Profile"
    description = "Test Password Profile"
    minimum_password_length = 8
    maximum_password_length = 16
    special_charset = "!#$%&()*+,-./:;<=>?@[\\]^_{|}~"
    
    at_least_one_lowercase = false
    at_least_one_uppercase = true
    at_least_one_digit = true
    no_consecutive_repeated_char = true
    at_least_one_special_char = true
    maximum_char_occurrence_count = 2
    
    first_character_type = "AnyChar"
    last_character_type = "AlphaNumericOnly"
    minimum_alphabetic_character_count = 1
    minimum_non_alphabetic_character_count = 1
}
