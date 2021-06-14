---
subcategory: "Common Attribute"
---

# Workflow approver attribute

workflow_approver is a common attribute in various resources.

## Example Usage

### Example of one level approval that assigned to a role

```terraform
data "centrify_role" "approvers" {
  name = "Access Approvers"
}

resource "centrify_account" "windows_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrify_system.windows1.id
    description = "Test Account for Windows"

    workflow_enabled = true
    workflow_approver {
      guid = data.centrify_role.approvers.id
      name = data.centrify_role.approvers.name
      type = "Role"
    }
}
```

### Example of one level approval that assigned to a Centrify Directory user

```terraform
data "centrify_user" "approver" {
  username = "approver@example.com"
}

resource "centrify_account" "windows_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrify_system.windows1.id
    description = "Test Account for Windows"

    workflow_enabled = true
    workflow_approver {
      guid = data.centrify_user.approver.id
      name = data.centrify_user.approver.username
      type = "User"
    }
}
```

### Example of one level approval that assigned to manager with backup approver

```terraform
resource "centrify_account" "windows_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrify_system.windows1.id
    description = "Test Account for Windows"

    workflow_enabled = true
    workflow_approver {
      type = "Manager"
      no_manager_action = "useBackup"
      backup_approver {
        guid = data.centrify_role.approvers.id
        name = data.centrify_role.approvers.name
        type = "Role"
      }
    }
}
```

### Example of two levels approval. First level to an AD user while second level to user's manager

```terraform
data "centrify_directoryservice" "demo_lab" {
    name = "demo.lab"
    type = "Active Directory"
}

// data source for AD user ad.user@demo.lab
data "centrify_directoryobject" "ad_user" {
    directory_services = [
        data.centrify_directoryservice.demo_lab.id
    ]
    name = "ad.user"
    object_type = "User"
}

resource "centrify_account" "windows_account" {
    name = "testaccount"
    credential_type = "Password"
    password = "xxxxxxxxxxxxxx"
    host_id = centrify_system.windows1.id
    description = "Test Account for Windows"

    workflow_enabled = true
    workflow_approver {
      guid = data.centrify_directoryobject.ad_user.id
      name = "ad.user@demo.lab"
      type = "User"
      options_selector = true // this attribute must be added to only one approval level if there are multiple levels
    }
    workflow_approver {
      type = "Manager"
      no_manager_action = "deny"
    }
}
```

## Argument Reference

Optional:

- `name` - (String) User or role name of the approver.
- `type` - (String) Type of the the approver. Can be set to `User`, `Role` or `Manager`.
- `guid` - (String) Object ID of approver. Applicable if `type` is either `User` or `Role`.
- `options_selector` - (Boolean) When there are multiple approval levels, set this attribute to `true` for only **one**.
- `no_manager_action` - (String) Action if user has no manager. Applicable if `type` is `Manager`. Can be set to `approve`, `deny` or `useBackup`.
- `backup_approver` - (Block List, Max: 1) The backup approver. (see [reference for `backup_approver`](#reference-for-backup_approver)).

## Reference for `backup_approver`

Optional:

- `name` - (String) User or role name of backup approver.
- `type` - (String) Type of the backup approver. Can be set to `User` or `Role`.
- `guid` - (String) Object ID of backup approver.
