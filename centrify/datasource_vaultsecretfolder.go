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
				Optional:    true,
				Description: "Description of an secret folder",
			},
			"parent_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parent folder path of an secret folder",
			},
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

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	if result["ID"] == nil {
		return fmt.Errorf("VaultSecretFolder ID is not set")
	}

	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	if result["Description"] != nil {
		d.Set("description", result["Description"].(string))
	}
	if result["ParentPath"] != nil {
		d.Set("parent_path", result["ParentPath"].(string))
	}

	return nil
}
