resource "centrifyvault_vaultsystem" "ciscoios1" {
    # System -> Settings menu related settings
    name = "Demo CiscoIOS 1"
    fqdn = "192.168.2.4"
    computer_class = "CiscoIOS"
    session_type = "Ssh"
    description = "Demo CiscoIOS 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
