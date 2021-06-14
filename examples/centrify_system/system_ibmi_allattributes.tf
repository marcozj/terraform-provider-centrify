data "centrify_passwordprofile" "ibmi_pw_pf" {
    name = "IBM i Profile"
}

resource "centrify_system" "ibmi2" {
    # System -> Settings menu related settings
    name = "Demo IBMi 2"
    fqdn = "192.168.2.4"
    computer_class = "IBMi"
    session_type = "Ssh"
    description = "Demo IBMi 2"
    port = 22
    system_timezone = "UTC"

    # System -> Policy menu related settings
    checkout_lifetime = 60
    allow_remote_access = true
    default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id

    # System -> Advanced menu related settings
    remove_user_on_session_end = true
    allow_multiple_checkouts = true
    enable_password_rotation = true
    password_rotate_interval = 60
    enable_password_rotation_after_checkin = true
    minimum_password_age = 90
    password_profile_id = data.centrify_passwordprofile.ibmi_pw_pf.id
    enable_password_history_cleanup = true
    password_historycleanup_duration = 100

    enable_sshkey_rotation = true
    sshkey_rotate_interval = 60
    minimum_sshkey_age = 90
    sshkey_algorithm = "RSA_2048"
    enable_sshkey_history_cleanup = true
    sshkey_historycleanup_duration = 120

    challenge_rule {
      authentication_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
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
        data.centrify_connector.connector1.id
    ]

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","ManageSession","Edit","Delete","OfflineRescue","AddAccount","UnlockAccount"]
    }

    sets = [
        data.centrify_manualset.test_set.id
    ]

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
