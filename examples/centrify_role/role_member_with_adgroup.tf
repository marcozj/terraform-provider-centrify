// data source for Active Directory domain demo.lab
data "centrify_directoryservice" "demo_lab" {
    // name is the actual Active Directory doman name
    name = "demo.lab"
    type = "Active Directory"
}

// Existing Active Directory group. Connector must be reachable to domain controller
// data source for AD group ad.group@demo.lab
data "centrify_directoryobject" "ad_group" {
    directory_services = [
        data.centrify_directoryservice.demo_lab.id
    ]
    name = "ad.group"
    object_type = "Group"
}

resource "centrify_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    // Active Directory group
    member {
        id = data.centrify_directoryobject.ad_group.id
        type = "Group"
    }
}