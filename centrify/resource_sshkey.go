package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSHKeyCreate,
		Read:   resourceSSHKeyRead,
		Update: resourceSSHKeyUpdate,
		Delete: resourceSSHKeyDelete,
		Exists: resourceSSHKeyExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the SSH Key",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the SSH Key",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "SSH private key",
			},
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Passphrase to use for encrypting the PrivateKey",
			},
			"key_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default SSH Key Challenge Profile",
			},
			// Add to Sets
			"sets": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				Set: schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Sets",
			},
			"permission":     getPermissionSchema(),
			"challenge_rule": getChallengeRulesSchema(),
		},
	}
}

func resourceSSHKeyExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking SSH Key exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSSHKey(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("SSH Key exists in tenant: %s", object.ID)
	return true, nil
}

func resourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading SSH Key: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewVaultSecret object and populate ID attribute
	object := vault.NewSSHKey(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SSH Key: %v", err)
	}
	//logger.Debugf("SSH Key from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceSSHKeyRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading SSH Key: %s", object.Name)
	return nil
}

func resourceSSHKeyCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SSH Key creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a SSH Key object and populate all attributes
	object := vault.NewSSHKey(client)
	err := createUpateGetSSHKeyData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating SSH Key: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("SSH Key ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("name")
	d.SetPartial("description")
	d.SetPartial("private_key")

	// 2nd step to update challenge login profile
	// Create API call doesn't set challenge profile so need to run update again
	if object.SSHKeysDefaultProfileID != "" || object.ChallengeRules != nil {
		object.PrivateKey = ""
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating SSH Key attribute: %v", err)
		}
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
	}

	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
		d.SetPartial("sets")
	}

	// add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting SSH Key permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of SSH Key completed: %s", object.Name)
	return resourceSSHKeyRead(d, m)
}

func resourceSSHKeyUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SSH Key update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewSSHKey(client)
	object.ID = d.Id()
	err := createUpateGetSSHKeyData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "description", "private_key", "default_profile_id", "challenge_rule") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating SSH Key attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("name")
		d.SetPartial("description")
		d.SetPartial("private_key")
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
	}

	if d.HasChange("sets") {
		old, new := d.GetChange("sets")
		// Remove old Sets
		for _, v := range flattenSchemaSetToStringSlice(old) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "remove")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error removing SSH Key from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding SSH Key to Set: %v", err)
			}
		}
		d.SetPartial("sets")
	}

	// Deal with Permissions
	if d.HasChange("permission") {
		old, new := d.GetChange("permission")
		// We don't want to care the details of changes
		// So, let's first remove the old permissions
		var err error
		if old != nil {
			// do not validate old values
			object.Permissions, err = expandPermissions(old, object.ValidPermissions, false)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(true)
			if err != nil {
				return fmt.Errorf("Error removing SSH Key permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding SSH Key permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of SSH Key completed: %s", object.Name)
	return resourceSSHKeyRead(d, m)
}

func resourceSSHKeyDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of SSH Key: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSSHKey(client)
	object.ID = d.Id()

	// Remove challenge profile first otherwise deletion will fail
	createUpateGetSSHKeyData(d, object)
	if object.SSHKeysDefaultProfileID != "" {
		object.SSHKeysDefaultProfileID = ""
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating SSH Key attribute: %v", err)
		}
	}

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting SSH Key: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of SSH Key completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetSSHKeyData(d *schema.ResourceData, object *vault.SSHKey) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("private_key"); ok {
		object.PrivateKey = v.(string)
	}
	if v, ok := d.GetOk("passphrase"); ok {
		object.Passphrase = v.(string)
	}
	if v, ok := d.GetOk("default_profile_id"); ok {
		object.SSHKeysDefaultProfileID = v.(string)
	}
	if v, ok := d.GetOk("sets"); ok {
		object.Sets = flattenSchemaSetToStringSlice(v)
	}
	// Permissions
	if v, ok := d.GetOk("permission"); ok {
		var err error
		object.Permissions, err = expandPermissions(v, object.ValidPermissions, true)
		if err != nil {
			return err
		}
	}
	// Challenge rules
	if v, ok := d.GetOk("challenge_rule"); ok {
		object.ChallengeRules = expandChallengeRules(v.([]interface{}))
		// Perform validations
		if err := validateChallengeRules(object.ChallengeRules); err != nil {
			return fmt.Errorf("Schema setting error: %s", err)
		}
	}

	return nil
}
