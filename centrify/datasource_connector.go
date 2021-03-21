package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceConnector() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConnectorRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Connector",
			},
			"machine_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rdp_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ad_proxy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_api_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldap_proxy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius_external_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Active", "Inactive"}, false),
			},
			"online": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vpc_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceConnectorRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding connector")
	client := m.(*restapi.RestClient)
	object := vault.NewConnector(client)
	object.Name = d.Get("name").(string)
	object.Status = d.Get("status").(string)
	object.Version = d.Get("version").(string)
	object.VpcIdentifier = d.Get("vpc_identifier").(string)
	object.VmIdentifier = d.Get("vm_identifier").(string)

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found connector: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("machine_name", result["MachineName"].(string))
	d.Set("ssh_service", result["SSHService"].(string))
	d.Set("rdp_service", result["RDPService"].(string))
	d.Set("ad_proxy", result["ADProxy"].(string))
	d.Set("app_gateway", result["AppGateway"].(string))
	d.Set("http_api_service", result["HttpAPIService"].(string))
	d.Set("ldap_proxy", result["LDAPProxy"].(string))
	d.Set("radius_service", result["RadiusService"].(string))
	d.Set("radius_external_service", result["RadiusExternalService"].(string))
	d.Set("version", result["Version"].(string))
	if result["VpcIdentifier"] != nil {
		d.Set("vpc_identifier", result["VpcIdentifier"].(string))
	}
	if result["VmIdentifier"] != nil {
		d.Set("vm_identifier", result["VmIdentifier"].(string))
	}

	return nil
}
