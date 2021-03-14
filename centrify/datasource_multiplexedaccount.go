package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceMultiplexedAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMultiplexedAccountRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the multiplexed account",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the multiplexed account",
			},
			"account1_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account2_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account1": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account2": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"active_account": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceMultiplexedAccountRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding multiplexed account")
	client := m.(*restapi.RestClient)
	object := vault.NewMultiplexedAccount(client)
	object.Name = d.Get("name").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving multiplexed account: %s", err)
	}

	//logger.Debugf("Found multiplexed account: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("description", result["Description"].(string))
	// RedRock/query doesn't return these attributes
	//d.Set("account1_id", result["RealAccount1ID"].(string))
	//d.Set("account2_id", result["RealAccount2ID"].(string))
	//d.Set("account1", result["RealAccount1"].(string))
	//d.Set("account2", result["RealAccount2"].(string))
	if result["ActiveAccount"] != nil {
		d.Set("active_account", result["ActiveAccount"].(string))
	}

	return nil
}
