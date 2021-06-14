
resource "centrify_webapp_saml" "saml_webapp" {
    name = "Test SAML Web App"
    template_name = "Generic SAML"
    description = "Test SAML Web Application"
    sp_config_method = 1
    sp_metadata_url = "https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml"
    //sp_metadata_xml = file("sp_meta.xml")
    /*
    sp_entity_id = "urn:federation:MicrosoftOnline"
    acs_url = "https://login.microsoftonline.com/login.srf"
    recipient_sameas_acs_url = true
    //recipient = "https://login.microsoftonline.com/login.srf"
    sign_assertion = true
    name_id_format = "unspecified"
    //sp_single_logout_url = "https://login.microsoftonline.com/logout.srf"
    //relay_state = "state"
    authn_context_class = "unspecified"
    */
    saml_attribute {
        name = "attribute1"
        value = "value1"
    }
    saml_attribute {
        name = "attribute2"
        value = "value2"
    }
    saml_response_script = "test;"
    username_strategy = "ADAttribute"
    username = "userprincipalname"

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
      guid = data.centrify_role.system_admin.id
      name = data.centrify_role.system_admin.name
      type = "Role"
    }
    
}

output "idp_meta_url" {
    value = centrify_webapp_saml.saml_webapp.idp_metadata_url
}

resource "centrify_webapp_saml" "awsconsole" {
    name = "Test AWS Console"
    template_name = "AWSConsoleSAML" // "Generic SAML", "AWSConsoleSAML", "ClouderaSAML", "CloudLock SAML", "ConfluenceServerSAML", "Dome9Saml", "GitHubEnterpriseSAML", "JIRACloudSAML", "JIRAServerSAML", "PaloAltoNetworksSAML", "SplunkOnPremSAML", "SumoLogicSAML"
    corp_identifier = "232545454522"
    description = "Test AWS Console"
    sp_config_method = 1

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
    
}
