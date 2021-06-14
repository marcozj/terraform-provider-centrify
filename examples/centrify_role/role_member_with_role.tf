// Existing role
data "centrify_role" "system_admin" {
    name = "System Administrator"
}

resource "centrify_role" "nested_role" {
    name = "Nested role"
    description = "Nested role."
}

resource "centrify_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    member {
        id = centrify_role.nested_role.id
        type = "Role"
    }

    member {
        id = data.centrify_role.system_admin.id
        type = "Role"
    }
}