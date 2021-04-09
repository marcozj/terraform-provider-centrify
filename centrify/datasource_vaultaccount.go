package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/keypairtype"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultAccountRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the account",
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
			"access_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS access key id",
			},
			"secret_access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "AWS secret access key",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password of the account",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "SSH private key",
			},
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Passphrase to use for encrypting the PrivateKey",
			},
			"key_pair_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      keypairtype.PrivateKey.String(),
				ValidateFunc: validation.StringInSlice([]string{keypairtype.PublicKey.String(), keypairtype.PrivateKey.String(), keypairtype.PuTTY.String()}, false),
			},
			"checkout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to checkout the password",
			},
			"checkin": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to checkin the password immediately after checkout",
			},
			"credential_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Either password or sshkey",
			},
			"credential_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sshkey_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of SSH key",
			},
			"is_admin_account": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is an administrative account",
			},
			"is_root_account": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is an root account for cloud provider",
			},
			"use_proxy_account": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use proxy account to manage this account",
			},
			"managed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If this account is managed",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the account",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Policy menu
			"checkout_lifetime": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Checkout lifetime (minutes)",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default password checkout profile id",
			},
			"access_secret_checkout_default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default secret access key checkout challenge rule id",
			},
			"access_secret_checkout_rule": getChallengeRulesSchema(),
			// Workflow
			"workflow_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"workflow_approver": getWorkflowApproversSchema(),
			"challenge_rule":    getChallengeRulesSchema(),
			"access_key":        getAccessKeySchema(),
		},
	}
}

func dataSourceVaultAccountRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding vault account")
	client := m.(*restapi.RestClient)
	object := vault.NewAccount(client)
	object.User = d.Get("name").(string)
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

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving vault account with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		switch k {
		case "challenge_rule", "access_secret_checkout_rule":
			d.Set(k, v.(map[string]interface{})["rule"])
		case "workflow_approvers":
			if object.WorkflowEnabled && v.(string) != "" {
				// convertWorkflowSchema expects "workflow_approvers" in format of {"WorkflowApprover":[{"Type":"Manager","NoManagerAction":"useBackup","BackupApprover":{"Guid":"xxxxxx_xxxx_xxxx_xxxxxxxxx","Name":"Infrastructure Owners","Type":"Role"}}]}
				// which matches ProxyWorkflowApprover struct
				wfschema, err := convertWorkflowSchema("{\"WorkflowApprover\":" + v.(string) + "}")
				if err != nil {
					return err
				}
				d.Set("workflow_approver", wfschema)
				d.Set(k, v)
			}
		default:
			d.Set(k, v)
		}
	}

	// Checkout credential
	if d.Get("checkout").(bool) {
		switch object.CredentialType {
		case "Password":
			pw, err := object.CheckoutPassword(d.Get("checkin").(bool))
			if err != nil {
				return err
			}
			d.Set("password", pw)
		case "SshKey":
			//object.CredentialID = result["CredentialId"].(string)
			sshkey := vault.NewSSHKey(client)
			sshkey.ID = object.CredentialID
			sshkey.KeyPairType = d.Get("key_pair_type").(string)
			sshkey.Passphrase = d.Get("passphrase").(string)
			sshkey.KeyFormat = "PEM"
			thekey, err := sshkey.RetriveSSHKey()
			if err != nil {
				return err
			}
			d.Set("private_key", thekey)
		case "AwsAccessKey":
			secretkey, err := object.RetrieveAccessKey(d.Get("access_key_id").(string))
			if err != nil {
				return err
			}
			d.Set("secret_access_key", secretkey)
		}
	}

	return nil
}
