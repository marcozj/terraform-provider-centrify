
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        mobile_device {
            allow_enrollment = true
            permit_non_compliant_device = true
            enable_invite_enrollment = true
            allow_notify_multi_devices = true
            enable_debug = true
            location_tracking = true
            force_fingerprint = true
            allow_fallback_pin = true
            require_passcode = true
            auto_lock_timeout = 15
            lock_app_on_exit = true
        }
    }
}