
data "centrifyvault_manualset" "test_set" {
  type = "Application"
  name = "Test Web Apps"
  subtype = "Web"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

// Data source to query existing Oauth web app
data "centrifyvault_webapp_oauth" "oauth_webapp" {
  name = "CentrifyCLI"
  application_id = "CentrifyCLI"
}

output "oauthapp_id" {
  value = data.centrifyvault_webapp_oauth.oauth_webapp.id
}
