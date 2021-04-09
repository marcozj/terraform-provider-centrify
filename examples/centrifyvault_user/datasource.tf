data "centrifyvault_user" "admin" {
    username = "admin@example.com"
}

output "id" {
  value = data.centrifyvault_user.admin.id
}
output "username" {
  value = data.centrifyvault_user.admin.username
}
output "email" {
  value = data.centrifyvault_user.admin.email
}
output "displayname" {
  value = data.centrifyvault_user.admin.displayname
}
output "password_never_expire" {
  value = data.centrifyvault_user.admin.password_never_expire
}
output "force_password_change_next" {
  value = data.centrifyvault_user.admin.force_password_change_next
}
output "oauth_client" {
  value = data.centrifyvault_user.admin.oauth_client
}
output "send_email_invite" {
  value = data.centrifyvault_user.admin.send_email_invite
}
output "description" {
  value = data.centrifyvault_user.admin.description
}
output "office_number" {
  value = data.centrifyvault_user.admin.office_number
}
output "home_number" {
  value = data.centrifyvault_user.admin.home_number
}
output "mobile_number" {
  value = data.centrifyvault_user.admin.mobile_number
}
output "redirect_mfa_user_id" {
  value = data.centrifyvault_user.admin.redirect_mfa_user_id
}
output "manager_username" {
  value = data.centrifyvault_user.admin.manager_username
}