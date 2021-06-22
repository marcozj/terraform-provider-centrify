data "centrify_cloudprovider" "my_aws" {
    name = "My AWS"
    cloud_account_id = "1234567890"
}

// Checkout an AWS IAM account access key
data "centrify_account" "myaccount" {
    name = "testiam1"
    cloudprovider_id = data.centrify_cloudprovider.my_aws.id
    access_key_id = "XXXXXXXXXXXXXX"
    checkout = true
}

// Checkout credential
output "secret_access_key" {
  value = data.centrify_account.myaccount.secret_access_key
}

output "id" {
    value = data.centrify_account.myaccount.id
}
output "name" {
    value = data.centrify_account.myaccount.name
}
output "host_id" {
    value = data.centrify_account.myaccount.host_id
}
output "domain_id" {
    value = data.centrify_account.myaccount.domain_id
}
output "database_id" {
    value = data.centrify_account.myaccount.database_id
}
output "cloudprovider_id" {
    value = data.centrify_account.myaccount.cloudprovider_id
}
output "access_key_id" {
    value = data.centrify_account.myaccount.access_key_id
}
output "credential_type" {
    value = data.centrify_account.myaccount.credential_type
}
output "credential_name" {
    value = data.centrify_account.myaccount.credential_name
}
output "key_pair_type" {
    value = data.centrify_account.myaccount.key_pair_type
}
output "sshkey_id" {
    value = data.centrify_account.myaccount.sshkey_id
}
output "is_admin_account" {
    value = data.centrify_account.myaccount.is_admin_account
}
output "is_root_account" {
    value = data.centrify_account.myaccount.is_root_account
}
output "use_proxy_account" {
    value = data.centrify_account.myaccount.use_proxy_account
}
output "managed" {
    value = data.centrify_account.myaccount.managed
}
output "description" {
    value = data.centrify_account.myaccount.description
}
output "status" {
    value = data.centrify_account.myaccount.status
}
output "checkout_lifetime" {
    value = data.centrify_account.myaccount.checkout_lifetime
}
output "default_profile_id" {
    value = data.centrify_account.myaccount.default_profile_id
}
output "access_secret_checkout_default_profile_id" {
    value = data.centrify_account.myaccount.access_secret_checkout_default_profile_id
}
output "access_secret_checkout_rule" {
    value = data.centrify_account.myaccount.access_secret_checkout_rule
}
output "workflow_enabled" {
    value = data.centrify_account.myaccount.workflow_enabled
}
output "workflow_approver" {
    value = data.centrify_account.myaccount.workflow_approver
}
output "challenge_rule" {
    value = data.centrify_account.myaccount.challenge_rule
}
output "access_key" {
    value = data.centrify_account.myaccount.access_key
}
