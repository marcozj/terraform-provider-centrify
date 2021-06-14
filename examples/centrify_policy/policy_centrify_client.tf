
resource "centrify_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrify_role.system_admin.id,
    ]
    
    settings {
        centrify_client {
            authentication_enabled = true
            default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
            allow_no_mfa_mech = true
            challenge_rule {
                authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
        }
    }
}