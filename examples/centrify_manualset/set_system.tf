resource "centrify_manualset" "my_systems" {
    name = "My Systems"
    type = "Server"
    description = "This Set contains my systems."

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
        rights = ["Grant","View","ManageSession","Edit","Delete","AgentAuth","OfflineRescue","AddAccount","UnlockAccount","ManagementAssignment","RequestZoneRole"]
    }
}