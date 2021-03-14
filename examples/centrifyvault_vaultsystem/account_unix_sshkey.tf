data "centrifyvault_sshkey" "testkey" {
    name = "My Test Key"
    key_pair_type = "PrivateKey"
    passphrase = ""
    key_format = "PEM"
    checkout = true
}

resource "centrifyvault_vaultaccount" "unix_account2" {
    name = "testaccount2"
    credential_type = "SshKey"
    sshkey_id = data.centrifyvault_sshkey.testkey.id
    host_id = centrifyvault_vaultsystem.unix1.id
    description = "Test Account for Unix"
    managed = false
    //default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
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