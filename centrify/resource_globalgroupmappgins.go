package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceGlobalGroupMappings_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMappingCreate,
		Read:   resourceGroupMappingRead,
		Update: resourceGroupMappingUpdate,
		Delete: resourceGroupMappingDelete,

		Schema:             getGroupMappingSchema(),
		DeprecationMessage: "resource centrifyvault_globalgroupmappings is deprecated will be removed in the future, use centrify_globalgroupmappings instead",
	}
}

func resourceGlobalGroupMappings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMappingCreate,
		Read:   resourceGroupMappingRead,
		Update: resourceGroupMappingUpdate,
		Delete: resourceGroupMappingDelete,

		Schema: getGroupMappingSchema(),
	}
}

func getGroupMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bulkupdate": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"mapping": {
			Type:     schema.TypeMap,
			Optional: true,
			//ConflictsWith: []string{"mapping"},
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		/*"mapping": {
			Type:          schema.TypeSet,
			Optional:      true,
			Set:           customGroupMappingHash,
			ConflictsWith: []string{"attribute_group"},
			Deprecated:    "use attribute_group instead",
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
		},*/
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
		return fmt.Errorf(" Error reading global group mappings: %v", err)
	}
	//logger.Debugf("Global group mapping from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceGroupMappingRead(): %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "attribute_group":
			mappings := make(map[string]interface{})
			for _, m := range v.([]interface{}) {
				var mapkey, mapvalue string
				for mk, mv := range m.(map[string]interface{}) {
					if mk == "attribute_value" {
						mapkey = mv.(string)
					}
					if mk == "group_name" {
						mapvalue = mv.(string)
					}
				}
				mappings[mapkey] = mapvalue
			}

			d.Set(k, mappings)
		default:
			d.Set(k, v)
		}
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

	var err error
	if object.BulkUpdate {
		err = object.Update()
	} else {
		err = object.Create()
	}
	if err != nil {
		return fmt.Errorf(" Error creating global group mappings: %v", err)
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
	if d.HasChanges("mapping", "attribute_group") {
		if object.BulkUpdate {
			err := object.Update()
			if err != nil {
				return fmt.Errorf(" Error updating global group mappings: %v", err)
			}
		} else {
			old, _ := d.GetChange("mapping")
			oldobject := vault.NewGroupMappings(client)
			oldobject.Mappings = expandGroupMappings(old)
			err := oldobject.Delete()
			if err != nil {
				return fmt.Errorf(" Error deleting global group mappings: %v", err)
			}

			err = object.Create()
			if err != nil {
				return fmt.Errorf(" Error adding global group mappings: %v", err)
			}
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
	var err error
	if object.BulkUpdate {
		object.Mappings = []vault.GroupMapping{}
		err = object.Update()
	} else {
		err = object.Delete()
	}
	if err != nil {
		return fmt.Errorf(" Error deleting global group mappings: %v", err)
	}

	d.SetId("")

	logger.Infof("Deletion of global group mappings completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGroupMappingsData(d *schema.ResourceData, object *vault.GroupMappings) {
	if v, ok := d.GetOk("bulkupdate"); ok {
		object.BulkUpdate = v.(bool)
	}
	if v, ok := d.GetOk("mapping"); ok && d.HasChange("mapping") {
		object.Mappings = expandAttributeGroup(v)
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

func expandAttributeGroup(v interface{}) []vault.GroupMapping {
	mappings := []vault.GroupMapping{}

	for k, v := range v.(map[string]interface{}) {
		mapping := vault.GroupMapping{}
		mapping.AttributeValue = k
		mapping.GroupName = v.(string)
		mappings = append(mappings, mapping)
	}
	logger.Debugf("Attribute group mappings: %+v", mappings)

	return mappings
}
