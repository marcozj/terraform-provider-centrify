
data "centrifyvault_manualset" "test_set" {
  type = "VaultDomain"
  name = "Test Set"
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("VaultDomain|%s", data.centrifyvault_manualset.test_set.id),
    ]
    
    settings {
        domain_set {
            // Domain Security
            checkout_lifetime = 60
            // Security Settings
            allow_multiple_checkouts = true
            enable_password_rotation = true
            password_rotate_interval = 91
            enable_password_rotation_after_checkin = true
            minimum_password_age = 70
            // Maintenance Settings
            enable_password_history_cleanup = true
            password_historycleanup_duration = 120
        }
    }
}