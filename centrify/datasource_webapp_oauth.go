package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceOauthWebApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOauthWebAppRead,

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
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application ID. Specify the name or 'target' that the mobile application uses to find this application.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Web App",
			},
			"oauth_profile": getOAuthProfileSchema(),
			"script": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Script to customize JWT token creation for this application",
			},
			"oidc_script": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceOauthWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Oauth webapp")
	client := m.(*restapi.RestClient)
	object := vault.NewOauthWebApp(client)
	object.Name = d.Get("name").(string)
	object.ApplicationID = d.Get("application_id").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving Oauth webapp with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
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

	return nil
}
