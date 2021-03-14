
data "centrifyvault_vaultaccount" "test_svc1" {
    name = "test_svc1"
    domain_id = data.centrifyvault_vaultdomain.example_com.id
}

data "centrifyvault_vaultaccount" "test_svc2" {
    name = "test_svc2"
    domain_id = data.centrifyvault_vaultdomain.example_com.id
}

resource "centrifyvault_multiplexedaccount" "testmultiplex" {
    name = "Account for TestWindowsService"
    description = "Multiplexed account for TestWindowsService"
    accounts = [
        data.centrifyvault_vaultaccount.test_svc1.id,
        data.centrifyvault_vaultaccount.test_svc2.id,
    ]
    permission {
        principal_id = data.centrifyvault_role.system_admin.id
        principal_name = data.centrifyvault_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","Edit","Delete"]
    }
}