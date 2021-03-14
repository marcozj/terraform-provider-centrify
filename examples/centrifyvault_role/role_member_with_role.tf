// Existing role
data "centrifyvault_role" "system_admin" {
    name = "System Administrator"
}

resource "centrifyvault_role" "nested_role" {
    name = "Nested role"
    description = "Nested role."
}

resource "centrifyvault_role" "test_role" {
    name = "Test role"
    description = "Test Role that has role as member."

    member {
        id = centrifyvault_role.nested_role.id
        type = "Role"
    }

    member {
        id = data.centrifyvault_role.system_admin.id
        type = "Role"
    }
}