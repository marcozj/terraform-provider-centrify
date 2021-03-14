data "centrifyvault_passwordprofile" "windows_pw_pf" {
    name = "Windows Profile"
}

# TODO: administrative_account_id setting
resource "centrifyvault_vaultsystem" "windows2" {
    # System -> Settings menu related settings
    name = "Demo Windows 2"
    fqdn = "192.168.2.3"
    computer_class = "Windows"
    session_type = "Rdp"
    description = "Demo Windows system 2"

    proxyuser = "admin@example.com"
    proxyuser_password = "xxxxxxxxxx"
    proxyuser_managed = false // When this is set to true, Centrify Vault tries to change the password immediately and may result in error if password change fails
    management_mode = "WinRMOverHttp"
    management_port = 5985
    system_timezone = "UTC"

    # System -> Policy menu related settings
    checkout_lifetime = 60
    allow_remote_access = true
    allow_rdp_clipboard = true
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    privilege_elevation_default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id

    # System -> Advanced menu related settings
    local_account_automatic_maintenance = true
    local_account_manual_unlock = true
    domain_id = data.centrifyvault_vaultdomain.example_com.id
    remove_user_on_session_end = true
    allow_multiple_checkouts = true
    enable_password_rotation = true
    password_rotate_interval = 60
    enable_password_rotation_after_checkin = true
    minimum_password_age = 90
    password_profile_id = data.centrifyvault_passwordprofile.windows_pw_pf.id
    enable_password_history_cleanup = true
    password_historycleanup_duration = 100

    # System -> Zone Role Workflow menu related settings
	  use_domainadmin_for_zonerole_workflow = true
	  enable_zonerole_workflow = true

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

    privilege_elevation_rule {
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

	# System -> Connectors menu related settings
	connector_list = [
        data.centrifyvault_connector.connector1.id
    ]

    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","ManageSession","Edit","Delete","AgentAuth","OfflineRescue","AddAccount","UnlockAccount","ManagementAssignment","RequestZoneRole"]
    }

    sets = [
        data.centrifyvault_manualset.test_set.id
    ]

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
