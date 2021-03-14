
data "centrifyvault_manualset" "test_set" {
  type = "DataVault"
  name = "Test Secrets"
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("DataVault|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        secret_set {
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            challenge_rule {
                authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
                rule {
                    filter = "IpAddress"
                    condition = "OpInCorpIpRange"
                }
            }
        }
    }
}