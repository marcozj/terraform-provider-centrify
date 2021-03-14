
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        radius {
            allow_radius = true
            require_challenges = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            send_vendor_attributes = true
            allow_external_radius = true
        }
    }
}