---
subcategory: "Access"
---

# centrifyvault_policy (Data Source)

This data source gets information of policy.

## Example Usage

```terraform
data "centrifyvault_policy" "Default_Policy" {
    name = "Default Policy"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policyorder.tf)

## Search Attributes

### Required

- `name` - (String) The name of the policy.

## Attributes Reference

- `id` - (String) ID of the policy.
- `name` - (String) The name of the policy.
- `link_type` - (String) Policy assignment type.
- `description` - (String) Description of the policy.
- `policy_assignment` - (Set of String) Policy assignment.
- `settings` - (Block List, Max: 1) Various settings of a policy. It can includ below settings.
  - `centrify_services` - (Block List, Max: 1) Settings in **Authentication -> Centrify Services** menu. Refer to [centrify_services](./policy_centrify_services.md) attribute for details.
  - `centrify_client` - (Block List, Max: 1) Settings in **Authentication -> Centrify Clients -> Login** menu. Refer to [centrify_client](./policy_centrify_client.md) attribute for details.
  - `centrify_css_server` - (Block List, Max: 1) Settings in **Authentication -> Centrify Server Suite Agents -> Linux, UNIX and Windows Servers** memu. Refer to [centrify_css_server](./policy_centrify_css_server.md) attribute for details.
  - `centrify_css_workstation` - (Block List, Max: 1) Settings in **Authentication -> Centrify Server Suite Agents -> Windows Workstations** menu. Refer to [centrify_css_workstation](./policy_centrify_css_workstation.md) attribute for details.
  - `centrify_css_elevation` - (Block List, Max: 1) Settings in **Authentication -> Centrify Server Suite Agents -> Privilege Elevation** menu. Refer to [centrify_css_elevation](./policy_centrify_css_elevation.md) attribute for details.
  - `self_service` - (Block List, Max: 1) Settings in **User Security -> Self Service** menu. Refer to [self_service](./policy_self_service.md) attribute for details.
  - `password_settings` - (Block List, Max: 1) Settings in **User Security -> Password Settings** menu. Refer to [password_settings](./policy_password_settings.md) attribute for details.
  - `oath_otp` (Block List, Max: 1) Settings in **User Security -> OATH OTP** menu. Refer to [oath_otp](./policy_oath_otp.md) attribute for details.
  - `radius` - (Block List, Max: 1) Settings in **User Security -> RADIUS** menu. Refer to [radius](./policy_radius.md) attribute for details.
  - `user_account` - (Block List, Max: 1) Settings in **User Security -> User Account** menu. Refer to [user_account](./policy_user_account.md) attribute for details.
  - `system_set` - (Block List, Max: 1) Settings in **Resouces -> Systems** menu. Refer to [system_set](./policy_system_set.md) attribute for details.
  - `database_set` - (Block List, Max: 1) Settings in **Resouces -> Databases** menu. Refer to [database_set](./policy_database_set.md) attribute for details.
  - `domain_set` - (Block List, Max: 1) Settings in **Resouces -> Domains** menu. Refer to [domain_set](./policy_domain_set.md) attribute for details.
  - `account_set` - (Block List, Max: 1) Settings in **Resouces -> Accounts** menu. Refer to [account_set](./policy_account_set.md) attribute for details.
  - `secret_set` - (Block List, Max: 1) Settings in **Resouces -> Secrets** menu. Refer to [secret_set](./policy_secret_set.md) attribute for details.
  - `sshkey_set` - (Block List, Max: 1) Settings in **Resouces -> SSH Keys** menu. Refer to [sshkey_set](./policy_sshkey_set.md) attribute for details.
  - `cloudproviders_set` - (Block List, Max: 1) Settings in **Resouces -> Cloud Providers** menu. Refer to [cloudproviders_set](./policy_cloudproviders_set.md) attribute for details.
  - `mobile_device` - (Block List, Max: 1) Settings in **Devices** menu. Refer to [mobile_device](./policy_mobile_device.md) attribute for details.