
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        centrify_services {
            authentication_enabled = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            // Session Parameters
            session_lifespan = 23
            allow_session_persist = true
            default_session_persist = true
            persist_session_lifespan = 30
            // Other Settings
            allow_iwa = true
            iwa_set_cookie = true
            iwa_satisfies_all = true
            use_certauth = true
            certauth_skip_challenge = true
            certauth_set_cookie = true
            certauth_satisfies_all = true
            allow_no_mfa_mech = true
            auth_rule_federated = false
            federated_satisfies_all = true
            block_auth_from_same_device = false
            continue_failed_sessions = true
            stop_auth_on_prev_failed = true
            remember_last_factor = true
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "DayOfWeek"
                    condition = "OpIsDayOfWeek"
                    value = "L,1,3,4"
                }
                rule {
                    filter = "Browser"
                    condition = "OpNotEqual"
                    value = "Firefox"
                }
                rule {
                    filter = "CountryCode"
                    condition = "OpNotEqual"
                    value = "GA"
                }
            }
        }
        
    }
    
}