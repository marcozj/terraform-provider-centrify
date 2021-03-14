resource "centrifyvault_role" "test_role" {
    name = "Test role"
    description = "Test Role with all attributes."
    adminrights = [
        "Admin Portal Login",
        "Application Management",
        "Computer Login and Privilege Elevation",
        "Device Management",
        "Federation Management",
        "MFA Unlock",
        "Privilege Elevation Management",
        "Privileged Access Service Administrator",
        "Privileged Access Service Power User",
        "Privileged Access Service User",
        "Radius Management",
        "Read Only Resource Management",
        "Read Only System Administration",
        "Register and Administer Connectors",
        "Report Management",
        "Role Management",
        "System Enrollment",
        "User Management",
    ]
}