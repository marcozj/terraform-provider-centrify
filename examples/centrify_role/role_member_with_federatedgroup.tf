// data source for Federated Direftory
data "centrify_directoryservice" "federated_dir" {
    // name must be "Federated Directory Service"
    name = "Federated Directory Service"
    // Available types are: "Centrify Directory", "Active Directory", "Federated Directory", "Google Directory", "LDAP Directory"
    type = "Federated Directory"
}

// Create new federated group with globalgroupmappings
resource "centrify_globalgroupmappings" "group_mappings" {
    bulkupdate = true
    mapping = {
        "Idp Group 1" = "Okta Infra Admins"
        "Idp Group 2" = "Azure PAS Users" // Assuming "Azure PAS Users" doesn't exist yet and will be created by this resource
    }
}

// New federated group created by centrify_globalgroupmappings
resource "centrify_federatedgroup" "fedgroup2" {
    name = centrify_globalgroupmappings.group_mappings.mapping["Idp Group 2"] // Reference to "Idp Group 2" map which returns "Azure PAS Users"
}

// Existing federated (virtual) group
data "centrify_federatedgroup" "fedgroup1" {
  name = "Okta Infra Admins"
}

resource "centrify_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    // Existing federated (virtual) group
    member {
        id = data.centrify_federatedgroup.fedgroup1.id
        type = "Group"
    }

    // New federated (virtual) group
    member {
        id = centrify_federatedgroup.fedgroup2.id
        type = "Group"
    }
}