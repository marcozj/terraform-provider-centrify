package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/webapp/accountmapping"
	"github.com/marcozj/golang-sdk/enum/webapp/oidc/applicationtemplate"

	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceOidcWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceOidcWebAppCreate,
		Read:   resourceOidcWebAppRead,
		Update: resourceOidcWebAppUpdate,
		Delete: resourceOidcWebAppDelete,
		Exists: resourceOidcWebAppExists,
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
					applicationtemplate.Generic.String(),
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
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application ID. Specify the name or 'target' that the mobile application uses to find this application.",
			},
			"oauth_profile": getOidcProfileSchema(),
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to generate OpenID Connect Authorization and UserInfo responses for this application",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					suppress := false
					if new == "" && old == "@GenericOpenIDConnect" {
						suppress = true
					}
					return suppress
				},
			},
			"oidc_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				}, false),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "userprincipalname",
				Description: "All users share the user name. Applicable if 'username_strategy' is 'Fixed'",
			},
			"user_map_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account mapping script. Applicable if 'username_strategy' is 'UseScript'",
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

func getOidcProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_secret": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The OpenID Client Secret for this Identity Provider",
				},
				"application_url": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Resource application URL. The OpenID Connect Service Provider URL",
				},
				"issuer": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The OpenID Connect Issuer URL for this application",
				},
				"client_id": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The OpenID Client ID for this Identity Provider",
				},
				"redirects": {
					Type:     schema.TypeSet,
					Required: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Authorized Redirect URIs. Redirect URI that the Service Provider will specify in the OpenID Connect request to Centrify",
				},
				// Tokens menu
				"token_lifetime": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Token lifetime",
				},
				"allow_refresh": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Issue refresh tokens",
				},
				"refresh_lifetime": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Refresh token lifetime",
				},
				"script": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Script to generate OpenID Connect Authorization and UserInfo responses for this application",
				},
			},
		},
	}
}

func resourceOidcWebAppExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Oidc WebApp exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewOidcWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Oidc WebApp exists in tenant: %s", object.ID)
	return true, nil
}

func resourceOidcWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Oidc WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewWebpp object and populate ID attribute
	object := vault.NewOidcWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading Oidc WebApp: %v", err)
	}
	//logger.Debugf("WebApp from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceOidcWebAppRead(): %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "oauth_profile":
			d.Set(k, []interface{}{v})
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
		default:
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading Oidc WebApp: %s", object.Name)
	return nil
}

func resourceOidcWebAppCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Oidc WebApp creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a WebApp object
	object := vault.NewOidcWebApp(client)
	object.TemplateName = d.Get("template_name").(string)
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating Oidc WebApp: %v", err)
	}
	if len(resp.Result) <= 0 {
		return fmt.Errorf("import application template returns incorrect result")
	}

	id := resp.Result[0].(map[string]interface{})["_RowKey"].(string)

	if id == "" {
		return fmt.Errorf("oidc WebApp ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Update attributes to complete creation
	err = createUpateGetOidcWebAppData(d, object)
	if err != nil {
		return err
	}

	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("error updating Oidc WebApp attribute: %v", err2)
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
			return fmt.Errorf("error setting Oidc WebApp permissions: %v", err)
		}
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Oidc WebApp completed: %s", object.Name)
	return resourceOidcWebAppRead(d, m)
}

func resourceOidcWebAppUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Oidc WebApp update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewOidcWebApp(client)
	object.ID = d.Id()
	// ClientId is gnerated value and must be supplied for update action,
	// so we need to read it from state and "inject" it into object
	if v, ok := d.GetOk("oauth_profile"); ok {
		profiles := v.([]interface{})
		if len(profiles) > 0 && profiles[0] != nil {
			d := profiles[0].(map[string]interface{})
			if v, ok := d["client_id"]; ok {
				object.OAuthProfile.ClientID = v.(string)
			}
		}
	}

	err := createUpateGetOidcWebAppData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "template_name", "description", "application_id", "oauth_profile", "script", "default_profile_id", "challenge_rule",
		"policy_script", "username_strategy", "ad_attribute", "username", "user_map_script", "workflow_enabled", "workflow_approver") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating Oauth WebApp attribute: %v", err)
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
				return fmt.Errorf("error removing Oidc WebApp from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("error adding Oidc WebApp to Set: %v", err)
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
				return fmt.Errorf("error adding Oidc WebApp permissions: %v", err)
			}
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of Oidc WebApp completed: %s", object.Name)
	return resourceOidcWebAppRead(d, m)
}

func resourceOidcWebAppDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Oidc WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewOidcWebApp(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("error deleting Oidc WebApp: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Oidc WebApp completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetOidcWebAppData(d *schema.ResourceData, object *vault.OidcWebApp) error {
	object.Name = d.Get("name").(string)
	object.TemplateName = d.Get("template_name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("application_id"); ok {
		object.ApplicationID = v.(string)
	}

	if v, ok := d.GetOk("oauth_profile"); ok {
		object.OAuthProfile = expandOidcProfile(v, object.OAuthProfile.ClientID)
	}
	if v, ok := d.GetOk("script"); ok {
		// This is annoying. "Script" attribute is used for update but "OpenIDConnectScript" attribute is used for read
		object.Script = v.(string)
		object.OpenIDConnectScript = v.(string)
	}
	// Specially handling for script
	// If "script" is changed from whatever to empty, assign "@GenericOpenIDConnect" to it
	if d.HasChange("script") {
		_, new := d.GetChange("script")
		if new == "" {
			object.Script = "@GenericOpenIDConnect"
		}
	}

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

func expandOidcProfile(v interface{}, clientid string) *vault.OidcProfile {
	options := v.([]interface{})
	if len(options) > 0 && options[0] != nil {
		d := options[0].(map[string]interface{})
		data := &vault.OidcProfile{}
		data.ClientID = clientid
		data.ClientSecret = d["client_secret"].(string)
		if v, ok := d["application_url"]; ok {
			data.Url = v.(string)
		}
		if v, ok := d["redirects"]; ok {
			data.Redirects = flattenSchemaSetToStringSlice(v)
		}
		data.TokenLifetime = d["token_lifetime"].(string)
		if v, ok := d["allow_refresh"]; ok {
			data.AllowRefresh = v.(bool)
		}
		if v, ok := d["refresh_lifetime"]; ok {
			data.RefreshLifetime = v.(string)
		}

		return data
	}
	return nil
}

/*
func flattenOAuthProfileData(v *vault.OidcProfile) []interface{} {
	data := map[string]interface{}{
		"client_id":        v.ClientID,
		"issuer":           v.Issuer,
		"client_secret":    v.ClientSecret,
		"application_url":  v.Url,
		"token_lifetime":   v.TokenLifetime,
		"allow_refresh":    v.AllowRefresh,
		"refresh_lifetime": v.RefreshLifetime,
		"redirects":        flattenStringSliceToSet(v.Redirects),
	}

	return []interface{}{data}
}

func flattenOidcProfileData(v *vault.OidcProfile) ([]interface{}, error) {
	schemamap, err := vault.GenerateSchemaMap(v)
	if err != nil {
		return nil, err
	}
	return []interface{}{schemamap}, nil
}
*/
