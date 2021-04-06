---
subcategory: "Applications"
---

# centrifyvault_webapp_saml (Data Source)

This data source gets information of SAML web application.

## Example Usage

```terraform
data "centrifyvault_webapp_saml" "saml_webapp" {
  name = "My SAML App"
}

output "idp_metadata_url" {
  value = data.centrifyvault_webapp_saml.saml_webapp.idp_metadata_url
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_saml)

## Search Attributes

### Required

- `name` - (String) Name of the SAML application.

### Optional

- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.
- `corp_identifier` - (String) AWS Account ID or Jira Cloud Subdomain. Applicable when `AWSConsoleSAML` or `JIRACloudSAML` template is used.
- `app_entity_id` - (String) Cloudera Entity ID or JIRA Cloud SP Entity ID. Applicable when `ClouderaSAML` or `JIRACloudSAML` template is used.
- `sp_entity_id` - (String) SP Entity ID, also known as SP Issuer, or Audience, is a value given by your Service Provider.

## Attributes Reference

- `id` - id of the SAML application.
- `name` - name property.
- `application_id` - application_id property.
- `corp_identifier` - corp_identifier property.
- `app_entity_id` - app_entity_id property.
- `sp_entity_id` - sp_entity_id property.
- `sp_metadata_url` - sp_metadata_url property.
- `sp_metadata_xml` - sp_metadata_xml property.
- `idp_metadata_url` - idp_metadata_url property.
