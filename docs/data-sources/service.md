---
subcategory: "Resources"
---

# centrifyvault_service (Data Source)

This data source gets information of Windows service.

## Example Usage

```terraform
data "centrifyvault_service" "testservice" {
    service_name = "TestWindowsService"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_service)

## Search Attributes

### Required

- `service_name` - (String) The name of the service to be managed.

## Attributes Reference

- `id` - (String) ID of the service.
- `service_type` - (String) Service type.
- `system_id` - (String) The ID of target Windows system where the service runs.
- `description` - (String) Description of the service.
- `enable_management` - (Boolean) Enable management of this service password.
- `admin_account_id` - (String) Administrative account id that used to manage the password for the service. Select a vaulted domain account to manage the password for this service. The account must be stored in the Privileged Access Service and have sufficient permissions to modify (rotate) the service account password. A managed domain account is recommended so it is rotated after each use.
- `multiplexed_account_id` - (String) The multiplexed account id to run the service. Select a multiplexed account to run the service. Multiplexed accounts are required to enable automated password rotation for services that run under a service or user account.
- `restart_service` - (Boolean) Restart Service when password is rotated.
- `restart_time_restriction` - (Boolean) Enforce restart time restrictions.
- `days_of_week` - (Set of String) Day of the week restart allowed.
- `restart_start_time` - (String) Start time of the time range restart is allowed.
- `restart_end_time` - (String) End time of the time range restart is allowed.
- `use_utc_time` - (Boolean) Whether to use UTC time.
