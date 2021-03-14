
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        user_account {
            allow_user_change_password = true
            password_change_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            show_fido2 = true
            fido2_prompt = "FIDO2 Key"
            fido2_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            show_otp = true
            otp_prompt = "Google Authenticator"
            otp_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            configure_security_questions = true
            prevent_dup_answers = false
            user_defined_questions = 3
            admin_defined_questions = 2
            min_char_in_answer = 2
            question_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            allow_phone_pin_change = true
            min_phone_pin_length = 6
            phone_pin_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            allow_mfa_redirect_change = true
            user_profile_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            default_language = "en"
        }
    }
}