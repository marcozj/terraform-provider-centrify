package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/computerclass"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultSystem() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultSystemRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the system",
			},
			"fqdn": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Hostname or IP address of the system",
				ValidateFunc: validation.NoZeroValues,
			},
			"computer_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the system",
				ValidateFunc: validation.StringInSlice([]string{
					computerclass.Windows.String(),
					computerclass.Unix.String(),
					computerclass.CiscoAsyncOS.String(),
					computerclass.CiscoIOS.String(),
					computerclass.CiscoNXOS.String(),
					computerclass.JuniperJunos.String(),
					computerclass.HPNonStop.String(),
					computerclass.IBMi.String(),
					computerclass.CheckPointGaia.String(),
					computerclass.PaloAltoPANOS.String(),
					computerclass.F5BIGIP.String(),
					computerclass.VMwareVMkernel.String(),
					computerclass.GenericSSH.String(),
					computerclass.CustomSSH.String(),
				}, false),
			},
		},
	}
}

func dataSourceVaultSystemRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding system")
	client := m.(*restapi.RestClient)
	object := vault.NewSystem(client)
	object.Name = d.Get("name").(string)
	object.FQDN = d.Get("fqdn").(string)
	if v, ok := d.GetOk("computer_class"); ok {
		object.ComputerClass = v.(string)
	}

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found system: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("fqdn", result["FQDN"].(string))
	d.Set("computer_class", result["ComputerClass"].(string))

	return nil
}
