data "centrify_account" "ad_admin" {
    name = "ad_admin"
    domain_id = data.centrify_domain.example_com.id
}

data "centrify_system" "member1" {
    name = "member1"
    fqdn = "member1.demo.lab"
    computer_class = "Windows"
}

data "centrify_manualset" "test_set" {
    type = "Subscriptions"
    name = "Test Set"
}

resource "centrify_service" "testservice" {
    service_name = "TestWindowsService"
    description = "Test Windows Service in member1"
    system_id = data.centrify_system.member1.id
    service_type = "WindowsService"
    enable_management = true
    admin_account_id = data.centrify_account.ad_admin.id
    multiplexed_account_id = centrify_multiplexedaccount.testmultiplex.id
    restart_service = true
    restart_time_restriction = true
    days_of_week = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
    restart_start_time = "09:00"
    restart_end_time = "10:00"
    use_utc_time = false
    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","Edit","Delete"]
    }
    sets = [
        data.centrify_manualset.test_set.id,
    ]
}