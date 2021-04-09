---
subcategory: "Settings"
---

# centrifyvault_passwordprofile (Data Source)

This data source gets information of password profile.

## Example Usage

```terraform
data "centrifyvault_passwordprofile" "win_profile" {
    name = "Windows Profile"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_passwordprofile)

## Search Attributes

### Required

- `name` - (String) The name of password profile.

### Optional

- `profile_type` - (String) The type of password profile.

## Attributes Reference

- `id` - id of the password profile.
- `name` - (String) The name of the password profile.
- `description` - (String) Description of password profile.
- `minimum_password_length` - (Number) Minimum password length.
- `maximum_password_length` - (Number) Maximum password length.
- `special_charset` - (String) Special Characters.
- `at_least_one_lowercase` - (Boolean) At least one lower-case alpha character.
- `at_least_one_uppercase` - (Boolean) At least one upper-case alpha character.
- `at_least_one_digit` - (Boolean) At least one digit.
- `no_consecutive_repeated_char` - (Boolean) No consecutive repeated characters.
- `at_least_one_special_char` - (Boolean) At least one special character.
- `maximum_char_occurrence_count` - (Number) Maximum character occurrence count.
- `first_character_type` - (String) A leading alpha or alphanumeric character.
- `last_character_type` - (String) A trailing alpha or alphanumeric character.
- `minimum_alphabetic_character_count` - (Number) Min number of alpha characters.
- `minimum_non_alphabetic_character_count`-  (Number) Min number of non-alpha characters.
