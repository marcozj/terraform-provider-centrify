---
subcategory: "Resources"
---

# centrifyvault_connector (Data Source)

This data source gets information of Centrify Connector.

## Example Usage

```terraform
data "centrifyvault_connector" "connector1" {
    name = "connector_host1" // Connector name registered in Centrify
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_connector)

## Search Attributes

### Required

- `name` - (String) Name of the Connector.

### Optional

- `status` - (String) Online status of the Connector. Can be set to `Active` or `Inactive`.
- `machine_name` - (String) Machine name of the Connector.
- `dns_host_name` - (String) DNS host name of the Connector.
- `forest` - (String) Forst name of the Connector.
- `version` - (String) Version number of the Connector.
- `vpc_identifier` - (String) AWS VPC Identifier. In a format of arn:aws:ec2:\<region name\>:\<VPC owner ID\>:vpc/\<VPC ID\>. For example: arn:aws:ec2:ap-southeast-1:012345678912:vpc/vpc-06e6bcae08ed8577c

## Attributes Reference

- `id` - (String) The ID of this resource.
- `name` - (String) name property.
- `machine_name` - (String) machine_name property.
- `dns_host_name` - (String) dns_host_name property.
- `forest` - (String) forest property.
- `ssh_service` - (String) ssh_service property.
- `rdp_service` - (String) rdp_service property.
- `ad_proxy` - (String) ad_proxy property.
- `app_gateway` - (String) app_gateway property.
- `http_api_service` - (String) http_api_service property.
- `ldap_proxy` - (String) ldap_proxy property.
- `radius_service` - (String) radius_service property.
- `radius_external_service` - (String) radius_external_service property.
- `version` - (String) version property.
- `vpc_identifier` - (String) vpc_identifier property.
