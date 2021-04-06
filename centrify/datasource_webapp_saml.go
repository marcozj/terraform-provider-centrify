package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func dataSourceSamlWebApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSamlWebAppRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Web App",
			},
			"corp_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS Account ID or Jira Cloud Subdomain",
			},
			"app_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloudera Entity ID or JIRA Cloud SP Entity ID",
			},
			"application_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Application ID. Specify the name or 'target' that the mobile application uses to find this application.",
			},
			"idp_metadata_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sp_metadata_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_metadata_xml": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service Provider metadata in XML format",
			},
			"sp_entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SP Entity ID, also known as SP Issuer, or Audience, is a value given by your Service Provider",
			},
		},
	}
}

func dataSourceSamlWebAppRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Finding Saml webapp")
	client := m.(*restapi.RestClient)
	object := vault.NewSamlWebApp(client)
	object.Name = d.Get("name").(string)
	if v, ok := d.GetOk("application_id"); ok {
		object.ServiceName = v.(string)
	}
	if v, ok := d.GetOk("corp_identifier"); ok {
		object.CorpIdentifier = v.(string)
	}
	if v, ok := d.GetOk("app_entity_id"); ok {
		object.AdditionalField1 = v.(string)
	}
	if v, ok := d.GetOk("sp_entity_id"); ok {
		object.Audience = v.(string)
	}

	// We can't use simple Query method because it doesn't return all attributes
	err := object.GetByName()
	if err != nil {
		return fmt.Errorf("error retrieving SAML webapp with name '%s': %s", object.Name, err)
	}
	d.SetId(object.ID)
	if object.Name != "" {
		d.Set("name", object.Name)
	}
	if object.ServiceName != "" {
		d.Set("application_id", object.ServiceName)
	}
	if object.CorpIdentifier != "" {
		d.Set("corp_identifier", object.CorpIdentifier)
	}
	if object.AdditionalField1 != "" {
		d.Set("app_entity_id", object.AdditionalField1)
	}
	if object.Audience != "" {
		d.Set("sp_entity_id", object.Audience)
	}
	if object.SpMetadataUrl != "" {
		d.Set("sp_metadata_url", object.SpMetadataUrl)
	}
	if object.SpMetadataXml != "" {
		d.Set("sp_metadata_xml", object.SpMetadataXml)
	}
	if object.IdpMetadataUrl != "" {
		d.Set("idp_metadata_url", object.IdpMetadataUrl)
	}

	return nil
}
