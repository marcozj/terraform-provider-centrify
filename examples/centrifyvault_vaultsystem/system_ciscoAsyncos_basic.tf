resource "centrifyvault_vaultsystem" "ciscoasyncos1" {
    # System -> Settings menu related settings
    name = "Demo CiscoAsyncOS 1"
    fqdn = "192.168.2.4"
    computer_class = "CiscoAsyncOS"
    session_type = "Ssh"
    description = "Demo CiscoAsyncOS 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
