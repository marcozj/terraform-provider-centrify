---
subcategory: "Resources"
---

# centrifyvault_vaultsystem (Data Source)

This data source gets information of system.

## Example Usage

```terraform
data "centrifyvault_vaultsystem" "demo_system" {
    name = "demosystem"
    fqdn = "demosystem.example.com"
    computer_class = "Unix"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultsystem)

## Search Attributes

### Required

- `name` - (String) The name of the system.
- `fqdn` - (String) Hostname or IP address of the system.

### Optional

- `computer_class` - (String) Type of the system. Can be set to `Windows`, `Unix`, `CiscoIOS`, `CiscoNXOS`, `JuniperJunos`, `HpNonStopOS`, `IBMi`, `CheckPointGaia`, `PaloAltoNetworksPANOS`, `F5NetworksBIGIP`, `CiscoAsyncOS`, `VMwareVMkernel`, `GenericSsh` or `CustomSsh`.

## Attributes Reference

- `id` - id of the system.
- `name` - name property.
- `fqdn` - fqdn property.
- `computer_class` - computer_class property.
