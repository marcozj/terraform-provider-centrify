---
subcategory: "Common Attribute"
---

# Zone role attribute

assigned_zonerole is a common attribute for `centrify_system` and `centrify_domain` resources.

## Example Usage

```terraform
resource "centrify_system" "windows" {
    name = "demo_windows"
    fqdn = "192.168.18.30"
    computer_class = "Windows"
    session_type = "Rdp"
    description = "Demo Windows system 1"

    domain_id = data.centrify_domain.example_com.id
    use_domainadmin_for_zonerole_workflow = true
    enable_zonerole_workflow = true
    
    use_domain_assignment_for_zoneroles = false
    assigned_zonerole {
      name = "Windows Login/Global"
    }
    assigned_zonerole {
      name = "cfyw-Windows System Admin/Windows Zone"
    }

    use_domain_assignment_for_zonerole_approvers = false
    assigned_zonerole_approver {
        guid = data.centrify_user.approver.id
        name = data.centrify_user.approver.username
        type = "User"
        options_selector = true
    }
    assigned_zonerole_approver {
        guid = data.centrify_role.system_admin.id
        name = data.centrify_role.system_admin.name
        type = "Role"
    }
}
```

## Argument Reference

Required:

- `name` - (String) Zone role name. In the format of "zone role name>/zone name". For example "Windows Login/Global". **Note:** The rest of optional attributes are automatically resolved and filled if specified role `name` is found.

Optional:

- `zone_description` - (String) Description of the zone.
- `zone_dn` - (String) Distinguished Name (DN) of the zone.
- `description` - (String) Description of the zone role.
- `zone_canonical_name` - (String) Cannoical name of the zone.
- `unix` - (Boolean) The zone role is for Unix.
- `windows` - (Boolean) The zone role is for Windows.
