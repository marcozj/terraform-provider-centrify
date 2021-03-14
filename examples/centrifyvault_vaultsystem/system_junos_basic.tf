resource "centrifyvault_vaultsystem" "junos1" {
    # System -> Settings menu related settings
    name = "Demo Juniper Junos 1"
    fqdn = "192.168.2.2"
    computer_class = "JuniperJunos"
    session_type = "Ssh"
    description = "Demo Juniper Junos system 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
