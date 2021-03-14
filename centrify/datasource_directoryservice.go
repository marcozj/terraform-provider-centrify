package centrify

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/directoryservice"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceDirectoryService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDirectoryServiceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Directory Service",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the Directory Service",
				ValidateFunc: validation.StringInSlice([]string{
					directoryservice.CentrifyDirectory.String(),
					directoryservice.ActiveDirectory.String(),
					directoryservice.FederatedDirectory.String(),
					directoryservice.GoogleDirectory.String(),
					directoryservice.LDAPDirectory.String(),
				}, false),
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Status of the Directory Service",
			},
		},
	}
}

func dataSourceDirectoryServiceRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding DirectoryService")
	client := m.(*restapi.RestClient)
	object := vault.NewDirectoryServices(client)

	err := object.Read()
	if err != nil {
		return fmt.Errorf("Error retrieving directory services: %s", err)
	}

	name := d.Get("name").(string)
	var dirtype string
	switch d.Get("type").(string) {
	case directoryservice.CentrifyDirectory.String():
		dirtype = "CDS"
	case directoryservice.ActiveDirectory.String():
		dirtype = "AdProxy"
	case directoryservice.FederatedDirectory.String():
		dirtype = "FDS"
	case directoryservice.GoogleDirectory.String():
		dirtype = "GDS"
	case directoryservice.LDAPDirectory.String():
		dirtype = "LdapProxy"
	}

	var results []vault.DirectoryService
	for _, v := range object.DirServices {
		if dirtype == v.Service && name == v.Config {
			results = append(results, v)
		}
	}
	if len(results) == 0 {
		return errors.New("Query returns 0 object")
	}
	if len(results) > 1 {
		return fmt.Errorf("Query returns too many objects (found %d, expected 1)", len(results))
	}

	var result = results[0]
	//logger.Debugf("Found connector: %+v", result)
	d.SetId(result.ID)
	d.Set("name", result.Config)
	d.Set("status", result.Status)
	d.Set("type", name)

	return nil
}
