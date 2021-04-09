data "centrifyvault_vaultsystem" "centos1" {
    name = "centos1"
    fqdn = "centos1.demo.lab"
    computer_class = "Unix"
}

// Checkout an account password and do not checkin immediately
data "centrifyvault_vaultaccount" "myaccount" {
    name = "testuser"
    host_id = data.centrifyvault_vaultsystem.centos1.id
    checkout = true
    checkin = false
}

// Checkout credential
output "password" {
    value = data.centrifyvault_vaultaccount.myaccount.password
}

output "id" {
    value = data.centrifyvault_vaultaccount.myaccount.id
}
output "name" {
    value = data.centrifyvault_vaultaccount.myaccount.name
}
output "host_id" {
    value = data.centrifyvault_vaultaccount.myaccount.host_id
}
output "domain_id" {
    value = data.centrifyvault_vaultaccount.myaccount.domain_id
}
output "database_id" {
    value = data.centrifyvault_vaultaccount.myaccount.database_id
}
output "cloudprovider_id" {
    value = data.centrifyvault_vaultaccount.myaccount.cloudprovider_id
}
output "access_key_id" {
    value = data.centrifyvault_vaultaccount.myaccount.access_key_id
}
output "credential_type" {
    value = data.centrifyvault_vaultaccount.myaccount.credential_type
}
output "credential_name" {
    value = data.centrifyvault_vaultaccount.myaccount.credential_name
}
output "key_pair_type" {
    value = data.centrifyvault_vaultaccount.myaccount.key_pair_type
}
output "sshkey_id" {
    value = data.centrifyvault_vaultaccount.myaccount.sshkey_id
}
output "is_admin_account" {
    value = data.centrifyvault_vaultaccount.myaccount.is_admin_account
}
output "is_root_account" {
    value = data.centrifyvault_vaultaccount.myaccount.is_root_account
}
output "use_proxy_account" {
    value = data.centrifyvault_vaultaccount.myaccount.use_proxy_account
}
output "managed" {
    value = data.centrifyvault_vaultaccount.myaccount.managed
}
output "description" {
    value = data.centrifyvault_vaultaccount.myaccount.description
}
output "status" {
    value = data.centrifyvault_vaultaccount.myaccount.status
}
output "checkout_lifetime" {
    value = data.centrifyvault_vaultaccount.myaccount.checkout_lifetime
}
output "default_profile_id" {
    value = data.centrifyvault_vaultaccount.myaccount.default_profile_id
}
output "access_secret_checkout_default_profile_id" {
    value = data.centrifyvault_vaultaccount.myaccount.access_secret_checkout_default_profile_id
}
output "access_secret_checkout_rule" {
    value = data.centrifyvault_vaultaccount.myaccount.access_secret_checkout_rule
}
output "workflow_enabled" {
    value = data.centrifyvault_vaultaccount.myaccount.workflow_enabled
}
output "workflow_approver" {
    value = data.centrifyvault_vaultaccount.myaccount.workflow_approver
}
output "challenge_rule" {
    value = data.centrifyvault_vaultaccount.myaccount.challenge_rule
}
output "access_key" {
    value = data.centrifyvault_vaultaccount.myaccount.access_key
}

