# Centrify Provider

The Centrify provider is used to interact with the resources in Centrify Platform. It also allows other Terraform providers to retrieve vaulted password or secret from Centrify Platform.

Use the navigation to the left to read about the available resources.

## Using The Provider

### Specifying Provider Requirements

Use special `terraform` configuration block type to configure some behaviors of Terraform itself, such as provider source and minimum version.

```terraform
terraform {
  required_providers {
    centrify = {
      source  = "marcozj/centrify"
    }
  }
}
```

### Configure Provider Credential

The provider needs to be configured with the proper credentials before it can be used.

#### Example Usage (OAuth client id and credential authentication)

```terraform
# Configure Centrify Provider to use OAuth client id and credential authentication
provider "centrify" {
    url = "https://<tenantid>.my.centrify.net"
    appid = "<YOUR APPLICATION ID>"
    scope = "<YOUR OAUTH2 SCOPE>"
    username = "<YOUR OAUTH2 CLIENT ID>"
    password = "<YOUR OAUTH2 CLIENT CREDENTIAL>"
}
```

#### Example Usage (OAuth2 token authentication)

```terraform
# Configure Centrify Provider to use OAuth2 token authentication
provider "centrify" {
    url = "https://<tenantid>.my.centrify.net"
    appid = "<YOUR APPLICATION ID>"
    scope = "<YOUR OAUTH2 SCOPE>"
    token = "<YOUR OAUTH2 TOKEN>"
}
```

#### Example Usage (DMC authentication)

```terraform
# Configure Centrify Provider to use DMC authentication
# The host on which terraform is run must have Centrify Client installed and enrolled into Centrify Platform
provider "centrify" {
    url = "https://<tenantid>.my.centrify.net"
    scope = "<YOUR DMC SCOPE>"
    use_dmc = true
}
```

## Provider Argument Reference

The Provider supports OAuth2 and DMC authentication methods.

- `url` - (Required) This is the cloud tenant or on-prem PAS URL, for example `https://abc1234.my.centrify.net`. It must be provided, but it can also be sourced from the `CENTRIFY_URL` environment variable.
- `appid` - (Optional) This is the OAuth application ID configured in Centrify Platform. It must be provided if `use_dmc` isn't set to true. It can also be sourced from the `CENTRIFY_APPID` environment variable.
- `scope` - (Required) This is either the OAuth or DMC scope. It must be provided, but it can also be sourced from the `CENTRIFY_SCOPE` environment variable.
- `token` - (Optional) This is the Oauth token. It can also be sourced from the `CENTRIFY_TOKEN` environment variable.
- `username` - (Optional) Authorized user to retrieve Oauth token. It can also be sourced from the `CENTRIFY_USERNAME` environment variable. If `token` is provided, this argument is ignored.
- `password` - (Optional) Authorized user's password for retrieving Oauth token. It can also be sourced from the `CENTRIFY_PASSWORD` environment variable. If `token` is provided, this argument is ignored.
- `use_dmc` - (Optional) Whether to use DMC authentication. It can also be sourced from the `CENTRIFY_USEDMC` environment variable. The default is `false`. If this is set to `true`, `appid`, `token`, `username` and `password` arguments are ingored.
- `skip_cert_verify` - (Optional) Whether to skip certificate validation. It is used for testing against on-prem PAS deployment which uses self-signed certificate. It can also be sourced from the `CENTRIFY_SKIPCERTVERIFY` environment variable. The default is `false`.
- `log_level` - (Optional) Log level. Can be set to `fatal`, `error`, `info`, or `debug`. It can also be sourced from `CENTRIFY_LOGLEVEL` environment variable. Default is `error`.
- `logpath` - (Optional) If specified, logging information is written to the file. It can also be sourced from `CENTRIFY_LOGPATH` environment variable.

## Supported Resources and Data Sources

|  Entity  |  Resource  |  Data Source  |
| ---- | ---- | --- |
| Directory Service | | [`centrify_directoryservice`](./data-sources/directoryservice.md) |
| Directory Object | | [`centrify_directoryobject`](./data-sources/directoryobject.md) |
| Global Group Mapping | [`centrify_globalgroupmappings`](./resources/globalgroupmappings.md) | |
| Centrify Directory User | [`centrify_user`](./resources/user.md) | [`centrify_user`](./data-sources/user.md) |
| Role | [`centrify_role`](./resources/role.md) | [`centrify_role`](./data-sources/role.md) |
| Authentication Profile | [`centrify_authenticationprofile`](./resources/authenticationprofile.md) | [`centrify_authenticationprofile`](./data-sources/authenticationprofile.md) |
| Password Profile | [`centrify_passwordprofile`](./resources/passwordprofile.md) | [`centrify_passwordprofile`](./data-sources/passwordprofile.md) |
| Connector | | [`centrify_connector`](./data-sources/connector.md) |
| System | [`centrify_system`](./resources/vaultsystem.md) | [`centrify_system`](./data-sources/vaultsystem.md) |
| Database | [`centrify_database`](./resources/vaultdatabase.md) | [`centrify_database`](./data-sources/vaultdatabase.md) |
| Domain | [`centrify_domain`](./resources/vaultdomain.md) | [`centrify_domain`](./data-sources/vaultdomain.md) |
| Domain Configuration | [`centrify_domainconfiguration`](./resources/vaultdomainconfiguration.md) | |
| Cloud Provider | [`centrify_cloudprovider`](./resources/cloudprovider.md) | [`centrify_cloudprovider`](./data-sources/cloudprovider.md) |
| Account | [`centrify_account`](./resources/vaultaccount.md) | [`centrify_account`](./data-sources/vaultaccount.md) |
| Multiplexed Account | [`centrify_multiplexedaccount`](./resources/multiplexedaccount.md) | [`centrify_multiplexedaccount`](./data-sources/multiplexedaccount.md) |
| Secret | [`centrify_secret`](./resources/vaultsecret.md) | [`centrify_secret`](./data-sources/vaultsecret.md) |
| Secret Folder | [`centrify_secretfolder`](./resources/vaultsecretfolder.md) | [`centrify_secretfolder`](./data-sources/vaultsecretfolder.md) |
| SSH Key | [`centrify_sshkey`](./resources/sshkey.md) | [`centrify_sshkey`](./data-sources/sshkey.md) |
| Windows Service | [`centrify_service`](./resources/service.md) | [`centrify_service`](./data-sources/service.md) |
| Generic Web App | [`centrify_webapp_generic`](./resources/webapp_generic.md) | [`centrify_webapp_generic`](./data-sources/webapp_generic.md) |
| SAML Web App | [`centrify_webapp_saml`](./resources/webapp_saml.md) | [`centrify_webapp_saml`](./data-sources/webapp_saml.md) |
| Oauth Web App | [`centrify_webapp_oauth`](./resources/webapp_oauth.md) | [`centrify_webapp_oauth`](./data-sources/webapp_oauth.md) |
| OpenID Connect Web App | [`centrify_webapp_oidc`](./resources/webapp_oidc.md) | [`centrify_webapp_oidc`](./data-sources/webapp_oidc.md) |
| Desktop App | [`centrify_desktopapp`](./resources/desktopapp.md) | [`centrify_desktopapp`](./data-sources/desktopapp.md) |
| Policy Order | [`centrify_policyorder`](./resources/policy.md) | |
| Policy | [`centrify_policy`](./resources/policy.md) | [`centrify_policy`](./data-sources/policy.md) |
| Global Workflow | [`centrify_globalworkflow`](./resources/globalworkflow.md) | |
