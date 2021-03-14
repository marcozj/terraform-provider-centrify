package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceConnector() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConnectorRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Connector",
			},
		},
	}
}

func dataSourceConnectorRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding connector")
	client := m.(*restapi.RestClient)
	object := vault.NewConnector(client)
	object.Name = d.Get("name").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found connector: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))

	return nil
}
