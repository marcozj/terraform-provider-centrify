---
subcategory: "Resources"
---

# centrifyvault_connector (Data Source)

This data source gets information of Centrify Connector.

## Example Usage

```terraform
data "centrifyvault_connector" "connector1" {
    name = "connector_host1" // Connector name registered in Centrify
}
```

## Search Attributes

### Required

- `name` - (String) Name of the Connector.

## Attributes Reference

- `id` - (String) The ID of this resource.
