
resource "centrify_webapp_oidc" "oidcapp" {
    name = "Test OIDC App"
    template_name = "Generic OpenID Connect"
    application_id = "TestOIDCApp" // No space
    description = "Test OIDC Application"
    //script = "afd23a;"
    //username_strategy = "ADAttribute" // ADAttribute, Fixed, UseScript
    //username = "userprincipalname"

    oauth_profile {
        client_secret = "mysecret"
        application_url = "https://example.com"
        redirects = ["https://example.net", "https://example.com"]
        token_lifetime = "8:00:00"
        allow_refresh = true
        refresh_lifetime = "251.00:00:00"
    }

    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
      rule {
        filter = "DayOfWeek"
        condition = "OpIsDayOfWeek"
        value = "L,1,3,4,5"
      }
    }
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

output "clientid" {
    value = centrify_webapp_oidc.oidcapp.oauth_profile[0].client_id
}
