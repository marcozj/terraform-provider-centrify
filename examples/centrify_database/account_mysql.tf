resource "centrify_account" "mssql_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    database_id = centrify_database.mssql.id
    description = "Test Account for MS SQL"
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
        rights = ["Grant","View","Checkout","Edit","Delete","UpdatePassword","RotatePassword"]
    }

    sets = [
        data.centrify_manualset.test_accounts.id
    ]
}