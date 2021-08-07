package centrify

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceRole_deprecated() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Exists: resourceRoleExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema:             getRoleSchema(),
		DeprecationMessage: "resource centrifyvault_role is deprecated will be removed in the future, use centrify_role instead",
	}
}

func resourceRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Update: resourceRoleUpdate,
		Delete: resourceRoleDelete,
		Exists: resourceRoleExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: getRoleSchema(),
	}
}

func getRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the role",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of an role",
		},
		"adminrights": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
func resourceRoleExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking role exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewRole(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Role exists in tenant: %s", object.ID)
	return true, nil
}

func resourceRoleRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading role: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a role object and populate ID attribute
	object := vault.NewRole(client)
	object.ID = d.Id()
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
	logger.Debugf("Generated Map for resourceRoleRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading role: %s", object.Name)
	return nil
}

func resourceRoleCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning Role creation: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)

	// Create a role object and populate all attributes
	object := vault.NewRole(client)
	createUpateGetRoleData(d, object)

	// Response contains only role id
	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf(" Error creating role: %v", err)
	}

	id := resp.Result["_RowKey"].(string)
	if id == "" {
		return fmt.Errorf(" The Role ID is not set")
	}
	d.SetId(id)

	// Need to populate ID attribute otherwise AssignAdminRights function will fail
	object.ID = id

	logger.Debugf("Role created: %s", object.Name)

	// Handle role members
	if len(object.Members) > 0 {
		resp, err := object.UpdateRoleMembers(object.Members, "Add")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error adding members to role: %v", err)
		}
	}

	// Assign admin rights
	if object.AdminRights != nil {
		resp, err := object.AssignAdminRights()
		if err != nil || !resp.Success {
			log.Fatalf("Error updating role admin rights: %v", err)
			return nil
		}
		logger.Debugf("Updated admin rights to: %v", object.AdminRights)
	}

	// Creation completed
	d.Partial(false)
	logger.Infof("Creation of role completed: %s", object.Name)
	return resourceRoleRead(d, m)
}

func resourceRoleUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning role update: %s", ResourceIDString(d))

	// Enable partial state mode
	d.Partial(true)

	client := m.(*restapi.RestClient)
	object := vault.NewRole(client)
	object.ID = d.Id()
	createUpateGetRoleData(d, object)

	if d.HasChanges("name", "description") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error updating role attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
	}

	// Deal with role members
	if d.HasChange("member") {
		old, new := d.GetChange("member")
		// Remove old members
		resp, err := object.UpdateRoleMembers(expandRoleMembers(old), "Delete")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error removing members from role: %v", err)
		}
		// Add new members
		resp, err = object.UpdateRoleMembers(expandRoleMembers(new), "Add")
		if err != nil || !resp.Success {
			return fmt.Errorf(" Error adding members to role: %v", err)
		}
	}

	// Deal with admin rights change
	if d.HasChange("adminrights") {
		// To update admin rights, we need to remove all existing ones first
		rights, err := object.GetAdminRights()
		if err != nil {
			return fmt.Errorf(" Error getting existing role admin rights: %v", err)
		}
		logger.Debugf("Removing existing admin rights: %v", rights)
		if rights != nil && len(rights) > 0 {
			resp, err := object.RemoveAdminRights(rights)
			if err != nil || !resp.Success {
				return fmt.Errorf(" Error removing existing role admin rights: %v", err)
			}
		}
		logger.Debugf("Removed existing admin rights: %v", rights)

		// Set new admin rights
		if d.Get("adminrights") != nil && d.Get("adminrights").(*schema.Set).Len() > 0 {
			tfAdminRights := d.Get("adminrights").(*schema.Set).List()
			adminrights := make([]string, len(tfAdminRights))
			for i, tfAdminRight := range tfAdminRights {
				adminrights[i] = tfAdminRight.(string)
			}
			object.AdminRights = adminrights
			logger.Debugf("Adding admin rights: %v", adminrights)

			resp, err := object.AssignAdminRights()
			if err != nil || !resp.Success {
				return fmt.Errorf(" Error updating role admin rights: %v", err)
			}
			logger.Debugf("Updated admin rights to: %v", adminrights)
		}
	}

	// We succeeded, disable partial mode. This causes Terraform to save all fields again.
	d.Partial(false)
	logger.Infof("Updating of role completed: %s", object.Name)
	return resourceRoleRead(d, m)
}

func resourceRoleDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of role: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewRole(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf(" Error deleting role: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of role completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGetRoleData(d *schema.ResourceData, object *vault.Role) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok && d.HasChange("description") {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("adminrights"); ok && d.HasChange("adminrights") {
		object.AdminRights = flattenSchemaSetToStringSlice(v)
	}
	if v, ok := d.GetOk("member"); ok && d.HasChange("member") {
		object.Members = expandRoleMembers(v)
	}

	return nil
}
