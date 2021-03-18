---
subcategory: "Resources"
---

# centrifyvault_cloudprovider (Data Source)

This data source gets information of authentication cloud provider.

## Example Usage

```terraform
data "centrifyvault_cloudprovider" "my_aws" {
    name = "My AWS"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_cloudprovider)

## Search Attributes

### Required

- `name` - (String) Name of the cloud provider.

### Optional

- `cloud_account_id` - (String) Account ID of the cloud provider.

## Attributes Reference

- `id` - (String) The ID of this resource.
- `name` - (String) Name property.
- `cloud_account_id` - (String) Account ID proerty.
