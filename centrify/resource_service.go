package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/servicetype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,
		Exists: resourceServiceExists,

		Schema: map[string]*schema.Schema{
			"system_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The target system id where the service runs",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the service",
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					servicetype.WindowsService.String(),
					servicetype.ScheduledTask.String(),
					servicetype.IISApplicationPool.String(),
				}, false),
			},
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the service to be managed.",
			},
			"enable_management": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable management of this service password",
			},
			"admin_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Administrative account id that used to manage the password for the service",
			},
			"multiplexed_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The multiplexed account id to run the service",
			},
			"restart_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restart Service when password is rotated",
			},
			"restart_time_restriction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enforce restart time restrictions",
			},
			"days_of_week": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Day of the week restart allowed",
			},
			"restart_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Start time of the time range restart is allowed",
			},
			"restart_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "End time of the time range restart is allowed",
			},
			"use_utc_time": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to use UTC time",
			},
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

func resourceServiceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking service exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewService(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Service exists in tenant: %s", object.ID)
	return true, nil
}

func resourceServiceRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading service: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewService object and populate ID attribute
	object := vault.NewService(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading service: %v", err)
	}
	//logger.Debugf("Service from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceServiceRead(): %+v", schemamap)
	for k, v := range schemamap {
		if k == "days_of_week" {
			// Convert "value1,value1" to schema.TypeSet
			d.Set("days_of_week", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		} else {
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading service: %s", object.Name)
	return nil
}

func resourceServiceCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning service creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a service object and populate all attributes
	object := vault.NewService(client)
	err := createUpateGetServiceData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating service: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("Service ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("service_name")
	d.SetPartial("description")
	d.SetPartial("system_id")
	d.SetPartial("service_type")
	d.SetPartial("service_name")
	d.SetPartial("enable_management")
	d.SetPartial("admin_account_id")
	d.SetPartial("multiplexed_account_id")
	d.SetPartial("restart_service")
	d.SetPartial("restart_time_restriction")
	d.SetPartial("days_of_week")
	d.SetPartial("restart_start_time")
	d.SetPartial("restart_end_time")
	d.SetPartial("use_utc_time")

	// 2nd step to add service to Sets
	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
		d.SetPartial("sets")
	}

	// add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting service permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of service completed: %s", object.Name)
	return resourceServiceRead(d, m)
}

func resourceServiceUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning service update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewService(client)
	object.ID = d.Id()
	err := createUpateGetServiceData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("service_name", "description", "system_id", "service_type", "enable_management", "admin_account_id", "multiplexed_account_id",
		"restart_service", "restart_time_restriction", "days_of_week", "restart_start_time", "restart_end_time", "use_utc_time") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating service attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("description")
		d.SetPartial("system_id")
		d.SetPartial("service_type")
		d.SetPartial("service_name")
		d.SetPartial("enable_management")
		d.SetPartial("admin_account_id")
		d.SetPartial("multiplexed_account_id")
		d.SetPartial("restart_service")
		d.SetPartial("restart_time_restriction")
		d.SetPartial("days_of_week")
		d.SetPartial("restart_start_time")
		d.SetPartial("restart_end_time")
		d.SetPartial("use_utc_time")
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
				return fmt.Errorf("Error removing Service from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding Service to Set: %v", err)
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
				return fmt.Errorf("Error removing service permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding servicepermissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of service completed: %s", object.Name)
	return resourceServiceRead(d, m)
}

func resourceServiceDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of service: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewService(client)
	object.ID = d.Id()

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting service: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of service completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetServiceData(d *schema.ResourceData, object *vault.Service) error {
	object.Name = d.Get("service_name").(string)
	object.SystemID = d.Get("system_id").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	object.ServiceType = d.Get("service_type").(string)
	if v, ok := d.GetOk("enable_management"); ok {
		object.EnableManagement = v.(bool)
	}
	if v, ok := d.GetOk("admin_account_id"); ok {
		object.AdminAccountID = v.(string)
	}
	if v, ok := d.GetOk("multiplexed_account_id"); ok {
		object.MultiplexedAccountID = v.(string)
	}
	if v, ok := d.GetOk("restart_service"); ok {
		object.RestartService = v.(bool)
	}
	if v, ok := d.GetOk("restart_time_restriction"); ok {
		object.RestartTimeRestriction = v.(bool)
	}
	if v, ok := d.GetOk("days_of_week"); ok {
		object.DaysOfWeek = flattenSchemaSetToString(v.(*schema.Set))
	}
	if v, ok := d.GetOk("restart_start_time"); ok {
		object.RestartStartTime = v.(string)
	}
	if v, ok := d.GetOk("restart_end_time"); ok {
		object.RestartEndTime = v.(string)
	}
	if v, ok := d.GetOk("use_utc_time"); ok {
		object.UseUTCTime = v.(bool)
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
