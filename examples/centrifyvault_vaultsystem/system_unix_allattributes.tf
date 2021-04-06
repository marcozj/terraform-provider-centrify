data "centrifyvault_passwordprofile" "unix_pw_pf" {
    name = "Unix Profile"
}
# TODO: administrative_account_id and local_account_automatic_maintenance setting
resource "centrifyvault_vaultsystem" "unix2" {
    # System -> Settings menu related settings
    name = "Demo Unix 2"
    fqdn = "192.168.2.2"
    computer_class = "Unix"
    session_type = "Ssh"
    description = "Demo Unix system 2"
    port = 22

    use_my_account = true
    proxyuser = "admin"
    proxyuser_password = "xxxxxxxxxx"
    proxyuser_managed = false // When this is set to true, Centrify Vault tries to change the password immediately and may result in error if password change fails
    system_timezone = "UTC"

    # System -> Policy menu related settings
    checkout_lifetime = 60
    allow_remote_access = true
    default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
    privilege_elevation_default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id

    # System -> Advanced menu related settings
    //local_account_automatic_maintenance = true
    domain_id = data.centrifyvault_vaultdomain.example_com.id
    allow_multiple_checkouts = true
    enable_password_rotation = true
    password_rotate_interval = 60
    enable_password_rotation_after_checkin = true
    minimum_password_age = 90
    password_profile_id = data.centrifyvault_passwordprofile.unix_pw_pf.id
    enable_password_history_cleanup = true
    password_historycleanup_duration = 100

    enable_sshkey_rotation = true
    sshkey_rotate_interval = 60
    minimum_sshkey_age = 90
    sshkey_algorithm = "RSA_2048"
    enable_sshkey_history_cleanup = true
    sshkey_historycleanup_duration = 120
    
    # System -> Workflow menu
    agent_auth_workflow_enabled = true
    agent_auth_workflow_approver {
        guid = data.centrifyvault_role.system_admin.id
        name = data.centrifyvault_role.system_admin.name
        type = "Role"
    }
    privilege_elevation_workflow_enabled = true
    privilege_elevation_workflow_approver {
        guid = data.centrifyvault_role.system_admin.id
        name = data.centrifyvault_role.system_admin.name
        type = "Role"
    }

    # System -> Zone Role Workflow menu related settings
	  // domain_id must be set
	  use_domainadmin_for_zonerole_workflow = true
	  enable_zonerole_workflow = true
     // Assing override zone roles
    use_domain_assignment_for_zoneroles = false
    assigned_zonerole {
      name = "UNIX Login/Global" // name is in format of "<zone role name>/<zone name>"
    }
    // Assign override zone role approver
    use_domain_assignment_for_zonerole_approvers = false
    assigned_zonerole_approver {
        guid = data.centrifyvault_role.system_admin.id
        name = data.centrifyvault_role.system_admin.name
        type = "Role"
    }

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
