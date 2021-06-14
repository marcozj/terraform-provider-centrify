resource "centrify_domain" "example_lab" {
    name = "example.lab"
    description = "example.lab domain"
    verify = false
    // Policy menu
    checkout_lifetime = 90
    // Advanced -> Security Settings
    allow_multiple_checkouts = true
    enable_password_rotation = true
    password_rotate_interval = 90
    enable_password_rotation_after_checkin = true
    minimum_password_age = 120
    password_profile_id = data.centrify_passwordprofile.domain_pw_pf.id
    // Advanced -> Maintenance Settings
    enable_password_history_cleanup = true
    password_historycleanup_duration = 100
    // Advanced -> Domain/Zone Tasks
    enable_zone_joined_check = true
    zone_joined_check_interval = 90
    enable_zonerole_cleanup = true
    zonerole_cleanup_interval = 6

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","UnlockAccount","AddAccount"]
    }

    sets = [
        data.centrify_manualset.test_set.id,
    ]

    connector_list = [
        data.centrify_connector.connector1.id
    ]
}
