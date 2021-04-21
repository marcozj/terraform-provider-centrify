package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceDesktopApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDesktopAppRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Desktop App",
			},
			"template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template name of the Desktop App",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Web App",
			},
			"application_host_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Application host",
			},
			"login_credential_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host login credential type",
			},
			"application_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Host login credential account",
			},
			"application_alias": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The alias name of the published RemoteApp program",
			},
			"command_line": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command line",
			},
			"command_parameter": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      customCommandParamHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the parameter",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the parameter",
						},
						"target_object_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of selected parameter value",
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default authentication profile ID",
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

func dataSourceDesktopAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding DesktopApp")
	client := m.(*restapi.RestClient)
	object := vault.NewDesktopApp(client)
	object.Name = d.Get("name").(string)

	// We can't use simple Query method because it doesn't return all attributes
	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving DesktopApp with name '%s': %s", object.Name, err)
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
