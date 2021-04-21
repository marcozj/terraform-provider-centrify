package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceVaultDomainConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultDomainConfigurationCreate,
		Read:   resourceVaultDomainConfigurationRead,
		Update: resourceVaultDomainConfigurationUpdate,
		Delete: resourceVaultDomainConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of domain",
			},
			// Advanced menu -> Administrative Account Settings
			"administrative_account_id": {
				Type:        schema.TypeString,
				Required:    true,
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
				Optional:    true,
				Sensitive:   true,
				Description: "Password of administrative account",
			},
			"auto_domain_account_maintenance": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Automatic Domain Account Maintenance",
			},
			"auto_local_account_maintenance": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Automatic Local Account Maintenance",
			},
			"manual_domain_account_unlock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Manual Domain Account Unlock",
			},
			"manual_local_account_unlock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable Manual Local Account Unlock",
			},
			"provisioning_admin_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provisioning Administrative Account ID (must be managed)",
			},
			"reconciliation_account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Reconciliation account name",
			},
			// Zone Role Workflow menu
			"enable_zonerole_workflow": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable zone role requests for this system",
			},
			"assigned_zonerole":          getZoneRoleSchema(),
			"assigned_zonerole_approver": getWorkflowApproversSchema(),
		},
	}
}

func resourceVaultDomainConfigurationRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Domain Configuration: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a Domain object and populate ID attribute
	object := vault.NewDomain(client)
	object.ID = d.Get("domain_id").(string)
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Domain: %v", err)
	}
	//logger.Debugf("Domain from tenant: %v", object)

	d.Set("administrative_account_id", object.AdminAccountID)
	d.Set("administrator_display_name", object.AdministratorDisplayName)
	d.Set("administrative_account_name", object.AdminAccountName)
	d.Set("auto_domain_account_maintenance", object.AutoDomainAccountMaintenance)
	d.Set("auto_local_account_maintenance", object.AutoLocalAccountMaintenance)
	d.Set("manual_domain_account_unlock", object.ManualDomainAccountUnlock)
	d.Set("manual_local_account_unlock", object.ManualLocalAccountUnlock)
	d.Set("provisioning_admin_id", object.ProvisioningAdminID)
	d.Set("reconciliation_account_name", object.ReconciliationAccountName)
	if object.ZoneRoleWorkflowApprovers != "" {
		wfschema, err := convertWorkflowSchema("{\"WorkflowApprover\":" + object.ZoneRoleWorkflowApprovers + "}")
		if err != nil {
			return err
		}
		d.Set("assigned_zonerole_approver", wfschema)
	}
	if object.ZoneRoleWorkflowRoles != "" {
		zrschema, err := convertZoneRoleSchema("{\"ZoneRoleWorkflowRole\":" + object.ZoneRoleWorkflowRoles + "}")
		if err != nil {
			return err
		}
		d.Set("assigned_zonerole", zrschema)
	}

	logger.Infof("Completed reading Domain Configuration: %s", object.Name)
	return nil
}

func resourceVaultDomainConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain Configuration: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a Domain object and populate all attributes
	object := vault.NewDomain(client)
	object.ID = d.Get("domain_id").(string)
	err := object.Read()
	// If the domain does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Domain: %v", err)
	}

	// Get the rest of attributes
	err = createUpateGetDomainConfigurationData(d, object)
	if err != nil {
		return err
	}

	// set administrative account
	err = object.SetAdminAccount()
	if err != nil {
		return fmt.Errorf("Error setting Domain administrative account: %v", err)
	}

	d.SetPartial("administrative_account_id")
	d.SetId(fmt.Sprintf("%s-configuration", object.ID))

	// Update Reconciliation Options and Unix/Linux Local Accounts settings
	_, err = object.Update()
	if err != nil {
		return fmt.Errorf("Error updating Domain: %v", err)
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Setting of Domain Configuration completed: %s", object.Name)
	return resourceVaultDomainConfigurationRead(d, m)
}

func resourceVaultDomainConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain Configuration update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewDomain(client)
	object.ID = d.Get("domain_id").(string)
	err := object.Read()
	// If the domain does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		//d.SetId("")
		return fmt.Errorf("Error reading Domain: %v", err)
	}

	err = createUpateGetDomainConfigurationData(d, object)
	if err != nil {
		return err
	}

	// Deal with administative account change first otherwise account maintenace options can't be set
	if d.HasChange("administrative_account_id") {
		err := object.SetAdminAccount()
		if err != nil {
			return fmt.Errorf("Error updating Domain administrative account: %v", err)
		}
		d.SetPartial("administrative_account_id")
	}

	if d.HasChanges("auto_domain_account_maintenance", "auto_local_account_maintenance", "manual_domain_account_unlock", "manual_local_account_unlock",
		"provisioning_admin_id", "reconciliation_account_name", "enable_zonerole_workflow", "assigned_zonerole", "assigned_zonerole_approver") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating Domain attribute: %v", err)
		}
		d.SetPartial("auto_domain_account_maintenance")
		d.SetPartial("auto_local_account_maintenance")
		d.SetPartial("manual_domain_account_unlock")
		d.SetPartial("manual_local_account_unlock")
		d.SetPartial("provisioning_admin_id")
		d.SetPartial("reconciliation_account_name")
		d.SetPartial("enable_zonerole_workflow")
		d.SetPartial("assigned_zonerole")
		d.SetPartial("assigned_zonerole_approver")
	}

	d.Partial(false)
	logger.Infof("Updating of Domain Configuration completed: %s", object.Name)
	return resourceVaultDomainConfigurationRead(d, m)
}

func resourceVaultDomainConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning removing of Domain Configuration: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDomain(client)
	object.ID = d.Get("domain_id").(string)
	err := object.Read()
	// If the domain does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Domain: %v", err)
	}
	object.AdminAccountID = ""
	object.ProvisioningAdminID = ""
	object.ReconciliationAccountName = ""
	object.AutoDomainAccountMaintenance = false
	object.AutoLocalAccountMaintenance = false
	object.ManualDomainAccountUnlock = false
	object.ManualLocalAccountUnlock = false
	object.ZoneRoleWorkflowEnabled = false
	object.ZoneRoleWorkflowRoles = ""
	object.ZoneRoleWorkflowApprovers = ""

	err = object.SetAdminAccount()
	if err != nil {
		return fmt.Errorf("Error setting Domain administrative account: %v", err)
	}

	resp, err := object.Update()
	if err != nil || !resp.Success {
		return fmt.Errorf("Error removing of Domain Configuration: %v", err)
	}

	d.SetId("")

	logger.Infof("Removing of Domain Configuration completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetDomainConfigurationData(d *schema.ResourceData, object *vault.Domain) error {
	if v, ok := d.GetOk("administrative_account_id"); ok {
		object.AdminAccountID = v.(string)
	}
	if v, ok := d.GetOk("administrative_account_name"); ok {
		object.AdminAccountName = v.(string)
	}
	if v, ok := d.GetOk("administrative_account_password"); ok {
		object.AdminAccountPassword = v.(string)
	}

	object.AutoDomainAccountMaintenance = d.Get("auto_domain_account_maintenance").(bool)
	object.AutoLocalAccountMaintenance = d.Get("auto_local_account_maintenance").(bool)
	object.ManualDomainAccountUnlock = d.Get("manual_domain_account_unlock").(bool)
	object.ManualLocalAccountUnlock = d.Get("manual_local_account_unlock").(bool)

	if v, ok := d.GetOk("provisioning_admin_id"); ok {
		object.ProvisioningAdminID = v.(string)
	}
	if v, ok := d.GetOk("reconciliation_account_name"); ok {
		object.ReconciliationAccountName = v.(string)
	}
	// Zone Role Workflow
	if v, ok := d.GetOk("enable_zonerole_workflow"); ok {
		object.ZoneRoleWorkflowEnabled = v.(bool)
	}
	if v, ok := d.GetOk("assigned_zonerole"); ok {
		object.ZoneRoleWorkflowRoleList = expandZoneRoles(v)
	}
	if v, ok := d.GetOk("assigned_zonerole_approver"); ok {
		object.ZoneRoleWorkflowApproverList = expandWorkflowApprovers(v.([]interface{})) // This is a slice
	}

	return nil
}
