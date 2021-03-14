package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceGlobalGroupMappings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMappingCreate,
		Read:   resourceGroupMappingRead,
		Update: resourceGroupMappingUpdate,
		Delete: resourceGroupMappingDelete,

		Schema: map[string]*schema.Schema{
			"mapping": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      customGroupMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Group attribute value",
						},
						"group_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Group name",
						},
					},
				},
			},
		},
	}
}

func resourceGroupMappingRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading global group mappings: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewGroupMappings(client)
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading global group mappings: %v", err)
	}
	//logger.Debugf("Manual Set from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceRoleRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	//d.Set("mapping", object.Mappings)

	logger.Infof("Completed reading global group mappings")
	return nil
}

func resourceGroupMappingCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning global group mappings creation: %s", ResourceIDString(d))

	d.SetId("centrifyvault_global_group_mappings")

	client := m.(*restapi.RestClient)
	object := vault.NewGroupMappings(client)

	createUpateGroupMappingsData(d, object)

	err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating global group mappings: %v", err)
	}

	// Creation completed
	logger.Infof("Creation of global group mappings completed: %s", d.Id())
	return resourceGroupMappingRead(d, m)
}

func resourceGroupMappingUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning global group mappings update: %s", ResourceIDString(d))

	d.SetId("centrifyvault_global_group_mappings")

	client := m.(*restapi.RestClient)
	object := vault.NewGroupMappings(client)

	createUpateGroupMappingsData(d, object)

	// If there is change, delete all then add
	if d.HasChanges("mapping") {
		old, _ := d.GetChange("mapping")
		oldobject := vault.NewGroupMappings(client)
		oldobject.Mappings = expandGroupMappings(old)
		err := oldobject.Delete()
		if err != nil {
			return fmt.Errorf("Error deleting global group mappings: %v", err)
		}

		err = object.Create()
		if err != nil {
			return fmt.Errorf("Error adding global group mappings: %v", err)
		}
	}

	return resourceGroupMappingRead(d, m)
}

func resourceGroupMappingDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of global group mappings: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object := vault.NewGroupMappings(client)
	// We need to fill the mappings so that they can be deleted one by one
	createUpateGroupMappingsData(d, object)
	err := object.Delete()
	if err != nil {
		return fmt.Errorf("Error deleting global group mappings: %v", err)
	}

	d.SetId("")

	logger.Infof("Deletion of global group mappings completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGroupMappingsData(d *schema.ResourceData, object *vault.GroupMappings) {
	if v, ok := d.GetOk("mapping"); ok {
		object.Mappings = expandGroupMappings(v)
	}
}

func expandGroupMappings(v interface{}) []vault.GroupMapping {
	mappings := []vault.GroupMapping{}

	for _, p := range v.(*schema.Set).List() {
		mapping := vault.GroupMapping{}
		mapping.AttributeValue = p.(map[string]interface{})["attribute_value"].(string)
		mapping.GroupName = p.(map[string]interface{})["group_name"].(string)
		mappings = append(mappings, mapping)
	}
	logger.Debugf("Group mappings: %+v", mappings)

	return mappings
}
