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

- `id` - (String) ID of the SAML application.
- `name` - (String) Name of the generic web application.
- `template_name` - (String) Web application template.
- `ur` - url property.
- `url` - (String) The URL of the application.
- `description` - (String) Description of the web application.
- `hostname_suffix` - (String) The host name suffix for the url of the login form, for example, acme.com.
- `username_field` - (String) The CSS Selector for the user name field in the login form, for example, input#login-username.
- `password_field` - (String) The CSS Selector for the password field in the login form, for example, input#login-password.
- `submit_field` - (String) The CSS Selector for the Submit button in the login form, for example, input#login-button. This entry is optional. It is required only if you cannot submit the form by pressing the enter key.
- `form_field` - (String) The CSS Selector for the form field of the login form, for example, form#loginForm.
- `additional_login_field` - (String) The CSS Selector for any Additional Login Field required to login besides username and password, such as Company name or Agency ID. For example, the selector could be input#login-company-id. This entry is required only if there is an additional login field besides username and password.
- `additional_login_field_value` - (String) The value for the Additional Login Field. For example, if there is an additional login field for the company name, enter the company name here. This entry is required if Additional Login Field is set.
- `selector_timeout` - (Int) Use this field to indicate the number of milliseconds to wait for the expected input selectors to load before timing out on failure. A zero or negative number means no timeout.
- `order` - (String) Use this field to specify the order of login if it is not username, password and submit.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default Profile (used if no conditions matched).
- `policy_script` - (String) Use script to specify authentication rules (configured rules are ignored). Conflicts with `challenge_rule`.
- `username_strategy` - (String) Account mapping method. Can be set to `ADAttribute`, `Fixed`, `SetByUser` or `UseScript`. Default is `ADAttribute`.
- `username` - (String) All users share the user name. Applicable if `username_strategy` is `ADAttribute` or `Fixed`. Default is `userprincipalname`.
- `password` - (String Sensitive) Password for all user share one name. Applicable if `username_strategy` is `Fixed`.
- `use_ad_login_pw` - (Boolean) Use the login password supplied by the user (Active Directory users only). Applicable if `username_strategy` is `ADAttribute`.
- `use_ad_login_pw_by_script` - (Boolan) Use the login password supplied by the user for account mapping script (Active Directory users only). Applicable if `username_strategy` is `UseScript`.
- `user_map_script` - (String) Account mapping script. Applicable if `username_strategy` is `UseScript`.
- `script` - (String) Script to log the user in to this application.
- `workflow_enabled` - (Boolean) Enable workflow for this application.
- `workflow_approver` - (Block List) List of approvers. Refer to [workflow_approver](./attribute_workflow_approver.md) attribute for details.
