package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourcePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the policy",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the policy",
			},
			"link_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link type of the policy",
			},
			"policy_assignment": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of roles or sets assigned to the policy",
			},
			"settings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"centrify_services":        getCentrifyServicesSchema(),
						"centrify_client":          getCentrifyClientSchema(),
						"centrify_css_server":      getCentrifyCSSServerSchema(),
						"centrify_css_workstation": getCentrifyCSSWorkstationSchema(),
						"centrify_css_elevation":   getCentrifyCSSElevationSchema(),
						"self_service":             getSelfServiceSchema(),
						"password_settings":        getPasswordSettingsSchema(),
						"oath_otp":                 getOATHOTPSchema(),
						"radius":                   getRadiusSchema(),
						"user_account":             getUserAccountSchema(),
						"system_set":               getSystemSetSchema(),
						"database_set":             getDatabaseAndDomainSetSchema(),
						"domain_set":               getDatabaseAndDomainSetSchema(),
						"account_set":              getAccountSetSchema(),
						"secret_set":               getSecretSetSchema(),
						"sshkey_set":               getSSHKeySetSchema(),
						"cloudproviders_set":       getCloudProvidersSchema(),
						"mobile_device":            getMobileDeviceSchema(),
					},
				},
			},
		},
	}
}

func dataSourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding policy")
	client := m.(*restapi.RestClient)
	object := vault.NewPolicy(client)
	object.Name = d.Get("name").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving policy with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "plink":
			// Handle plink content. In schema, following attributes are in root level but they are sub map section
			d.Set("link_type", object.Plink.LinkType)
			d.Set("policy_assignment", object.Plink.Params)
		case "settings":
			// Handle settings content.
			service := make(map[string]interface{})
			// convert each service map into []interface{}
			for service_key, service_value := range v.(map[string]interface{}) {
				processed_service_value := make(map[string]interface{})
				// convert challenge_rule map into []interface{}
				for attribute_key, attribute_value := range service_value.(map[string]interface{}) {
					switch attribute_key {
					case "challenge_rule", "access_secret_checkout_rule", "privilege_elevation_rule":
						processed_service_value[attribute_key] = attribute_value.(map[string]interface{})["rule"]
					case "admin_user_password":
						processed_service_value[attribute_key] = []interface{}{attribute_value}
					default:
						processed_service_value[attribute_key] = attribute_value
					}
				}
				service[service_key] = []interface{}{processed_service_value}
			}
			d.Set(k, []interface{}{service})
		default:
			d.Set(k, v)
		}
	}
	/*
		result, err := object.Query("name")
		if err != nil {
			return fmt.Errorf("error retrieving policy with name '%s': %s", object.Name, err)
		}

		//logger.Debugf("Found user: %+v", result)
		d.SetId(result["ID"].(string))
		d.Set("description", result["Description"].(string))
		d.Set("link_type", result["LinkType"].(string))
	*/
	return nil
}
