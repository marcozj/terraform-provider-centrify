package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceSamlWebApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSamlWebAppRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Web App",
			},
			"template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template name of the Web App",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
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
			"sp_metadata_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sp_metadata_xml": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service Provider metadata in XML format",
			},
			"sp_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SP Entity ID, also known as SP Issuer, or Audience, is a value given by your Service Provider",
			},
			"acs_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Assertion Consumer Service (ACS) URL",
			},
			"recipient_sameas_acs_url": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Recipient is same as ACS URL",
			},
			"recipient": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service Provider recipient if it is different from ACS URL",
			},
			"sign_assertion": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Sign assertion if true, otherwise sign response",
			},
			"name_id_format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "This is the Format attribute value in the <NameID> element in SAML Response. Select the NameID Format that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'.",
			},
			"sp_single_logout_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Single Logout URL",
			},
			"relay_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If your Service Provider specifies a Relay State value to use, specify it here.",
			},
			"authn_context_class": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Select the Authentication Context Class that your Service Provider specifies to use. If SP does not specify one, select 'unspecified'.",
			},
			// SAML Response menu
			"saml_attribute": getSamlAttributeSchema(),
			"saml_script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Javascript used to produce custom logic for SAML response",
			},
			// Policy menu
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default authentication profile ID",
			},
			// Account Mapping menu
			"username_strategy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account mapping",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All users share the user name. Applicable if 'username_strategy' is 'Fixed'",
			},
			"user_map_script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account mapping script. Applicable if 'username_strategy' is 'UseScript'",
			},
			// Workflow
			"workflow_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"workflow_approver": getWorkflowApproversSchema(),
			"challenge_rule":    getChallengeRulesSchema(),
			"policy_script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Use script to specify authentication rules (configured rules are ignored)",
			},
		},
	}
}

func dataSourceSamlWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Saml webapp")
	client := m.(*restapi.RestClient)
	object := vault.NewSamlWebApp(client)
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("application_id"); ok {
		object.ServiceName = v.(string)
	}
	if v, ok := d.GetOk("corp_identifier"); ok {
		object.CorpIdentifier = v.(string)
	}
	if v, ok := d.GetOk("app_entity_id"); ok {
		object.AdditionalField1 = v.(string)
	}
	if v, ok := d.GetOk("sp_entity_id"); ok {
		object.Audience = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving SAML webapp with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "challenge_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		case "workflow_settings":
			if v.(string) != "" {
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

	return nil
}
