
data "centrify_manualset" "test_set" {
  type = "VaultDatabase"
  name = "Test Set"
}

resource "centrify_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Collection"
    policy_assignment = [
        format("VaultDatabase|%s", data.centrify_manualset.test_set.id),
    ]
    
    settings {
        database_set {
            // Account Security
            checkout_lifetime = 60
            // Security Settings
            allow_multiple_checkouts = true
            enable_password_rotation = true
            password_rotate_interval = 90
            enable_password_rotation_after_checkin = true
            minimum_password_age = 70
            // Maintenance Settings
            enable_password_history_cleanup = true
            password_historycleanup_duration = 120
        }
    }
}