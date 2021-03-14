// data source for Federated Direftory
data "centrifyvault_directoryservice" "federated_dir" {
    // name must be "Federated Directory Service"
    name = "Federated Directory Service"
    // Avaiable types are: "Centrify Directory", "Active Directory", "Federated Directory", "Google Directory", "LDAP Directory"
    type = "Federated Directory"
}

// data source for existing federated group
data "centrifyvault_directoryobject" "federated_group" {
    directory_services = [
        data.centrifyvault_directoryservice.federated_dir.id
    ]
    name = "Azure PAS Users"
    object_type = "Group"
}

// Create new federated group with globalgroupmappings
resource "centrifyvault_globalgroupmappings" "group_mapping" {
  mapping {
    attribute_value = "IdP Group"
    group_name      = "IdP PAS Admin" // Assuming "IdP PAS Admin" doesn't exist yet
  }
}

// data source for newly created federated group
data "centrifyvault_directoryobject" "idp_pas_admin" {
    directory_services = [
        data.centrifyvault_directoryservice.federated_dir.id
    ]
    name = "IdP PAS Admin"
    object_type = "Group"

    // Need to wait for "IdP PAS Admin" created first
    depends_on = [
        centrifyvault_globalgroupmappings.group_mapping,
    ]
}

resource "centrifyvault_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    // Existing federated (virtual) group
    member {
        id = data.centrifyvault_directoryobject.federated_group.id
        type = "Group"
    }

    // New federated (virtual) group
    member {
        id = data.centrifyvault_directoryobject.idp_pas_admin.id
        type = "Group"
    }
}