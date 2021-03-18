---
subcategory: "Policy Configuration"
---

# mobile_device attribute

**mobile_device** is a sub attribute in settings attribute within **centrifyvault_policy** Resource.

## Example Usage

```terraform
resource "centrifyvault_policy" "test_policy" {
    name = "Test Policy"
    description = "Test Policy"
    link_type = "Role"
    policy_assignment = [
        data.centrifyvault_role.system_admin.id,
    ]
    
    settings {
        mobile_device {
            allow_enrollment = true
            permit_non_compliant_device = true
            enable_invite_enrollment = true
            allow_notify_multi_devices = true
            enable_debug = true
            location_tracking = true
            force_fingerprint = true
            allow_fallback_pin = true
            require_passcode = true
            auto_lock_timeout = 15
            lock_app_on_exit = true
        }
    }
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policy_mobile_device.tf)

## Argument Reference

Optional:

- `allow_enrollment` - (Boolean) Permit device registration.
- `permit_non_compliant_device` - (Boolean) Permit non-compliant devices to register. This policy must be set to Yes to bypass the Google services SafetyNet check and allow registration of the device. **Notes:** This policy is enforced by the registration application running on the device. If the user uses web registration instead of an application, this policy is not enforced. This policy is not supported on OS X and Android devices earlier than version 2.3.
- `enable_invite_enrollment` - (Boolean) Enable invite based registration. You must enable it to allow users to register their devices using the system generated QR code.
- `allow_notify_multi_devices` - (Boolean) Allow user notifications on multiple devices.
- `enable_debug` - (Boolean) Enable debug logging. There are two logging modes on devices: regular - the default setting - and debug logging. Use this policy to turn on the debug logging mode.
- `location_tracking` - (Boolean) Report mobile device location.
- `force_fingerprint` - (Boolean) Enforce fingerprint scan for Mobile Authenticator. Enable it to require that users provide a finger print scan to use mobile authenticator. Using the associated policy option, users can alternatively use the client application PIN for access.
  - `allow_fallback_pin` - (Boolean) Allow App PIN. Enable it to allow users to access the mobile authenticator code using finger print or the client application PIN.
- `require_passcode` - (Boolean) Require client application passcode on device. Enable it to require a passcode to open the client application, disalbe it to allow opening the client application without a passcode. **Important:** You must enable it to enable other client application passcode policies.
  - `auto_lock_timeout` - (Number) Auto-Lock (minutes). Set the number of minutes of inactivity before the client application is locked. **Important:** The "Require client application passcode on device" policy must be set to Yes to enforce this policy. Valid values are `1`, `2`, `5`, `15` or `30`.
  - `lock_app_on_exit` - (Boolean) Lock on exit. Enable it to require a passcode to open the client application after the client has been closed, disable it to allow opening the client application without a passcode. **Important:** The "Require client application passcode on device" policy must be set to Yes to enforce this policy.
