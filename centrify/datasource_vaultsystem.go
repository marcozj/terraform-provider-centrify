package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/computerclass"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultSystem() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultSystemRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the system",
			},
			"fqdn": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Hostname or IP address of the system",
				ValidateFunc: validation.NoZeroValues,
			},
			"computer_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the system",
				ValidateFunc: validation.StringInSlice([]string{
					computerclass.Windows.String(),
					computerclass.Unix.String(),
					computerclass.CiscoAsyncOS.String(),
					computerclass.CiscoIOS.String(),
					computerclass.CiscoNXOS.String(),
					computerclass.JuniperJunos.String(),
					computerclass.HPNonStop.String(),
					computerclass.IBMi.String(),
					computerclass.CheckPointGaia.String(),
					computerclass.PaloAltoPANOS.String(),
					computerclass.F5BIGIP.String(),
					computerclass.VMwareVMkernel.String(),
					computerclass.GenericSSH.String(),
					computerclass.CustomSSH.String(),
				}, false),
			},
			"session_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login session type that the system supports",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the system",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port that used to connect to the system",
			},
			"use_my_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable Use My Account",
			},
			"management_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Management mode of the system. For Windows only",
			},
			"management_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Management port for account management. For Windows, F5, PAN-OS and VMKernel only",
			},
			"system_timezone": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "System time zone",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxyuser": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxyuser_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"proxyuser_managed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			// System -> Policy menu related settings
			"checkout_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the number of minutes that a checked out password is valid.",
			},
			"allow_remote_access": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow access from a public network (web client only)",
			},
			"allow_rdp_clipboard": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow RDP client to sync local clipboard with remote session",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default System Login Profile (used if no conditions matched)",
			},
			"privilege_elevation_default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default Privilege Elevation Profile (used if no conditions matched)",
			},
			// System -> Advanced menu related settings
			"local_account_automatic_maintenance": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Local Account Automatic Maintenance",
			},
			"local_account_manual_unlock": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Local Account Manual Unlock",
			},
			"domain_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AD domain that this system belongs to",
			},
			"remove_user_on_session_end": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Remove local accounts upon session termination - Windows only ",
			},
			"allow_multiple_checkouts": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow multiple password checkouts for this system",
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
			"enable_sshkey_rotation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic SSH key rotation",
			},
			"sshkey_rotate_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SSH key rotation interval (days)",
			},
			"minimum_sshkey_age": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum SSH Key Age (days)",
			},
			"sshkey_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SSH Key Generation Algorithm",
			},
			"enable_sshkey_history_cleanup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic SSH key history cleanup",
			},
			"sshkey_historycleanup_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SSH key history cleanup (days)",
			},
			// Workflow - Agent Auth and Privilege Elevation
			"agent_auth_workflow_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"agent_auth_workflow_approver": getWorkflowApproversSchema(),
			"privilege_elevation_workflow_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"privilege_elevation_workflow_approver": getWorkflowApproversSchema(),
			// System -> Zone Role Workflow menu related settings
			"use_domainadmin_for_zonerole_workflow": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use Domain Administrator Account for Zone Role Workflow operations",
			},
			"enable_zonerole_workflow": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable zone role requests for this system",
			},
			"use_domain_assignment_for_zoneroles": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Assignable Zone Roles - Use domain assignments",
			},
			"assigned_zonerole": getZoneRoleSchema(),
			"use_domain_assignment_for_zonerole_approvers": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Approver list - Use domain assignments",
			},
			"assigned_zonerole_approver": getWorkflowApproversSchema(),
			// System -> Connectors menu related settings
			"connector_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of Connectors",
			},
			"challenge_rule":           getChallengeRulesSchema(),
			"privilege_elevation_rule": getChallengeRulesSchema(),
		},
	}
}

func dataSourceVaultSystemRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding system")
	client := m.(*restapi.RestClient)
	object := vault.NewSystem(client)
	object.Name = d.Get("name").(string)
	object.FQDN = d.Get("fqdn").(string)
	if v, ok := d.GetOk("computer_class"); ok {
		object.ComputerClass = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving system with name '%s': %s", object.Name, err)
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
		case "challenge_rule", "privilege_elevation_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		case "agent_auth_workflow_approver", "privilege_elevation_workflow_approver":
			d.Set(k, processBackupApproverSchema(v))
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
