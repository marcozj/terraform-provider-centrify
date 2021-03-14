resource "centrifyvault_user" "testuser" {
    username = "testuser@example.com"
    email = "testuser@example.com"
    displayname = "Test User"
    description = "Test user"
    password = "xxxxxxxxxxxx"
    confirm_password = "xxxxxxxxxxxx"
    password_never_expire = true
    force_password_change_next = true
    oauth_client = false
    send_email_invite = true
    office_number = "+00 00000000"
    home_number = "+00 00000000"
    mobile_number = "+00 00000000"
    redirect_mfa_user_id = data.centrifyvault_user.admin.id
    manager_username = "admin@example.com"
    
}