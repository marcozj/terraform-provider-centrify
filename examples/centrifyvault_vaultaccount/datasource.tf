data "centrifyvault_vaultsystem" "centos1" {
    name = "centos1"
    fqdn = "centos1.demo.lab"
    computer_class = "Unix"
}

data "centrifyvault_cloudprovider" "my_aws" {
    name = "My AWS"
}

// Checkout an account password and do not checkin immediately
data "centrifyvault_vaultaccount" "centos1_local_account" {
    name = "clocal_account"
    host_id = data.centrifyvault_vaultsystem.centos1.id
    checkout = true
    checkin = false
}

output "local_account_password" {
  value = data.centrifyvault_vaultaccount.centos1_local_account.password
}

// Checkout SSH key
data "centrifyvault_vaultaccount" "sshuser" {
    name = "sshuser"
    host_id = data.centrifyvault_vaultsystem.centos1.id
    key_pair_type = "PrivateKey"
    //passphrase = ""
    checkout = true
}

output "sshuser_sshkey" {
  value = data.centrifyvault_vaultaccount.sshuser.private_key
}

// Retrieve IAM account secret key
data "centrifyvault_vaultaccount" "iamuser" {
    name = "terraformuser"
    cloudprovider_id = data.centrifyvault_cloudprovider.my_aws.id
    access_key_id = "XXXXXXXXXXX"
    checkout = true
}

output "iamuser_secretkey" {
  value = data.centrifyvault_vaultaccount.iamuser.secret_access_key
}
