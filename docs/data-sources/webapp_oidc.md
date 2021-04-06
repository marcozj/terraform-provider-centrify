---
subcategory: "Applications"
---

# centrifyvault_webapp_oidc (Data Source)

This data source gets information of OpenID Connect web application.

## Example Usage

```terraform
data "centrifyvault_webapp_oidc" "oidc_webapp" {
  name = "OpenID App"
  application_id = "OpenID pp"
}

output "oidcapp_clienid" {
  value = data.centrifyvault_webapp_oidc.oidc_webapp.oauth_profile[0].client_id
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_oidc)

## Search Attributes

### Required

- `name` - (String) Name of the Oauth application.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.

## Attributes Reference

- `id` - id of the SAML application.
- `name` - name property.
- `application_id` - application_id property.
- `description` - description property.
- `oauth_profile` - oauth_profile property.
  - `client_secret` - client_secret property.
  - `application_url` - application_url property.
  - `issuer` - issuer property.
  - `client_id` - client_id property.
  - `redirects` - redirects property.
  - `token_lifetime` - token_lifetime property.
  - `allow_refresh` - allow_refresh property.
  - `refresh_lifetime` - refresh_lifetime property.
  - `script` - script property.
