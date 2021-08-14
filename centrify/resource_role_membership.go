package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceRoleMembership_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleMembershipCreate,
		Read:   resourceRoleMembershipRead,
		Update: resourceRoleMembershipUpdate,
		Delete: resourceRoleMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema:             getRoleMembershipSchema(),
		DeprecationMessage: "resource centrifyvault_role_membership is deprecated will be removed in the future, use centrify_role_membership instead",
	}
}

func resourceRoleMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleMembershipCreate,
		Read:   resourceRoleMembershipRead,
		Update: resourceRoleMembershipUpdate,
		Delete: resourceRoleMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getRoleMembershipSchema(),
	}
}

func getRoleMembershipSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"role_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "ID of the role",
		},
		"member": {
			Type:     schema.TypeSet,
			Optional: true,
			Set:      customRoleMemberHash,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "ID of the member",
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
						Description: "Name of the member",
					},
					"type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Type of the member",
						ValidateFunc: validation.StringInSlice([]string{
							"User",
							"Group",
							"Role",
						}, false),
					},
				},
			},
		},
	}
}

func resourceRoleMembershipRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading role membership: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a role object and populate ID attribute
	object := vault.NewRoleMembership(client)
	object.ID = d.Id()
	object.RoleID = d.Get("role_id").(string)
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf(" Error reading role: %v", err)
	}
	logger.Debugf("Role from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceRoleMembershipRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading role membership: %s", object.Name)
	return nil
}

func resourceRoleMembershipCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning role membership creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a role object and populate all attributes
	object := vault.NewRoleMembership(client)
	createUpateGetRoleMembershipData(d, object)

	// Handle role members
	if len(object.Members) > 0 {
		resp, err := object.UpdateRoleMembers(object.Members, "Add")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error adding members to role: %v", err)
		}
	}

	//d.SetId(d.Get("name").(string))
	d.SetId(object.RoleID)
	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of role membership completed: %s", object.Name)
	return resourceRoleMembershipRead(d, m)
}

func resourceRoleMembershipUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning role membership update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewRoleMembership(client)
	object.ID = d.Id()
	createUpateGetRoleMembershipData(d, object)

	// Deal with role members
	if d.HasChange("member") {
		old, new := d.GetChange("member")
		// Remove old members
		resp, err := object.UpdateRoleMembers(expandRoleMembers(old), "Delete")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Failed to remove members from role: %v", err)
		}
		// Add new members
		resp, err = object.UpdateRoleMembers(expandRoleMembers(new), "Add")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Failed to add members to role: %v", err)
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of role membership completed: %s", object.Name)
	return resourceRoleMembershipRead(d, m)
}

func resourceRoleMembershipDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of role membership: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewRoleMembership(client)
	object.ID = d.Id()
	createUpateGetRoleMembershipData(d, object)
	// Handle role members
	if len(object.Members) > 0 {
		resp, err := object.UpdateRoleMembers(object.Members, "Delete")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Failed to remove members from role: %v", err)
		}
	}

	d.SetId("")
	logger.Infof("Deletion of role membership completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetRoleMembershipData(d *schema.ResourceData, object *vault.RoleMembership) error {
	object.RoleID = d.Get("role_id").(string)
	object.ID = object.RoleID
	if v, ok := d.GetOk("member"); ok {
		object.Members = expandRoleMembers(v)
	}

	return nil
}
