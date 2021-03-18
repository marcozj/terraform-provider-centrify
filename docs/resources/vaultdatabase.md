---
subcategory: "Resources"
---

# centrifyvault_vaultdatabase (Resource)

This resource allows you to create/update/delete database.

## Example Usage

```terraform
resource "centrifyvault_vaultdatabase" "mssql" {
  # Database -> Settings menu related settings
  name           = "My MS SQL"
  hostname       = "mssql.example.com"
  database_class = "SQLServer"
  instance_name  = "MYINSTANCE"
  description    = "MS SQL Database"
  port           = 1433
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdatabase)

## Argument Reference

### Required

- `name` - (String) Name of the Database.
- `database_class` - (String) Type of the Database. Can be set to `SQLServer`, `Oracle` or `SAPAse`.
- `hostname` - (String) Hostname or IP address of the Database.

### Optional

- `skip_reachability_test` - (Boolean) Verify Database Settings
- `port` - (Number) Port number that used to connect to the Database.
- `instance_name` - (String) Instance name of MS SQL Database. Required if `database_class` is `SQLServer`.
- `service_name` - (String) Service name of Oracle database. Required if `database_class` is `Oracle`.
- `description` - (String) Description of the Database.
- `checkout_lifetime` - (Number) Specifies the number of minutes that a checked out password is valid. Range between `15` to `2147483647`.
- `allow_multiple_checkouts` - (Boolean) Allow multiple password checkouts for this database. Specifies whether multiple users can have the same database account password checked out at the same time.
- `enable_password_rotation` - (Boolean) Enable periodic password rotation. Specifies whether managed password should be rotated periodically.
- `password_rotate_interval` - (Number) Password rotation interval (days). Rotates managed passwords automatically at the interval you specify. Range between `1` to `2147483647`.
- `enable_password_rotation_after_checkin` - (Boolean) Enable password rotation after checkin. Specifies whether managed password should be rotated after it's checked in.
- `minimum_password_age` - (Number) Minimum Password Age (days). Minimum amount of days old a password must be before it is rotated. Range between `0` to `2147483647`.
- `password_profile_id` - (String) Password complexity profile id.
- `enable_password_history_cleanup` - (Boolean) Enable periodic password history cleanup. Specifies whether retired passwords should be deleted periodically.
- `password_historycleanup_duration` - (Number) Password history cleanup (days). Deletes retired passwords automatically that were last modified either equal to or greater than the number of days specified here. Range between `90` to `2147483647`.
- `connector_list` (Set of String) List of Connector IDs. Refer to [connector_list](./attribute_connector_list.md) attribute for details.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.
