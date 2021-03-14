resource "centrifyvault_vaultsystem" "windows1" {
    name = "Demo Windows 1"
    fqdn = "192.168.2.3"
    computer_class = "Windows"
    session_type = "Rdp"
    description = "Demo Windows system 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
