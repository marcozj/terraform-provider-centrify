package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/workflowtype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourceGlobalWorkflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobalWorkflowCreate,
		Read:   resourceGlobalWorkflowRead,
		Update: resourceGlobalWorkflowUpdate,
		Delete: resourceGlobalWorkflowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					workflowtype.AccountWorkflow.String(),
					workflowtype.AgentAuthWorkflow.String(),
					workflowtype.SecretsWorkflow.String(),
					workflowtype.PrivilegeElevationWorkflow.String(),
				}, false),
			},
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable workflow for all accounts/systems/secrets",
						},
						"default_options": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"approver": getWorkflowApproversSchema(),
					},
				},
			},
		},
	}
}

func resourceGlobalWorkflowRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading global workflow: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object, err := vault.NewGlobalWorkflow(client, d.Get("type").(string))
	if err != nil {
		return err
	}
	object.ID = d.Id()
	err = object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("error reading global workflow %v", err)
	}

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceGlobalWorkflowRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading global workflow: %s", object.Type)
	return nil
}

func resourceGlobalWorkflowCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning global workflow creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object, err := vault.NewGlobalWorkflow(client, d.Get("type").(string))
	if err != nil {
		return err
	}

	createUpateGlobalWorkflowData(d, object)

	_, err = object.Update()
	if err != nil {
		return fmt.Errorf("error creating global workflow: %v", err)
	}

	id := "centrifyvault_global_workflow_" + d.Get("type").(string)
	d.SetId(id)
	object.ID = id

	// Update completed
	logger.Infof("Creation of global workflow completed: %s", d.Id())
	return resourceGlobalWorkflowRead(d, m)
}

func resourceGlobalWorkflowUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning global workflow update: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object, err := vault.NewGlobalWorkflow(client, d.Get("type").(string))
	if err != nil {
		return err
	}
	object.ID = d.Id()

	createUpateGlobalWorkflowData(d, object)

	if d.HasChanges("type", "settings") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("error updating global workflow: %v", err)
		}
	}

	logger.Infof("Updating of global workflow completed: %s", object.Type)
	return resourceGlobalWorkflowRead(d, m)
}

func resourceGlobalWorkflowDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning disabling of global workflow: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object, err := vault.NewGlobalWorkflow(client, d.Get("type").(string))
	if err != nil {
		return err
	}
	object.ID = d.Id()
	err = object.Delete()
	if err != nil {
		return fmt.Errorf("error disabling global workflow: %v", err)
	}

	d.SetId("")

	logger.Infof("Disabling of global workflow completed: %s", ResourceIDString(d))
	return nil
}

func createUpateGlobalWorkflowData(d *schema.ResourceData, object *vault.GlobalWorkflow) {
	object.Type = d.Get("type").(string)
	if v, ok := d.GetOk("settings"); ok {
		settings := v.([]interface{})
		if len(settings) > 0 && settings[0] != nil {
			d := settings[0].(map[string]interface{})
			object.Settings = &vault.GlobalWorkflowSetting{
				Enabled:        d["enabled"].(bool),
				DefaultOptions: d["default_options"].(string),
			}
			if v, ok := d["approver"]; ok {
				object.Settings.ApproverList = expandWorkflowApprovers(v.([]interface{}))
			}
		}
	}
}
