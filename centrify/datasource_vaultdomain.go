package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultDomainRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the domain",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the domain",
			},
			"forest_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Policy menu related settings
			"checkout_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Checkout lifetime (minutes)",
			},
			// Advanced -> Security Settings
			"allow_multiple_checkouts": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow multiple password checkouts per AD account added for this domain",
			},
			"enable_password_rotation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic password rotation",
			},
			"password_rotate_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Password rotation interval (days)",
			},
			"enable_password_rotation_after_checkin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable password rotation after checkin",
			},
			"minimum_password_age": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Password Age (days)",
			},
			"password_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password complexity profile id",
			},
			// Advanced -> Maintenance Settings
			"enable_password_history_cleanup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic password history cleanup",
			},
			"password_historycleanup_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Password history cleanup (days)",
			},
			// Advanced -> Domain/Zone Tasks
			"enable_zone_joined_check": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic domain/zone joined check",
			},
			"zone_joined_check_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Domain/zone joined check interval (minutes)",
			},
			"enable_zonerole_cleanup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic removal of expired zone role assignments",
			},
			"zonerole_cleanup_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Expired zone role assignment removal interval (hours)",
			},
			// Domain -> Connectors menu related settings
			"connector_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of Connectors",
			},
			// Advanced menu -> Administrative Account Settings
			"administrative_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of administrative account",
			},
			"administrator_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"administrative_account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of administrative account",
			},
			"administrative_account_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password of administrative account",
			},
			"auto_domain_account_maintenance": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable Automatic Domain Account Maintenance",
			},
			"auto_local_account_maintenance": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable Automatic Local Account Maintenance",
			},
			"manual_domain_account_unlock": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable Manual Domain Account Unlock",
			},
			"manual_local_account_unlock": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable Manual Local Account Unlock",
			},
			"provisioning_admin_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Provisioning Administrative Account ID (must be managed)",
			},
			"reconciliation_account_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reconciliation account name",
			},
			// Zone Role Workflow menu
			"enable_zonerole_workflow": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable zone role requests for this system",
			},
			"assigned_zonerole":          getZoneRoleSchema(),
			"assigned_zonerole_approver": getWorkflowApproversSchema(),
		},
	}
}

func dataSourceVaultDomainRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding domain")
	client := m.(*restapi.RestClient)
	object := vault.NewDomain(client)
	object.Name = d.Get("name").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving domain with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "connector_list":
			// Convert "value1,value1" to schema.TypeSet
			d.Set("connector_list", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		case "assigned_zoneroles":
			zrschema, err := convertZoneRoleSchema("{\"ZoneRoleWorkflowRole\":" + v.(string) + "}")
			if err != nil {
				return err
			}
			d.Set("assigned_zonerole", zrschema)
			d.Set(k, v)
		case "assigned_zonerole_approvers":
			// convertWorkflowSchema expects "assigned_zonerole_approvers" in format of {"WorkflowApprover":[{"Type":"Manager","NoManagerAction":"useBackup","BackupApprover":{"Guid":"xxxxxx_xxxx_xxxx_xxxxxxxxx","Name":"Infrastructure Owners","Type":"Role"}}]}
			// which matches ProxyWorkflowApprover struct
			wfschema, err := convertWorkflowSchema("{\"WorkflowApprover\":" + v.(string) + "}")
			if err != nil {
				return err
			}
			d.Set("assigned_zonerole_approver", wfschema)
			d.Set(k, v)
		default:
			d.Set(k, v)
		}
	}

	return nil
}
