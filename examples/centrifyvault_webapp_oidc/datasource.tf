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

data "centrifyvault_role" "approvers" {
  name = "Approvers Role"
}

// Data source to query existing OIDC web app
data "centrifyvault_webapp_oidc" "oidc_webapp" {
  name = "OpenID Connect App"
  application_id = "OpenIDConnectApp"
}

output "oidcapp_id" {
  value = data.centrifyvault_webapp_oidc.oidc_webapp.id
}
output "oidcapp_appid" {
  value = data.centrifyvault_webapp_oidc.oidc_webapp.application_id
}
output "oidcapp_clienid" {
  value = data.centrifyvault_webapp_oidc.oidc_webapp.oauth_profile[0].client_id
}
