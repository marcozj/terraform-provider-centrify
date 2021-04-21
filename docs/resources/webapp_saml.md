---
subcategory: "Applications"
---

# centrifyvault_webapp_saml (Resource)

This resource allows you to create/update/delete SAML web application.

## Example Usage

```terraform
resource "centrifyvault_webapp_saml" "saml_webapp" {
    name = "SAML Web App"
    template_name = "Generic SAML"
    description = "SAML Web Application"
    sp_config_method = 1
    sp_metadata_url = "https://nexus.microsoftonline-p.com/federationmetadata/saml20/federationmetadata.xml"

    saml_attribute {
        name = "attribute1"
        value = "value1"
    }
    saml_attribute {
        name = "attribute2"
        value = "value2"
    }
    saml_response_script = "test;"
    username_strategy = "ADAttribute"
    ad_attribute = "userprincipalname"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_webapp_saml)

## Argument Reference

### Required

- `name` - (String) Name of the SAML application.
- `template_name` - (String) SAML application template. Can be set to `Generic SAML`, `AWSConsoleSAML`, `ClouderaSAML`, `CloudLock SAML`, `ConfluenceServerSAML`, `Dome9Saml`, `GitHubEnterpriseSAML`, `JIRACloudSAML`, `JIRAServerSAML`, `PaloAltoNetworksSAML`, `SplunkOnPremSAML` or `SumoLogicSAML`.

### Optional

- `description` - (String) Description of the SAML application.
- `corp_identifier` - (String) AWS Account ID or Jira Cloud Subdomain. Applicable when `AWSConsoleSAML` or `JIRACloudSAML` template is used.
- `app_entity_id` - (String) Cloudera Entity ID or JIRA Cloud SP Entity ID. Applicable when `ClouderaSAML` or `JIRACloudSAML` template is used.
- `application_id` - (String) Application ID. Specify the name or 'target' that the mobile application uses to find this application.
- `sp_config_method` - (Int) Configuration method for Service Provider. To use manual configuration, set this to `0`. To use metadata configuration, set this to `1`.
- If `sp_config_method` is set to `1`, specify following arguments:
  - `sp_metadata_url` - (String) Service Provider metadata URL. When this is sepcified, Service Provider metadata is automatically loaded from URL and `sp_metadata_xml` is ignore.
  - `sp_metadata_xml` - (String) The metadata provided by Service Provider.
- If `sp_config_method` is set to `0`, specify following arguments:
  - `sp_entity_id` - (String) SP Entity ID, also known as SP Issuer, or Audience, is a value given by your Service Provider.
  - `acs_url` - (String) Assertion Consumer Service (ACS) URL.
  - `recipient_sameas_acs_url` - (Boolean) Recipient is same as ACS URL. Default is `true`.
  - `recipient` - (String) Service Provider recipient if it is different from ACS URL.
  - `sign_assertion` - (Boolean) Sign assertion if true, otherwise sign response. Default is `true`.
  - `name_id_format` - (String) This is the Format attribute value in the \<NameID\> element in SAML Response. Select the NameID Format that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'. Can be set to `unspecified`, `emailAddress`, `transient`, `persistent`, `entity`, `kerberos`, `WindowsDomainQualifiedName` or `X509SubjectName`. Default is `unspecified`.
  - `sp_single_logout_url` - (String) Single Logout URL.
  - `relay_state` - (String) If your Service Provider specifies a Relay State value to use, specify it here.
  - `authn_context_class` - (String) Select the Authentication Context Class that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'. Can be set to `unspecified`, `PasswordProtectedTransport`, `AuthenticatedTelephony`, `InternetProtocol`, `InternetProtocolPassword`, `Kerberos`,`MobileOneFactorContract`, `MobileOneFactorUnregistered`, `MobileTwoFactorContract`, `MobileTwoFactorUnregistered`, `NomadTelephony`, `Password`, `PersonalTelephony`, `PGP`, `PreviousSession`, `SecureRemotePassword`, `Smartcard`, `SmartcardPKI`, `SoftwarePKI`, `SPKI`, `Telephony`, `TimeSyncToken`, `TLSClient`, `X509` or `XMLDSig`. Default is `unspecified`.
- `saml_attribute` - (Block Set) (see [reference for `saml_attribute`](#reference-for-saml_attribute)).
- `saml_response_script` - (String) Javascript used to produce custom logic for SAML response.
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

## Reference for `saml_attribute`

Optional:

- `name` - (String) Name of the attribute.
- `vaule` - (String) Value of the attribute.

## Import

SAML Application can be imported using the resource `id`, e.g.

```shell
terraform import centrifyvault_webapp_saml.example xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```

**Limitation:** `permission` and `sets` aren't support in import process.
