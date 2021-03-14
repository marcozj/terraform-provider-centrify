package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username in loginid@suffix format",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address",
			},
			"displayname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Display name",
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding user")
	client := m.(*restapi.RestClient)
	object := vault.NewUser(client)
	object.Name = d.Get("username").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found user: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("username", result["Username"].(string))
	d.Set("email", result["Email"].(string))
	d.Set("displayname", result["DisplayName"].(string))

	return nil
}
