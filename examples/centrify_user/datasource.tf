data "centrify_user" "admin" {
    username = "admin@centrify.com.207"
}

output "id" {
  value = data.centrify_user.admin.id
}
output "username" {
  value = data.centrify_user.admin.username
}
output "email" {
  value = data.centrify_user.admin.email
}
output "displayname" {
  value = data.centrify_user.admin.displayname
}
output "password_never_expire" {
  value = data.centrify_user.admin.password_never_expire
}
output "force_password_change_next" {
  value = data.centrify_user.admin.force_password_change_next
}
output "oauth_client" {
  value = data.centrify_user.admin.oauth_client
}
output "send_email_invite" {
  value = data.centrify_user.admin.send_email_invite
}
output "description" {
  value = data.centrify_user.admin.description
}
output "office_number" {
  value = data.centrify_user.admin.office_number
}
output "home_number" {
  value = data.centrify_user.admin.home_number
}
output "mobile_number" {
  value = data.centrify_user.admin.mobile_number
}
output "redirect_mfa_user_id" {
  value = data.centrify_user.admin.redirect_mfa_user_id
}
output "manager_username" {
  value = data.centrify_user.admin.manager_username
}