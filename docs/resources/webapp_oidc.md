---
subcategory: "Applications"
---

# centrifyvault_webapp_oidc (Resource)

This resource allows you to create/update/delete OpenID Connect web application.

## Example Usage

```terraform
resource "centrifyvault_webapp_oidc" "oidcapp" {
    name = "Test OIDC App"
    template_name = "Generic OpenID Connect"
    application_id = "TestOIDCApp" // No space
    description = "Test OIDC Application"

    oauth_profile {
        client_secret = "mysecret"
        application_url = "https://example.com"
        redirects = ["https://example.net", "https://example.com"]
        token_lifetime = "8:00:00"
        allow_refresh = true
        refresh_lifetime = "251.00:00:00"
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_oidc)

## Argument Reference

### Required

- `name` - (String) Name of the ODIC application.
- `template_name` - (String) ODIC application template. Can be set to `Generic OpenID Connect`.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application. **Note:** No white space is allowed.

### Optional

- `description` - (String) Description of the ODIC application.
- `oauth_profile` - (Block List, Max 1) (see [reference for `oauth_profile`](#reference-for-oauth_profile)).
- `script` - (String) Script to generate OpenID Connect Authorization and UserInfo responses for this application.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched). Default is `AlwaysAllowed`.
- `policy_script` - (String) Use script to specify authentication rules (configured rules are ignored). Conflicts with `challenge_rule`.
- `username_strategy` - (String) Account mapping method. Can be set to `ADAttribute`, `Fixed` or `UseScript`. Default is `ADAttribute`.
- `username` - (String) All users share the user name. Applicable if `username_strategy` is `Fixed`.
- `user_map_script` - (String) Account mapping script. Applicable if `username_strategy` is `UseScript`.
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.

## Reference for `oauth_profile`

Required:

- `client_secret` - (String) The OpenID Client Secret for this Identity Provider.
- `application_url` - (String) Resource application URL. The OpenID Connect Service Provider URL.
- `redirects` - (Set of String) List of allowed redirects.
- `token_lifetime` - (String) Token lifetime. It is "d.hh:mm:ss" format. For example, "5:00:00" means 5 hours.

Optional:

- `allow_refresh` - (Boolean) Issue refresh tokens.
- `refresh_lifetime` - (String) Refresh token lifetime. It is "d.hh:mm:ss" format. For example, "365.00:00:00" means 365 days.
- `script` - (String) Script to generate OpenID Connect Authorization and UserInfo responses for this application.

## Import

OpenID Connect Application can be imported using the resource `id`, e.g.

```shell
terraform import centrifyvault_webapp_oidc.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` and `sets` aren't support in import process.
