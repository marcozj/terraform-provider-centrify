data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_manualset" "test_set" {
  type = "SshKeys"
  name = "Test Set"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

resource "centrifyvault_sshkey" "test_key" {
  name = "Test Key"
    description = "Test RSA key"
    private_key = file("rsa.key") // rsa.key file must exist
    passphrase = ""
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    
    sets = [
        data.centrifyvault_manualset.test_set.id,
    ]
    
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
      rule {
        filter = "Browser"
        condition = "OpEqual"
        value = "Chrome"
      }
    }
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "Browser"
        condition = "OpEqual"
        value = "Chrome"
      }
    }

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","Retrieve"]
    }
}