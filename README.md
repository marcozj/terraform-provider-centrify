# Terraform Provider for Centrify Platform

The Terraform Provider for Centrify Platform is a Terraform plugin that allows other Terraform providers to retrieve vaulted password or secret from Centrify Platform. It also enables full configuration management of Centrify Platform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or higher
- [Go](https://golang.org/doc/install) 1.13 or higher (to build the provider plugin)

## Building The Provider

### The GOPATH environment variable

The GOPATH environment variable specifies the location of your workspace. It defaults to a directory named go inside your home directory, so $HOME/go on Unix, and %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows.
The command go env GOPATH prints the effective current GOPATH; it prints the default location if the environment variable is unset.
If you have not set GOPATH, you can substitute $HOME/go in those commands or else run:

```sh
$ export GOPATH=$(go env GOPATH)
```

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-centrifyvault`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers
$ cd $GOPATH/src/github.com/terraform-providers
$ git clone https://github.com/marcozj/terraform-provider-centrifyvault terraform-provider-centrifyvault
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-centrifyvault
$ make build
```

To install the provider in your home directory

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-centrifyvault
$ make install
```

## Using The Provider

The provider needs to be configured with the proper credentials before it can be used. Refer to [provider document](./docs/index.md) page for details.

## Example Usage

You can use Centrify Terraform Provider to configure Centrify platform including creation/modification/deletion of user, role, system, account, etc. It also allows other Terraform providers to retrieve vaulted password or secret from Centrify platform.

Refer to **Supported Resources and Data Sources** section in [provider document](./docs/index.md) page for details of supported configurations and [example](./examples/) usage.

For example, this is how to [create a Windows system](./examples/centrify_system/system_windows_basic.tf) in Centrify Platform. This is how to [retrieve vaulted credentials](./examples/centrify_account/datasource_password.tf).
