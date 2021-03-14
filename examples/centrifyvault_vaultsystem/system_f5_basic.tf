resource "centrifyvault_vaultsystem" "f5bigip1" {
    # System -> Settings menu related settings
    name = "Demo F5 1"
    fqdn = "192.168.2.4"
    computer_class = "F5NetworksBIGIP"
    session_type = "Ssh"
    description = "Demo F5 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
