/*
resource "centrifyvault_vaultsystem" "testsystem1" {
    # System -> Settings menu related settings
    name = "Windows 01"
    fqdn = "192.168.2.3"
    computer_class = "Windows"
    session_type = "Rdp"
    description = "Windows system 1"
    //port = 22
    //use_my_account = false
    //management_mode = "Smb"
    //system_timezone = "UTC-08"
    //proxyuser = "admin@example.com"
    //proxyuser_password = "xxxxxxxxxx"
    //proxyuser_managed = false // When this is set to true, Centrify Vault tries to change the password immediately and may result in error if password change fails

    # System -> Policy menu related settings
    //checkout_lifetime = 60
    //allow_remote_access = false
    //allow_rdp_clipboard = true
    default_profile_id = data.centrifyvault_authenticationprofile.step_up_auth_pf.id

    # System -> Advanced menu related settings
    //local_account_automatic_maintenance = true
    //local_account_manual_unlock = true
    //domain_id = data.centrifyvault_vaultdomain.demo_lab.id
    //allow_multiple_checkouts = true
    //enable_password_rotation = true
    //password_rotate_interval = 60
    //enable_password_rotation_after_checkin = true
    //minimum_password_age = 90
    //password_profile_id = data.centrifyvault_passwordprofile.test_pw_pf1.id
    //enable_password_history_cleanup = true
    //password_historycleanup_duration = 100

    # System -> Zone Role Workflow menu related settings
	//use_domainadmin_for_zonerole_workflow = true
	//enable_zonerole_workflow = true

	# System -> Connectors menu related settings
    
	//connector_list = [
    //    data.centrifyvault_connector.XXXXXXXXXX.id,
    //    data.centrifyvault_connector.dc01.id
    //]

    sets = [
    //    centrifyvault_manualset.all_systems.id,
    //    data.centrifyvault_manualset.lab_systems.id
    ]

    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.step_up_auth_pf.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }
    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.step_up_auth_pf.id
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


    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}


resource "centrifyvault_vaultsystem" "testsystem2" {
    name = "Linux 01"
    fqdn = "192.168.2.4"
    computer_class = "Unix"
    session_type = "Ssh"
    description = "Linux system 1"
    port = 22
    use_my_account = true

    permission {
        principal_id = data.centrifyvault_role.sso_role.id
        principal_name = data.centrifyvault_role.sso_role.name
        principal_type = "Role"
        rights = [
          "Grant",
          "View",
          "RequestZoneRole"
        ]
    }

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}


resource "centrifyvault_vaultsystem" "testsystem3" {
    name = "Cisco 01"
    fqdn = "192.168.2.5"
    computer_class = "CiscoIOS"
    session_type = "Ssh"
    description = "Cisco system 1"

    permission {
        principal_id = data.centrifyvault_role.sso_role.id
        principal_name = data.centrifyvault_role.sso_role.name
        principal_type = "Role"
        rights = [
          "Grant",
          "View",
          "Edit",
          "ManageSession",
        ]
    }

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
*/