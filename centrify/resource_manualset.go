package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/settype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceManualSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceManualSetCreate,
		Read:   resourceManualSetRead,
		Update: resourceManualSetUpdate,
		Delete: resourceManualSetDelete,
		Exists: resourceManualSetExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the manual set",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of an manual set",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				// Server -> Systems
				// Subscriptions -> Services
				// DataVault -> Secrets
				Description: "Type of set. Valid values are: Server, VaultAccount, VaultDatabase, VaultDomain, DataVault, SshKeys, Subscriptions, Application, ResourceProfiles",
				ValidateFunc: validation.StringInSlice([]string{
					settype.System.String(),
					settype.Account.String(),
					settype.Database.String(),
					settype.Domain.String(),
					settype.Secret.String(),
					settype.SSHKey.String(),
					settype.Service.String(),
					settype.Application.String(),
					settype.ResourceProfile.String(),
					settype.CloudProvider.String(),
				}, false),
			},
			"subtype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SubObjectType for application. Valid values are: Web and Desktop",
				ValidateFunc: validation.StringInSlice([]string{
					"Web",
					"Desktop",
				}, false),
			},
			"permission":        getPermissionSchema(),
			"member_permission": getPermissionSchema(),
		},
	}
}

func resourceManualSetExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking Manual Set exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewManualSet(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Manual Set exists in tenant: %s", object.ID)
	return true, nil
}

func resourceManualSetRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading Manual Set: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a Manual Set object and populate ID attribute
	object := vault.NewManualSet(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Manual Set: %v", err)
	}
	//logger.Debugf("Manual Set from tenant: %v", object)

	d.Set("name", object.Name)
	d.Set("description", object.Description)

	logger.Infof("Completed reading Manual Set: %s", object.Name)
	return nil
}

func resourceManualSetCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Manual Set creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a manual set object and populate all attributes
	object, err := vault.NewManualSetWithType(client, d.Get("type").(string))
	if err != nil {
		return err
	}

	err = createUpateGetManualSetData(d, object)
	if err != nil {
		return err
	}

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating Manual Set: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("Manual Set ID is not set")
	}
	d.SetId(id)
	// Creation partially completed
	d.SetPartial("name")
	d.SetPartial("type")
	d.SetPartial("subtype")
	d.SetPartial("description")
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Handle Set permissions
	if _, ok := d.GetOk("permission"); ok {

		_, err = object.SetPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting Manual Set permissions: %v", err)
		}
		d.SetPartial("permission")
	}

	// Handle Set member permissions
	if _, ok := d.GetOk("member_permission"); ok {
		_, err = object.SetMemberPermissions(false)
		if err != nil {
			return fmt.Errorf("Error setting Manual Set member permissions: %v", err)
		}
		d.SetPartial("member_permission")
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of Manual Set completed: %s", object.Name)
	return resourceManualSetRead(d, m)
}

func resourceManualSetUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Manual Set update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object, err := vault.NewManualSetWithType(client, d.Get("type").(string))
	if err != nil {
		return err
	}
	// Both ID and Name must be set
	object.ID = d.Id()
	err = createUpateGetManualSetData(d, object)
	if err != nil {
		return err
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "description") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating Manual Set attribute: %v", err)
		}
		//logger.Debugf("Updated attributes to: %v", object)
		d.SetPartial("name")
		d.SetPartial("description")
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
				return fmt.Errorf("Error removing Manual Set permissions: %v", err)
			}
		}

		if new != nil {
			object.Permissions, err = expandPermissions(new, object.ValidPermissions, true)
			if err != nil {
				return err
			}
			_, err = object.SetPermissions(false)
			if err != nil {
				return fmt.Errorf("Error adding Manual Set permissions: %v", err)
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
				return fmt.Errorf("Error removing Manual Set member permissions: %v", err)
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
				return fmt.Errorf("Error adding Manual Set member permissions: %v", err)
			}
		}
		d.SetPartial("member_permission")
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of Manual Set completed: %s", object.Name)
	return resourceManualSetRead(d, m)
}

func resourceManualSetDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of Manual Set: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewManualSet(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting Manual Set: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of Manual Set completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetManualSetData(d *schema.ResourceData, object *vault.ManualSet) error {
	object.Name = d.Get("name").(string)
	object.ObjectType = d.Get("type").(string)
	object.SubObjectType = d.Get("subtype").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
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

	return nil
}
