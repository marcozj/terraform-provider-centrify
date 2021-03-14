package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceVaultDomainReconciliation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultDomainReconciliationCreate,
		Read:   resourceVaultDomainReconciliationRead,
		Update: resourceVaultDomainReconciliationUpdate,
		Delete: resourceVaultDomainReconciliationDelete,

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
		},
	}
}

func resourceVaultDomainReconciliationRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Domain reconciliation settings: %s", ResourceIDString(d))
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

	logger.Infof("Completed reading Domain reconciliation settings: %s", object.Name)
	return nil
}

func resourceVaultDomainReconciliationCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain reconciliation settings: %s", ResourceIDString(d))

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
	err = createUpateGetDomainReconciliationData(d, object)
	if err != nil {
		return err
	}

	// set administrative account
	err = object.SetAdminAccount()
	if err != nil {
		return fmt.Errorf("Error setting Domain administrative account: %v", err)
	}

	d.SetPartial("administrative_account_id")
	d.SetId(fmt.Sprintf("%s-reconciliation", object.ID))

	// Update Reconciliation Options and Unix/Linux Local Accounts settings
	_, err = object.Update()
	if err != nil {
		return fmt.Errorf("Error updating Domain: %v", err)
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Setting of Domain reconciliation completed: %s", object.Name)
	return resourceVaultDomainReconciliationRead(d, m)
}

func resourceVaultDomainReconciliationUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain reconciliation update: %s", ResourceIDString(d))

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

	err = createUpateGetDomainReconciliationData(d, object)
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
		"provisioning_admin_id", "reconciliation_account_name") {
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
	}

	d.Partial(false)
	logger.Infof("Updating of Domain reconciliation completed: %s", object.Name)
	return resourceVaultDomainReconciliationRead(d, m)
}

func resourceVaultDomainReconciliationDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning removing of Domain reconciliation settings: %s", ResourceIDString(d))
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

	err = object.SetAdminAccount()
	if err != nil {
		return fmt.Errorf("Error setting Domain administrative account: %v", err)
	}

	resp, err := object.Update()
	if err != nil || !resp.Success {
		return fmt.Errorf("Error removing of Domain reconciliation settings: %v", err)
	}

	d.SetId("")

	logger.Infof("Removing of Domain reconciliation setting completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetDomainReconciliationData(d *schema.ResourceData, object *vault.Domain) error {
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

	return nil
}
