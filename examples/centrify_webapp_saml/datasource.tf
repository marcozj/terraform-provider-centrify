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

// Data source to query existing SAML web app
data "centrify_webapp_saml" "saml_webapp" {
  name = "My SAML App"
}

output "id" {
    value = data.centrify_webapp_saml.saml_webapp.id
}
output "name" {
    value = data.centrify_webapp_saml.saml_webapp.name
}
output "template_name" {
    value = data.centrify_webapp_saml.saml_webapp.template_name
}
output "description" {
    value = data.centrify_webapp_saml.saml_webapp.description
}
output "corp_identifier" {
    value = data.centrify_webapp_saml.saml_webapp.corp_identifier
}
output "app_entity_id" {
    value = data.centrify_webapp_saml.saml_webapp.app_entity_id
}
output "application_id" {
    value = data.centrify_webapp_saml.saml_webapp.application_id
}
output "idp_metadata_url" {
    value = data.centrify_webapp_saml.saml_webapp.idp_metadata_url
}
output "sp_metadata_url" {
    value = data.centrify_webapp_saml.saml_webapp.sp_metadata_url
}
output "sp_metadata_xml" {
    value = data.centrify_webapp_saml.saml_webapp.sp_metadata_xml
}
output "sp_entity_id" {
    value = data.centrify_webapp_saml.saml_webapp.sp_entity_id
}
output "acs_url" {
    value = data.centrify_webapp_saml.saml_webapp.acs_url
}
output "recipient_sameas_acs_url" {
    value = data.centrify_webapp_saml.saml_webapp.recipient_sameas_acs_url
}
output "recipient" {
    value = data.centrify_webapp_saml.saml_webapp.recipient
}
output "sign_assertion" {
    value = data.centrify_webapp_saml.saml_webapp.sign_assertion
}
output "name_id_format" {
    value = data.centrify_webapp_saml.saml_webapp.name_id_format
}
output "sp_single_logout_url" {
    value = data.centrify_webapp_saml.saml_webapp.sp_single_logout_url
}
output "relay_state" {
    value = data.centrify_webapp_saml.saml_webapp.relay_state
}
output "authn_context_class" {
    value = data.centrify_webapp_saml.saml_webapp.authn_context_class
}
output "saml_attribute" {
    value = data.centrify_webapp_saml.saml_webapp.saml_attribute
}
output "saml_script" {
    value = data.centrify_webapp_saml.saml_webapp.saml_script
}
output "default_profile_id" {
    value = data.centrify_webapp_saml.saml_webapp.default_profile_id
}
output "username_strategy" {
    value = data.centrify_webapp_saml.saml_webapp.username_strategy
}
output "username" {
    value = data.centrify_webapp_saml.saml_webapp.username
}
output "user_map_script" {
    value = data.centrify_webapp_saml.saml_webapp.user_map_script
}
output "workflow_enabled" {
    value = data.centrify_webapp_saml.saml_webapp.workflow_enabled
}
output "workflow_approver" {
    value = data.centrify_webapp_saml.saml_webapp.workflow_approver
}
output "challenge_rule" {
    value = data.centrify_webapp_saml.saml_webapp.challenge_rule
}
output "policy_script" {
    value = data.centrify_webapp_saml.saml_webapp.policy_script
}