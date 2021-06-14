
resource "centrify_secretfolder" "level1_folder" {
    name = "Level 1 Folder"
    description = "Level 1 Folder"
    // Note: When parent folder is enforced with challenge authentication, deletion of subfolder and secret in it will fail
    //default_profile_id = data.centrify_authenticationprofile.newdevice_auth_pf.id
    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","Add"]
    }

    member_permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Edit","Delete","RetrieveSecret"]
    }
}

resource "centrify_secretfolder" "level2_folder" {
    name = "Level 2 Folder"
    description = "Level 2 Folder"
    parent_id = centrify_secretfolder.level1_folder.id
}

resource "centrify_secretfolder" "level3_folder" {
    name = "Level 3 Folder"
    description = "Level 3 Folder"
    parent_id = centrify_secretfolder.level2_folder.id
}
