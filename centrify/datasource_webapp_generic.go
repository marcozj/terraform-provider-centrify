package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceGenericWebApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGenericWebAppRead,

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
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname_suffix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The host name suffix for the url of the login form, for example, acme.com.",
			},
			"username_field": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CSS Selector for the user name field in the login form, for example, input#login-username.",
			},
			"password_field": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CSS Selector for the password field in the login form, for example, input#login-password.",
			},
			"submit_field": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CSS Selector for the Submit button in the login form, for example, input#login-button. This entry is optional. It is required only if you cannot submit the form by pressing the enter key.",
			},
			"form_field": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CSS Selector for the form field of the login form, for example, form#loginForm.",
			},
			"additional_login_field": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CSS Selector for any Additional Login Field required to login besides username and password, such as Company name or Agency ID. For example, the selector could be input#login-company-id. This entry is required only if there is an additional login field besides username and password.",
			},
			"additional_login_field_value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The value for the Additional Login Field. For example, if there is an additional login field for the company name, enter the company name here. This entry is required if Additional Login Field is set.",
			},
			"selector_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Use this field to indicate the number of milliseconds to wait for the expected input selectors to load before timing out on failure. A zero or negative number means no timeout.",
			},
			"order": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Use this field to specify the order of login if it is not username, password and submit.",
			},
			"script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Script to log the user in to this application",
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
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for all user share one name",
			},
			"use_ad_login_pw": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use the login password supplied by the user (Active Directory users only)",
			},
			"use_ad_login_pw_by_script": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use the login password supplied by the user for account mapping script (Active Directory users only)",
			},
			"user_map_script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account mapping script",
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

func dataSourceGenericWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Generic webapp")
	client := m.(*restapi.RestClient)
	object := vault.NewGenericWebApp(client)
	object.Name = d.Get("name").(string)

	// We can't use simple Query method because it doesn't return all attributes
	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving Generic webapp with name '%s': %s", object.Name, err)
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
