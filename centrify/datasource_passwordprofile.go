package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourcePasswordProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePasswordProfileRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of password profile",
			},
			"profile_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of password profile",
			},
			// computed attributes
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of password profile",
			},
			"minimum_password_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Minimum password length",
			},
			"maximum_password_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum password length",
			},
			"at_least_one_lowercase": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "At least one lower-case alpha character",
			},
			"at_least_one_uppercase": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "At least one upper-case alpha character",
			},
			"at_least_one_digit": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "At least one digit",
			},
			"no_consecutive_repeated_char": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "No consecutive repeated characters",
			},
			"at_least_one_special_char": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "At least one special character",
			},
			"maximum_char_occurrence_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum character occurrence count",
			},
			"special_charset": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Special Characters",
			},
			"first_character_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A leading alpha or alphanumeric character",
			},
			"last_character_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A trailing alpha or alphanumeric character",
			},
			"minimum_alphabetic_character_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Min number of alpha characters",
			},
			"minimum_non_alphabetic_character_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Min number of non-alpha characters",
			},
		},
	}
}

func dataSourcePasswordProfileRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding password profile")
	client := m.(*restapi.RestClient)
	object := vault.NewPasswordProfile(client)
	object.Name = d.Get("name").(string)
	object.ProfileType = d.Get("profile_type").(string)

	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving password profile with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	//logger.Debugf("Generated Map: %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	return nil
}
