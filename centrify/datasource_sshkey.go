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

func dataSourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the ssh key",
			},
			"key_pair_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Which key to retrieve from the pair, must be either PublicKey, PrivateKey, or PPK",
				ValidateFunc: validation.StringInSlice([]string{
					keypairtype.PublicKey.String(),
					keypairtype.PrivateKey.String(),
					keypairtype.PuTTY.String(),
				}, false),
			},
			"key_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "PEM",
				Description: "KeyFormat to retrieve the key in - only works for PublicKey",
			},
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Passphrase to use for decrypting the PrivateKey",
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"checkout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to retrieve SSH Key",
			},
			"ssh_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Content of the SSH Key",
			},
			"default_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default SSH Key Challenge Profile",
			},
			"challenge_rule": getChallengeRulesSchema(),
		},
	}
}

func dataSourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding SSH Key")
	client := m.(*restapi.RestClient)
	object := vault.NewSSHKey(client)
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("key_pair_type"); ok {
		object.KeyPairType = v.(string)
	}
	if v, ok := d.GetOk("passphrase"); ok {
		object.Passphrase = v.(string)
	}
	if v, ok := d.GetOk("key_format"); ok {
		object.KeyFormat = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving SSH Key with name '%s': %s", object.Name, err)
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
		default:
			d.Set(k, v)
		}
	}

	// Retrieve SSH Key
	if d.Get("checkout").(bool) {
		thekey, err := object.RetriveSSHKey()
		if err != nil {
			return fmt.Errorf("error checking out SSH Key with name '%s': %s", object.Name, err)
		}
		d.Set("ssh_key", thekey)
	}

	return nil
}
