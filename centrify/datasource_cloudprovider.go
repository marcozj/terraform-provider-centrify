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
				Optional:    true,
				Description: "Account ID of the cloud provider",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud provider",
			},
		},
	}
}

func dataSourceCloudProviderRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding CloudProvider")
	client := m.(*restapi.RestClient)
	object := vault.NewCloudProvider(client)
	object.CloudAccountID = d.Get("cloud_account_id").(string)
	object.Name = d.Get("name").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found CloudProvider: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("cloud_account_id", result["CloudAccountId"].(string))

	return nil
}
