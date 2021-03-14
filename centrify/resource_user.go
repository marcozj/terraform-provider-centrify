package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Exists: resourceUserExists,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username in loginid@suffix format",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address",
			},
			"displayname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Display name",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the user",
			},
			"confirm_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the user",
			},
			"password_never_expire": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Password never expires",
			},
			"force_password_change_next": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Require password change at next login",
			},
			"oauth_client": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Is OAuth confidential client",
			},
			"send_email_invite": {
				Type:     schema.TypeBool,
				Optional: true,
				//Default:     true,
				Description: "Send email invite for user profile setup",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the user",
			},
			"office_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"home_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mobile_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_mfa_user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Redirect multi factor authentication to a different user account (UUID value)",
			},
			"manager_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Username of the manager",
			},
			// Add to roles
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				Set: schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Add to list of Roles",
			},
		},
	}
}

func resourceUserExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking user exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewUser(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("User exists in tenant: %s", object.ID)
	return true, nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading user: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a NewUser object and populate ID attribute
	object := vault.NewUser(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading user: %v", err)
	}
	//logger.Debugf("User from tenant: %+v", object)
	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceUserRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading user: %s", object.Name)
	return nil
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning user creation: %s", ResourceIDString(d))
	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a NewUser object and populate all attributes
	object := vault.NewUser(client)
	createUpateGetUserData(d, object)

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("User ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	d.SetPartial("username")
	d.SetPartial("email")
	d.SetPartial("displayname")
	d.SetPartial("password_never_expire")
	d.SetPartial("orce_password_change_next")
	d.SetPartial("oauth_client")
	d.SetPartial("office_number")
	d.SetPartial("home_number")
	d.SetPartial("mobile_number")
	d.SetPartial("redirect_mfa_user_id")
	d.SetPartial("manager_username")

	// 2rd step to add system to Sets
	if len(object.Roles) > 0 {
		for _, v := range object.Roles {
			roleObj := vault.NewRole(client)
			roleObj.ID = v
			resp, err := roleObj.UpdateMembers([]string{object.ID}, "Add", "Users")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding user to role: %v", err)
			}
		}
		d.SetPartial("roles")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of user completed: %s", object.Name)
	return resourceUserRead(d, m)
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning user update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewUser(client)
	object.ID = d.Id()
	createUpateGetUserData(d, object)

	// Deal with normal attribute changes first
	if d.HasChanges("name", "email", "displayname", "password_never_expire", "force_password_change_next", "oauth_client", "send_email_invite",
		"description", "office_number", "home_number", "mobile_number", "redirect_mfa_user_id", "manager_username") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating VaultAccount attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
		d.SetPartial("name")
		d.SetPartial("email")
		d.SetPartial("displayname")
		d.SetPartial("password_never_expire")
		d.SetPartial("force_password_change_next")
		d.SetPartial("oauth_client")
		d.SetPartial("send_email_invite")
		d.SetPartial("description")
		d.SetPartial("office_number")
		d.SetPartial("home_number")
		d.SetPartial("mobile_number")
		d.SetPartial("redirect_mfa_user_id")
		d.SetPartial("manager_username")
	}

	// Change password
	if d.HasChange("password") {
		resp, err := object.ChangePassword()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating user password: %v", err)
		}
		d.SetPartial("password")
	}

	// Deal with role member
	if d.HasChange("roles") {
		old, new := d.GetChange("roles")
		// Remove old roles
		for _, v := range flattenSchemaSetToStringSlice(old) {
			roleObj := vault.NewRole(client)
			roleObj.ID = v
			resp, err := roleObj.UpdateMembers([]string{object.ID}, "Delete", "Users")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error removing user from role: %v", err)
			}
		}
		// Add new roles
		for _, v := range flattenSchemaSetToStringSlice(new) {
			roleObj := vault.NewRole(client)
			roleObj.ID = v
			resp, err := roleObj.UpdateMembers([]string{object.ID}, "Add", "Users")
			if err != nil || !resp.Success {
				return fmt.Errorf("Error adding user to role: %v", err)
			}
		}
		d.SetPartial("roles")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of user completed: %s", object.Name)
	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of user: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewUser(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting user: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of user completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetUserData(d *schema.ResourceData, object *vault.User) error {
	object.Name = d.Get("username").(string)
	if v, ok := d.GetOk("email"); ok {
		object.Mail = v.(string)
	}
	if v, ok := d.GetOk("displayname"); ok {
		object.DisplayName = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		object.Password = v.(string)
	}
	if v, ok := d.GetOk("confirm_password"); ok {
		object.ConfirmPassword = v.(string)
	}
	if v, ok := d.GetOk("password_never_expire"); ok {
		object.PasswordNeverExpire = v.(bool)
	}
	if v, ok := d.GetOk("force_password_change_next"); ok {
		object.ForcePasswordChangeNext = v.(bool)
	}
	if v, ok := d.GetOk("oauth_client"); ok {
		object.OauthClient = v.(bool)
	}
	if v, ok := d.GetOk("send_email_invite"); ok {
		object.SendEmailInvite = v.(bool)
	}
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("office_number"); ok {
		object.OfficeNumber = v.(string)
	}
	if v, ok := d.GetOk("home_number"); ok {
		object.HomeNumber = v.(string)
	}
	if v, ok := d.GetOk("mobile_number"); ok {
		object.MobileNumber = v.(string)
	}
	//if v, ok := d.GetOk("redirect_mfa"); ok {
	//		object.RedirectMFA = v.(bool)
	//}
	if v, ok := d.GetOk("redirect_mfa_user_id"); ok {
		object.RedirectMFAUserID = v.(string)
	}
	if v, ok := d.GetOk("manager_username"); ok {
		object.ReportsTo = v.(string)
	}
	if v, ok := d.GetOk("roles"); ok {
		object.Roles = flattenSchemaSetToStringSlice(v)
	}

	return nil
}
