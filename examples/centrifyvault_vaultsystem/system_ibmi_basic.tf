resource "centrifyvault_vaultsystem" "ibmi1" {
    # System -> Settings menu related settings
    name = "Demo IBMi 1"
    fqdn = "192.168.2.4"
    computer_class = "IBMi"
    session_type = "Ssh"
    description = "Demo IBMi 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
