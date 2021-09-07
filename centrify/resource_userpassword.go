package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceUserPassword_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPasswordCreate,
		Read:   resourceUserPasswordRead,
		Update: resourceUserPasswordUpdate,
		Delete: resourceUserPasswordDelete,

		Schema:             getUserPasswordSchema(),
		DeprecationMessage: "resource centrifyvault_userpassword is deprecated will be removed in the future, use centrify_userpassword instead",
	}
}

func resourceUserPassword() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPasswordCreate,
		Read:   resourceUserPasswordRead,
		Update: resourceUserPasswordUpdate,
		Delete: resourceUserPasswordDelete,

		Schema: getUserPasswordSchema(),
	}
}

func getUserPasswordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_uuid": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The uuid of Centrify Directory User",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "New password of the user",
		},
	}
}

func resourceUserPasswordRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading user password: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewUser object and populate ID attribute
	object := vault.NewUser(client)
	object.ID = d.Get("user_uuid").(string)
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading user: %v", err)
	}
	//logger.Debugf("User from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceUserPasswordRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading user password: %s", object.Name)
	return nil
}

func resourceUserPasswordCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning user password creation: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewUser object and populate all attributes
	object := vault.NewUser(client)
	object.ID = d.Get("user_uuid").(string)
	createUpateGetUserPasswordData(d, object)

	// Change password
	if d.HasChange("password") {
		resp, err := object.ChangePassword()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating user password: %v", err)
		}
	}

	d.SetId(object.ID)

	logger.Infof("Creation of user password completed: %s", object.Name)
	return resourceUserPasswordRead(d, m)
}

func resourceUserPasswordUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning user password update: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object := vault.NewUser(client)
	object.ID = d.Id()
	createUpateGetUserPasswordData(d, object)

	// Change password
	if d.HasChange("password") {
		resp, err := object.ChangePassword()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating user password: %v", err)
		}
	}

	logger.Infof("Updating of user password completed: %s", object.Name)
	return resourceUserPasswordRead(d, m)
}

func resourceUserPasswordDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of user: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewUser(client)
	object.ID = d.Id()

	d.SetId("")
	logger.Infof("Deletion of user password completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetUserPasswordData(d *schema.ResourceData, object *vault.User) error {
	object.Password = d.Get("password").(string)

	return nil
}
