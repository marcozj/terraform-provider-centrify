resource "centrifyvault_vaultsystem" "genericssh1" {
    # System -> Settings menu related settings
    name = "Demo Generic SSH 1"
    fqdn = "192.168.2.4"
    computer_class = "GenericSsh"
    session_type = "Ssh"
    description = "Demo Generic SSH 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
