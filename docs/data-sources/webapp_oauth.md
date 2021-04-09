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

- `id` - (String) ID of the OAuth application.
- `template_name` - (String) OAuth application template.
- `name` - (String) Name of the OAuth application.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.
- `description` - (String) Description of the OAuth application.
- `oauth_profile` - (Block List, Max 1) (see [reference for `oauth_profile`](#reference-for-oauth_profile)).
- `oidc_script` - (String) Script to customize JWT token creation for this application.

## Reference for `oauth_profile`

- `token_type` - (String) Token type.
- `allowed_auth` - (Set of String) List of allowed authentication methods.
- `token_lifetime` - (String) Token lifetime. It is "d.hh:mm:ss" format. For example, "5:00:00" means 5 hours.
- `clientid_type` - (String) ClientID type. Can be set to `0` or `1`. `0` for Anything or List, `1` for Confidential.
- `issuer` - (String) OAuth server issuer. Applicalbe to OAuth Server only.
- `audience` - (String) OAuth server audience. Applicalbe to OAuth Server only.
- `allowed_clients` - (Set of String) List of allowed clients.
- `must_oauth_client` - (Boolean) Must be OAuth Client. Appliable if `clientid_type` is `1`.
- `redirects` - (Set of String) List of allowed redirects.
- `allow_refresh` - (Boolean) Issue refresh tokens.
- `refresh_lifetime` - (String) Refresh token lifetime. It is "d.hh:mm:ss" format. For example, "365.00:00:00" means 365 days.
- `confirm_authorization` - (Boolean) User must confirm authorization request.
- `allow_scope_select` - (Boolan) Allow scope selection.
- `scope` - (Block Set) List of Oauth scope. (see [reference for `scope`](#reference-for-scope)).
- `client_id` - (String) OIDC client id.

## Reference for `scope`

Required:

- `name` - (String) Name of the scope.
- `description` - (String) Description of the scope.
- `allowed_rest_apis` - (Set of String) List of allowed REST APIs in Regex format.
