data "centrify_user" "admin" {
    username = "admin@example.com"
}

// data source for Active Directory domain demo.lab
data "centrify_directoryservice" "demo_lab" {
    // name is the actual Active Directory doman name
    name = "demo.lab"
    type = "Active Directory"
}

// Existing Active Directory user. Connector must be reachable to domain controller
// data source for AD user ad.user@demo.lab
data "centrify_directoryobject" "ad_user" {
    directory_services = [
        data.centrify_directoryservice.demo_lab.id
    ]
    name = "alex.foster@demo.lab"
    object_type = "User"
}

// data source for Federated Direftory
data "centrify_directoryservice" "federated_dir" {
    // name must be "Federated Directory Service"
    name = "Federated Directory Service"
    // Available types are: "Centrify Directory", "Active Directory", "Federated Directory", "Google Directory", "LDAP Directory"
    type = "Federated Directory"
}

// data source for federated user
data "centrify_directoryobject" "federated_user" {
    directory_services = [
        data.centrify_directoryservice.federated_dir.id
    ]
    name = "alex.foster@democorp.club"
    object_type = "User"
}

// Existing federated (virtual) group
data "centrify_federatedgroup" "fedgroup1" {
  name = "Okta Infra Admins"
}

// Existing role whose membership to be managed
data "centrify_role" "testrole" {
    name = "Test Role"
}

// Role membership for exsting role
resource "centrify_role_membership" "testrolemembers" {
  role_id = data.centrify_role.testrole.id

  // Centrify Directory user
  member {
    id = data.centrify_user.admin.id
    type = "User"
  }

  // Active Directory user
  member {
    id = data.centrify_directoryobject.ad_user.id
    type = "User"
  }

  // Federated user
  member {
    id = data.centrify_directoryobject.federated_user.id
    type = "User"
  }

  // Existing federated (virtual) group
  member {
    id = data.centrify_federatedgroup.fedgroup1.id
    type = "Group"
  }
}
