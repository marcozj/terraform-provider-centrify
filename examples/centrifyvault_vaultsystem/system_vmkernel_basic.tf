resource "centrifyvault_vaultsystem" "vmkernel1" {
    # System -> Settings menu related settings
    name = "Demo VMware VMkernel 1"
    fqdn = "192.168.2.4"
    computer_class = "VMwareVMkernel"
    session_type = "Ssh"
    description = "Demo VMware VMkernel 1"

    lifecycle {
      ignore_changes = [
        computer_class,
        session_type,
        ]
    }
}
