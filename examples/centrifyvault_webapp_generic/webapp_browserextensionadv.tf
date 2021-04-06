
resource "centrifyvault_webapp_generic" "browserextadvapp" {
    name = "Test Browser Extension Advanced App"
    template_name = "GenericBrowserExtensionScript"
    description = "Test Browser Extension Advanced Application"
    url = "https://www.google.com"

    username_strategy = "Fixed"
    //use_ad_login_pw = true
    username = "username"
    password = "password"
    //use_ad_login_pw_by_script = true
    //user_map_script = "test;"
    script = "test;"

    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }
    //policy_script = "test;"

    sets = [
        data.centrifyvault_manualset.test_set.id
    ]

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Run"]
    }

    workflow_enabled = true
    workflow_approver {
      type = "Manager"
      no_manager_action = "useBackup"
      backup_approver {
        guid = data.centrifyvault_role.approvers.id
        name = data.centrifyvault_role.approvers.name
        type = "Role"
      }
    }
    workflow_approver {
      guid = data.centrifyvault_role.system_admin.id
      name = data.centrifyvault_role.system_admin.name
      type = "Role"
      options_selector = true
    }
    
}
