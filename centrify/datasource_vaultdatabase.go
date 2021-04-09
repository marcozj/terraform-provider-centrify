package centrify

import (
	"fmt"
	"strings"

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
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the Database",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port that used to connect to the Database",
			},
			"checkout_lifetime": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the number of minutes that a checked out password is valid.",
			},
			// Database -> Advanced menu related settings
			"allow_multiple_checkouts": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Allow multiple password checkouts for this database",
			},
			"enable_password_rotation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic password rotation",
			},
			"password_rotate_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Password rotation interval (days)",
			},
			"enable_password_rotation_after_checkin": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable password rotation after checkin",
			},
			"minimum_password_age": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum Password Age (days)",
			},
			"password_profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password complexity profile id",
			},
			"enable_password_history_cleanup": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable periodic password history cleanup",
			},
			"password_historycleanup_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Password history cleanup (days)",
			},
			// Database -> Connectors menu related settings
			"connector_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of Connectors",
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
	if v, ok := d.GetOk("instance_name"); ok {
		object.InstanceName = v.(string)
	}
	if v, ok := d.GetOk("service_name"); ok {
		object.ServiceName = v.(string)
	}

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving database with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		if k == "connector_list" {
			// Convert "value1,value1" to schema.TypeSet
			d.Set("connector_list", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		} else {
			d.Set(k, v)
		}
	}

	return nil
}
