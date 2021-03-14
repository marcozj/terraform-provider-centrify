data "centrifyvault_sshkey" "testkey" {
    name = "testkey"
    key_pair_type = "PrivateKey"
    passphrase = ""
    key_format = "PEM"
    checkout = true
}

output "testkey_sshkey" {
  value = data.centrifyvault_sshkey.testkey.ssh_key
}