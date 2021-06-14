data "centrify_sshkey" "testkey" {
    name = "testkey"
    key_pair_type = "PrivateKey"
    passphrase = ""
    key_format = "PEM"
    checkout = true
}

output "id" {
  value = data.centrify_sshkey.testkey.id
}
output "name" {
  value = data.centrify_sshkey.testkey.name
}
output "key_pair_type" {
  value = data.centrify_sshkey.testkey.key_pair_type
}
output "key_format" {
  value = data.centrify_sshkey.testkey.key_format
}
output "description" {
  value = data.centrify_sshkey.testkey.description
}
output "key_type" {
  value = data.centrify_sshkey.testkey.key_type
}
output "default_profile_id" {
  value = data.centrify_sshkey.testkey.default_profile_id
}
output "challenge_rule" {
  value = data.centrify_sshkey.testkey.challenge_rule
}
output "testkey_sshkey" {
  value = data.centrify_sshkey.testkey.ssh_key
}
