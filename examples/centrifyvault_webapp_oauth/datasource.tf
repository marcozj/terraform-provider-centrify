
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

output "id" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.id
}
output "name" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.name
}
output "template_name" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.template_name
}
output "application_id" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.application_id
}
output "description" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.description
}
output "oauth_profile" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.oauth_profile
}
output "oidc_script" {
    value = data.centrifyvault_webapp_oauth.oauth_webapp.oidc_script
}
