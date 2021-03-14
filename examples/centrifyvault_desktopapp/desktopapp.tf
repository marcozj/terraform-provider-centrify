resource "centrifyvault_desktopapp" "test_desktopapp" {
    name = "Test Desktop App"
    template_name = "GenericDesktopApplication"
    description = "Test Desktop Application"
    application_host_id = data.centrifyvault_vaultsystem.apphost.id
    login_credential_type = "SharedAccount"
    application_account_id = data.centrifyvault_vaultaccount.shared_account.id
    application_alias = "pas_desktopapp"
    
    command_line = "--ini=ini\\web_myapp.ini --username={user.User} --password={user.Password}"
    command_parameter {
        name = "system"
        type = "Server"
        target_object_id = data.centrifyvault_vaultsystem.my_app.id
    }
    command_parameter {
        name = "user"
        type = "VaultAccount"
        target_object_id = data.centrifyvault_vaultaccount.admin.id
    }
    
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }
    
    sets = [
        data.centrifyvault_manualset.test_set.id
    ]

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Run"]
    }
}
