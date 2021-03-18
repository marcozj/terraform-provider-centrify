---
subcategory: "Settings"
---

# centrifyvault_authenticationprofile (Resource)

This resource allows you to create/update/delete Authentication Profile.

## Example Usage

```terraform
resource "centrifyvault_authenticationprofile" "twofa" {
    name = "2FA Authentication Profile"
    challenges = [
        "UP",
        "OTP,PF,SMS,EMAIL,OATH,RADIUS,U2F,SQ",
    ]
    additional_data {
        number_of_questions = 1
    }
    pass_through_duration = 0
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_authenticationprofile)

## Argument Reference

### Required

- `name` - (String) The name of the authenticaiton profile.
- `challenges` - (List of String) Authentication mechanisms for challenges. Minimum 1 item and maximum 3 items. Each item contains comma separated list of authenticaiton mechanism. For example "UP,OTP,SMS,EMAIL,OATH". Valid short form of authenticaiton mechanism are `UP`, `OTP`, `PF`, `SMS`, `EMAIL`, `OATH`, `RADIUS`, `U2F` and `SQ`. (Password:"UP", MobileAuthenticator:"OTP", PhoneCall:"PF", SMS:"SMS", EmailConfirmationCode:"EMAIL", OATH_OTP:"OATH", Radius:"RADIUS", FIDO2:"U2F", SecurityQuestions:"SQ")

### Optional

- `pass_through_duration` - (Number) Challenge Pass-Through Duration in minutes. Default is `30`.
- `additional_data` (Block List, Max: 1) (see [below reference for additional_data](#reference-for-additional_data))

## [Reference for `additional_data`]

Optional:

- `number_of_questions` - (Number) Number of questions user must answer. Range between `0` to `10`.
