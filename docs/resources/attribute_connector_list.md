---
subcategory: "Common Attribute"
---

# connector_list attribute

connector_list is a common attribute in various resources.

## Example Usage

```terraform
data "centrifyvault_connector" "connector1" {
    name = "connector_host1" // Connector name registered in Centrify
}

data "centrifyvault_connector" "connector2" {
    name = "connector_host2" Connector name registered in Centrify
}

resource "centrifyvault_vaultsystem" "linuxsystem" {
    name = "Test Linux"
    fqdn = "192.168.2.1"
    computer_class = "Unix"
    session_type = "Ssh"
    description = "Linux system"
    port = 22

    connector_list = [
        data.centrifyvault_connector.connector1.id,
        data.centrifyvault_connector.connector2.id
    ]
}
```

## Argument Reference

- `connector_list` (Set of String) List of Connector IDs.
