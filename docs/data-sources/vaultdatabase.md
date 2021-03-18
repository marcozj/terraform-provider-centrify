---
subcategory: "Resources"
---

# centrifyvault_vaultdatabase (Data Source)

This data source gets information of database.

## Example Usage

```terraform
data "centrifyvault_vaultdatabase" "sql-centrifysuite" {
    name = "SQL-CENTRIFYSUITE"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_vaultdatabase)

## Search Attributes

### Required

- `name` - (String) Name of the Database.

### Optional

- `database_class` - (String) Type of the Database. Can be set to `SQLServer`, `Oracle` or `SAPAse`.
- `hostname` - (String) Hostname or IP address of the Database.

## Attributes Reference

- `id` - id of the database.
- `name` - name property.
- `hostname` - hostname property.
- `database_class` - database_class property.
- `port` - port property.
- `instance_name` - instance_name property.
- `service_name` - service_name property.
