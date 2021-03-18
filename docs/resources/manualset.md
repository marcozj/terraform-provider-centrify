---
subcategory: "Resources"
---

# centrifyvault_manualset (Resource)

This resource allows you to create/update/delete manual Set.

## Example Usage

```terraform
resource "centrifyvault_manualset" "all_systems" {
    name = "All Systems"
    type = "Server"
    description = "This Set contains all systems."

    permission {
        principal_id = centrifyvault_role.lab_infra_admin.id
        principal_name = centrifyvault_role.lab_infra_admin.name
        principal_type = "Role"
        rights = ["Grant","View"]
    }

    member_permission {
        principal_id = centrifyvault_role.lab_infra_admin.id
        principal_name = centrifyvault_role.lab_infra_admin.name
        principal_type = "Role"
        rights = ["Grant","ManageSession"]
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_manualset)

## Argument Reference

### Required

- `name` - (String) The name of the manual set.
- `type` - (String) Type of set. Can be set to `Server`, `VaultAccount`, `VaultDatabase`, `VaultDomain`, `DataVault`, `SshKeys`, `Subscriptions`, `Application`, or `ResourceProfiles`.

### Optional

- `description` - (String) Description of an manual set.
- `subtype` - (String) SubObjectType for application. Can be set to `Web` or `Desktop`. Only applicable if type is `Application`.
- `permission` - (Block Set) Set permissions. Refer to [permission attribute](./attribute_permission.md) for details.
- `member_permission` - (Block Set) Set member permissions. Refer to [member_permission attribute](./attribute_permission.md) for details. Each type of Set has different member_permission values and you can find them in [examples](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_manualset).
