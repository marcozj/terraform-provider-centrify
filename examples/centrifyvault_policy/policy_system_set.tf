
data "centrifyvault_manualset" "test_set" {
  type = "Server"
  name = "Test Set"
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("Server|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        system_set {
            // Account Policy
            checkout_lifetime = 60
            // System Policy
            allow_remote_access = true
            allow_rdp_clipboard = true
            local_account_automatic_maintenance = true
            local_account_manual_unlock = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            privilege_elevation_default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            // Security Settings
            remove_user_on_session_end = true
            allow_multiple_checkouts = true
            enable_password_rotation = true
            password_rotate_interval = 80
            enable_password_rotation_after_checkin = true
            minimum_password_age = 30
            minimum_sshkey_age = 30
            enable_sshkey_rotation = true
            sshkey_rotate_interval = 90
            sshkey_algorithm = "RSA_2048"
            // Maintenance Settings
            enable_password_history_cleanup = true
            password_historycleanup_duration = 120
            enable_sshkey_history_cleanup = true
            sshkey_historycleanup_duration = 120
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
            privilege_elevation_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
        }
    }
}