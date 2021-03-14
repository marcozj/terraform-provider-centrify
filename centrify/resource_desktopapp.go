package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/desktopapp/applicationtemplate"
	"github.com/marcozj/golang-sdk/enum/desktopapp/cmdparamtype"
	"github.com/marcozj/golang-sdk/enum/desktopapp/logincredential"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceDesktopApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceDesktopAppCreate,
		Read:   resourceDesktopAppRead,
		Update: resourceDesktopAppUpdate,
		Delete: resourceDesktopAppDelete,
		Exists: resourceDesktopAppExists,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Template name of the Desktop App",
				ValidateFunc: validation.StringInSlice([]string{
					applicationtemplate.Generic.String(),
					applicationtemplate.SQLServerManagementStudio.String(),
					applicationtemplate.Toad.String(),
					applicationtemplate.VSphereClient.String(),
				}, false),
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Desktop App",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the Desktop App",
			},
			"application_host_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application host",
			},
			"login_credential_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host login credential type",
				ValidateFunc: validation.StringInSlice([]string{
					logincredential.UserADCredential.String(),
					logincredential.PromptForCredential.String(),
					logincredential.SelectAlternativeAccount.String(),
					logincredential.SharedAccount.String(),
				}, false),
			},
			"application_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host login credential account",
			},
			"application_alias": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The alias name of the published RemoteApp program",
			},
			"command_line": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Command line",
			},
			"command_parameter": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      customCommandParamHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the parameter",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the parameter",
							ValidateFunc: validation.StringInSlice([]string{
								cmdparamtype.Integer.String(),
								cmdparamtype.Date.String(),
								cmdparamtype.String.String(),
								cmdparamtype.Account.String(),
								cmdparamtype.CloudProivder.String(),
								cmdparamtype.Database.String(),
								cmdparamtype.Device.String(),
								cmdparamtype.Domain.String(),
								cmdparamtype.ResourceProfile.String(),
								cmdparamtype.Role.String(),
								cmdparamtype.Secret.String(),
								cmdparamtype.Service.String(),
								cmdparamtype.SSHKey.String(),
								cmdparamtype.System.String(),
								cmdparamtype.User.String(),
							}, false),
						},
						"target_object_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of selected parameter value",
						},
					},
				},
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AlwaysAllowed", // It must to be "--", "AlwaysAllowed", "-1" or UUID of authen profile
				Description: "Default authentication profile ID",
			},
			"sets": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Sets",
			},
			"permission":     getPermissionSchema(),
			"challenge_rule": getChallengeRulesSchema(),
			"policy_script": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"challenge_rule"},
				Description:   "Use script to specify authentication rules (configured rules are ignored)",
			},
		},
	}
}

func resourceDesktopAppExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking DesktopApp exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDesktopApp(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("DesktopApp exists in tenant: %s", object.ID)
	return true, nil
}

func resourceDesktopAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading DesktopApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewVaultSecret object and populate ID attribute
	object := vault.NewDesktopApp(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading DesktopApp: %v", err)
	}
	//logger.Debugf("DesktopApp from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceDesktopAppRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading DesktopApp: %s", object.Name)
	return nil
}

func resourceDesktopAppCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning DesktopApp creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a DesktopApp object
	object := vault.NewDesktopApp(client)
	object.TemplateName = d.Get("template_name").(string)
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating DesktopApp: %v", err)
	}
	if len(resp.Result) <= 0 {
		return fmt.Errorf("Import application template returns incorrect result")
	}

	id := resp.Result[0].(map[string]interface{})["_RowKey"].(string)

	if id == "" {
		return fmt.Errorf("DesktopApp ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Update attributes to complete creation
	err = getUpateGetDesktopAppData(d, object)
	if err != nil {
		return err
	}

	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("Error updating DesktopApp attribute: %v", err2)
	}

	d.SetPartial("name")
	d.SetPartial("template_name")
	d.SetPartial("description")
	d.SetPartial("application_host_id")
	d.SetPartial("login_credential_type")
	d.SetPartial("application_account_id")
	d.SetPartial("application_alias")
	d.SetPartial("command_line")
	d.SetPartial("command_parameter")
	d.SetPartial("default_profile_id")
	d.SetPartial("challenge_rule")
	d.SetPartial("policy_script")

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
			return fmt.Errorf("Error setting DesktopApp permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of DesktopApp completed: %s", object.Name)
	return resourceDesktopAppRead(d, m)
}

func resourceDesktopAppUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning DesktopApp update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewDesktopApp(client)
	object.ID = d.Id()
	err := getUpateGetDesktopAppData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "template_name", "description", "application_host_id", "login_credential_type", "application_account_id", "application_alias",
		"command_line", "command_parameter", "default_profile_id", "challenge_rule", "policy_script") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating DesktopApp attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("name")
		d.SetPartial("template_name")
		d.SetPartial("description")
		d.SetPartial("application_host_id")
		d.SetPartial("login_credential_type")
		d.SetPartial("application_account_id")
		d.SetPartial("application_alias")
		d.SetPartial("command_line")
		d.SetPartial("command_parameter")
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
		d.SetPartial("policy_script")
	}

	if d.HasChange("sets") {
		old, new := d.GetChange("sets")
		// Remove old Sets
		for _, v := range flattenSchemaSetToStringSlice(old) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "remove")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error removing DesktopApp from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding DesktopApp to Set: %v", err)
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
				return fmt.Errorf("Error removing DesktopApp permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding DesktopApp permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of DesktopApp completed: %s", object.Name)
	return resourceDesktopAppRead(d, m)
}

func resourceDesktopAppDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of DesktopApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewDesktopApp(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting DesktopApp: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of DesktopApp completed: %s", ResourceIDString(d))
	return nil
}

func getUpateGetDesktopAppData(d *schema.ResourceData, object *vault.DesktopApp) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("application_host_id"); ok {
		object.DesktopAppRunHostID = v.(string)
	}
	if v, ok := d.GetOk("login_credential_type"); ok {
		object.DesktopAppRunAccountType = v.(string)
	}
	if v, ok := d.GetOk("application_account_id"); ok {
		object.DesktopAppRunAccountID = v.(string)
	}
	if v, ok := d.GetOk("application_alias"); ok {
		object.DesktopAppProgramName = v.(string)
	}
	if v, ok := d.GetOk("command_line"); ok {
		object.DesktopAppCmdline = v.(string)
	}
	if v, ok := d.GetOk("command_parameter"); ok {
		object.DesktopAppParams = expandCommandParams(v)
	}
	if v, ok := d.GetOk("default_profile_id"); ok {
		object.DefaultAuthProfile = v.(string)
	}
	if v, ok := d.GetOk("policy_script"); ok {
		object.PolicyScript = v.(string)
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
