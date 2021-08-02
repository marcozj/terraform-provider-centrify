package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
)

var logPath string

// Provider returns a schema.Provider for Centrify Platform.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_URL", "VAULT_URL"}, ""),
				Description: "Centrify Platform URL",
			},
			"appid": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_APPID", "VAULT_APPID"}, ""),
				Description: "Application ID",
			},
			"scope": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_SCOPE", "VAULT_SCOPE"}, ""),
				Description: "OAuth2 scope",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_USERNAME", "VAULT_USERNAME"}, ""),
				Description: "Username",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_PASSWORD", "VAULT_PASSWORD"}, ""),
				Description: "Password",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_TOKEN", "VAULT_TOKEN"}, ""),
				Description: "OAuth or DMC token",
			},
			"use_dmc": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_USEDMC", "VAULT_USEDMC"}, false),
				Description: "Whether to use DMC",
			},
			"logpath": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_LOGPATH", "VAULT_LOGPATH"}, ""),
				Description: "Path of log file",
			},
			"skip_cert_verify": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_SKIPCERTVERIFY", "VAULT_SKIPCERTVERIFY"}, false),
				Description: "Whether to skip certification verification",
			},
			"log_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Logging level",
				Default:     "error",
				ValidateFunc: validation.StringInSlice([]string{
					"fatal",
					"error",
					"info",
					"debug",
				}, false),
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"CENTRIFY_LOGLEVEL", "VAULT_LOGLEVEL"}, "Error"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"centrifyvault_user":                  dataSourceUser_deprecated(),
			"centrifyvault_role":                  dataSourceRole_deprecated(),
			"centrifyvault_policy":                dataSourcePolicy_deprecated(),
			"centrifyvault_manualset":             dataSourceManualSet_deprecated(),
			"centrifyvault_passwordprofile":       dataSourcePasswordProfile_deprecated(),
			"centrifyvault_authenticationprofile": dataSourceAuthenticationProfile_deprecated(),
			"centrifyvault_connector":             dataSourceConnector_deprecated(),
			"centrifyvault_vaultdomain":           dataSourceDomain_deprecated(),
			"centrifyvault_vaultsystem":           dataSourceSystem_deprecated(),
			"centrifyvault_vaultdatabase":         dataSourceDatabase_deprecated(),
			"centrifyvault_vaultaccount":          dataSourceAccount_deprecated(),
			"centrifyvault_vaultsecret":           dataSourceSecret_deprecated(),
			"centrifyvault_vaultsecretfolder":     dataSourceSecretFolder_deprecated(),
			"centrifyvault_sshkey":                dataSourceSSHKey_deprecated(),
			"centrifyvault_desktopapp":            dataSourceDesktopApp_deprecated(),
			"centrifyvault_directoryservice":      dataSourceDirectoryService_deprecated(),
			"centrifyvault_directoryobject":       dataSourceDirectoryObject_deprecated(),
			"centrifyvault_multiplexedaccount":    dataSourceMultiplexedAccount_deprecated(),
			"centrifyvault_service":               dataSourceService_deprecated(),
			"centrifyvault_cloudprovider":         dataSourceCloudProvider_deprecated(),
			"centrifyvault_webapp_saml":           dataSourceSamlWebApp_deprecated(),
			"centrifyvault_webapp_oauth":          dataSourceOauthWebApp_deprecated(),
			"centrifyvault_webapp_oidc":           dataSourceOidcWebApp_deprecated(),
			"centrifyvault_webapp_generic":        dataSourceGenericWebApp_deprecated(),
			// Change centrifyvault_* centrify_*
			"centrify_user":                  dataSourceUser(),
			"centrify_role":                  dataSourceRole(),
			"centrify_policy":                dataSourcePolicy(),
			"centrify_manualset":             dataSourceManualSet(),
			"centrify_passwordprofile":       dataSourcePasswordProfile(),
			"centrify_authenticationprofile": dataSourceAuthenticationProfile(),
			"centrify_connector":             dataSourceConnector(),
			"centrify_domain":                dataSourceDomain(),
			"centrify_system":                dataSourceSystem(),
			"centrify_database":              dataSourceDatabase(),
			"centrify_account":               dataSourceAccount(),
			"centrify_secret":                dataSourceSecret(),
			"centrify_secretfolder":          dataSourceSecretFolder(),
			"centrify_sshkey":                dataSourceSSHKey(),
			"centrify_desktopapp":            dataSourceDesktopApp(),
			"centrify_directoryservice":      dataSourceDirectoryService(),
			"centrify_directoryobject":       dataSourceDirectoryObject(),
			"centrify_multiplexedaccount":    dataSourceMultiplexedAccount(),
			"centrify_service":               dataSourceService(),
			"centrify_cloudprovider":         dataSourceCloudProvider(),
			"centrify_webapp_saml":           dataSourceSamlWebApp(),
			"centrify_webapp_oauth":          dataSourceOauthWebApp(),
			"centrify_webapp_oidc":           dataSourceOidcWebApp(),
			"centrify_webapp_generic":        dataSourceGenericWebApp(),
			"centrify_federatedgroup":        dataSourceFederatedGroup(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"centrifyvault_user":                      resourceUser_deprecated(),
			"centrifyvault_role":                      resourceRole_deprecated(),
			"centrifyvault_policyorder":               resourcePolicyLinks_deprecated(),
			"centrifyvault_policy":                    resourcePolicy_deprecated(),
			"centrifyvault_manualset":                 resourceManualSet_deprecated(),
			"centrifyvault_passwordprofile":           resourcePasswordProfile_deprecated(),
			"centrifyvault_authenticationprofile":     resourceAuthenticationProfile_deprecated(),
			"centrifyvault_vaultdomain":               resourceDomain_deprecated(),
			"centrifyvault_vaultdomainreconciliation": resourceDomainConfiguration_deprecated(), // To be removed
			"centrifyvault_vaultdomainconfiguration":  resourceDomainConfiguration_deprecated(),
			"centrifyvault_vaultsystem":               resourceSystem_deprecated(),
			"centrifyvault_vaultdatabase":             resourceDatabase_deprecated(),
			"centrifyvault_vaultaccount":              resourceAccount_deprecated(),
			"centrifyvault_vaultsecret":               resourceSecret_deprecated(),
			"centrifyvault_vaultsecretfolder":         resourceSecretFolder_deprecated(),
			"centrifyvault_sshkey":                    resourceSSHKey_deprecated(),
			"centrifyvault_desktopapp":                resourceDesktopApp_deprecated(),
			"centrifyvault_multiplexedaccount":        resourceMultiplexedAccount_deprecated(),
			"centrifyvault_service":                   resourceService_deprecated(),
			"centrifyvault_cloudprovider":             resourceCloudProvider_deprecated(),
			"centrifyvault_globalgroupmappings":       resourceGlobalGroupMappings_deprecated(),
			"centrifyvault_globalworkflow":            resourceGlobalWorkflow_deprecated(),
			"centrifyvault_webapp_saml":               resourceSamlWebApp_deprecated(),
			"centrifyvault_webapp_oauth":              resourceOauthWebApp_deprecated(),
			"centrifyvault_webapp_oidc":               resourceOidcWebApp_deprecated(),
			"centrifyvault_webapp_generic":            resourceGenericWebApp_deprecated(),
			// Change centrifyvault_* centrify_*
			"centrify_user":                  resourceUser(),
			"centrify_role":                  resourceRole(),
			"centrify_policyorder":           resourcePolicyLinks(),
			"centrify_policy":                resourcePolicy(),
			"centrify_manualset":             resourceManualSet(),
			"centrify_passwordprofile":       resourcePasswordProfile(),
			"centrify_authenticationprofile": resourceAuthenticationProfile(),
			"centrify_domain":                resourceDomain(),
			"centrify_domainreconciliation":  resourceDomainConfiguration(), // To be removed
			"centrify_domainconfiguration":   resourceDomainConfiguration(),
			"centrify_system":                resourceSystem(),
			"centrify_database":              resourceDatabase(),
			"centrify_account":               resourceAccount(),
			"centrify_secret":                resourceSecret(),
			"centrify_secretfolder":          resourceSecretFolder(),
			"centrify_sshkey":                resourceSSHKey(),
			"centrify_desktopapp":            resourceDesktopApp(),
			"centrify_multiplexedaccount":    resourceMultiplexedAccount(),
			"centrify_service":               resourceService(),
			"centrify_cloudprovider":         resourceCloudProvider(),
			"centrify_globalgroupmappings":   resourceGlobalGroupMappings(),
			"centrify_globalworkflow":        resourceGlobalWorkflow(),
			"centrify_webapp_saml":           resourceSamlWebApp(),
			"centrify_webapp_oauth":          resourceOauthWebApp(),
			"centrify_webapp_oidc":           resourceOidcWebApp(),
			"centrify_webapp_generic":        resourceGenericWebApp(),
			"centrify_federatedgroup":        resourceFederatedGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:            d.Get("url").(string),
		AppID:          d.Get("appid").(string),
		Scope:          d.Get("scope").(string),
		Username:       d.Get("username").(string),
		Password:       d.Get("password").(string),
		Token:          d.Get("token").(string),
		UseDMC:         d.Get("use_dmc").(bool),
		LogPath:        d.Get("logpath").(string),
		SkipCertVerify: d.Get("skip_cert_verify").(bool),
		LogLevel:       d.Get("log_level").(string),
	}
	switch config.LogLevel {
	case "fatal":
		logger.SetLevel(logger.LevelFatal)
	case "error":
		logger.SetLevel(logger.LevelError)
	case "info":
		logger.SetLevel(logger.LevelInfo)
	case "debug":
		logger.SetLevel(logger.LevelDebug)
	}

	logPath = config.LogPath

	if config.LogPath != "" {
		logger.SetLogPath(config.LogPath)
		logger.EnableErrorStackTrace()
	}

	logger.Infof("Starting provider configuration...")
	if err := config.Valid(); err != nil {
		return nil, err
	}

	restClient, err := config.getClient()

	if err != nil {
		return nil, fmt.Errorf("failed to authenticate to Centrify Platform: %v", err)
	}
	logger.Infof("Connected to Centrify Platform %s", config.URL)

	return restClient, nil
}
