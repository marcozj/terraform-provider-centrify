---
subcategory: "Applications"
---

# centrifyvault_webapp_oauth (Resource)

This resource allows you to create/update/delete Oauth web application.

## Example Usage

```terraform
resource "centrifyvault_webapp_oauth" "oauthclient" {
    name = "Test OAuth Client App"
    template_name = "OAuth2ServerClient"
    application_id = "TestOAuthClient"
    description = "Test OAuth Client Application"
    
    oauth_profile {
        clientid_type = 1
        must_oauth_client = true
        redirects = ["https://example.net", "https://example.com"]
        token_type = "JwtRS256"
        allowed_auth = ["ClientCreds", "ResourceCreds"]
        token_lifetime = "8:00:00"
        allow_refresh = true
        refresh_lifetime = "250.00:00:00"
        confirm_authorization = true
        allow_scope_select = true
        scope {
            name = "cli"
            description = "Used for CLI call"
            allowed_rest_apis = ["/SaasManage/GetApplication", "/RedRock/query"]
        }
        scope {
            name = "aapm"
            description = "Used for AAPM calls"
            allowed_rest_apis = [".*"]
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_oauth)

## Argument Reference

### Required

- `name` - (String) Name of the OAuth application.
- `template_name` - (String) OAuth application template. Can be set to `OAuth2ServerClient` or `OAuth2Server`.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application. **Note:** No white space is allowed.

### Optional

- `description` - (String) Description of the OAuth application.
- `oauth_profile` - (Block List, Max 1) (see [reference for `oauth_profile`](#reference-for-oauth_profile)).
- `script` - (String) Script to customize JWT token creation for this application.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Reference for `oauth_profile`

Required:

- `token_type` - (String) Token type. Can be `JwtRS256` or `Opaque`.
- `allowed_auth` - (Set of String) List of allowed authentication methods. List of moethods include `AuthorizationCode`, `Implicit`,`ClientCreds` and `ResourceCreds`.
- `token_lifetime` - (String) Token lifetime. It is "d.hh:mm:ss" format. For example, "5:00:00" means 5 hours.

Optional:

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

## Reference for `scope`

Required:

- `name` - (String) Name of the scope.

Optional:

- `description` - (String) Description of the scope.
- `allowed_rest_apis` - (Set of String) List of allowed REST APIs in Regex format.

## Import

OAuth2 Application can be imported using the resource `id`, e.g.

```shell
terraform import centrifyvault_webapp_oauth.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```
