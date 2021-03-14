package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

/***** TO DO **********
To determine when to use host_id, database_id or domain_id
***********************/
func resourceVaultAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVaultAccountCreate,
		Read:   resourceVaultAccountRead,
		Update: resourceVaultAccountUpdate,
		Delete: resourceVaultAccountDelete,
		Exists: resourceVaultAccountExists,

		Schema: map[string]*schema.Schema{
			// Settings menu
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the account",
			},
			"credential_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Either password or sshkey",
				ValidateFunc: validation.StringInSlice([]string{
					"Password",
					"SshKey",
					"AwsAccessKey",
				}, false),
			},
			"sshkey_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"password", "checkout_lifetime", "default_profile_id"},
				Description:   "ID of SSH key",
			},
			"password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"sshkey_id"},
				Description:   "Password of the account",
			},
			"host_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"domain_id", "database_id", "cloudprovider_id"},
				Description:   "ID of the system it belongs to",
			},
			"domain_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"host_id", "database_id", "cloudprovider_id"},
				Description:   "ID of the domain it belongs to",
			},
			"database_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"domain_id", "host_id", "cloudprovider_id"},
				Description:   "ID of the database it belongs to",
			},
			"cloudprovider_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"domain_id", "host_id", "database_id"},
				Description:   "ID of the cloud provider it belongs to",
			},
			// Optional attributes
			"is_admin_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this is an administrative account",
			},
			"is_root_account": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this is an root account for cloud provider",
			},
			"use_proxy_account": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"sshkey_id", "database_id", "domain_id", "cloudprovider_id"},
				Description:   "Use proxy account to manage this account",
			},
			"managed": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"cloudprovider_id"},
				Description:   "If this account is managed",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the account",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Policy menu
			"checkout_lifetime": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"sshkey_id"},
				Description:   "Checkout lifetime (minutes)",
				ValidateFunc:  validation.IntBetween(15, 2147483647),
			},
			"default_profile_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"sshkey_id"},
				Description:   "Default password checkout profile id",
			},
			"access_secret_checkout_default_profile_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"sshkey_id", "host_id", "domain_id", "database_id", "default_profile_id", "challenge_rule"},
				Description:   "Default secret access key checkout challenge rule id",
			},
			"access_secret_checkout_rule": getChallengeRulesSchema(),
			// Add to Sets
			"sets": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Sets",
			},
			"permission":     getPermissionSchema(),
			"challenge_rule": getChallengeRulesSchema(),
			"access_key":     getAccessKeySchema(),
		},
	}
}

func resourceVaultAccountExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking VaultAccount exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewAccount(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("VaultAccount exists in tenant: %s", object.ID)
	return true, nil
}

func resourceVaultAccountRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading VaultAccount: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewAccount object and populate ID attribute
	object := vault.NewAccount(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VaultAccount: %v", err)
	}
	//logger.Debugf("VaultAccount from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceVaultAccountRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading VaultAccount: %s", object.Name)
	return nil
}

func resourceVaultAccountCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultAccount creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)

	// Create a VaultAccount object and populate all attributes
	object := vault.NewAccount(client)
	err := createUpateGetAccountData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating VaultAccount: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("VaultAccount ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// 2nd step to update password checkout profile
	// Create API call doesn't set challenge profile so need to run update again
	if object.PasswordCheckoutDefaultProfile != "" {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultAccount attribute: %v", err)
		}
		d.SetPartial("default_profile_id")
	}

	// Add to Sets
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
			return fmt.Errorf("Error setting VaultAccount permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// set as admin account
	if object.IsAdminAccount {
		err := object.SetAdminAccount(object.IsAdminAccount)
		if err != nil {
			return fmt.Errorf("Error setting VaultAccount as administrative account: %v", err)
		}
		d.SetPartial("is_admin_account")
	}

	// add IAM account access key
	if len(object.AccessKeys) > 0 {
		logger.Debugf("Adding access key...")
		for _, v := range object.AccessKeys {
			err := object.SafeAddAccessKey(v)
			if err != nil {
				return fmt.Errorf("Error adding access key %s : %v", v.AccessKeyID, err)
			}
		}
		d.SetPartial("access_key")

	}

	// Creation completed
	logger.Infof("Creation of VaultAccount completed: %s", object.User)
	return resourceVaultAccountRead(d, m)
}

func resourceVaultAccountUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning VaultAccount update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewAccount(client)
	object.ID = d.Id()
	err := createUpateGetAccountData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "credential_type", "host_id", "domain_id", "database_id", "cloudprovider_id", "sshkey_id", "description",
		"use_proxy_account", "managed", "checkout_lifetime", "default_profile_id", "challenge_rule") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultAccount attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("name")
		d.SetPartial("credential_type")
		d.SetPartial("sshkey_id")
		d.SetPartial("host_id")
		d.SetPartial("domain_id")
		d.SetPartial("database_id")
		d.SetPartial("cloudprovider_id")
		d.SetPartial("description")
		d.SetPartial("use_proxy_account")
		d.SetPartial("managed")
		d.SetPartial("checkout_lifetime")
		d.SetPartial("default_profile_id")
		d.SetPartial("challenge_rule")
	}

	// Deal with Set member
	if d.HasChange("sets") {
		old, new := d.GetChange("sets")
		// Remove old Sets
		for _, v := range flattenSchemaSetToStringSlice(old) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "remove")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error removing VaultAccount from Set: %v", err)
			}
		}
		// Add new Sets
		for _, v := range flattenSchemaSetToStringSlice(new) {
			setObj := vault.NewManualSet(client)
			setObj.ID = v
			setObj.ObjectType = object.SetType
			resp, err := setObj.UpdateSetMembers([]string{object.ID}, "add")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding VaultAccount to Set: %v", err)
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
				return fmt.Errorf("Error removing VaultAccount permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding VaultAccount permissions: %v", err)
			}
		}
		d.SetPartial("permission")
	}

	// Change password
	if d.HasChange("password") {
		resp, err := object.ChangePassword()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultAccount password: %v", err)
		}
		d.SetPartial("password")
	}

	// Handle admin account
	if d.HasChange("is_admin_account") {
		err := object.SetAdminAccount(object.IsAdminAccount)
		if err != nil {
			return fmt.Errorf("Error setting VaultAccount as administrative account: %v", err)
		}
		d.SetPartial("is_admin_account")
	}

	// Deal with access key
	if d.HasChange("access_key") {
		old, new := d.GetChange("access_key")
		// Remove the old access keys
		m := old.(*schema.Set).List()
		if m != nil {
			for _, v := range m {
				id := v.(map[string]interface{})["id"].(string)
				keyid := v.(map[string]interface{})["access_key_id"].(string)
				if id != "" {
					err := object.DeleteAccessKey(id)
					if err != nil {
						return fmt.Errorf("Error deleting access key %s : %v", keyid, err)
					}
					logger.Debugf("Deleted old key: %+v", keyid)
				}
			}
		}
		// Add the new access keys
		m = new.(*schema.Set).List()
		if m != nil {
			for _, v := range m {
				keyid := v.(map[string]interface{})["access_key_id"].(string)
				secretkey := v.(map[string]interface{})["secret_access_key"].(string)
				key := vault.AccessKey{}
				key.AccessKeyID = keyid
				key.SecretAccessKey = secretkey
				if keyid != "" {
					err := object.SafeAddAccessKey(key)
					if err != nil {
						return fmt.Errorf("Error adding access key %s : %v", keyid, err)
					}
					logger.Debugf("Added new key: %+v", keyid)
				}
			}
		}

		d.SetPartial("access_key")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of VaultAccount completed: %s", object.Name)
	return resourceVaultAccountRead(d, m)
}

func resourceVaultAccountDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of VaultAccount: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewAccount(client)
	object.ID = d.Id()
	// check if this is an admin account. If so, clear it first otherwise deletion will fail
	if v, ok := d.GetOk("is_admin_account"); ok {
		if v, ok := d.GetOk("host_id"); ok {
			object.Host = v.(string)
		}
		if v.(bool) {
			err := object.SetAdminAccount(false)
			if err != nil {
				return fmt.Errorf("Error clearing VaultAccount as administrative account: %v", err)
			}
		}
	}

	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting VaultAccount: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of VaultAccount completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetAccountData(d *schema.ResourceData, object *vault.Account) error {
	object.User = d.Get("name").(string)
	if v, ok := d.GetOk("credential_type"); ok {
		object.CredentialType = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		object.Password = v.(string)
	}
	if v, ok := d.GetOk("sshkey_id"); ok {
		object.SSHKeyID = v.(string)
	}
	if v, ok := d.GetOk("host_id"); ok {
		object.Host = v.(string)
	}
	if v, ok := d.GetOk("domain_id"); ok {
		object.DomainID = v.(string)
	}
	if v, ok := d.GetOk("database_id"); ok {
		object.DatabaseID = v.(string)
	}
	if v, ok := d.GetOk("cloudprovider_id"); ok {
		object.CloudProviderID = v.(string)
	}
	// Optional attributes
	if v, ok := d.GetOk("is_admin_account"); ok {
		object.IsAdminAccount = v.(bool)
	}
	if v, ok := d.GetOk("is_root_account"); ok {
		object.IsRootAccount = v.(bool)
	}
	if v, ok := d.GetOk("use_proxy_account"); ok {
		object.UseWheel = v.(bool)
	}
	if v, ok := d.GetOk("managed"); ok {
		object.IsManaged = v.(bool)
	}
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("checkout_lifetime"); ok {
		object.DefaultCheckoutTime = v.(int)
	}
	if v, ok := d.GetOk("default_profile_id"); ok {
		object.PasswordCheckoutDefaultProfile = v.(string)
	}
	if v, ok := d.GetOk("sets"); ok {
		object.Sets = flattenSchemaSetToStringSlice(v)
	}
	if v, ok := d.GetOk("access_key"); ok {
		object.AccessKeys = expandAccessKeys(v)
	}
	if v, ok := d.GetOk("access_secret_checkout_default_profile_id"); ok {
		object.AccessSecretCheckoutDefaultProfile = v.(string)
	}
	// Permissions
	if v, ok := d.GetOk("permission"); ok {
		var err error
		object.ResolveValidPermissions()
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
	// Secret Access Key checkout Challenge rules
	if v, ok := d.GetOk("access_secret_checkout_rule"); ok {
		object.AccessSecretCheckoutRules = expandChallengeRules(v.([]interface{}))
		// Perform validations
		if err := validateChallengeRules(object.ChallengeRules); err != nil {
			return fmt.Errorf("Schema setting error: %s", err)
		}
	}

	// Perform validations
	if err := object.ValidateCredentialType(); err != nil {
		logger.Errorf("there is error: %s", err)
		return fmt.Errorf("Schema setting error: %s", err)
	}
	return nil
}
