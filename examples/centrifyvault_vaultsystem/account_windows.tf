resource "centrifyvault_vaultaccount" "windows_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrifyvault_vaultsystem.windows1.id
    description = "Test Account for Windows"
    use_proxy_account = false
    checkout_lifetime = 70
    managed = false
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }

    // Enable account workflow
    workflow_enabled = true
    workflow_approver {
      guid = data.centrifyvault_role.system_admin.id
      name = data.centrifyvault_role.system_admin.name
      type = "Role"
    }
    
    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Checkout","Login","FileTransfer","Edit","Delete","UpdatePassword","WorkspaceLogin","RotatePassword"]
    }

    sets = [
        data.centrifyvault_manualset.test_accounts.id
    ]
}