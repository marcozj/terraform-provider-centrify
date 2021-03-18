---
subcategory: "Resources"
---

# centrifyvault_service (Resource)

This resource allows you to create/update/delete service.

## Example Usage

```terraform
resource "centrifyvault_service" "testservice" {
    service_name = "TestWindowsService"
    description = "Test Windows Service in member1"
    system_id = data.centrifyvault_vaultsystem.member1.id
    service_type = "WindowsService"
    enable_management = true
    admin_account_id = data.centrifyvault_vaultaccount.ad_admin.id
    multiplexed_account_id = centrifyvault_multiplexedaccount.testmultiplex.id
    restart_service = true
    restart_time_restriction = true
    days_of_week = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
    restart_start_time = "09:00"
    restart_end_time = "10:00"
    use_utc_time = false
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_service)

### Required

- `service_name` - (String) The name of the service to be managed.
- `service_type` - (String) Service type. Can be set to `WindowsService`, `ScheduledTask`, or `IISApplicationPool`.
- `system_id` - (String) The ID of target Windows system where the service runs.

### Optional

- `description` - (String) Description of the service.
- `enable_management` - (Boolean) Enable management of this service password. **Note**: the Windows system must be up and service must exists.
- `admin_account_id` - (String) Administrative account id that used to manage the password for the service. Select a vaulted domain account to manage the password for this service. The account must be stored in the Privileged Access Service and have sufficient permissions to modify (rotate) the service account password. A managed domain account is recommended so it is rotated after each use.
- `multiplexed_account_id` - (String) The multiplexed account id to run the service. Select a multiplexed account to run the service. Multiplexed accounts are required to enable automated password rotation for services that run under a service or user account.
- `restart_service` - (Boolean) Restart Service when password is rotated.
- `restart_time_restriction` - (Boolean) Enforce restart time restrictions.
- `days_of_week` - (Set of String) Day of the week restart allowed.
- `restart_start_time` - (String) Start time of the time range restart is allowed.
- `restart_end_time` - (String) End time of the time range restart is allowed.
- `use_utc_time` - (Boolean) Whether to use UTC time.
- `permission` - (Block Set) Domain permissions. Refer to [permission](./attribute_permission.md) attribute for details.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.
