package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourcePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the policy",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the policy",
			},
			"link_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Link type of the policy",
				ValidateFunc: validation.StringInSlice([]string{
					"Global",
					"Role",
					"Collection",
					"Inactive",
				}, false),
			},
			"policy_assignment": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of roles or sets assigned to the policy",
			},
		},
	}
}

func dataSourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding policy")
	client := m.(*restapi.RestClient)
	object := vault.NewPolicy(client)
	object.Name = d.Get("name").(string)

	result, err := object.Query("name")
	if err != nil {
		return fmt.Errorf("Error retrieving vault object: %s", err)
	}

	//logger.Debugf("Found user: %+v", result)
	d.SetId(result["ID"].(string))
	d.Set("description", result["Description"].(string))
	d.Set("link_type", result["LinkType"].(string))

	return nil
}
