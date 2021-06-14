data "centrify_secret" "test_secret" {
    secret_name = "testsecret"
    checkout = true
}

// Checkout secret
output "secret_text" {
    value = data.centrify_secret.test_secret.secret_text
}

output "id" {
    value = data.centrify_secret.test_secret.id
}
output "secret_name" {
    value = data.centrify_secret.test_secret.secret_name
}
output "folder_id" {
    value = data.centrify_secret.test_secret.folder_id
}
output "description" {
    value = data.centrify_secret.test_secret.description
}
output "parent_path" {
    value = data.centrify_secret.test_secret.parent_path
}
output "type" {
    value = data.centrify_secret.test_secret.type
}
output "default_profile_id" {
    value = data.centrify_secret.test_secret.default_profile_id
}
output "workflow_enabled" {
    value = data.centrify_secret.test_secret.workflow_enabled
}
output "workflow_approver" {
    value = data.centrify_secret.test_secret.workflow_approver
}
output "challenge_rule" {
    value = data.centrify_secret.test_secret.challenge_rule
}

// Existing secret folder at top level
data "centrify_secretfolder" "level1_folder" {
    name = "Level 1 Folder"
}

// Existing secret folder at 2nd level
data "centrify_secretfolder" "level2_folder" {
    name = "Level 2 Folder"
    parent_path = "Level 1 Folder"
}

// Existing secret folder at 3rd level
data "centrify_secretfolder" "level3_folder" {
    name = "Level 3 Folder"
    parent_path = "Level 1 Folder\\Level 2 Folder"
}

output "id" {
    value = data.centrify_secretfolder.level1_folder.id
}
output "name" {
    value = data.centrify_secretfolder.level1_folder.name
}
output "description" {
    value = data.centrify_secretfolder.level1_folder.description
}
output "parent_path" {
    value = data.centrify_secretfolder.level1_folder.parent_path
}
output "parent_id" {
    value = data.centrify_secretfolder.level1_folder.parent_id
}
output "default_profile_id" {
    value = data.centrify_secretfolder.level1_folder.default_profile_id
}
output "challenge_rule" {
    value = data.centrify_secretfolder.level1_folder.challenge_rule
}