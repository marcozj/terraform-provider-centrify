resource "centrifyvault_vaultsystem" "cisconxos1" {
    # System -> Settings menu related settings
    name = "Demo CiscoNXOS 1"
    fqdn = "192.168.2.4"
    computer_class = "CiscoNXOS"
    session_type = "Ssh"
    description = "Demo CiscoNXOS 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
