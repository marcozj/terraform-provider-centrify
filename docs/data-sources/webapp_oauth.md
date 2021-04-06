---
subcategory: "Applications"
---

# centrifyvault_webapp_oauth (Data Source)

This data source gets information of Oauth web application.

## Example Usage

```terraform
data "centrifyvault_webapp_oauth" "oauth_webapp" {
  name = "CentrifyCLI"
  application_id = "CentrifyCLI"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_oauth)

## Search Attributes

### Required

- `name` - (String) Name of the Oauth application.

### Optional

- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.

## Attributes Reference

- `id` - id of the SAML application.
- `name` - name property.
- `application_id` - application_id property.
