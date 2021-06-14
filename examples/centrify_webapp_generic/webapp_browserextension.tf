
resource "centrify_webapp_generic" "browserextapp" {
    name = "Test Browser Extension App"
    template_name = "Generic Browser Extension"
    description = "Test Browser Extension Application"
    url = "https://www.google.com"

    username_strategy = "Fixed"
    //use_ad_login_pw = true
    username = "username"
    password = "password"
    //use_ad_login_pw_by_script = true
    //user_map_script = "test;"

    hostname_suffix = "amazon.com"
    username_field = "input#resolving_input"
    password_field = "input[type='password']"
    submit_field = "input#signInSubmit-input"
    form_field = "form#ap_signin_form"
    //additional_login_field = 
    //additional_login_field_value = 
    selector_timeout = 10
    order = "[[\"fill\",\"username\"],[\"click\",\"button#next_button\"],[\"sleep\",\"1000\"],[\"fillEnter\",\"password\"],[\"waitForNewPage\"],[\"fillEnter\",\"password\"]]"

    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }
    //policy_script = "test;"

    sets = [
        data.centrify_manualset.test_set.id
    ]

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Run"]
    }

    workflow_enabled = true
    workflow_approver {
      type = "Manager"
      no_manager_action = "useBackup"
      backup_approver {
        guid = data.centrify_role.approvers.id
        name = data.centrify_role.approvers.name
        type = "Role"
      }
    }
    workflow_approver {
      guid = data.centrify_role.system_admin.id
      name = data.centrify_role.system_admin.name
      type = "Role"
      options_selector = true
    }
    
}
