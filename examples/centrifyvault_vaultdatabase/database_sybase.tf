data "centrifyvault_passwordprofile" "ase_pw_pf" {
  name = "SAP ASE Profile"
}

resource "centrifyvault_vaultdatabase" "sybasedb" {
  # Database -> Settings menu related settings
  name           = "My Sybase DB"
  hostname       = "sapase.example.com"
  database_class = "SAPAse"
  description    = "SAP Sybase Database"
  port           = 3638

  # Database -> Policy menu related settings
  checkout_lifetime = 60

  # Database -> Advanced menu related settings
  allow_multiple_checkouts               = true
  enable_password_rotation               = true
  password_rotate_interval               = 60
  enable_password_rotation_after_checkin = true
  minimum_password_age                   = 90
  password_profile_id                    = data.centrifyvault_passwordprofile.ase_pw_pf.id
  enable_password_history_cleanup        = true
  password_historycleanup_duration       = 100

  # System -> Connectors menu related settings
  connector_list = [
    data.centrifyvault_connector.connector1.id
  ]

  sets = [
    data.centrifyvault_manualset.test_set.id
  ]

  permission {
    principal_id   = data.centrifyvault_role.system_admin.id
    principal_name = data.centrifyvault_role.system_admin.name
    principal_type = "Role"
    rights = ["Grant","View","Edit","Delete"]
  }

  lifecycle {
    ignore_changes = [
      database_class
    ]
  }
}
