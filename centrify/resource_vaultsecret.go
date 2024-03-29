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

func resourceSecret_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Update: resourceSecretUpdate,
		Delete: resourceSecretDelete,
		Exists: resourceSecretExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema:             getSecretSchema(),
		DeprecationMessage: "resource centrifyvault_vaultsecret is deprecated will be removed in the future, use centrify_secret instead",
	}
}

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretCreate,
		Read:   resourceSecretRead,
		Update: resourceSecretUpdate,
		Delete: resourceSecretDelete,
		Exists: resourceSecretExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getSecretSchema(),
	}
}

func getSecretSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		// Workflow
		"workflow_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		//"workflow_default_options": getWorkflowDefaultOptionsSchema(),
		"workflow_approver": getWorkflowApproversSchema(),
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
	}
}

func resourceSecretExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Secret exist: %s", ResourceIDString(d))
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

	logger.Infof("Secret exists in tenant: %s", object.ID)
	return true, nil
}

func resourceSecretRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Secret: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewSecret object and populate ID attribute
	object := vault.NewSecret(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf(" Error reading Secret: %v", err)
	}
	//logger.Debugf("Secret from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceSecretRead(): %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "challenge_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		case "workflow_approver":
			d.Set(k, processBackupApproverSchema(v))
		default:
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading Secret: %s", object.Name)
	return nil
}

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Secret creation: %s", ResourceIDString(d))

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
		return fmt.Errorf(" Error creating Secret: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf(" Secret ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// 2nd step to update password checkout profile
	// Create API call doesn't set challenge profile so need to run update again
	err = getUpateGetSecretData(d, object)
	if err != nil {
		return err
	}

	if object.DataVaultDefaultProfile != "" || object.ChallengeRules != nil {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error updating Secret attribute: %v", err)
		}
	}

	if len(object.Sets) > 0 {
		err := object.AddToSetsByID(object.Sets)
		if err != nil {
			return err
		}
	}

	// add permissions
	if _, ok := d.GetOk("permission"); ok {
		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf(" Error setting secret permissions: %v", err)
		}
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Secret completed: %s", object.SecretName)
	return resourceSecretRead(d, m)
}

func resourceSecretUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Secret update: %s", ResourceIDString(d))

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
	if d.HasChanges("secret_name", "description", "secret_text", "folder_id", "type", "parent_path", "default_profile_id", "challenge_rule",
		"workflow_enabled", "workflow_approver") {
		// Special handling for default_profile_id. Whenever there is change, default_profile_id must be set otherwise default profile setting will be removed
		if v, ok := d.GetOk("default_profile_id"); ok && !d.HasChange("default_profile_id") {
			object.DataVaultDefaultProfile = v.(string)
		}

		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error updating Secret attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
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
				return fmt.Errorf(" Error removing secret from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf(" Error adding secret to Set: %v", err)
			}
		}
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
				return fmt.Errorf(" Error removing secret permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf(" Error adding secret permissions: %v", err)
			}
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of Secret completed: %s", object.Name)
	return resourceSecretRead(d, m)
}

func resourceSecretDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Secret: %s", ResourceIDString(d))
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
			return fmt.Errorf(" Error updating Secret attribute: %v", err)
		}
	}

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf(" Error deleting Secret: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Secret completed: %s", ResourceIDString(d))
	return nil
}

func getCreateSecretData(d *schema.ResourceData, object *vault.Secret) error {
	object.SecretName = d.Get("secret_name").(string)
	if v, ok := d.GetOk("description"); ok && d.HasChange("description") {
		object.Description = v.(string)
	}
	object.Type = d.Get("type").(string)
	if v, ok := d.GetOk("secret_text"); ok && d.HasChange("secret_text") {
		object.SecretText = v.(string)
	}
	if v, ok := d.GetOk("folder_id"); ok && d.HasChange("folder_id") {
		object.FolderID = v.(string)
	}
	if v, ok := d.GetOk("parent_path"); ok && d.HasChange("parent_path") {
		object.ParentPath = v.(string)
	}
	// Workflow
	if v, ok := d.GetOk("workflow_enabled"); ok {
		object.WorkflowEnabled = v.(bool)
	}
	//if v, ok := d.GetOk("workflow_default_options"); ok {
	//	object.WorkflowDefaultOptions = expandWorkflowDefaultOptions(v.(interface{}))
	//}
	if v, ok := d.GetOk("workflow_approver"); ok {
		object.WorkflowApprovers = expandWorkflowApprovers(v.([]interface{})) // This is a slice
	}

	return nil
}

func getUpateGetSecretData(d *schema.ResourceData, object *vault.Secret) error {
	getCreateSecretData(d, object)

	if v, ok := d.GetOk("default_profile_id"); ok && d.HasChange("default_profile_id") {
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
	if v, ok := d.GetOk("challenge_rule"); ok && d.HasChange("challenge_rule") {
		object.ChallengeRules = expandChallengeRules(v.([]interface{}))
		// Perform validations
		if err := validateChallengeRules(object.ChallengeRules); err != nil {
			return fmt.Errorf(" Schema setting error: %s", err)
		}
	}

	return nil
}
