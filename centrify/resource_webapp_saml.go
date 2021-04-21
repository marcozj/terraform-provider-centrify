package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/webapp/accountmapping"
	"github.com/marcozj/golang-sdk/enum/webapp/saml/applicationtemplate"
	"github.com/marcozj/golang-sdk/enum/webapp/saml/configurationmethod"

	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceSamlWebApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSamlWebAppCreate,
		Read:   resourceSamlWebAppRead,
		Update: resourceSamlWebAppUpdate,
		Delete: resourceSamlWebAppDelete,
		Exists: resourceSamlWebAppExists,
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
					applicationtemplate.SAML.String(),
					applicationtemplate.AWSConsole.String(),
					applicationtemplate.Cloudera.String(),
					applicationtemplate.CloudLock.String(),
					applicationtemplate.ConfluenceServer.String(),
					applicationtemplate.Dome9.String(),
					applicationtemplate.GitHubEnterprise.String(),
					applicationtemplate.JIRACloud.String(),
					applicationtemplate.JIRAServer.String(),
					applicationtemplate.PaloAltoNetworks.String(),
					applicationtemplate.SplunkOnPrem.String(),
					applicationtemplate.SumoLogic.String(),
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
			"corp_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS Account ID or Jira Cloud Subdomain",
			},
			"app_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloudera Entity ID or JIRA Cloud SP Entity ID",
			},
			"application_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Application ID. Specify the name or 'target' that the mobile application uses to find this application.",
			},
			"idp_metadata_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Trust menu
			"sp_metadata_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_config_method": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Service Provider configuration method: metadata or manual configuration",
				ValidateFunc: validation.IntInSlice([]int{int(configurationmethod.Manual), int(configurationmethod.MetaData)}),
			},
			"sp_metadata_xml": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service Provider metadata in XML format",
				// When Service Provider Configuration is set to use metadata and sp_metadata_url is used, this attribute value
				// is automatically filled. Therefore we want to ignore this attribute during update action.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					suppress := false
					if v, ok := d.GetOk("sp_config_method"); ok {
						sp_metadata_url := d.Get("sp_metadata_url")
						sp_config_method := v.(int)
						if sp_config_method == 1 && sp_metadata_url != "" {
							suppress = true
						}
					}
					return suppress
				},
			},
			"sp_entity_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "SP Entity ID, also known as SP Issuer, or Audience, is a value given by your Service Provider",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"acs_url": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Assertion Consumer Service (ACS) URL",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"recipient_sameas_acs_url": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          true,
				Description:      "Recipient is same as ACS URL",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"recipient": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Service Provider recipient if it is different from ACS URL",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"sign_assertion": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          true,
				Description:      "Sign assertion if true, otherwise sign response",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"name_id_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "unspecified",
				Description: "This is the Format attribute value in the <NameID> element in SAML Response. Select the NameID Format that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'.",
				ValidateFunc: validation.StringInSlice([]string{
					"unspecified",
					"emailAddress",
					"transient",
					"persistent",
					"entity",
					"kerberos",
					"WindowsDomainQualifiedName",
					"X509SubjectName",
				}, false),
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"sp_single_logout_url": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Single Logout URL",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"relay_state": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "If your Service Provider specifies a Relay State value to use, specify it here.",
				DiffSuppressFunc: samlAttributeSuppress,
			},
			"authn_context_class": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "unspecified",
				Description:      "Select the Authentication Context Class that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'.",
				DiffSuppressFunc: samlAttributeSuppress,
				ValidateFunc: validation.StringInSlice([]string{
					"unspecified",
					"PasswordProtectedTransport",
					"AuthenticatedTelephony",
					"InternetProtocol",
					"InternetProtocolPassword",
					"Kerberos",
					"MobileOneFactorContract",
					"MobileOneFactorUnregistered",
					"MobileTwoFactorContract",
					"MobileTwoFactorUnregistered",
					"NomadTelephony",
					"Password",
					"PersonalTelephony",
					"PGP",
					"PreviousSession",
					"SecureRemotePassword",
					"Smartcard",
					"SmartcardPKI",
					"SoftwarePKI",
					"SPKI",
					"Telephony",
					"TimeSyncToken",
					"TLSClient",
					"X509",
					"XMLDSig",
				}, false),
			},
			// SAML Response menu
			"saml_attribute": getSamlAttributeSchema(),
			"saml_response_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Javascript used to produce custom logic for SAML response",
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

func getSamlAttributeSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Map attributes from your source directory to SAML attributes that should be included in the SAML response for this application.",
		Set:         customSamlAttributeHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

// This DiffSuppressFunc function is used by few attributes when Service Provider Configuration is set to use metadata.
// In this case, those attribute values are derived from metadata. We want to ignore these attributes during update.
func samlAttributeSuppress(k, old, new string, d *schema.ResourceData) bool {
	suppress := false
	if v, ok := d.GetOk("sp_config_method"); ok {
		if v.(int) == 1 {
			suppress = true
		}
	}
	return suppress
}

func resourceSamlWebAppExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking SAML WebApp exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSamlWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("SAML WebApp exists in tenant: %s", object.ID)
	return true, nil
}

func resourceSamlWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading SAML WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewWebpp object and populate ID attribute
	object := vault.NewSamlWebApp(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading SAML WebApp: %v", err)
	}
	//logger.Debugf("WebApp from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceSamlWebAppRead(): %+v", schemamap)
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
		default:
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading SAML WebApp: %s", object.Name)
	return nil
}

func resourceSamlWebAppCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SAML WebApp creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a WebApp object
	object := vault.NewSamlWebApp(client)
	object.TemplateName = d.Get("template_name").(string)
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating SAML WebApp: %v", err)
	}
	if len(resp.Result) <= 0 {
		return fmt.Errorf("import application template returns incorrect result")
	}

	id := resp.Result[0].(map[string]interface{})["_RowKey"].(string)

	if id == "" {
		return fmt.Errorf("SAML WebApp ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Update attributes to complete creation
	err = createUpateGetSamlWebAppData(d, object)
	if err != nil {
		return err
	}

	resp2, err2 := object.Update()
	if err2 != nil || !resp2.Success {
		return fmt.Errorf("error updating SAML WebApp attribute: %v", err2)
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
			return fmt.Errorf("error setting SAML WebApp permissions: %v", err)
		}
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of SAML WebApp completed: %s", object.Name)
	return resourceSamlWebAppRead(d, m)
}

func resourceSamlWebAppUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SAML WebApp update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewSamlWebApp(client)
	object.ID = d.Id()
	err := createUpateGetSamlWebAppData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "template_name", "description", "corp_identifier", "app_entity_id", "application_id", "sp_config_method", "sp_metadata_xml", "sp_entity_id",
		"acs_url", "recipient_sameas_acs_url", "recipient", "sign_assertion", "name_id_format", "sp_single_logout_url", "encrypt_assertion",
		"relay_state", "authn_context_class", "saml_attribute", "saml_response_script", "default_profile_id", "challenge_rule",
		"policy_script", "username_strategy", "ad_attribute", "username", "user_map_script", "workflow_enabled", "workflow_approver") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating SAML WebApp attribute: %v", err)
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
				return fmt.Errorf("error removing SAML WebApp permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("error adding SAML WebApp permissions: %v", err)
			}
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of WebApp completed: %s", object.Name)
	return resourceSamlWebAppRead(d, m)
}

func resourceSamlWebAppDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of SAML WebApp: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSamlWebApp(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("error deleting SAML WebApp: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of SAML WebApp completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetSamlWebAppData(d *schema.ResourceData, object *vault.SamlWebApp) error {
	object.Name = d.Get("name").(string)
	object.TemplateName = d.Get("template_name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("corp_identifier"); ok {
		object.CorpIdentifier = v.(string)
	}
	if v, ok := d.GetOk("app_entity_id"); ok {
		object.AdditionalField1 = v.(string)
	}

	if v, ok := d.GetOk("application_id"); ok {
		object.ServiceName = v.(string)
	}
	if v, ok := d.GetOk("sp_config_method"); ok {
		object.SpConfigMethod = v.(int)
	}
	if v, ok := d.GetOk("sp_metadata_url"); ok {
		object.SpMetadataUrl = v.(string)
	}
	if v, ok := d.GetOk("sp_metadata_xml"); ok {
		object.SpMetadataXml = v.(string)
	}
	if v, ok := d.GetOk("sp_entity_id"); ok {
		object.Audience = v.(string)
	}
	if v, ok := d.GetOk("acs_url"); ok {
		object.ACS_Url = v.(string)
	}
	if v, ok := d.GetOk("recipient_sameas_acs_url"); ok {
		object.RecipientSameAsAcsUrl = v.(bool)
	}
	if v, ok := d.GetOk("recipient"); ok {
		object.Recipient = v.(string)
	}
	if v, ok := d.GetOk("sign_assertion"); ok {
		object.WantAssertionsSigned = v.(bool)
	}
	if v, ok := d.GetOk("name_id_format"); ok {
		object.NameIDFormat = v.(string)
	}
	if v, ok := d.GetOk("sp_single_logout_url"); ok {
		object.SpSingleLogoutUrl = v.(string)
	}
	if v, ok := d.GetOk("relay_state"); ok {
		object.RelayState = v.(string)
	}
	if v, ok := d.GetOk("authn_context_class"); ok {
		object.AuthnContextClass = v.(string)
	}
	if v, ok := d.GetOk("saml_response_script"); ok {
		object.SamlResponseScript = v.(string)
	}
	if v, ok := d.GetOk("saml_attribute"); ok {
		object.SamlAttributes = expandSamlAttributes(v)
	}
	if v, ok := d.GetOk("username_strategy"); ok {
		object.UserNameStrategy = v.(string)
	}
	//if v, ok := d.GetOk("ad_attribute"); ok {
	//	object.ADAttribute = v.(string)
	//}
	if v, ok := d.GetOk("username"); ok {
		object.Username = v.(string)
	}
	if v, ok := d.GetOk("user_map_script"); ok {
		object.UserMapScript = v.(string)
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

func expandSamlAttributes(v interface{}) []vault.SamlAttribute {
	m := v.(*schema.Set).List()
	var attributes []vault.SamlAttribute
	for _, lrv := range m {
		attribute := vault.SamlAttribute{}
		attribute.Name = lrv.(map[string]interface{})["name"].(string)
		attribute.Value = lrv.(map[string]interface{})["value"].(string)
		attributes = append(attributes, attribute)
	}

	return attributes
}
