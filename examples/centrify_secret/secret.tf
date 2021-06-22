data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrify_manualset" "test_set" {
  type = "DataVault"
  name = "Test Set"
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

resource "centrify_secret" "test_secret" {
    secret_name = "Test Secret"
    description = "Test Secret"
    secret_text = "xxxxxxxxxxxxx"
    type = "Text"
    folder_id = centrify_secretfolder.level2_folder.id
    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    
    sets = [
        data.centrify_manualset.test_set.id,
    ]

    workflow_enabled = true
    workflow_approver {
      guid = data.centrify_role.system_admin.id
      name = data.centrify_role.system_admin.name
      type = "Role"
    }
    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","RetrieveSecret"]
    }

    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }

    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
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


