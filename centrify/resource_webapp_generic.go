package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/webapp/accountmapping"
	"github.com/marcozj/golang-sdk/enum/webapp/generic/applicationtemplate"

	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceGenericWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceGenericWebAppCreate,
		Read:   resourceGenericWebAppRead,
		Update: resourceGenericWebAppUpdate,
		Delete: resourceGenericWebAppDelete,
		Exists: resourceGenericWebAppExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Template name of the Web App",
				ValidateFunc: validation.StringInSlice([]string{
					applicationtemplate.Bookmark.String(),
					applicationtemplate.BrowserExtension.String(),
					applicationtemplate.BrowserExtensionAdvanced.String(),
					applicationtemplate.NTLMBasic.String(),
					applicationtemplate.UserPassword.String(),
				}, false),
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Web App",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the Web App",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URL of the application",
			},
			// For browser extension web app
			"hostname_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host name suffix for the url of the login form, for example, acme.com.",
			},
			"username_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CSS Selector for the user name field in the login form, for example, input#login-username.",
			},
			"password_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CSS Selector for the password field in the login form, for example, input#login-password.",
			},
			"submit_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CSS Selector for the Submit button in the login form, for example, input#login-button. This entry is optional. It is required only if you cannot submit the form by pressing the enter key.",
			},
			"form_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CSS Selector for the form field of the login form, for example, form#loginForm.",
			},
			"additional_login_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CSS Selector for any Additional Login Field required to login besides username and password, such as Company name or Agency ID. For example, the selector could be input#login-company-id. This entry is required only if there is an additional login field besides username and password.",
			},
			"additional_login_field_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value for the Additional Login Field. For example, if there is an additional login field for the company name, enter the company name here. This entry is required if Additional Login Field is set.",
			},
			"selector_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 60000),
				Description:  "Use this field to indicate the number of milliseconds to wait for the expected input selectors to load before timing out on failure. A zero or negative number means no timeout.",
			},
			"order": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Use this field to specify the order of login if it is not username, password and submit.",
			},
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to log the user in to this application",
			},
			// Policy menu
			"default_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AlwaysAllowed", // It must to be "--", "AlwaysAllowed", "-1" or UUID of authen profile
				Description: "Default authentication profile ID",
			},
			// Account Mapping menu
			"username_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ADAttribute",
				Description: "Account mapping",
				ValidateFunc: validation.StringInSlice([]string{
					accountmapping.ADAttribute.String(),
					accountmapping.SharedAccount.String(),
					accountmapping.UseScript.String(),
					accountmapping.SetByUser.String(),
				}, false),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "userprincipalname",
				Description: "All users share the user name. Applicable if 'username_strategy' is 'Fixed'",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password for all user share one name",
			},
			"use_ad_login_pw": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use the login password supplied by the user (Active Directory users only)",
			},
			"use_ad_login_pw_by_script": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use the login password supplied by the user for account mapping script (Active Directory users only)",
			},
			"user_map_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account mapping script",
			},

			// Workflow
			"workflow_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"workflow_approver": getWorkflowApproversSchema(),
			"workflow_settings": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceGenericWebAppExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Generic WebApp exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewGenericWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Generic WebApp exists in tenant: %s", object.ID)
	return true, nil
}

func resourceGenericWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Generic WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewWebpp object and populate ID attribute
	object := vault.NewGenericWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading Generic WebApp: %v", err)
	}
	//logger.Debugf("WebApp from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceGenericWebAppRead(): %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "challenge_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		case "workflow_settings":
			if object.WorkflowEnabled && v.(string) != "" {
				wfschema, err := convertWorkflowSchema(v.(string))
				if err != nil {
					return err
				}
				d.Set("workflow_approver", wfschema)
				d.Set(k, v)
			}
		case "user_map_script":
			// Another annoying thing. UserMapScript can't seem to be removed once it set previously.
			// So, ignore it if UserNameStrategy isn't "UseScript"
			if object.UserNameStrategy == accountmapping.UseScript.String() {
				d.Set(k, v)
			}
		default:
			// Password value from read operation returns encrypted string which is different from clear text string in local state.
			// This causes apply action to update password. So, ignore password attribute
			if k != "password" {
				d.Set(k, v)
			}
		}
	}

	logger.Infof("Completed reading Generic WebApp: %s", object.Name)
	return nil
}

func resourceGenericWebAppCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Generic WebApp creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a WebApp object
	object := vault.NewGenericWebApp(client)
	object.TemplateName = d.Get("template_name").(string)
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating Generic WebApp: %v", err)
	}
	if len(resp.Result) <= 0 {
		return fmt.Errorf("import application template returns incorrect result")
	}

	id := resp.Result[0].(map[string]interface{})["_RowKey"].(string)

	if id == "" {
		return fmt.Errorf("generic WebApp ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Update attributes to complete creation
	err = createUpateGetGenericWebAppData(d, object)
	if err != nil {
		return err
	}

	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("error updating Generic WebApp attribute: %v", err2)
	}

	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
	}

	// add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("error setting Generic WebApp permissions: %v", err)
		}
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Generic WebApp completed: %s", object.Name)
	return resourceGenericWebAppRead(d, m)
}

func resourceGenericWebAppUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Generic WebApp update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewGenericWebApp(client)
	object.ID = d.Id()

	err := createUpateGetGenericWebAppData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "template_name", "description", "url", "hostname_suffix", "username_field", "password_field",
		"submit_field", "form_field", "additional_login_field", "additional_login_field_value", "selector_timeout",
		"order", "script", "default_profile_id", "challenge_rule", "policy_script", "username_strategy", "ad_attribute", "username",
		"use_ad_login_pw", "password", "use_ad_login_pw_by_script", "user_map_script", "workflow_enabled", "workflow_approver") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating Generic WebApp attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
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
				return fmt.Errorf("error removing Generic WebApp from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("error adding Generic WebApp to Set: %v", err)
			}
		}
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
				return fmt.Errorf("error removing Oauth WebApp permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("error adding Generic WebApp permissions: %v", err)
			}
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of Generic WebApp completed: %s", object.Name)
	return resourceGenericWebAppRead(d, m)
}

func resourceGenericWebAppDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Generic WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewGenericWebApp(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("error deleting Generic WebApp: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Generic WebApp completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetGenericWebAppData(d *schema.ResourceData, object *vault.GenericWebApp) error {
	object.Name = d.Get("name").(string)
	object.TemplateName = d.Get("template_name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("url"); ok {
		object.Url = v.(string)
	}

	if d.HasChanges("hostname_suffix", "username_field", "password_field", "submit_field", "form_field",
		"additional_login_field", "additional_login_field_value", "selector_timeout", "order", "script") {
		err := object.ResetAppScript()
		if err != nil {
			return err
		}
	}

	if v, ok := d.GetOk("hostname_suffix"); ok {
		object.HostNameSuffix = v.(string)
	}
	if v, ok := d.GetOk("username_field"); ok {
		object.UsernameField = v.(string)
	}
	if v, ok := d.GetOk("password_field"); ok {
		object.PasswordField = v.(string)
	}
	if v, ok := d.GetOk("submit_field"); ok {
		object.SubmitField = v.(string)
	}
	if v, ok := d.GetOk("form_field"); ok {
		object.FormField = v.(string)
	}
	if v, ok := d.GetOk("additional_login_field"); ok {
		object.CorpIdField = v.(string)
	}
	if v, ok := d.GetOk("additional_login_field_value"); ok {
		object.CorpIdentifier = v.(string)
	}
	if v, ok := d.GetOk("selector_timeout"); ok {
		object.SelectorTimeout = v.(int)
	}
	if v, ok := d.GetOk("order"); ok {
		object.Order = v.(string)
	}

	if v, ok := d.GetOk("script"); ok {
		object.Script = v.(string)
	}
	/*
		// Specially handling for script
		// If "script" is changed from whatever to empty, call resetappscript api
		if d.HasChange("script") {
			_, new := d.GetChange("script")
			if new == "" {
				err := object.ResetAppScript()
				if err != nil {
					return err
				}
			}
		}
	*/
	// Account mapping
	if v, ok := d.GetOk("username_strategy"); ok {
		object.UserNameStrategy = v.(string)
	}
	if v, ok := d.GetOk("username"); ok {
		object.Username = v.(string)
	}
	if v, ok := d.GetOk("user_map_script"); ok {
		object.UserMapScript = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		object.Password = v.(string)
	}
	if v, ok := d.GetOk("use_ad_login_pw"); ok {
		object.UseLoginPwAdAttr = v.(bool)
	}
	if v, ok := d.GetOk("use_ad_login_pw_by_script"); ok {
		object.UseLoginPwUseScript = v.(bool)
	}
	// Policy
	if v, ok := d.GetOk("default_profile_id"); ok {
		object.DefaultAuthProfile = v.(string)
	}

	if v, ok := d.GetOk("policy_script"); ok {
		object.PolicyScript = v.(string)
	}
	if v, ok := d.GetOk("sets"); ok {
		object.Sets = flattenSchemaSetToStringSlice(v)
	}
	// Workflow
	if v, ok := d.GetOk("workflow_enabled"); ok {
		object.WorkflowEnabled = v.(bool)
	}
	if v, ok := d.GetOk("workflow_approver"); ok {
		object.WorkflowApproverList = expandWorkflowApprovers(v.([]interface{})) // This is a slice
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
			return fmt.Errorf("schema setting error: %s", err)
		}
	}

	return nil
}
