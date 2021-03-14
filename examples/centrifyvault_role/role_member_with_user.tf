// Existing Centrify Directory user
data "centrifyvault_user" "admin" {
    username = "admin@example.com"
}

// data source for Active Directory domain demo.lab
data "centrifyvault_directoryservice" "demo_lab" {
    // name is the actual Active Directory doman name
    name = "demo.lab"
    type = "Active Directory"
}

// Existing Active Directory user. Connector must be reachable to domain controller
// data source for AD user ad.user@demo.lab
data "centrifyvault_directoryobject" "ad_user" {
    directory_services = [
        data.centrifyvault_directoryservice.demo_lab.id
    ]
    name = "ad.user"
    object_type = "User"
}

// data source for Federated Direftory
data "centrifyvault_directoryservice" "federated_dir" {
    // name must be "Federated Directory Service"
    name = "Federated Directory Service"
    // Avaiable types are: "Centrify Directory", "Active Directory", "Federated Directory", "Google Directory", "LDAP Directory"
    type = "Federated Directory"
}

// data source for federated user
data "centrifyvault_directoryobject" "federated_user" {
    directory_services = [
        data.centrifyvault_directoryservice.federated_dir.id
    ]
    name = "federated.user@democorp.club"
    object_type = "User"
}

resource "centrifyvault_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    // Centrify Directory user
    member {
        id = data.centrifyvault_user.admin.id
        type = "User"
    }

    // Active Directory user
    member {
        id = data.centrifyvault_directoryobject.ad_user.id
        type = "User"
    }

    // Federated user
    member {
        id = data.centrifyvault_directoryobject.federated_user.id
        type = "User"
    }
}