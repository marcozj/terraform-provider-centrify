
data "centrifyvault_manualset" "test_set" {
  type = "CloudProviders"
  name = "Test Set"
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("CloudProviders|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        cloudproviders_set {
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
            enable_interactive_password_rotation = true
            prompt_change_root_password = true
            enable_password_rotation_reminders = true
            password_rotation_reminder_duration = 20
        }
    }
}