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

- `name` - (String) Name of the OIDC application.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.

## Attributes Reference

- `id` - (String) ID of the OIDC application.
- `name` - (String) Name of the OIDC application.
- `template_name` - (String) ODIC application template.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.
- `description` - (String) Description of the ODIC application.
- `oauth_profile` - (Block List, Max 1) (see [reference for `oauth_profile`](#reference-for-oauth_profile)).
- `oidc_script` - (String) Script to generate OpenID Connect Authorization and UserInfo responses for this application.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched).
- `policy_script` - (String) Use script to specify authentication rules (configured rules are ignored).
- `username_strategy` - (String) Account mapping method.
- `username` - (String) All users share the user name. Applicable if `username_strategy` is `Fixed`.
- `user_map_script` - (String) Account mapping script. Applicable if `username_strategy` is `UseScript`.
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.

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
