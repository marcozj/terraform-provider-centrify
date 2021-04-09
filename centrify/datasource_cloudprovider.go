package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceCloudProvider() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudProviderRead,

		Schema: map[string]*schema.Schema{
			"cloud_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account ID of the cloud provider",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud provider",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the cloud provider",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the cloud provider",
			},
			"enable_interactive_password_rotation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable interactive password rotation",
			},
			"prompt_change_root_password": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Prompt to change root password every login and password checkin",
			},
			"enable_password_rotation_reminders": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable password rotation reminders",
			},
			"password_rotation_reminder_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum number of days since last rotation to trigger a reminder",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default Root Account Login Profile (used if no conditions matched)",
			},
			"challenge_rule": getChallengeRulesSchema(),
		},
	}
}

func dataSourceCloudProviderRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding CloudProvider")
	client := m.(*restapi.RestClient)
	object := vault.NewCloudProvider(client)
	object.CloudAccountID = d.Get("cloud_account_id").(string)
	object.Name = d.Get("name").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving CloudProvider with name '%s': %s", object.Name, err)
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
		default:
			d.Set(k, v)
		}
	}

	return nil
}
