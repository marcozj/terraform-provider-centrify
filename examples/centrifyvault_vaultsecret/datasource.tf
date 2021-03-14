data "centrifyvault_vaultsecret" "test_secret" {
    secret_name = "testsecret"
    checkout = true
}

output "test_secret" {
    value = data.centrifyvault_vaultsecret.test_secret.secret_text
}

// Existing secret folder at top level
data "centrifyvault_vaultsecretfolder" "level1_folder" {
    name = "Level 1 Folder"
}

// Existing secret folder at 2nd level
data "centrifyvault_vaultsecretfolder" "level2_folder" {
    name = "Level 2 Folder"
    parent_path = "Level 1 Folder"
}

// Existing secret folder at 3rd level
data "centrifyvault_vaultsecretfolder" "level3_folder" {
    name = "Level 3 Folder"
    parent_path = "Level 1 Folder\\Level 2 Folder"
}

