
data "centrify_account" "test_svc1" {
    name = "test_svc1"
    domain_id = data.centrify_domain.example_com.id
}

data "centrify_account" "test_svc2" {
    name = "test_svc2"
    domain_id = data.centrify_domain.example_com.id
}

resource "centrify_multiplexedaccount" "testmultiplex" {
    name = "Account for TestWindowsService"
    description = "Multiplexed account for TestWindowsService"
    accounts = [
        data.centrify_account.test_svc1.id,
        data.centrify_account.test_svc2.id,
    ]
    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","Edit","Delete"]
    }
}