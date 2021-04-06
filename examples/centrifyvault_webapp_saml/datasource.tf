data "centrifyvault_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrifyvault_manualset" "test_set" {
  type = "Application"
  name = "Test Web Apps"
  subtype = "Web"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

// Data source to query existing SAML web app
data "centrifyvault_webapp_saml" "saml_webapp" {
  name = "My SAML App"
}

output "corp_identifier" {
  value = data.centrifyvault_webapp_saml.saml_webapp.corp_identifier
}

output "idp_metadata_url" {
  value = data.centrifyvault_webapp_saml.saml_webapp.idp_metadata_url
}
