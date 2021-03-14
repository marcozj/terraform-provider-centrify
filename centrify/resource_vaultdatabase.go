package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/databaseclass"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceVaultDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultDatabaseCreate,
		Read:   resourceVaultDatabaseRead,
		Update: resourceVaultDatabaseUpdate,
		Delete: resourceVaultDatabaseDelete,
		Exists: resourceVaultDatabaseExists,

		Schema: map[string]*schema.Schema{
			// Database -> Settings menu related settings
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Database",
			},
			"hostname": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Hostname or IP address of the Database",
				ValidateFunc: validation.NoZeroValues,
			},
			"database_class": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the Database",
				ValidateFunc: validation.StringInSlice([]string{
					databaseclass.SQLServer.String(),
					databaseclass.Oracle.String(),
					databaseclass.SAPASE.String(),
				}, false),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the Database",
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Port that used to connect to the Database",
				ValidateFunc: validation.IsPortNumber,
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name of MS SQL Database",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service name of Oracle database",
			},
			"skip_reachability_test": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Verify Database Settings",
			},
			// Database -> Policy menu related settings
			"checkout_lifetime": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Specifies the number of minutes that a checked out password is valid.",
				ValidateFunc: validation.IntBetween(15, 2147483647),
			},
			// Database -> Advanced menu related settings
			"allow_multiple_checkouts": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Allow multiple password checkouts for this database",
			},
			"enable_password_rotation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable periodic password rotation",
			},
			"password_rotate_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Password rotation interval (days)",
				ValidateFunc: validation.IntBetween(1, 2147483647),
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
				Type:     schema.TypeString,
				Optional: true,
				//Computed:    true, // we want to remove this setting if it is not set so do not set to computed
				Description: "Password complexity profile id",
			},
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
			// Database -> Connectors menu related settings
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
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Sets",
			},
			"permission": getPermissionSchema(),
		},
	}
}

func resourceVaultDatabaseExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking VaultDatabase exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDatabase(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("VaultDatabase exists in tenant: %s", object.ID)
	return true, nil
}

func resourceVaultDatabaseRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading VaultDatabase: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a VaultDatabase object and populate ID attribute
	object := vault.NewDatabase(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VaultDatabase: %v", err)
	}
	//logger.Debugf("VaultDatabase from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceVaultDatabaseRead(): %+v", schemamap)
	for k, v := range schemamap {
		if k == "connector_list" {
			// Convert "value1,value1" to schema.TypeSet
			d.Set("connector_list", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		} else {
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading VaultDatabase: %s", object.Name)
	return nil
}

func resourceVaultDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultDatabase creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a Database object and populate all attributes
	object := vault.NewDatabase(client)
	err := createUpateGetVaultDatabaseData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating VaultDatabase: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("VaultDatabase ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("name")
	d.SetPartial("hostname")
	d.SetPartial("database_class")
	d.SetPartial("description")

	// 2nd step to update VaultDatabase login profile
	// Create API call doesn't set VaultDatabase login profile so need to run update again
	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("Error updating VaultDatabase attribute: %v", err2)
	}
	d.SetPartial("password_profile_id")

	// 3rd step to add VaultDatabase to Sets
	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
		d.SetPartial("sets")
	}

	// 4th step to add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting VaultDatabase permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of VaultDatabase completed: %s", object.Name)
	return resourceVaultDatabaseRead(d, m)
}

func resourceVaultDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultDatabase update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewDatabase(client)

	object.ID = d.Id()
	err := createUpateGetVaultDatabaseData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "hostname", "description", "port", "database_class", "checkout_lifetime", "allow_multiple_checkouts",
		"enable_password_rotation", "password_rotate_interval", "enable_password_rotation_after_checkin", "minimum_password_age", "password_profile_id",
		"enable_password_history_cleanup", "password_historycleanup_duration",
		"choose_connector", "connector_list") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultDatabase attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
		d.SetPartial("name")
		d.SetPartial("hostname")
		d.SetPartial("database_class")
		d.SetPartial("description")
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
				return fmt.Errorf("Error removing VaultDatabase from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding VaultDatabase to Set: %v", err)
			}
		}
		d.SetPartial("sets")
	}

	// Deal with Permissions
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
				return fmt.Errorf("Error removing VaultDatabase permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding VaultDatabase permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	d.Partial(false)
	logger.Infof("Updating of VaultDatabase completed: %s", object.Name)
	return resourceVaultDatabaseRead(d, m)
}

func resourceVaultDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of VaultDatabase: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDatabase(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting VaultDatabase: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of VaultDatabase completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetVaultDatabaseData(d *schema.ResourceData, object *vault.Database) error {
	// Database -> Settings menu related settings
	object.Name = d.Get("name").(string)
	object.FQDN = d.Get("hostname").(string)
	object.DatabaseClass = d.Get("database_class").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("port"); ok {
		object.Port = v.(int)
	}
	if v, ok := d.GetOk("instance_name"); ok {
		object.InstanceName = v.(string)
	}
	if v, ok := d.GetOk("service_name"); ok {
		object.ServiceName = v.(string)
	}
	if v, ok := d.GetOk("skip_reachability_test"); ok {
		object.SkipReachabilityTest = v.(bool)
	}
	// Database -> Policy menu related settings
	if v, ok := d.GetOk("checkout_lifetime"); ok {
		object.DefaultCheckoutTime = v.(int)
	}
	// Database -> Advanced menu related settings
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
	if v, ok := d.GetOk("enable_password_history_cleanup"); ok {
		object.AllowPasswordHistoryCleanUp = v.(bool)
	}
	if v, ok := d.GetOk("password_historycleanup_duration"); ok {
		object.PasswordHistoryCleanUpDuration = v.(int)
	}
	// Database -> Connectors menu related settings
	if v, ok := d.GetOk("connector_list"); ok {
		object.ProxyCollectionList = flattenSchemaSetToString(v.(*schema.Set))
	}
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
	// Verify database type
	if object.DatabaseClass == "SQLServer" && object.InstanceName == "" {
		return fmt.Errorf("instance_name must be provided for SQLServer database type")
	}
	if object.DatabaseClass == "Oracle" && object.ServiceName == "" {
		return fmt.Errorf("service_name must be provided for Oracle database type")
	}

	return nil
}
