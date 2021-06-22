resource "centrify_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
    template_name = "GenericDesktopApplication"
    description = "Test Desktop Application"
    application_host_id = data.centrify_system.apphost.id
    login_credential_type = "SharedAccount"
    application_account_id = data.centrify_account.shared_account.id
    application_alias = "pas_desktopapp"
    
    command_line = "--ini=ini\\web_myapp.ini --username={user.User} --password={user.Password}"
    command_parameter {
        name = "system"
        type = "Server"
        target_object_id = data.centrify_system.my_app.id
        value = data.centrify_system.my_app.name
    }
    command_parameter {
        name = "user"
        type = "VaultAccount"
        target_object_id = data.centrify_account.admin.id
        value = data.centrify_account.admin.name
    }
    
    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }
    
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
      guid = data.centrify_user.approver.id
      name = data.centrify_user.approver.username
      type = "User"
      options_selector = true // this attribute must be added to only one approver if there are multiple levels
    }
    workflow_approver {
      guid = data.centrify_role.system_admin.id
      name = data.centrify_role.system_admin.name
      type = "Role"
    }
}
