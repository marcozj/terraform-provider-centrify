data "centrify_role" "system_admin" {
    name = "System Administrator"
}

resource "centrify_user" "testuser" {
    username = "testuser@example.com"
    email = "testuser@example.com"
    displayname = "Test User"
    description = "Test user"
    password = "xxxxxxxxxxxx"
    confirm_password = "xxxxxxxxxxxx"
    password_never_expire = true
    roles = [
        data.centrify_role.system_admin.id
    ]

    lifecycle {
      ignore_changes = [
        password,
        confirm_password,
        ]
    }
}