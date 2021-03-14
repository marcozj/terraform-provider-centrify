package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/marcozj/golang-sdk/enum/databaseclass"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceVaultDatabase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVaultDatabaseRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Database",
			},
			"hostname": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Hostname or IP address of the Database",
				ValidateFunc: validation.NoZeroValues,
			},
			"database_class": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the Database",
				ValidateFunc: validation.StringInSlice([]string{
					databaseclass.SQLServer.String(),
					databaseclass.Oracle.String(),
					databaseclass.SAPASE.String(),
				}, false),
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Port that used to connect to the Database",
				ValidateFunc: validation.IsPortNumber,
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name of MS SQL Database",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service name of Oracle database",
			},
		},
	}
}

func dataSourceVaultDatabaseRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding database")
	client := m.(*restapi.RestClient)
	object := vault.NewDatabase(client)
	object.Name = d.Get("name").(string)
	object.FQDN = d.Get("hostname").(string)
	if v, ok := d.GetOk("database_class"); ok {
		object.DatabaseClass = v.(string)
	}

	result, err := object.Query()
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	d.SetId(result["ID"].(string))
	d.Set("name", result["Name"].(string))
	d.Set("hostname", result["FQDN"].(string))
	d.Set("database_class", result["DatabaseClass"].(string))
	d.Set("port", int(result["Port"].(float64)))
	if result["InstanceName"] != nil {
		d.Set("instance_name", result["InstanceName"].(string))
	}
	if result["ServiceName"] != nil {
		d.Set("service_name", result["ServiceName"].(string))
	}

	return nil
}
