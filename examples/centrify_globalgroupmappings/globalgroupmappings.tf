resource "centrify_globalgroupmappings" "group_mappings" {
    bulkupdate = true
    mapping = {
        "Idp Group 1" = "Okta PAS Admin"
        "Idp Group 2" = "Azure PAS Users"
    }
}
