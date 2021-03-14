
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        self_service {
            // Password Reset
            account_selfservice_enabled = true
            password_reset_enabled = true
            pwreset_allow_for_aduser = true
            pwreset_with_cookie_only = true
            login_after_reset = true
            pwreset_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            max_reset_attempts = 5
            // Account Unlock
            account_unlock_enabled = true
            unlock_allow_for_aduser = true
            unlock_with_cookie_only = true
            show_locked_message = true
            unlock_auth_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            // Active Directory Self Service Settings
            use_ad_admin = false
            ad_admin_user = "ad_admin"
            admin_user_password {
                //type = "SafeString"
                value = "xxxxxxxxxxx"
            }
            // Additional Policy Parameters
            max_reset_allowed = 6
            max_time_allowed = 50
        }
    }
}