package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultSecretRead,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the secret",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the secret",
			},
			"secret_text": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Content of the secret",
			},
			"folder_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the folder where the secret is located",
			},
			"parent_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path of parent folder",
			},
			"checkout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve secret content",
			},
		},
	}
}

func dataSourceVaultSecretRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding vault secret")
	client := m.(*restapi.RestClient)
	object := vault.NewSecret(client)
	object.SecretName = d.Get("secret_name").(string)
	if v, ok := d.GetOk("folder_id"); ok {
		object.FolderID = v.(string)
	}

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found secret: %+v", result)
	object.ID = result["ID"].(string)
	d.SetId(object.ID)
	d.Set("secret_name", result["SecretName"].(string))
	if result["Description"] != nil {
		d.Set("description", result["Description"].(string))
	}
	if result["ParentPath"] != nil {
		d.Set("parent_path", result["ParentPath"].(string))
	}
	if result["FolderId"] != nil {
		d.Set("folder_id", result["FolderId"].(string))
	}

	if d.Get("checkout").(bool) {
		text, err := object.CheckoutSecret()
		if err != nil {
			return fmt.Errorf("Error retrieving secret content: %s", err)
		}
		d.Set("secret_text", text)
	}

	return nil
}
