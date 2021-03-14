package centrify

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
)

var logPath string

// Provider returns a schema.Provider for Centrify Vault.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_URL", ""),
				Description: "Centrify Vault URL",
			},
			"appid": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_APPID", ""),
				Description: "Application ID",
			},
			"scope": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_SCOPE", ""),
				Description: "OAuth2 scope",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_USERNAME", ""),
				Description: "Username",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_PASSWORD", ""),
				Description: "Password",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_TOKEN", ""),
				Description: "OAuth or DMC token",
			},
			"use_dmc": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_USEDMC", false),
				Description: "Whether to use DMC",
			},
			"logpath": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_LOGPATH", ""),
				Description: "Path of log file",
			},
			"skip_cert_verify": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VAULT_SKIPCERTVERIFY", false),
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
				DefaultFunc: schema.EnvDefaultFunc("VAULT_LOGLEVEL", "Error"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"centrifyvault_user":                  dataSourceUser(),
			"centrifyvault_role":                  dataSourceRole(),
			"centrifyvault_policy":                dataSourcePolicy(),
			"centrifyvault_manualset":             dataSourceManualSet(),
			"centrifyvault_passwordprofile":       dataSourcePasswordProfile(),
			"centrifyvault_authenticationprofile": dataSourceAuthenticationProfile(),
			"centrifyvault_connector":             dataSourceConnector(),
			"centrifyvault_vaultdomain":           dataSourceVaultDomain(),
			"centrifyvault_vaultsystem":           dataSourceVaultSystem(),
			"centrifyvault_vaultdatabase":         dataSourceVaultDatabase(),
			"centrifyvault_vaultaccount":          dataSourceVaultAccount(),
			"centrifyvault_vaultsecret":           dataSourceVaultSecret(),
			"centrifyvault_vaultsecretfolder":     dataSourceVaultSecretFolder(),
			"centrifyvault_sshkey":                dataSourceSSHKey(),
			"centrifyvault_directoryservice":      dataSourceDirectoryService(),
			"centrifyvault_directoryobject":       dataSourceDirectoryObject(),
			"centrifyvault_multiplexedaccount":    dataSourceMultiplexedAccount(),
			"centrifyvault_cloudprovider":         dataSourceCloudProvider(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"centrifyvault_user":                      resourceUser(),
			"centrifyvault_role":                      resourceRole(),
			"centrifyvault_policyorder":               resourcePolicyLinks(),
			"centrifyvault_policy":                    resourcePolicy(),
			"centrifyvault_manualset":                 resourceManualSet(),
			"centrifyvault_passwordprofile":           resourcePasswordProfile(),
			"centrifyvault_authenticationprofile":     resourceAuthenticationProfile(),
			"centrifyvault_vaultdomain":               resourceVaultDomain(),
			"centrifyvault_vaultdomainreconciliation": resourceVaultDomainReconciliation(),
			"centrifyvault_vaultsystem":               resourceVaultSystem(),
			"centrifyvault_vaultdatabase":             resourceVaultDatabase(),
			"centrifyvault_vaultaccount":              resourceVaultAccount(),
			"centrifyvault_vaultsecret":               resourceVaultSecret(),
			"centrifyvault_vaultsecretfolder":         resourceVaultSecretFolder(),
			"centrifyvault_sshkey":                    resourceSSHKey(),
			"centrifyvault_desktopapp":                resourceDesktopApp(),
			"centrifyvault_multiplexedaccount":        resourceMultiplexedAccount(),
			"centrifyvault_service":                   resourceService(),
			"centrifyvault_cloudprovider":             resourceCloudProvider(),
			"centrifyvault_globalgroupmappings":       resourceGlobalGroupMappings(),
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
		return nil, fmt.Errorf("Failed to authenticate to Centrify Vault: %v", err)
	}
	logger.Infof("Connected to Centrify Vault %s", config.URL)

	return restClient, nil
}
