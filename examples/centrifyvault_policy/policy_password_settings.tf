
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        password_settings {
            // Password Requirements
            min_length = 12
            max_length = 24
            require_digit = true
            require_mix_case = true
            require_symbol = true
            // Display Requirements
            show_password_complexity = true
            complexity_hint = "Whatever ......."
            // Additional Requirements
            no_of_repeated_char_allowed = 2
            check_weak_password = true
            allow_include_username = true
            allow_include_displayname = true
            require_unicode = true
            // Password Age
            min_age_in_days = 10
            max_age_in_days = 90
            password_history = 10
            expire_soft_notification = 35
            expire_hard_notification = 72
            expire_notification_mobile = true
            // Capture Settings
            bad_attempt_threshold = 5
            capture_window = 20
            lockout_duration = 30
        }
    }
}