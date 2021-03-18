---
subcategory: "Resources"
---

# centrifyvault_manualset (Data Source)

This data source gets information of manual Set. Only works for user created Set but not built-in Set.

## Example Usage

```terraform
data "centrifyvault_manualset" "lab_systems" {
    type = "Server"
    name = "LAB Systems"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_manualset)

## Search Attributes

### Required

- `name` - (String) The name of the manual Set.
- `type` - (String) Type of set. Can be set to `Server`, `VaultAccount`, `VaultDatabase`, `VaultDomain`, `DataVault`, `SshKeys`, `Subscriptions`, `Application`, or `ResourceProfiles`.

### Optional

- `subtype` - (String) SubObjectType for application. Can be set to `Web` or `Desktop`. Only applicable if type is `Application`.

## Attributes Reference

- `id` - id of the manual set.
- `name` - name property.
- `description` - (String) description property.
