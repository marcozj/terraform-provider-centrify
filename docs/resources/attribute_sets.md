---
subcategory: "Common Attribute"
---

# Sets attribute

Sets is a common attribute in various resources.

## Example Usage

```terraform
resource "centrifyvault_manualset" "my_systems" {
    name = "My Systems"
    type = "Server"
    description = "This Set contains my systems."
}

data "centrifyvault_manualset" "lab_systems" {
    type = "Server"
    name = "LAB Systems"
}

resource "centrifyvault_vaultsystem" "linuxsystem" {
    name = "Test Linux"
    fqdn = "192.168.2.1"
    computer_class = "Unix"
    session_type = "Ssh"
    description = "Linux system"
    port = 22

    sets = [
        centrifyvault_manualset.my_systems.id,
        data.centrifyvault_manualset.lab_systems.id
    ]
}
```

## Argument Reference

- `sets` (Set of String) List of Set IDs the resource belongs to.
