package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultSecretRead,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the secret",
			},
			"parent_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path of parent folder",
			},
			"checkout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve secret content",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the secret",
			},
			"folder_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the folder where the secret is located",
			},
			"secret_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Content of the secret",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Either Text or File",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default Secret Challenge Profile (used if no conditions matched)",
			},
			// Workflow
			"workflow_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"workflow_approver": getWorkflowApproversSchema(),
			"challenge_rule":    getChallengeRulesSchema(),
		},
	}
}

func dataSourceVaultSecretRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding vault secret")
	client := m.(*restapi.RestClient)
	object := vault.NewSecret(client)
	object.SecretName = d.Get("secret_name").(string)
	if v, ok := d.GetOk("parent_path"); ok {
		object.ParentPath = v.(string)
	}
	if v, ok := d.GetOk("folder_id"); ok {
		object.FolderID = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving secret with name '%s': %s", object.SecretName, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
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

	if d.Get("checkout").(bool) {
		text, err := object.CheckoutSecret()
		if err != nil {
			return fmt.Errorf("error checking out secret content with name '%s': %s", object.SecretName, err)
		}
		d.Set("secret_text", text)
	}

	return nil
}
