# RELEASE NOTES

## 0.1.6 (May 15, 2021)

IMPROVEMENTS:

- New `bulkupdate` argument for `centrifyvault_globalgroupmappings` resource.

## 0.1.5 (April 21, 2021)

FEATURES:

- **New Data Resource:** `centrifyvault_desktopapp`
- **New Data Resource:** `centrifyvault_service`

IMPROVEMENTS:

- Add support for importing resources

## 0.1.4 (April 9, 2021)

IMPROVEMENTS:

- Expose more attributes reference for all data source types.

BUG FIXES:

- `centrifyvault_connector` data source fail to run when Connector is not installed on AD joined machine.

## 0.1.3 (April 6, 2021)

FEATURES:

- **New Resource:** `centrifyvault_globalworkflow`
- **New Resource:** `centrifyvault_webapp_generic`
- **New Resource:** `centrifyvault_webapp_saml`
- **New Resource:** `centrifyvault_webapp_oauth`
- **New Resource:** `centrifyvault_webapp_oidc`
- **New Data Resource:** `centrifyvault_webapp_generic`
- **New Data Resource:** `centrifyvault_webapp_saml`
- **New Data Resource:** `centrifyvault_webapp_oauth`
- **New Data Resource:** `centrifyvault_webapp_oidc`

IMPROVEMENTS:

- `centrifyvault_connector` data source
  - Add `status`, `forest`, `version` and `vpc_identifier` search attributes
  - Add more attribute references.
- `centrifyvault_vaultaccount` resource
  - Add `workflow_enabled` and `workflow_approver` attributes
- `centrifyvault_desktopapp` resource
  - Add `workflow_enabled` and `workflow_approver` attributes
- `centrifyvault_vaultsecret` resource
  - Add `workflow_enabled` and `workflow_approver` attributes
- `centrifyvault_vaultsystem` resource
  - Add Agent Auth and Privilege Elevation Workflow related attributes: `agent_auth_workflow_enabled`, `agent_auth_workflow_approver`, `privilege_elevation_workflow_enabled` and `privilege_elevation_workflow_approver`.
  - Add Zone Role Workflow related attributes: `use_domainadmin_for_zonerole_workflow`, `enable_zonerole_workflow`, `use_domain_assignment_for_zoneroles`, `assigned_zonerole`, `use_domain_assignment_for_zonerole_approvers` and `assigned_zonerole_approver`.
- Replace `centrifyvault_vaultdomainreconciliation` with `centrifyvault_vaultdomainconfiguration`
  - Add `enable_zonerole_workflow`, `assigned_zonerole` and `assigned_zonerole_approver` attribute to `centrifyvault_vaultdomainconfiguration`
- `centrifyvault_vaultdomain` resource
  - Rename `enable_zone_role_cleanup` to `enable_zonerole_cleanup`
  - Rename `zone_role_cleanup_interval` to `zonerole_cleanup_interval`
  - Add `parent_id` and `forest_id` attributes
- Improve error message for all data sources when object can't be found.

BUG FIXES:

- `centrifyvault_vaultaccount` resource: `password` attribute causes apply action always update resource.
- Detect `challenge_rule` change in tenant for these Terraform managed resources: `resource_desktopapp`, `resource_sshkey`, `resource_vaultaccount`, `resource_vaultcloudprovider`, `resource_vaultsecret` and `resource_vaultsystem`.

## 0.1.2 (March 18, 2021)

BUG FIXES:

- Documentation links and file layout.

## 0.1.1 (March 14, 2021)

BUG FIXES:

- `connector_list` attribute in for `centrifyvault_vaultsystem` resource doesn't take effect during creation.
- Documentation links.

## 0.1.0 (March 14, 2021)

- Initial release
