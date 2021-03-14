---
page_title: "Centrify Provider"
description: |-
  The Centrify provider is used to interact with the resources in Centrify Vault.
---

# Centrify Provider

The Centrify provider is used to interact with the resources in Centrify Vault. It also allows other Terraform providers to retrieve vaulted password or secret from Centrify Vault.

Use the navigation to the left to read about the available resources.

## Using The Provider

### Specifying Provider Requirements**

Use special `terraform` configuration block type to configure some behaviors of Terraform itself, such as provider source and minimum version.

```terraform
terraform {
  required_providers {
    centrifyvault = {
      source  = "marcozj/centrifyvault"
    }
  }
}
```

### Configure Provider Credential

The provider needs to be configured with the proper credentials before it can be used.

#### Example Usage (OAuth client id and credential authentication)

```terraform
# Configure CentrifyVault Provider to use OAuth client id and credential authentication
provider "centrifyvault" {
    url = "https://<tenantid>.my.centrify.net"
    appid = "<YOUR APPLICATION ID>"
    scope = "<YOUR OAUTH2 SCOPE>"
    username = "<YOUR OAUTH2 CLIENT ID>"
    password = "<YOUR OAUTH2 CLIENT CREDENTIAL>"
}
```

#### Example Usage (OAuth2 token authentication)

```terraform
# Configure CentrifyVault Provider to use OAuth2 token authentication
provider "centrifyvault" {
    url = "https://<tenantid>.my.centrify.net"
    appid = "<YOUR APPLICATION ID>"
    scope = "<YOUR OAUTH2 SCOPE>"
    token = "<YOUR OAUTH2 TOKEN>"
}
```

#### Example Usage (DMC authentication)

```terraform
# Configure CentrifyVault Provider to use DMC authentication
# The host on which terraform is run must have Centrify Client installed and enrolled into Centrify Vault
provider "centrifyvault" {
    url = "https://<tenantid>.my.centrify.net"
    scope = "<YOUR DMC SCOPE>"
    use_dmc = true
}
```

## Provider Argument Reference

The Provider supports OAuth2 and DMC authentication methods.

- `url` - (Required) This is the cloud tenant or on-prem PAS URL, for example `https://abc1234.my.centrify.net`. It must be provided, but it can also be sourced from the `VAULT_URL` environment variable.
- `appid` - (Optional) This is the OAuth application ID configured in Centrify Vault. It must be provided if `use_dmc` isn't set to true. It can also be sourced from the `VAULT_APPID` environment variable.
- `scope` - (Required) This is either the OAuth or DMC scope. It must be provided, but it can also be sourced from the `VAULT_SCOPE` environment variable.
- `token` - (Optional) This is the Oauth token. It can also be sourced from the `VAULT_TOKEN` environment variable.
- `username` - (Optional) Authorized user to retrieve Oauth token. It can also be sourced from the `VAULT_USERNAME` environment variable. If `token` is provided, this argument is ignored.
- `password` - (Optional) Authorized user's password for retrieving Oauth token. It can also be sourced from the `VAULT_PASSWORD` environment variable. If `token` is provided, this argument is ignored.
- `use_dmc` - (Optional) Whether to use DMC authentication. It can also be sourced from the `VAULT_USEDMC` environment variable. The default is `false`. If this is set to `true`, `appid`, `token`, `username` and `password` arguments are ingored.
- `skip_cert_verify` - (Optional) Whether to skip certificate validation. It is used for testing against on-prem PAS deployment which uses self-signed certificate. It can also be sourced from the `VAULT_SKIPCERTVERIFY` environment variable. The default is `false`.
- `log_level` - (Optional) Log level. Can be set to `fatal`, `error`, `info`, or `debug`. It can also be sourced from `VAULT_LOGLEVEL` environment variable. Default is `error`.
- `logpath` - (Optional) If specified, logging information is written to the file. It can also be sourced from `VAULT_LOGPATH` environment variable.

