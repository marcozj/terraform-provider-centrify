package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultSecretFolder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultSecretFolderRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the secret folder",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of an secret folder",
			},
			"parent_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parent folder path of an secret folder",
			},
			"parent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Parent folder ID of an secret folder",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default Secret Challenge Profile (used if no conditions matched)",
			},
			"challenge_rule": getChallengeRulesSchema(),
		},
	}
}

func dataSourceVaultSecretFolderRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding VaultSecretFolder")
	client := m.(*restapi.RestClient)
	object := vault.NewSecretFolder(client)
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("parent_path"); ok {
		object.ParentPath = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving secret folder with name '%s': %s", object.Name, err)
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
