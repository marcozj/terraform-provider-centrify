resource "centrifyvault_vaultsystem" "hpnonstop1" {
    # System -> Settings menu related settings
    name = "Demo HP Nonstop 1"
    fqdn = "192.168.2.2"
    computer_class = "HpNonStopOS"
    session_type = "Ssh"
    description = "Demo HP Nonstop system 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