## Supported Resources and Data Sources

|  Entity  |  Resource  |  Data Source  |
| ---- | ---- | --- |
| Directory Service | | [`centrifyvault_directoryservice`](./data-sources/centrifyvault_directoryservice.md) |
| Directory Object | | [`centrifyvault_directoryobject`](./data-sources/centrifyvault_directoryobject.md) |
| Global Group Mapping | [`centrifyvault_globalgroupmappings`](./resources/centrifyvault_globalgroupmappings.md) | |
| Centrify Directory User | [`centrifyvault_user`](./resources/centrifyvault_user.md) | [`centrifyvault_user`](./data-sources/centrifyvault_user.md) |
| Role | [`centrifyvault_role`](./resources/centrifyvault_role.md) | [`centrifyvault_role`](./data-sources/centrifyvault_role.md) |
| Authentication Profile | [`centrifyvault_authenticationprofile`](./resources/centrifyvault_authenticationprofile.md) | [`centrifyvault_authenticationprofile`](./data-sources/centrifyvault_authenticationprofile.md) |
| Password Profile | [`centrifyvault_passwordprofile`](./resources/centrifyvault_passwordprofile.md) | [`centrifyvault_passwordprofile`](./data-sources/centrifyvault_passwordprofile.md) |
| System | [`centrifyvault_vaultsystem`](./resources/centrifyvault_vaultsystem.md) | [`centrifyvault_vaultsystem`](./data-sources/centrifyvault_vaultsystem.md) |
| Database | [`centrifyvault_vaultdatabase`](./resources/centrifyvault_vaultdatabase.md) | [`centrifyvault_vaultdatabase`](./data-sources/centrifyvault_vaultdatabase.md) |
| Domain | [`centrifyvault_vaultdomain`](./resources/centrifyvault_vaultdomain.md) | [`centrifyvault_vaultdomain`](./data-sources/centrifyvault_vaultdomain.md) |
| Domain Reconciliation | [`centrifyvault_vaultdomainreconciliation`](./resources/centrifyvault_vaultdomainreconciliation.md) | |
| CloudProviders | [`centrifyvault_cloudprovider`](./resources/centrifyvault_cloudprovider.md) | [`centrifyvault_cloudprovider`](./data-sources/centrifyvault_cloudprovider.md) |
| Account | [`centrifyvault_vaultaccount`](./resources/centrifyvault_vaultaccount.md) | [`centrifyvault_vaultaccount`](./data-sources/centrifyvault_vaultaccount.md) |
| Multiplexed Account | [`centrifyvault_multiplexedaccount`](./resources/centrifyvault_multiplexedaccount.md) | [`centrifyvault_multiplexedaccount`](./data-sources/centrifyvault_multiplexedaccount.md) |
| Secret | [`centrifyvault_vaultsecret`](./resources/centrifyvault_vaultsecret.md) | [`centrifyvault_vaultsecret`](./data-sources/centrifyvault_vaultsecret.md) |
| Secret Folder | [`centrifyvault_vaultsecretfolder`](./resources/centrifyvault_vaultsecretfolder.md) | [`centrifyvault_vaultsecretfolder`](./data-sources/centrifyvault_vaultsecretfolder.md) |
| SSH Key | [`centrifyvault_sshkey`](./resources/centrifyvault_sshkey.md) | [`centrifyvault_sshkey`](./data-sources/centrifyvault_sshkey.md) |
| Windows Service | [`centrifyvault_service`](./resources/centrifyvault_service.md) | |
| Desktop App | [`centrifyvault_desktopapp`](./resources/centrifyvault_desktopapp.md) | |
| Policy | [`centrifyvault_policyorder`](./resources/centrifyvault_policy.md) | |
| Policy | [`centrifyvault_policy`](./resources/centrifyvault_policy.md) | [`centrifyvault_policy`](./data-sources/centrifyvault_policy.md) |