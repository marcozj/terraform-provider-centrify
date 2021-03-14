package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceVaultDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultDomainCreate,
		Read:   resourceVaultDomainRead,
		Update: resourceVaultDomainUpdate,
		Delete: resourceVaultDomainDelete,
		Exists: resourceVaultDomainExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// Settings menu related settings
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The domain name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the domain",
			},
			"verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to verify the Domain upon creation",
			},
			// Policy menu related settings
			"checkout_lifetime": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Checkout lifetime (minutes)",
				ValidateFunc: validation.IntBetween(15, 2147483647),
			},
			/*
				// Advanced menu -> Administrative Account Settings
				"administrative_account_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "ID of administrative account",
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
					Description: "Enable Automatic Domain Account Maintenance",
				},
				"auto_local_account_maintenance": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable Automatic Local Account Maintenance",
				},
				"manual_domain_account_unlock": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable Manual Domain Account Unlock",
				},
				"manual_local_account_unlock": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable Manual Local Account Unlock",
				},
			*/
			// Advanced -> Security Settings
			"allow_multiple_checkouts": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Allow multiple password checkouts per AD account added for this domain",
			},
			"enable_password_rotation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable periodic password rotation",
			},
			"password_rotate_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Password rotation interval (days)",
			},
			"enable_password_rotation_after_checkin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable password rotation after checkin",
			},
			"minimum_password_age": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Minimum Password Age (days)",
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"password_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password complexity profile id",
			},
			// Advanced -> Maintenance Settings
			"enable_password_history_cleanup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable periodic password history cleanup",
			},
			"password_historycleanup_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Password history cleanup (days)",
				ValidateFunc: validation.IntBetween(90, 2147483647),
			},
			// Advanced -> Domain/Zone Tasks
			"enable_zone_joined_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable periodic domain/zone joined check",
			},
			"zone_joined_check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1440,
				Description:  "Domain/zone joined check interval (minutes)",
				ValidateFunc: validation.IntBetween(1, 2147483647),
			},
			"enable_zone_role_cleanup": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable periodic removal of expired zone role assignments",
			},
			"zone_role_cleanup_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      6,
				Description:  "Expired zone role assignment removal interval (hours)",
				ValidateFunc: validation.IntBetween(1, 2147483647),
			},
			// System -> Connectors menu related settings
			"connector_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of Connectors",
			},
			// Add to Sets
			"sets": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				Set: schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Sets",
			},
			"permission": getPermissionSchema(),
		},
	}
}

func resourceVaultDomainExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Domain exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDomain(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Domain exists in tenant: %s", object.ID)
	return true, nil
}

func resourceVaultDomainRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Domain: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a Domain object and populate ID attribute
	object := vault.NewDomain(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Domain: %v", err)
	}
	//logger.Debugf("Domain from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceVaultDomainRead(): %+v", schemamap)
	for k, v := range schemamap {
		if k == "connector_list" {
			// Convert "value1,value1" to schema.TypeSet
			d.Set("connector_list", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		} else {
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading Domain: %s", object.Name)
	return nil
}

func resourceVaultDomainCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a Domain object and populate all attributes
	object := vault.NewDomain(client)
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("verify"); ok {
		object.VerifyDomain = v.(bool)
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating Domain: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("Domain ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("name")
	d.SetPartial("description")

	// Get the rest of attributes
	err = createUpateGetDomainData(d, object)
	if err != nil {
		return err
	}
	/*
		// 2nd step, set administrative account
		if object.AdminAccountID != "" {
			err := object.SetAdminAccount()
			if err != nil {
				return fmt.Errorf("Error setting Domain administrative account: %v", err)
			}
		}
	*/
	// 3nd step, update domain after creation
	_, err = object.Update()
	if err != nil {
		return fmt.Errorf("Error updating Domain: %v", err)
	}

	// 4rd step to add system to Sets
	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
		d.SetPartial("sets")
	}

	// 5th step to add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting Domain permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Domain completed: %s", object.Name)
	return resourceVaultDomainRead(d, m)
}

func resourceVaultDomainUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Domain update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewDomain(client)

	object.ID = d.Id()
	err := createUpateGetDomainData(d, object)
	if err != nil {
		return err
	}

	// Deal with Permissions first. To set admin account, AddAccount permission is required
	if d.HasChange("permission") {
		old, new := d.GetChange("permission")
		// We don't want to care the details of changes
		// So, let's first remove the old permissions
		var err error
		if old != nil {
			// do not validate old values
			object.Permissions, err = expandPermissions(old, object.ValidPermissions, false)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(true)
			if err != nil {
				return fmt.Errorf("Error removing Domain permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding Domain permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}
	/*
		// Deal with administative account change first otherwise account maintenace options can't be set
		if d.HasChange("administrative_account_id") {
			err := object.SetAdminAccount()
			if err != nil {
				return fmt.Errorf("Error updating Domain administrative account: %v", err)
			}
			d.SetPartial("administrative_account_id")
		}
	*/
	// Deal with normal attribute changes first
	if d.HasChanges("name", "description", "checkout_lifetime", "allow_multiple_checkouts", "enable_password_rotation", "password_rotate_interval",
		"enable_password_rotation_after_checkin", "minimum_password_age", "password_profile_id", "enable_password_history_cleanup",
		"password_historycleanup_duration", "enable_zone_joined_check", "zone_joined_check_interval", "enable_zone_role_cleanup",
		"zone_role_cleanup_interval", "connector_list") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating Domain attribute: %v", err)
		}
		//logger.Debugf("Updated attributes to: %+v", object)
		d.SetPartial("name")
		d.SetPartial("description")
		d.SetPartial("checkout_lifetime")
		d.SetPartial("allow_multiple_checkouts")
		d.SetPartial("enable_password_rotation")
		d.SetPartial("password_rotate_interval")
		d.SetPartial("enable_password_rotation_after_checkin")
		d.SetPartial("minimum_password_age")
		d.SetPartial("password_profile_id")
		d.SetPartial("enable_password_history_cleanup")
		d.SetPartial("password_historycleanup_duration")
		d.SetPartial("enable_zone_joined_check")
		d.SetPartial("zone_joined_check_interval")
		d.SetPartial("enable_zone_role_cleanup")
		d.SetPartial("zone_role_cleanup_interval")
		d.SetPartial("connector_list")
	}

	// Deal with Set member
	if d.HasChange("sets") {
		old, new := d.GetChange("sets")
		// Remove old Sets
		for _, v := range flattenSchemaSetToStringSlice(old) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "remove")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error removing System from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding System to Set: %v", err)
			}
		}
		d.SetPartial("sets")
	}

	d.Partial(false)
	logger.Infof("Updating of Domain completed: %s", object.Name)
	return resourceVaultDomainRead(d, m)
}

func resourceVaultDomainDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Domain: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDomain(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting Domain: %v", err)
	}

	if resp != nil && resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Domain completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetDomainData(d *schema.ResourceData, object *vault.Domain) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	// Policy menu related settings
	if v, ok := d.GetOk("checkout_lifetime"); ok {
		object.DefaultCheckoutTime = v.(int)
	}
	/*
		// Advanced menu -> Administrative Account Settings
		if v, ok := d.GetOk("administrative_account_id"); ok {
			object.AdminAccountID = v.(string)
		}

		if v, ok := d.GetOk("administrative_account_name"); ok {
			object.AdminAccountName = v.(string)
		}
		if v, ok := d.GetOk("administrative_account_domain"); ok {
			object.AdminAccountDomain = v.(string)
		}
		if v, ok := d.GetOk("administrative_account_password"); ok {
			object.AdminAccountPassword = v.(string)
		}

		if v, ok := d.GetOk("auto_domain_account_maintenance"); ok {
			object.AutoDomainAccountMaintenance = v.(bool)
		}
		if v, ok := d.GetOk("auto_local_account_maintenance"); ok {
			object.AutoLocalAccountMaintenance = v.(bool)
		}
		if v, ok := d.GetOk("manual_domain_account_unlock"); ok {
			object.ManualDomainAccountUnlock = v.(bool)
		}
		if v, ok := d.GetOk("manual_local_account_unlock"); ok {
			object.ManualLocalAccountUnlock = v.(bool)
		}
	*/
	// Advanced -> Security Settings
	if v, ok := d.GetOk("allow_multiple_checkouts"); ok {
		object.AllowMultipleCheckouts = v.(bool)
	}
	if v, ok := d.GetOk("enable_password_rotation"); ok {
		object.AllowPasswordRotation = v.(bool)
	}
	if v, ok := d.GetOk("password_rotate_interval"); ok {
		object.PasswordRotateDuration = v.(int)
	}
	if v, ok := d.GetOk("enable_password_rotation_after_checkin"); ok {
		object.AllowPasswordRotationAfterCheckin = v.(bool)
	}
	if v, ok := d.GetOk("minimum_password_age"); ok {
		object.MinimumPasswordAge = v.(int)
	}
	if v, ok := d.GetOk("password_profile_id"); ok {
		object.PasswordProfileID = v.(string)
	}
	// Advanced -> Maintenance Settings
	if v, ok := d.GetOk("enable_password_history_cleanup"); ok {
		object.AllowPasswordHistoryCleanUp = v.(bool)
	}
	if v, ok := d.GetOk("password_historycleanup_duration"); ok {
		object.PasswordHistoryCleanUpDuration = v.(int)
	}
	// Advanced -> Domain/Zone Tasks
	if v, ok := d.GetOk("enable_zone_joined_check"); ok {
		object.AllowRefreshZoneJoined = v.(bool)
	}
	if v, ok := d.GetOk("zone_joined_check_interval"); ok {
		object.RefreshZoneJoinedIntervalMinutes = v.(int)
	}
	if v, ok := d.GetOk("enable_zone_role_cleanup"); ok {
		object.AllowZoneRoleCleanup = v.(bool)
	}
	if v, ok := d.GetOk("zone_role_cleanup_interval"); ok {
		object.ZoneRoleCleanupIntervalHours = v.(int)
	}
	// System -> Connectors menu related settings
	if v, ok := d.GetOk("connector_list"); ok {
		object.ProxyCollectionList = flattenSchemaSetToString(v.(*schema.Set))
	}
	// Sets
	if v, ok := d.GetOk("sets"); ok {
		object.Sets = flattenSchemaSetToStringSlice(v)
	}
	// Permissions
	if v, ok := d.GetOk("permission"); ok {
		var err error
		object.Permissions, err = expandPermissions(v, object.ValidPermissions, true)
		if err != nil {
			return err
		}
	}

	return nil
}
