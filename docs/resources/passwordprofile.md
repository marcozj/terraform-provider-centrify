---
subcategory: "Settings"
---

# centrifyvault_passwordprofile (Resource)

This resource allows you to create/update/delete password profile.

## Example Usage

```terraform
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
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_passwordprofile)

## Argument Reference

### Required

- `name` - (String) The name of the password profile.
- `description` - (String) Description of password profile.
- `minimum_password_length` - (Number) Minimum password length. Range between `4` to `128`.
- `maximum_password_length` - (Number) Maximum password length. Range between `8` to `128`.
- `special_charset` - (String) Special Characters.

### Optional

- `at_least_one_lowercase` - (Boolean) At least one lower-case alpha character. Default is `true`.
- `at_least_one_uppercase` - (Boolean) At least one upper-case alpha character. Default is `true`.
- `at_least_one_digit` - (Boolean) At least one digit. Default is `true`.
- `no_consecutive_repeated_char` - (Boolean) No consecutive repeated characters.
- `at_least_one_special_char` - (Boolean) At least one special character. Default is `true`.
- `maximum_char_occurrence_count` - (Number) Maximum character occurrence count. Range between `1` to `128`.
- `first_character_type` - (String) A leading alpha or alphanumeric character. Can be set to `AnyChar`, `AlphaOnly`, or `AlphaNumericOnly`. Default is `AnyChar`.
- `last_character_type` - (String) A trailing alpha or alphanumeric character. Can be set to `AnyChar`, `AlphaOnly`, or `AlphaNumericOnly`. Default is `AnyChar`.
- `minimum_alphabetic_character_count` - (Number) Min number of alpha characters. Range between `1` to `128`.
- `minimum_non_alphabetic_character_count`-  (Number) Min number of non-alpha characters. Range between `1` to `128`.
