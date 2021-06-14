package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceSecretFolder_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretFolderCreate,
		Read:   resourceSecretFolderRead,
		Update: resourceSecretFolderUpdate,
		Delete: resourceSecretFolderDelete,
		Exists: resourceSecretFolderExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema:             getSecretFolderSchema(),
		DeprecationMessage: "resource centrifyvault_vaultsecretfolder is deprecated will be removed in the future, use centrify_secretfolder instead",
	}
}

func resourceSecretFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretFolderCreate,
		Read:   resourceSecretFolderRead,
		Update: resourceSecretFolderUpdate,
		Delete: resourceSecretFolderDelete,
		Exists: resourceSecretFolderExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getSecretFolderSchema(),
	}
}

func getSecretFolderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the secret folder",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of an secret folder",
		},
		"parent_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Parent folder ID of an secret folder",
		},
		"parent_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Parent folder path of an secret folder",
		},
		"default_profile_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Default Secret Challenge Profile (used if no conditions matched)",
		},
		"permission":        getPermissionSchema(),
		"member_permission": getPermissionSchema(),
		"challenge_rule":    getChallengeRulesSchema(),
	}
}

func resourceSecretFolderExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking SecretFolder exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSecretFolder(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("SecretFolder exists in tenant: %s", object.ID)
	return true, nil
}

func resourceSecretFolderRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading SecretFolder: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewSecretFolder object and populate ID attribute
	object := vault.NewSecretFolder(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading SecretFolder: %v", err)
	}
	//logger.Debugf("SecretFolder from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceSecretFolderRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading SecretFolder: %s", object.Name)
	return nil
}

func resourceSecretFolderCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SecretFolder creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a SecretFolder object and populate all attributes
	object := vault.NewSecretFolder(client)
	err := getCreateSecretFolderData(d, object)
	if err != nil {
		return err
	}
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("error creating SecretFolder: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("SecretFolder ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("name")
	d.SetPartial("description")
	d.SetPartial("parent_id")

	// 2nd step to update SecretFolder login profile
	// Create API call doesn't set SecretFolder login profile so need to run update again
	err = getUpdateSecretFolderData(d, object)
	if err != nil {
		return err
	}

	if object.CollectionMembersDefaultProfile != "" || object.ChallengeRules != nil {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating SecretFolder attribute: %v", err)
		}
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
	}

	// Handle Set permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("error setting SecretFolder permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Handle Set member permissions
	if _, ok := d.GetOk("member_permission"); ok {
		_, err = object.SetMemberPermissions(false)
		if err != nil {
			return fmt.Errorf("error setting SecretFolder member permissions: %v", err)
		}
		d.SetPartial("member_permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of SecretFolder completed: %s", object.Name)
	return resourceSecretFolderRead(d, m)
}

func resourceSecretFolderUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning SecretFolder update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewSecretFolder(client)
	object.ID = d.Id()
	err := getUpdateSecretFolderData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "description", "default_profile_id", "challenge_rule") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating SecretFolder attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("name")
		d.SetPartial("description")
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
	}

	if d.HasChange("parent_id") {
		_, new := d.GetChange("parent_id")
		object.ParentID = new.(string)
		resp, err := object.MoveFolder()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating SecretFolder attribute: %v", err)
		}
		d.SetPartial("parent_id")
	}

	// Deal with permission changes
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
				return fmt.Errorf("error removing SecretFolder permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("error adding SecretFolder permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// Deal with member permission changes
	if d.HasChange("member_permission") {
		old, new := d.GetChange("member_permission")
		// We don't want to care the details of changes
		// So, let's first remove the old permissions
		if old != nil {
			var err error
			object.MemberPermissions, err = expandPermissions(old, object.ValidMemberPermissions, false)
			if err != nil {
				return err
			}
			_, err = object.SetMemberPermissions(true)
			if err != nil {
				return fmt.Errorf("error removing SecretFolder member permissions: %v", err)
			}
		}

		if new != nil {
			var err error
			object.MemberPermissions, err = expandPermissions(new, object.ValidMemberPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetMemberPermissions(false)
			if err != nil {
				return fmt.Errorf("error adding SecretFolder member permissions: %v", err)
			}
		}
		d.SetPartial("member_permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of SecretFolder completed: %s", object.Name)
	return resourceSecretFolderRead(d, m)
}

func resourceSecretFolderDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of SecretFolder: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSecretFolder(client)
	object.ID = d.Id()

	// Remove challenge profile first otherwise deletion will fail
	err := getUpdateSecretFolderData(d, object)
	if err != nil {
		return err
	}
	if object.CollectionMembersDefaultProfile != "" || object.ChallengeRules != nil {
		object.CollectionMembersDefaultProfile = ""
		object.ChallengeRules = nil
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating SecretFolder attribute: %v", err)
		}
	}

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("error deleting SecretFolder: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of SecretFolder completed: %s", ResourceIDString(d))
	return nil
}

func getCreateSecretFolderData(d *schema.ResourceData, object *vault.SecretFolder) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("parent_id"); ok {
		object.ParentID = v.(string)
	}

	return nil
}

func getUpdateSecretFolderData(d *schema.ResourceData, object *vault.SecretFolder) error {
	getCreateSecretFolderData(d, object)

	if v, ok := d.GetOk("default_profile_id"); ok {
		object.CollectionMembersDefaultProfile = v.(string)
	}
	if v, ok := d.GetOk("permission"); ok {
		var err error
		object.Permissions, err = expandPermissions(v, object.ValidPermissions, true)
		if err != nil {
			return err
		}
	}
	if v, ok := d.GetOk("member_permission"); ok {
		var err error
		object.MemberPermissions, err = expandPermissions(v, object.ValidMemberPermissions, true)
		if err != nil {
			return err
		}
	}
	// Challenge rules
	if v, ok := d.GetOk("challenge_rule"); ok {
		object.ChallengeRules = expandChallengeRules(v.([]interface{}))
		// Perform validations
		if err := validateChallengeRules(object.ChallengeRules); err != nil {
			return fmt.Errorf(" Schema setting error: %s", err)
		}
	}

	return nil
}
