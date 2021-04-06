package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceOauthWebApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOauthWebAppRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Web App",
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application ID. Specify the name or 'target' that the mobile application uses to find this application.",
			},
		},
	}
}

func dataSourceOauthWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Oauth webapp")
	client := m.(*restapi.RestClient)
	object := vault.NewOauthWebApp(client)
	object.Name = d.Get("name").(string)
	object.ApplicationID = d.Get("application_id").(string)

	// We can't use simple Query method because it doesn't return all attributes
	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving Oauth webapp with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)
	if object.Name != "" {
		d.Set("name", object.Name)
	}
	if object.ApplicationID != "" {
		d.Set("application_id", object.ApplicationID)
	}

	return nil
}
