data "centrifyvault_user" "approver" {
  username = "admin@example.com"
}

data "centrifyvault_role" "system_admin" {
  name = "System Administrator"
}

// Account Workflow
resource "centrifyvault_globalworkflow" "account_wf" {
    type = "wf"
    settings {
        enabled = true
        approver {
            type = "Manager"
            no_manager_action = "useBackup"
            backup_approver {
                guid = data.centrifyvault_user.approver.id
                name = data.centrifyvault_user.approver.username
                type = "User"
            }
        }
        approver {
            guid = data.centrifyvault_role.system_admin.id
            name = data.centrifyvault_role.system_admin.name
            type = "Role"
        }
    }
}

// Agent Auth Workflow
resource "centrifyvault_globalworkflow" "agentauth_wf" {
    type = "agentAuthWorkflow"
    settings {
        enabled = true
        approver {
            type = "Manager"
            no_manager_action = "useBackup"
            backup_approver {
                guid = data.centrifyvault_user.approver.id
                name = data.centrifyvault_user.approver.username
                type = "User"
            }
        }
        approver {
            guid = data.centrifyvault_role.system_admin.id
            name = data.centrifyvault_role.system_admin.name
            type = "Role"
        }
    }
}

// Secrets Workflow
resource "centrifyvault_globalworkflow" "secrets_wf" {
    type = "secretsWorkflow"
    settings {
        enabled = true
        approver {
            type = "Manager"
            no_manager_action = "useBackup"
            backup_approver {
                guid = data.centrifyvault_user.approver.id
                name = data.centrifyvault_user.approver.username
                type = "User"
            }
        }
        approver {
            guid = data.centrifyvault_role.system_admin.id
            name = data.centrifyvault_role.system_admin.name
            type = "Role"
        }
    }
}

// Privilege Elevation Workflow
resource "centrifyvault_globalworkflow" "elevation_wf" {
    type = "privilegeElevationWorkflow"
    settings {
        enabled = true
        approver {
            type = "Manager"
            no_manager_action = "useBackup"
            backup_approver {
                guid = data.centrifyvault_user.approver.id
                name = data.centrifyvault_user.approver.username
                type = "User"
            }
        }
        approver {
            guid = data.centrifyvault_role.system_admin.id
            name = data.centrifyvault_role.system_admin.name
            type = "Role"
        }
    }
}