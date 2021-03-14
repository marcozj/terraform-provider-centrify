package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultDomainRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the domain",
			},
		},
	}
}

func dataSourceVaultDomainRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding domain")
	client := m.(*restapi.RestClient)
	object := vault.NewDomain(client)
	object.Name = d.Get("name").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found domain: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))

	return nil
}
