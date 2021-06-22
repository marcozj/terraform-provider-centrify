resource "centrify_account" "ibmi_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrify_system.ibmi1.id
    description = "Test Account for IBMi"
    use_proxy_account = false
    checkout_lifetime = 70
    managed = false
    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Checkout","Login","FileTransfer","Edit","Delete","UpdatePassword","WorkspaceLogin","RotatePassword"]
    }

    sets = [
        data.centrify_manualset.test_accounts.id
    ]
}