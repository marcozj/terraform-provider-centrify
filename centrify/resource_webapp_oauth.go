package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/webapp/oauth/applicationtemplate"
	"github.com/marcozj/golang-sdk/enum/webapp/oauth/clientidtype"
	"github.com/marcozj/golang-sdk/enum/webapp/oauth/tokentype"

	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceOauthWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceOauthWebAppCreate,
		Read:   resourceOauthWebAppRead,
		Update: resourceOauthWebAppUpdate,
		Delete: resourceOauthWebAppDelete,
		Exists: resourceOauthWebAppExists,
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
					applicationtemplate.OAuth2Client.String(),
					applicationtemplate.OAuth2Server.String(),
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
			"oauth_profile": getOAuthProfileSchema(),
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Script to customize JWT token creation for this application",
			},
			"oidc_script": {
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
			"permission": getPermissionSchema(),
		},
	}
}

func getOAuthProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"clientid_type": {
					Type:         schema.TypeInt,
					Required:     true,
					Description:  "ClientID type",
					ValidateFunc: validation.IntInSlice([]int{int(clientidtype.AnythingOrList), int(clientidtype.Confidential)}),
				},
				"issuer": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "OAuth server issuer",
				},
				"audience": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "OAuth server audience",
				},
				"target_is_us": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"allowed_clients": {
					Type:     schema.TypeSet,
					Optional: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Allowed clients",
				},
				"must_oauth_client": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Must be OAuth Client",
				},
				"redirects": {
					Type:     schema.TypeSet,
					Optional: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Redirects",
				},
				// Tokens menu
				"token_type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "ClientID type",
					ValidateFunc: validation.StringInSlice([]string{
						tokentype.JwtRS256.String(),
						tokentype.Opaque.String(),
					}, false),
				},
				"allowed_auth": {
					Type:     schema.TypeSet,
					Required: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Authentication methods",
				},
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
				// Scope menu
				"confirm_authorization": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "User must confirm authorization request",
				},
				"allow_scope_select": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow scope selection",
				},
				"scope": getOAuthScopeSchema(),
			},
		},
	}
}

func getOAuthScopeSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Scope definitions",
		Set:         customOauthScopeHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"allowed_rest_apis": {
					Type:     schema.TypeSet,
					Optional: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Allowed REST APIs",
				},
			},
		},
	}
}

func resourceOauthWebAppExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Oauth WebApp exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewOauthWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Oauth WebApp exists in tenant: %s", object.ID)
	return true, nil
}

func resourceOauthWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Oauth WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewWebpp object and populate ID attribute
	object := vault.NewOauthWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading Oauth WebApp: %v", err)
	}
	//logger.Debugf("WebApp from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceOauthWebAppRead(): %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "oauth_profile":
			profile := make(map[string]interface{})
			for attribute_key, attribute_value := range v.(map[string]interface{}) {
				switch attribute_key {
				case "allowed_auth":
					profile[attribute_key] = schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(attribute_value.(string), ",")))
				default:
					profile[attribute_key] = attribute_value
				}
			}
			d.Set(k, []interface{}{profile})
		case "challenge_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		default:
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading Oauth WebApp: %s", object.Name)
	return nil
}

func resourceOauthWebAppCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Oauth WebApp creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a WebApp object
	object := vault.NewOauthWebApp(client)
	object.TemplateName = d.Get("template_name").(string)
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating Oauth WebApp: %v", err)
	}
	if len(resp.Result) <= 0 {
		return fmt.Errorf("import application template returns incorrect result")
	}

	id := resp.Result[0].(map[string]interface{})["_RowKey"].(string)

	if id == "" {
		return fmt.Errorf("oauth WebApp ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Update attributes to complete creation
	err = createUpateGetOauthWebAppData(d, object)
	if err != nil {
		return err
	}

	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("error updating Oauth WebApp attribute: %v", err2)
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
			return fmt.Errorf("error setting Oauth WebApp permissions: %v", err)
		}
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Oauth WebApp completed: %s", object.Name)
	return resourceOauthWebAppRead(d, m)
}

func resourceOauthWebAppUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Oauth WebApp update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewOauthWebApp(client)
	object.ID = d.Id()
	err := createUpateGetOauthWebAppData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "template_name", "description", "application_id", "oauth_profile", "script") {
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
				return fmt.Errorf("error removing SAML WebApp from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("error adding SAML WebApp to Set: %v", err)
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
				return fmt.Errorf("error adding Oauth WebApp permissions: %v", err)
			}
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of Oauth WebApp completed: %s", object.Name)
	return resourceOauthWebAppRead(d, m)
}

func resourceOauthWebAppDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Oauth WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewOauthWebApp(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("error deleting Oauth WebApp: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Oauth WebApp completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetOauthWebAppData(d *schema.ResourceData, object *vault.OauthWebApp) error {
	object.Name = d.Get("name").(string)
	object.TemplateName = d.Get("template_name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("application_id"); ok {
		object.ApplicationID = v.(string)
	}

	if v, ok := d.GetOk("oauth_profile"); ok {
		object.OAuthProfile = expandOAuthProfile(v)
	}
	if v, ok := d.GetOk("script"); ok {
		object.Script = v.(string)
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

	return nil
}

func expandOAuthProfile(v interface{}) *vault.OAuthProfile {
	options := v.([]interface{})
	if len(options) > 0 && options[0] != nil {
		d := options[0].(map[string]interface{})
		data := &vault.OAuthProfile{}
		data.ClientIDType = d["clientid_type"].(int)
		if v, ok := d["Issuer"]; ok {
			data.Issuer = v.(string)
		}
		if v, ok := d["audience"]; ok {
			data.Audience = v.(string)
		}
		if v, ok := d["allowed_clients"]; ok {
			data.AllowedClients = flattenSchemaSetToStringSlice(v)
		}
		if v, ok := d["must_oauth_client"]; ok {
			data.MustBeOauthClient = v.(bool)
		}
		if v, ok := d["redirects"]; ok {
			data.Redirects = flattenSchemaSetToStringSlice(v)
		}
		data.TokenType = d["token_type"].(string)
		if v, ok := d["allowed_auth"]; ok {
			data.AllowedAuth = flattenSchemaSetToString(v.(*schema.Set))
		}
		data.TokenLifetime = d["token_lifetime"].(string)
		if v, ok := d["allow_refresh"]; ok {
			data.AllowRefresh = v.(bool)
		}
		if v, ok := d["refresh_lifetime"]; ok {
			data.RefreshLifetime = v.(string)
		}
		if v, ok := d["confirm_authorization"]; ok {
			data.ConfirmAuthorization = v.(bool)
		}
		if v, ok := d["allow_scope_select"]; ok {
			data.AllowScopeSelect = v.(bool)
		}
		if v, ok := d["scope"]; ok {
			data.KnownScopes = expandOAuthScope(v)
		}

		return data
	}
	return nil
}

func expandOAuthScope(v interface{}) []vault.OAuthScope {
	m := v.(*schema.Set).List()
	var scopes []vault.OAuthScope
	for _, lrv := range m {
		scope := vault.OAuthScope{}
		scope.Name = lrv.(map[string]interface{})["name"].(string)
		scope.Description = lrv.(map[string]interface{})["description"].(string)
		v := lrv.(map[string]interface{})["allowed_rest_apis"].(interface{})
		scope.AllowedRestAPIs = flattenSchemaSetToStringSlice(v)

		scopes = append(scopes, scope)
	}

	return scopes
}
