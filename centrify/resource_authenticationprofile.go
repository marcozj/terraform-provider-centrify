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

func resourceAuthenticationProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationProfileCreate,
		Read:   resourceAuthenticationProfileRead,
		Update: resourceAuthenticationProfileUpdate,
		Delete: resourceAuthenticationProfileDelete,
		Exists: resourceAuthenticationProfileExists,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UUID of the authenticaiton profile",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the authenticaiton profile",
			},
			"challenges": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 2,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Authentication mechanisms for challenges",
			},
			"additional_data": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_questions": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "Number of questions user must answer",
							ValidateFunc: validation.IntBetween(0, 10),
						},
					},
				},
			},
			"pass_through_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "Challenge Pass-Through Duration",
			},
		},
	}
}

func resourceAuthenticationProfileExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking authentication profile exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewAuthenticationProfile(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Authentication profile exists in tenant: %s", object.ID)
	return true, nil
}

func resourceAuthenticationProfileRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading authentication profile: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a authentication profile object and populate ID attribute
	object := vault.NewAuthenticationProfile(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading authentication profile: %v", err)
	}
	//logger.Debugf("Authentication profile from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourceAuthenticationProfileRead(): %+v", schemamap)
	for k, v := range schemamap {
		if k == "additional_data" {
			d.Set(k, flattenAdditionalData(object.AdditionalData))
		} else {
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading authentication profile: %s", object.Name)
	return nil
}

func resourceAuthenticationProfileDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of authentication profile: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewAuthenticationProfile(client)
	object.ID = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting authentication profile: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of authentication profile completed: %s", ResourceIDString(d))
	return nil
}

func resourceAuthenticationProfileCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning authentication profile creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)

	// Create a authentication profile object and populate all attributes
	object := vault.NewAuthenticationProfile(client)
	createUpateGetAutheProfileData(d, object)

	resp, err := object.Create()
	if err != nil {
		return fmt.Errorf("Error creating authentication profile: %v", err)
	}

	id := resp.Result["Uuid"].(string)
	if id == "" {
		return fmt.Errorf("Authentication profile ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Creation completed
	logger.Infof("Creation of authentication profile completed: %s", object.Name)
	return resourceAuthenticationProfileRead(d, m)
}

func resourceAuthenticationProfileUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning authentication profile update: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object := vault.NewAuthenticationProfile(client)

	object.ID = d.Id()
	object.UUID = d.Id()
	createUpateGetAutheProfileData(d, object)

	// Deal with normal attribute changes first
	if d.HasChanges("name", "challenges", "pass_through_duration", "additional_data") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating authentication profile attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
	}

	logger.Infof("Updating of authentication profile completed: %s", object.Name)
	return resourceAuthenticationProfileRead(d, m)
}

func createUpateGetAutheProfileData(d *schema.ResourceData, object *vault.AuthenticationProfile) error {
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("pass_through_duration"); ok {
		object.DurationInMinutes = v.(int)
	}
	if v, ok := d.GetOk("additional_data"); ok {
		object.AdditionalData = expandAdditionalData(v)
	}
	if v, ok := d.GetOk("challenges"); ok {
		object.Challenges = flattenTypeListToSlice(v.([]interface{}))
	}

	return nil
}

func expandAdditionalData(v interface{}) *vault.AdditionalData {
	var adData *vault.AdditionalData
	d := v.([]interface{})[0].(map[string]interface{})
	adData = &vault.AdditionalData{
		NumberOfQuestions: d["number_of_questions"].(int),
	}

	return adData
}

func flattenAdditionalData(v *vault.AdditionalData) []interface{} {
	adData := map[string]interface{}{
		"number_of_questions": v.NumberOfQuestions,
	}

	return []interface{}{adData}
}
