resource "centrifyvault_vaultsystem" "checkpoint1" {
    # System -> Settings menu related settings
    name = "Demo Checkpoint 1"
    fqdn = "192.168.2.4"
    computer_class = "CheckPointGaia"
    session_type = "Ssh"
    description = "Demo Checkpoint 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
