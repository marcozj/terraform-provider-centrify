resource "centrifyvault_globalgroupmappings" "group_mappings" {
    mapping {
        attribute_value = "Idp Group 1"
        group_name = "Okta PAS Admin"
    }
    mapping {
        attribute_value = "Idp Group 2"
        group_name = "Azure PAS Users"
    }
}
