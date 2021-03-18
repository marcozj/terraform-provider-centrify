---
subcategory: "Common Attribute"
---

# challenge_rule attribute

challenge_rule is a common attribute in various resources.

## Example Usage

```terraform
resource "centrifyvault_vaultsystem" "win_system" {
    name = "WindowsServer"
    fqdn = "192.168.2.3"
    computer_class = "Windows"
    session_type = "Rdp"
    description = "My Windows system"

    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.xxx.id
      rule {
        filter = "IpAddress"
        condition = "OpInCorpIpRange"
      }
    }

    challenge_rule {
      authentication_profile_id = data.centrifyvault_authenticationprofile.xxx.id
      rule {
        filter = "DayOfWeek"
        condition = "OpIsDayOfWeek"
        value = "L,1,3,4,5"
      }
      rule {
        filter = "Browser"
        condition = "OpNotEqual"
        value = "Firefox"
      }
      rule {
        filter = "CountryCode"
        condition = "OpNotEqual"
        value = "GA"
      }
    }

}
```

## Argument Reference

- `authentication_profile_id` - (String) Authentication Profile ID (if all conditions met).
- `rule` - (Block Set) (see [Rule Argument Reference](#rule-argument-reference))

### Rule Argument Reference

#### Required

- `filter` - (String) Rule filter. Can be `IpAddress`, `IdentityCookie`, `DayOfWeek`, `Date`, `DateRange`, `Time`, `DeviceOs`, `Browser`, `CountryCode`, or `Zso`.
- `condition` - (String) Rule condition. Can be `OpInCorpIpRange`, `OpNotInCorpIpRange`, `OpExists`, `OpNotExists`, `OpIsDayOfWeek`, `OpLessThan`, `OpGreaterThan`, `OpBetween`, `OpEqual`, `OpNotEqual`, `OpIs`, `OpIsNot`, `OpHeader` or `OpArgument`.

#### Optional

- `value` - (String) Rule vaule.
