// data source for Active Directory domain demo.lab
data "centrifyvault_directoryservice" "demo_lab" {
    // name is the actual Active Directory doman name
    name = "demo.lab"
    type = "Active Directory"
}

// Existing Active Directory group. Connector must be reachable to domain controller
// data source for AD group ad.group@demo.lab
data "centrifyvault_directoryobject" "ad_group" {
    directory_services = [
        data.centrifyvault_directoryservice.demo_lab.id
    ]
    name = "ad.group"
    object_type = "Group"
}

resource "centrifyvault_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    // Active Directory group
    member {
        id = data.centrifyvault_directoryobject.ad_group.id
        type = "Group"
    }
}