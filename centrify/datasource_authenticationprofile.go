package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAuthenticationProfileRead,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "UUID of the authentication profile",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the authentication profile",
			},
			"pass_through_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Pass through duration of the authentication profile",
			},
		},
	}
}

func dataSourceAuthenticationProfileRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding authentication profile")
	client := m.(*restapi.RestClient)
	object := vault.NewAuthenticationProfile(client)
	object.Name = d.Get("name").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found authentication profile: %+v", result)
	d.SetId(result["Uuid"].(string))
	d.Set("uuid", result["Uuid"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("pass_through_duration", int(result["DurationInMinutes"].(float64)))

	return nil
}
