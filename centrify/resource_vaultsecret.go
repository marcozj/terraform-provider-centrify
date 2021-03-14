package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/secrettype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceVaultSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultSecretCreate,
		Read:   resourceVaultSecretRead,
		Update: resourceVaultSecretUpdate,
		Delete: resourceVaultSecretDelete,
		Exists: resourceVaultSecretExists,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the secret",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the secret",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Either Text or File",
				ValidateFunc: validation.StringInSlice([]string{
					secrettype.Text.String(),
					//secrettype.File.String(),
				}, false),
			},
			"secret_text": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Content of the secret",
			},
			"folder_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the folder where the secret is located",
			},
			"parent_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Path of parent folder",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default Secret Challenge Profile (used if no conditions matched)",
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

func resourceVaultSecretExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking VaultSecret exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSecret(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("VaultSecret exists in tenant: %s", object.ID)
	return true, nil
}

func resourceVaultSecretRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading VaultSecret: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewSecret object and populate ID attribute
	object := vault.NewSecret(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VaultSecret: %v", err)
	}
	//logger.Debugf("VaultSecret from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceVaultSecretRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading VaultSecret: %s", object.Name)
	return nil
}

func resourceVaultSecretCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultSecret creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a Secret object and populate all attributes
	object := vault.NewSecret(client)
	err := getCreateSecretData(d, object)
	if err != nil {
		return err
	}
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating VaultSecret: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("VaultSecret ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("secret_name")
	d.SetPartial("description")
	d.SetPartial("secret_text")
	d.SetPartial("type")
	d.SetPartial("folder_id")
	d.SetPartial("parent_path")

	// 2nd step to update password checkout profile
	// Create API call doesn't set challenge profile so need to run update again
	err = getUpateGetSecretData(d, object)
	if err != nil {
		return err
	}

	if object.DataVaultDefaultProfile != "" || object.ChallengeRules != nil {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultAccount attribute: %v", err)
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
			return fmt.Errorf("Error setting secret permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of VaultSecret completed: %s", object.SecretName)
	return resourceVaultSecretRead(d, m)
}

func resourceVaultSecretUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultSecret update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewSecret(client)
	object.ID = d.Id()
	err := getUpateGetSecretData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("secret_name", "description", "secret_text", "folder_id", "type", "parent_path", "default_profile_id", "challenge_rule") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultSecret attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("secret_name")
		d.SetPartial("description")
		d.SetPartial("secret_text")
		d.SetPartial("folder_id")
		d.SetPartial("type")
		d.SetPartial("parent_path")
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
				return fmt.Errorf("Error removing secret from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding secret to Set: %v", err)
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
				return fmt.Errorf("Error removing secret permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding secret permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of VaultSecret completed: %s", object.Name)
	return resourceVaultSecretRead(d, m)
}

func resourceVaultSecretDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of VaultSecret: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewSecret(client)
	object.ID = d.Id()

	// Remove challenge profile first otherwise deletion will fail
	err := getUpateGetSecretData(d, object)
	if err != nil {
		return err
	}
	if object.DataVaultDefaultProfile != "" || object.ChallengeRules != nil {
		object.DataVaultDefaultProfile = ""
		object.ChallengeRules = nil
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultSecret attribute: %v", err)
		}
	}

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting VaultSecret: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of VaultSecret completed: %s", ResourceIDString(d))
	return nil
}

func getCreateSecretData(d *schema.ResourceData, object *vault.Secret) error {
	object.SecretName = d.Get("secret_name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		object.Type = v.(string)
	}
	if v, ok := d.GetOk("secret_text"); ok {
		object.SecretText = v.(string)
	}
	if v, ok := d.GetOk("folder_id"); ok {
		object.FolderID = v.(string)
	}
	if v, ok := d.GetOk("parent_path"); ok {
		object.ParentPath = v.(string)
	}

	return nil
}

func getUpateGetSecretData(d *schema.ResourceData, object *vault.Secret) error {
	getCreateSecretData(d, object)

	if v, ok := d.GetOk("default_profile_id"); ok {
		object.DataVaultDefaultProfile = v.(string)
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
