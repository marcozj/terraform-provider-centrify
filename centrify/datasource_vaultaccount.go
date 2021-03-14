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

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found account: %+v", result)
	object.ID = result["ID"].(string)
	d.SetId(object.ID)
	d.Set("name", result["Name"].(string))
	if result["Host"] != nil {
		d.Set("host_id", result["Host"].(string))
	}
	if result["DomainID"] != nil {
		d.Set("domain_id", result["DomainID"].(string))
	}
	if result["DatabaseID"] != nil {
		d.Set("database_id", result["DatabaseID"].(string))
	}
	if result["CloudProviderId"] != nil {
		d.Set("cloudprovider_id", result["CloudProviderId"].(string))
	}
	credtype := result["CredentialType"].(string)

	// Checkout credential
	if d.Get("checkout").(bool) {
		switch credtype {
		case "Password":
			pw, err := object.CheckoutPassword(d.Get("checkin").(bool))
			if err != nil {
				return err
			}
			d.Set("password", pw)
		case "SshKey":
			object.CredentialID = result["CredentialId"].(string)
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
