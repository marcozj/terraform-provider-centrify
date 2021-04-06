---
subcategory: "Applications"
---

# centrifyvault_webapp_generic (Data Source)

This data source gets information of Bookmark, Browser Extension, NTLM and Basic, User-Password web applications.

## Example Usage

```terraform
data "centrifyvault_webapp_generic" "generic_webapp" {
  name = "Generic App"
}

output "generic_webapp_url" {
  value = data.centrifyvault_webapp_generic.generic_webapp.url
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_generic)

## Search Attributes

### Required

- `name` - (String) Name of the Oauth application.

## Attributes Reference

- `id` - id of the SAML application.
- `name` - name property.
- `ur` - url property.
