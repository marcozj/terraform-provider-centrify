package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRoleRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the role",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of an role",
			},
			"adminrights": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      customRoleMemberHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the member",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the member",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the member",
						},
					},
				},
			},
		},
	}
}

func dataSourceRoleRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding role")
	client := m.(*restapi.RestClient)
	object := vault.NewRole(client)
	object.Name = d.Get("name").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving role with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	return nil
}
