package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceFederatedGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFederatedGroupCreate,
		Read:   resourceFederatedGroupRead,
		// Update isn't supported because there is only name attribute
		// Change of name attribute results in re-creation of group
		//Update: resourceFederatedGroupUpdate,
		Delete: resourceFederatedGroupDelete,
		Exists: resourceFederatedGroupExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getFederatedGroupSchema(),
	}
}

func getFederatedGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The name of the federated group",
		},
	}
}

func resourceFederatedGroupExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking federated group exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewFederatedGroup(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Federated group exists in tenant: %s", object.ID)
	return true, nil
}

func resourceFederatedGroupRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading federated group: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a role object and populate ID attribute
	object := vault.NewFederatedGroup(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading federated group: %v", err)
	}
	logger.Debugf("Federated group from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceFederatedGroupRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading federated group: %s", object.Name)
	return nil
}

func resourceFederatedGroupCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning federated group creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)

	// Create a role object and populate all attributes
	object := vault.NewFederatedGroup(client)
	createUpateGetFederatedGroupData(d, object)

	// Response contains only federated group id
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating federated group: %v", err)
	}

	id := resp
	if id == "" {
		return fmt.Errorf("the federated group ID is not set")
	}
	d.SetId(id)
	object.ID = id

	logger.Infof("Creation of federated group completed: %s", object.Name)
	return resourceFederatedGroupRead(d, m)
}

func resourceFederatedGroupDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Deletion of federated group isn't supported. Update terraform state only but not upstream application")
	d.SetId("")

	return nil
}

func createUpateGetFederatedGroupData(d *schema.ResourceData, object *vault.FederatedGroup) error {
	object.Name = d.Get("name").(string)

	return nil
}
