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

output "generic_webapp_url" {
  value = data.centrifyvault_webapp_generic.generic_webapp.url
}
