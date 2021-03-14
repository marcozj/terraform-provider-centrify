
resource "centrifyvault_cloudprovider" "demo_aws" {
    type = "Aws"
    name = "Demo AWS"
    cloud_account_id = "123456789000"
    description = "Demo AWS"
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    enable_interactive_password_rotation = true
    prompt_change_root_password = true
    enable_password_rotation_reminders = true
    password_rotation_reminder_duration = 20

    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
      rule {
        filter = "DayOfWeek"
        condition = "OpIsDayOfWeek"
        value = "L,1,3,4,5"
      }
      rule {
        filter = "Browser"
		condition = "OpNotEqual"
        value = "Firefox"
      }
      rule {
        filter = "CountryCode"
		condition = "OpNotEqual"
        value = "GA"
      }
    }

    sets = [
        data.centrifyvault_manualset.test_set.id
    ]

    permission {
        principal_id   = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete"]
    }

}
