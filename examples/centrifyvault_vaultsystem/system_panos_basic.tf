resource "centrifyvault_vaultsystem" "panos1" {
    # System -> Settings menu related settings
    name = "Demo Palo Alto PAN-OS 1"
    fqdn = "192.168.2.4"
    computer_class = "PaloAltoNetworksPANOS"
    session_type = "Ssh"
    description = "Demo Palo Alto PAN-OS 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
