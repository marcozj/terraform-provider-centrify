package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/cloudprovidertype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceCloudProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudProviderCreate,
		Read:   resourceCloudProviderRead,
		Update: resourceCloudProviderUpdate,
		Delete: resourceCloudProviderDelete,
		Exists: resourceCloudProviderExists,

		Schema: map[string]*schema.Schema{
			"cloud_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account ID of the cloud provider",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud provider",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the cloud provider",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the cloud provider",
				ValidateFunc: validation.StringInSlice([]string{
					cloudprovidertype.AWS.String(),
				}, false),
			},
			"enable_interactive_password_rotation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable interactive password rotation",
			},
			"prompt_change_root_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Prompt to change root password every login and password checkin",
			},
			"enable_password_rotation_reminders": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable password rotation reminders",
			},
			"password_rotation_reminder_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Minimum number of days since last rotation to trigger a reminder",
				ValidateFunc: validation.IntBetween(1, 2147483647),
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default Root Account Login Profile (used if no conditions matched)",
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
			"permission":     getPermissionSchema(),
			"challenge_rule": getChallengeRulesSchema(),
		},
	}
}

func resourceCloudProviderExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking CloudProvider exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewCloudProvider(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("CloudProvider exists in tenant: %s", object.ID)
	return true, nil
}

func resourceCloudProviderRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading CloudProvider: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a System object and populate ID attribute
	object := vault.NewCloudProvider(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading System: %v", err)
	}
	//logger.Debugf("System from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceCloudProviderRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading CloudProvider: %s", object.Name)
	return nil
}

func resourceCloudProviderCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning CloudProvider creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a CloudProvider object and populate all attributes
	object := vault.NewCloudProvider(client)
	err := createUpateGetCloudProviderData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating CloudProvider: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("CloudProvider ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("name")
	d.SetPartial("cloud_account_id")
	d.SetPartial("description")

	// 2nd step to update CloudProvider login profile
	// Create API call doesn't set CloudProvider login profile so need to run update again
	if object.LoginDefaultProfile != "" {
		logger.Debugf("Update login profile for CloudProvider creation: %s", ResourceIDString(d))
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating System attribute: %v", err)
		}
		d.SetPartial("default_profile_id")
	}

	// 3rd step to add CloudProvider to Sets
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
			return fmt.Errorf("Error setting CloudProvider permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of CloudProvider completed: %s", object.Name)
	return resourceCloudProviderRead(d, m)
}

func resourceCloudProviderUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning CloudProvider update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewCloudProvider(client)

	object.ID = d.Id()
	err := createUpateGetCloudProviderData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "cloud_account_id", "description", "enable_interactive_password_rotation", "prompt_change_root_password",
		"enable_password_rotation_reminders", "password_rotation_reminder_duration", "default_profile_id", "challenge_rule") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating CloudProvider attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
		d.SetPartial("name")
		d.SetPartial("cloud_account_id")
		d.SetPartial("description")
		d.SetPartial("enable_interactive_password_rotation")
		d.SetPartial("prompt_change_root_password")
		d.SetPartial("enable_password_rotation_reminders")
		d.SetPartial("password_rotation_reminder_duration")
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
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
				return fmt.Errorf("Error removing CloudProvider from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding CloudProvider to Set: %v", err)
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
				return fmt.Errorf("Error removing CloudProvider permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding CloudProvider permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	d.Partial(false)
	logger.Infof("Updating of CloudProvider completed: %s", object.Name)
	return resourceCloudProviderRead(d, m)
}

func resourceCloudProviderDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of CloudProvider: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewCloudProvider(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting CloudProvider: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of CloudProvider completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetCloudProviderData(d *schema.ResourceData, object *vault.CloudProvider) error {
	// System -> Settings menu related settings
	object.Name = d.Get("name").(string)
	object.CloudAccountID = d.Get("cloud_account_id").(string)
	object.Type = d.Get("type").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("enable_interactive_password_rotation"); ok {
		object.EnableUnmanagedPasswordRotation = v.(bool)
	}
	if v, ok := d.GetOk("prompt_change_root_password"); ok {
		object.EnableUnmanagedPasswordRotationPrompt = v.(bool)
	}
	if v, ok := d.GetOk("enable_password_rotation_reminders"); ok {
		object.EnableUnmanagedPasswordRotationReminder = v.(bool)
	}
	if v, ok := d.GetOk("password_rotation_reminder_duration"); ok {
		object.UnmanagedPasswordRotationReminderDuration = v.(int)
	}
	if v, ok := d.GetOk("default_profile_id"); ok {
		object.LoginDefaultProfile = v.(string)
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
	// Challenge rules
	if v, ok := d.GetOk("challenge_rule"); ok {
		object.ChallengeRules = expandChallengeRules(v.([]interface{}))
		// Perform validations
		if err := validateChallengeRules(object.ChallengeRules); err != nil {
			return fmt.Errorf("Schema setting error: %s", err)
		}
	}

	return nil
}
