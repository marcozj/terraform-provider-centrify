data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrify_manualset" "test_set" {
  type = "Application"
  name = "Test Web Apps"
  subtype = "Web"
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

data "centrify_role" "approvers" {
  name = "Approvers Role"
}

// Data source to query existing OIDC web app
data "centrify_webapp_oidc" "oidc_webapp" {
  name = "OpenID Connect App"
  application_id = "OpenIDConnectApp"
}

output "id" {
    value = data.centrify_webapp_oidc.oidc_webapp.id
}
output "name" {
    value = data.centrify_webapp_oidc.oidc_webapp.name
}
output "application_id" {
    value = data.centrify_webapp_oidc.oidc_webapp.application_id
}
output "template_name" {
    value = data.centrify_webapp_oidc.oidc_webapp.template_name
}
output "oauth_profile" {
    value = data.centrify_webapp_oidc.oidc_webapp.oauth_profile
}
output "oidcapp_clienid" {
  value = data.centrify_webapp_oidc.oidc_webapp.oauth_profile[0].client_id
}
output "oidc_script" {
    value = data.centrify_webapp_oidc.oidc_webapp.oidc_script
}
output "default_profile_id" {
    value = data.centrify_webapp_oidc.oidc_webapp.default_profile_id
}
output "username_strategy" {
    value = data.centrify_webapp_oidc.oidc_webapp.username_strategy
}
output "username" {
    value = data.centrify_webapp_oidc.oidc_webapp.username
}
output "user_map_script" {
    value = data.centrify_webapp_oidc.oidc_webapp.user_map_script
}
output "workflow_enabled" {
    value = data.centrify_webapp_oidc.oidc_webapp.workflow_enabled
}
output "workflow_approver" {
    value = data.centrify_webapp_oidc.oidc_webapp.workflow_approver
}
output "challenge_rule" {
    value = data.centrify_webapp_oidc.oidc_webapp.challenge_rule
}
output "policy_script" {
    value = data.centrify_webapp_oidc.oidc_webapp.policy_script
}

