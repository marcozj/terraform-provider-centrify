package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceFederatedGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFederatedGroupRead,

		Schema: getDSFederatedGroupSchema(),
	}
}

func getDSFederatedGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the federated group",
		},
	}
}

func dataSourceFederatedGroupRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding federated group")
	client := m.(*restapi.RestClient)
	object := vault.NewFederatedGroup(client)
	object.Name = d.Get("name").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving federated group with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	return nil
}
