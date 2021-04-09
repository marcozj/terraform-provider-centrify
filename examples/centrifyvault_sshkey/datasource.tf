data "centrifyvault_sshkey" "testkey" {
    name = "testkey"
    key_pair_type = "PrivateKey"
    passphrase = ""
    key_format = "PEM"
    checkout = true
}

output "id" {
  value = data.centrifyvault_sshkey.testkey.id
}
output "name" {
  value = data.centrifyvault_sshkey.testkey.name
}
output "key_pair_type" {
  value = data.centrifyvault_sshkey.testkey.key_pair_type
}
output "key_format" {
  value = data.centrifyvault_sshkey.testkey.key_format
}
output "description" {
  value = data.centrifyvault_sshkey.testkey.description
}
output "key_type" {
  value = data.centrifyvault_sshkey.testkey.key_type
}
output "default_profile_id" {
  value = data.centrifyvault_sshkey.testkey.default_profile_id
}
output "challenge_rule" {
  value = data.centrifyvault_sshkey.testkey.challenge_rule
}
output "testkey_sshkey" {
  value = data.centrifyvault_sshkey.testkey.ssh_key
}