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
  name = "LAB Infrastructure Owners"
}

// Data source to query existing generic web app
data "centrifyvault_webapp_generic" "generic_webapp" {
  name = "Generic App"
}

output "id" {
    value = data.centrifyvault_webapp_generic.generic_webapp.id
}
output "name" {
    value = data.centrifyvault_webapp_generic.generic_webapp.name
}
output "description" {
    value = data.centrifyvault_webapp_generic.generic_webapp.description
}
output "url" {
    value = data.centrifyvault_webapp_generic.generic_webapp.url
}
output "hostname_suffix" {
    value = data.centrifyvault_webapp_generic.generic_webapp.hostname_suffix
}
output "username_field" {
    value = data.centrifyvault_webapp_generic.generic_webapp.username_field
}
output "password_field" {
    value = data.centrifyvault_webapp_generic.generic_webapp.password_field
}
output "submit_field" {
    value = data.centrifyvault_webapp_generic.generic_webapp.submit_field
}
output "form_field" {
    value = data.centrifyvault_webapp_generic.generic_webapp.form_field
}
output "additional_login_field" {
    value = data.centrifyvault_webapp_generic.generic_webapp.additional_login_field
}
output "additional_login_field_value" {
    value = data.centrifyvault_webapp_generic.generic_webapp.additional_login_field_value
}
output "selector_timeout" {
    value = data.centrifyvault_webapp_generic.generic_webapp.selector_timeout
}
output "order" {
    value = data.centrifyvault_webapp_generic.generic_webapp.order
}
output "script" {
    value = data.centrifyvault_webapp_generic.generic_webapp.script
}
output "default_profile_id" {
    value = data.centrifyvault_webapp_generic.generic_webapp.default_profile_id
}
output "username_strategy" {
    value = data.centrifyvault_webapp_generic.generic_webapp.username_strategy
}
output "username" {
    value = data.centrifyvault_webapp_generic.generic_webapp.username
}
output "use_ad_login_pw" {
    value = data.centrifyvault_webapp_generic.generic_webapp.use_ad_login_pw
}
output "use_ad_login_pw_by_script" {
    value = data.centrifyvault_webapp_generic.generic_webapp.use_ad_login_pw_by_script
}
output "user_map_script" {
    value = data.centrifyvault_webapp_generic.generic_webapp.user_map_script
}
output "workflow_enabled" {
    value = data.centrifyvault_webapp_generic.generic_webapp.workflow_enabled
}
output "workflow_approver" {
    value = data.centrifyvault_webapp_generic.generic_webapp.workflow_approver
}
output "challenge_rule" {
    value = data.centrifyvault_webapp_generic.generic_webapp.challenge_rule
}
output "policy_script" {
    value = data.centrifyvault_webapp_generic.generic_webapp.policy_script
}
