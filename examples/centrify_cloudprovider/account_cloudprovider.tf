resource "centrify_account" "aws_root_account" {
    name = "rootaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    cloudprovider_id = centrify_cloudprovider.demo_aws.id
    is_root_account = true
    description = "Test Root Account for AWS"
    checkout_lifetime = 70
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
        rights = ["Grant","View","Checkout","Login","Edit","Delete","UpdatePassword","RotatePassword"]
    }

    sets = [
        data.centrify_manualset.test_accounts.id
    ]
}

resource "centrify_account" "iam_account" {
    name = "iamuser"
    credential_type = "AwsAccessKey"
    cloudprovider_id = centrify_cloudprovider.demo_aws.id
    is_root_account = false
    description = "Test IAM Account for AWS"

    access_key {
      access_key_id = "XXXXXXXXXX"
      secret_access_key = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    }

    access_secret_checkout_default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    access_secret_checkout_rule {
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
        rights = ["Grant","View","Edit","Delete","Checkout"]
    }

    sets = [
        data.centrify_manualset.test_accounts.id
    ]
}