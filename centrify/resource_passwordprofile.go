package centrify

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func resourcePasswordProfile() *schema.Resource {
	return &schema.Resource{

		Create: resourcePasswordProfileCreate,
		Read:   resourcePasswordProfileRead,
		Update: resourcePasswordProfileUpdate,
		Delete: resourcePasswordProfileDelete,
		Exists: resourcePasswordProfileExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the password profile",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of password profile",
			},
			"minimum_password_length": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Minimum password length",
				ValidateFunc: validation.IntBetween(4, 128),
			},
			"maximum_password_length": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Maximum password length",
				ValidateFunc: validation.IntBetween(8, 128),
			},
			"at_least_one_lowercase": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "At least one lower-case alpha character",
			},
			"at_least_one_uppercase": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "At least one upper-case alpha character",
			},
			"at_least_one_digit": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "At least one digit",
			},
			"no_consecutive_repeated_char": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "No consecutive repeated characters",
			},
			"at_least_one_special_char": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "At least one special character",
			},
			"maximum_char_occurrence_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Maximum character occurrence count",
				ValidateFunc: validation.IntBetween(1, 128),
			},
			"special_charset": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Special Characters",
			},
			"first_character_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AnyChar",
				Description: "A leading alpha or alphanumeric character",
				ValidateFunc: validation.StringInSlice([]string{
					"AnyChar",
					"AlphaOnly",
					"AlphaNumericOnly",
				}, false),
			},
			"last_character_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "AnyChar",
				Description: "A trailing alpha or alphanumeric character",
				ValidateFunc: validation.StringInSlice([]string{
					"AnyChar",
					"AlphaOnly",
					"AlphaNumericOnly",
				}, false),
			},
			"minimum_alphabetic_character_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Min number of alpha characters",
				ValidateFunc: validation.IntBetween(1, 128),
			},
			"minimum_non_alphabetic_character_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "Min number of non-alpha characters",
				ValidateFunc: validation.IntBetween(1, 128),
			},
		},
	}
}

func resourcePasswordProfileExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking password profile exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewPasswordProfile(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Authentication password exists in tenant: %s", object.ID)
	return true, nil
}

func resourcePasswordProfileRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading password profile: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a password profile object and populate ID attribute
	object := vault.NewPasswordProfile(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading password profile: %v", err)
	}
	//logger.Debugf("password profile from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourcePasswordProfileRead(): %+v", schemamap)
	for k, v := range schemamap {
		d.Set(k, v)
	}

	logger.Infof("Completed reading password profile: %s", object.Name)
	return nil
}

func resourcePasswordProfileDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of password profile: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewPasswordProfile(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting password profile: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of password profile completed: %s", ResourceIDString(d))
	return nil
}

func resourcePasswordProfileCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning password profile creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)

	// Create a password profile object and populate all attributes
	object := vault.NewPasswordProfile(client)
	createUpateGetPasswordProfileData(d, object)

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating password profile: %v", err)
	}

	id := resp.Result
	if id == "" {
		return fmt.Errorf("Password profile ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Creation completed
	logger.Infof("Creation of password profile completed: %s", object.Name)
	return resourcePasswordProfileRead(d, m)
}

func resourcePasswordProfileUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning password profile update: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object := vault.NewPasswordProfile(client)

	object.ID = d.Id()
	createUpateGetPasswordProfileData(d, object)

	// Deal with normal attribute changes first
	if d.HasChanges("name", "description", "minimum_password_length", "maximum_password_length", "at_least_one_lowercase", "at_least_one_uppercase",
		"at_least_one_digit", "no_consecutive_repeated_char", "at_least_one_special_char", "maximum_char_occurrence_count", "special_charset",
		"first_character_type", "last_character_type", "minimum_alphabetic_character_count", "minimum_non_alphabetic_character_count") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating password profile attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
	}

	logger.Infof("Updating of password profile completed: %s", object.Name)
	return resourcePasswordProfileRead(d, m)
}

func createUpateGetPasswordProfileData(d *schema.ResourceData, object *vault.PasswordProfile) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("minimum_password_length"); ok {
		object.MinimumPasswordLength = v.(int)
	}
	if v, ok := d.GetOk("maximum_password_length"); ok {
		object.MaximumPasswordLength = v.(int)
	}
	if v, ok := d.GetOk("at_least_one_lowercase"); ok {
		object.AtLeastOneLowercase = v.(bool)
	}
	if v, ok := d.GetOk("at_least_one_uppercase"); ok {
		object.AtLeastOneUppercase = v.(bool)
	}
	if v, ok := d.GetOk("at_least_one_digit"); ok {
		object.AtLeastOneDigit = v.(bool)
	}
	if v, ok := d.GetOk("no_consecutive_repeated_char"); ok {
		object.ConsecutiveCharRepeatAllowed = v.(bool)
	}
	if v, ok := d.GetOk("at_least_one_special_char"); ok {
		object.AtLeastOneSpecial = v.(bool)
	}
	if v, ok := d.GetOk("maximum_char_occurrence_count"); ok {
		object.MaximumCharOccurrenceCount = v.(int)
	}
	if v, ok := d.GetOk("special_charset"); ok {
		object.SpecialCharSet = v.(string)
	}
	if v, ok := d.GetOk("first_character_type"); ok {
		object.FirstCharacterType = v.(string)
	}
	if v, ok := d.GetOk("last_character_type"); ok {
		object.LastCharacterType = v.(string)
	}
	if v, ok := d.GetOk("minimum_alphabetic_character_count"); ok {
		object.MinimumAlphabeticCharacterCount = v.(int)
	}
	if v, ok := d.GetOk("minimum_non_alphabetic_character_count"); ok {
		object.MinimumNonAlphabeticCharacterCount = v.(int)
	}

	return nil
}
