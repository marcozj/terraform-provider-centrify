data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_manualset" "test_set" {
  type = "DataVault"
  name = "Test Set"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

resource "centrifyvault_vaultsecret" "test_secret" {
    secret_name = "Test Secret"
    description = "Test Secret"
    secret_text = "xxxxxxxxxxxxx"
    type = "Text"
    folder_id = centrifyvault_vaultsecretfolder.level2_folder.id
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    
    sets = [
        data.centrifyvault_manualset.test_set.id,
    ]

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","RetrieveSecret"]
    }

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
        value = "L,1,3,4,5"
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


