
data "centrifyvault_manualset" "test_set" {
  type = "VaultAccount"
  name = "Test Accounts"
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("VaultAccount|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        account_set {
            checkout_lifetime = 60
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
            access_secret_checkout_dfault_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            access_secret_checkout_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
        }
    }
}