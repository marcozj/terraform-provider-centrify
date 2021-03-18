---
subcategory: "Access"
---

# centrifyvault_policy & centrifyvault_policyorder (Resource)

These resources allows you to create/update/delete policy.
When creates a policy using `centrifyvault_policy`, it must be added to `centrifyvault_policyorder` together with existing policies and place it at desired order.

## Example Usage

This example creates a policy named "Test Policy" and place it before existing "Default Policy".

```terraform
data "centrifyvault_policy" "Default_Policy" {
    name = "Default Policy"
}

resource "centrifyvault_policyorder" "policy_order" {
    policy_order = [
        centrifyvault_policy.test_policy.id,
        data.centrifyvault_policy.Default_Policy.id,
    ]
}

resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        centrify_services {
            authentication_enabled = true
            default_profile_id = data.centrifyvault_authenticationprofile.newdevice_auth_pf.id
            session_lifespan = 23
            allow_session_persist = true
            default_session_persist = true
            persist_session_lifespan = 30
        }
        
        oath_otp {
            allow_otp = true
        }
    }
    
}
```

More examples for `centrifyvault_policyorder` can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policyorder.tf)
More examples for `centrifyvault_policy` can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/)

## Argument Reference for centrifyvault_policyorder

### Required (centrifyvault_policyorder)

- `policy_order` - (List of String) List of policy IDs.

## Argument Reference for centrifyvault_policy

### Required (centrifyvault_policy)

- `name` - (String) The name of the policy.
- `link_type` - (String) Policy assignment type. Can be set to `Global`, `Role`, `Collection` or `Inactive`.

### Optional (centrifyvault_policy)

- `description` - (String) Description of the policy.
- `policy_assignment` - (Set of String) Policy assignment. List of role Is or set IDs assigned to the policy. For role, it is simply list of IDs. For set, it follows following format.
  
  ```terraform
    policy_assignment = [
        "Server|@All Systems", // Built-in Set
        "Server|<system set id>", // Custom Set
    ]
  ```

  ```terraform
    policy_assignment = [
        "VaultDatabase|@SQL Server", // Built-in Set
        "VaultDatabase|<database set id>", // Custom Set
    ]
  ```

  ```terraform
    policy_assignment = [
        "VaultDomain|@All Domains", // Built-in Set
        "VaultDomain|<domain set id>", // Custom Set
    }
  ```

  ```terraform
    policy_assignment = [
        "VaultAccount|@Database Accounts", // Built-in Set
        "VaultAccount|<account set id>", // Custom Set
    ]
  ```

  ```terraform
    policy_assignment = [
        "DataVault|@Text Generic Secrets", // Built-in Set
        "DataVault|<secret set id>", // Custom Set
    ]
  ```

  ```terraform
    policy_assignment = [
        "SshKeys|@Managed SshKeys", // Built-in Set
        "SshKeys|<sshkey set id>", // Custom Set
    ]
  ```

  ```terraform
    policy_assignment = [
        "CloudProviders|@Favorite CloudProviders", // Built-in Set
        "CloudProviders|<cloud provider set id>", // Custom Set
    ]
  ```

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
