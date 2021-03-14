
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

resource "centrifyvault_authenticationprofile" "stepup" {
    name = "Step-up Authentication Profile"
    challenges = [
        "OTP,PF,SMS,EMAIL,OATH,RADIUS,U2F,SQ",
    ]
    additional_data {
        number_of_questions = 1
    }
    pass_through_duration = 0
}
