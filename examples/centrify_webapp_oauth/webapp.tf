
resource "centrify_webapp_oauth" "oauthclient" {
    name = "Test OAuth Client App"
    template_name = "OAuth2ServerClient" // OAuth2ServerClient or OAuth2Server
    application_id = "TestOAuthClient" // No space
    description = "Test OAuth Client Application"
    
    oauth_profile {
        clientid_type = 1 // 0 for Anything or List, 1 for Confidential
        //allowed_clients = ["client1", "client2"]
        must_oauth_client = true
        redirects = ["https://example.net", "https://example.com"]
        token_type = "JwtRS256"
        allowed_auth = ["ClientCreds", "ResourceCreds"]
        token_lifetime = "8:00:00"
        allow_refresh = true
        refresh_lifetime = "250.00:00:00"
        confirm_authorization = true
        allow_scope_select = true
        scope {
            name = "cli"
            description = "Used for CLI call"
            allowed_rest_apis = ["/SaasManage/GetApplication", "/RedRock/query"]
        }
        scope {
            name = "aapm"
            description = "Used for AAPM calls"
            allowed_rest_apis = [".*"]
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
}


resource "centrify_webapp_oauth" "oauthserver" {
    name = "Test OAuth Server App"
    template_name = "OAuth2Server" // OAuth2ServerClient or OAuth2Server
    application_id = "TestOAuthServer" // No space
    description = "Test OAuth Server Application"
    
    oauth_profile {
        clientid_type = 0 // anything, list, confidential
        allowed_clients = ["client1", "client2"]
        must_oauth_client = true
        redirects = ["https://example.net", "https://example.com"]
        token_type = "JwtRS256"
        allowed_auth = ["AuthorizationCode", "Implicit"]
        audience = "My:Audience"
        token_lifetime = "5:00:00"
        allow_refresh = true
        refresh_lifetime = "365.00:00:00"
        confirm_authorization = true
        allow_scope_select = true
        scope {
            name = "cli"
            description = "Used for CLI call"
        }
        scope {
            name = "aapm"
            description = "Used for AAPM calls"
        }
    }
    script = "// custom setClaim\nsetClaim('Department', LoginUser.Get('department'));"
    
    sets = [
        data.centrify_manualset.test_set.id
    ]

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Run"]
    }
}