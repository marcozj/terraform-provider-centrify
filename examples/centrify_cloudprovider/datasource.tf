data "centrify_cloudprovider" "my_aws" {
    name = "My AWS"
    cloud_account_id = "1234567890"
}

output "id" {
    value = data.centrify_cloudprovider.my_aws.id
}
output "cloud_account_id" {
    value = data.centrify_cloudprovider.my_aws.cloud_account_id
}
output "description" {
    value = data.centrify_cloudprovider.my_aws.description
}
output "type" {
    value = data.centrify_cloudprovider.my_aws.type
}
output "enable_interactive_password_rotation" {
    value = data.centrify_cloudprovider.my_aws.enable_interactive_password_rotation
}
output "prompt_change_root_password" {
    value = data.centrify_cloudprovider.my_aws.prompt_change_root_password
}
output "enable_password_rotation_reminders" {
    value = data.centrify_cloudprovider.my_aws.enable_password_rotation_reminders
}
output "password_rotation_reminder_duration" {
    value = data.centrify_cloudprovider.my_aws.password_rotation_reminder_duration
}
output "default_profile_id" {
    value = data.centrify_cloudprovider.my_aws.default_profile_id
}
output "challenge_rule" {
    value = data.centrify_cloudprovider.my_aws.challenge_rule
}


data "centrify_authenticationprofile" "newdevice_auth_pf" {
    name = "Default New Device Login Profile"
}

data "centrify_manualset" "test_set" {
  type = "CloudProviders"
  name = "Test Set"
}

data "centrify_role" "system_admin" {
  name = "System Administrator"
}

data "centrify_manualset" "test_accounts" {
  type = "VaultAccount"
  name = "Test Accounts"
}