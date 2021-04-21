package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
			"system_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The target system id where the service runs",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the service",
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the service to be managed.",
			},
			"enable_management": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable management of this service password",
			},
			"admin_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Administrative account id that used to manage the password for the service",
			},
			"multiplexed_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The multiplexed account id to run the service",
			},
			"restart_service": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Restart Service when password is rotated",
			},
			"restart_time_restriction": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enforce restart time restrictions",
			},
			"days_of_week": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Day of the week restart allowed",
			},
			"restart_start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Start time of the time range restart is allowed",
			},
			"restart_end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "End time of the time range restart is allowed",
			},
			"use_utc_time": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to use UTC time",
			},
		},
	}
}

func dataSourceServiceRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Service")
	client := m.(*restapi.RestClient)
	object := vault.NewService(client)
	object.Name = d.Get("service_name").(string)

	// We can't use simple Query method because it doesn't return all attributes
	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving Service with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		if k == "days_of_week" {
			// Convert "value1,value1" to schema.TypeSet
			d.Set("days_of_week", schema.NewSet(schema.HashString, StringSliceToInterface(strings.Split(v.(string), ","))))
		} else {
			d.Set(k, v)
		}
	}

	return nil
}
