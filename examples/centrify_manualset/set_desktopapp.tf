resource "centrify_manualset" "my_desktopapps" {
    name = "My Desktop Apps"
    type = "Application"
    subtype = "Desktop"
    description = "This Set contains my desktop applications."

    permission {
        principal_id = data.centrify_user.admin.id
        principal_name = data.centrify_user.admin.id
        principal_type = "User"
        rights = ["Grant","View","Edit","Delete"]
    }

    permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View"]
    }

    member_permission {
        principal_id = data.centrify_role.system_admin.id
        principal_name = data.centrify_role.system_admin.name
        principal_type = "Role"
        rights = ["Grant","View","Run"]
    }

    lifecycle {
      ignore_changes = [
        type,
        subtype,
      ]
    }
}