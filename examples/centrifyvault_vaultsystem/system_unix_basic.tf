resource "centrifyvault_vaultsystem" "unix1" {
    # System -> Settings menu related settings
    name = "Demo Unix 1"
    fqdn = "192.168.2.2"
    computer_class = "Unix"
    session_type = "Ssh"
    description = "Demo Unix system 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
