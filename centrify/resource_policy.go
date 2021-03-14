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

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,
		Exists: resourcePolicyExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the policy",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the policy",
			},
			"link_type": {
				Type:        schema.TypeString,
				Required:    true,
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
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"centrify_services":        getCentrifyServicesSchema(),
						"centrify_client":          getCentrifyClientSchema(),
						"centrify_css_server":      getCentrifyCSSServerSchema(),
						"centrify_css_workstation": getCentrifyCSSWorkstationSchema(),
						"centrify_css_elevation":   getCentrifyCSSElevationSchema(),
						"self_service":             getSelfServiceSchema(),
						"password_settings":        getPasswordSettingsSchema(),
						"oath_otp":                 getOATHOTPSchema(),
						"radius":                   getRadiusSchema(),
						"user_account":             getUserAccountSchema(),
						"system_set":               getSystemSetSchema(),
						"database_set":             getDatabaseAndDomainSetSchema(),
						"domain_set":               getDatabaseAndDomainSetSchema(),
						"account_set":              getAccountSetSchema(),
						"secret_set":               getSecretSetSchema(),
						"sshkey_set":               getSSHKeySetSchema(),
						"cloudproviders_set":       getCloudProvidersSchema(),
						"mobile_device":            getMobileDeviceSchema(),
					},
				},
			},
		},
	}
}

func resourcePolicyExists(d *schema.ResourceData, m interface{}) (bool, error) {
	logger.Infof("Checking policy exist: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewPolicy(client)
	object.ID = d.Id()
	err := object.Read()

	if err != nil {
		if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}

	logger.Infof("Policy exists in tenant: %s", object.ID)
	return true, nil
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Reading policy: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	// Create a policy object and populate ID attribute
	object := vault.NewPolicy(client)
	object.ID = d.Id()
	err := object.Read()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading policy: %v", err)
	}
	//logger.Debugf("Policy from tenant: %v", object)

	schemamap, err := vault.GenerateSchemaMap(object)
	if err != nil {
		return err
	}
	logger.Debugf("Generated Map for resourcePolicyRead(): %+v", schemamap)
	for k, v := range schemamap {
		if k == "plink" {
			// Handle plink content. In schema, following attributes are in root level but they are sub map section
			d.Set("link_type", object.Plink.LinkType)
			d.Set("params", object.Plink.Params)
		} else {
			d.Set(k, v)
		}
	}

	logger.Infof("Completed reading policy: %s", object.Name)
	return nil
}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning deletion of policy: %s", ResourceIDString(d))
	client := m.(*restapi.RestClient)

	object := vault.NewPolicy(client)
	object.ID = d.Id()
	//object.Path = d.Id()
	resp, err := object.Delete()

	// If the resource does not exist, inform Terraform. We want to immediately
	// return here to prevent further processing.
	if err != nil {
		return fmt.Errorf("Error deleting policy: %v", err)
	}

	if resp.Success {
		d.SetId("")
	}

	logger.Infof("Deletion of policy completed: %s", ResourceIDString(d))
	return nil
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning policy creation: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)

	// Create a policy object and populate all attributes
	object := vault.NewPolicy(client)
	err := createUpateGetPolicyData(d, object)
	if err != nil {
		return fmt.Errorf("Error constructing policy data: %v", err)
	}

	_, err = object.Create()
	if err != nil {
		return fmt.Errorf("Error creating policy: %v", err)
	}

	// Creation API call doesn't return ID. ID isn't in UUID format but in "/Policy/<name>" format. So, set it manaully instead.
	id := "/Policy/" + object.Name
	if id == "" {
		return fmt.Errorf("Policy ID is not set")
	}
	d.SetId(id)
	// Need to populate ID attribute for subsequence processes
	object.ID = id

	// Creation completed
	logger.Infof("Creation of policy completed: %s", object.Name)
	return resourcePolicyRead(d, m)
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	logger.Infof("Beginning policy update: %s", ResourceIDString(d))

	client := m.(*restapi.RestClient)
	object := vault.NewPolicy(client)

	object.ID = d.Id()
	object.Plink.ID = d.Id()
	err := createUpateGetPolicyData(d, object)
	if err != nil {
		return fmt.Errorf("Error constructing policy data: %v", err)
	}

	// Deal with normal attribute changes first
	if d.HasChanges("name", "description", "link_type", "policy_assignment", "settings") {
		resp, err := object.Update()
		if err != nil || !resp.Success {
			return fmt.Errorf("Error updating policy attribute: %v", err)
		}
		logger.Debugf("Updated attributes to: %+v", object)
	}

	logger.Infof("Updating of policy completed: %s", object.Name)
	return resourcePolicyRead(d, m)
}

func createUpateGetPolicyData(d *schema.ResourceData, object *vault.Policy) error {
	object.Name = d.Get("name").(string)
	object.Plink.LinkType = d.Get("link_type").(string)
	if v, ok := d.GetOk("description"); ok {
		object.Description = v.(string)
	}
	if v, ok := d.GetOk("policy_assignment"); ok {
		object.Plink.Params = flattenSchemaSetToStringSlice(v)
	}
	if v, ok := d.GetOk("settings"); ok {
		settings := v.([]interface{})
		if len(settings) > 0 && settings[0] != nil {
			var err error
			object.Settings, err = expandSettingsData(v)
			if err != nil {
				return fmt.Errorf("Schema setting error: %s", err)
			}
			// Perform validations
			if err := object.ValidateSettings(); err != nil {
				return fmt.Errorf("Schema setting error: %s", err)
			}
		}
	}

	return nil
}

func expandSettingsData(v interface{}) (*vault.PolicySettings, error) {
	var err error
	s := v.([]interface{})
	//data := &settings{}
	if len(s) > 0 && s[0] != nil {
		d := s[0].(map[string]interface{})
		data := &vault.PolicySettings{
			//CentrifyServices:       expendCentrifyServicesData(d["centrify_services"]),
			//CentrifyClient:         expandCentrifyClientData(d["centrify_client"]),
			//CentrifyCSSServer:      expandCentrifyCSSServerData(d["centrify_css_server"]),
			//CentrifyCSSWorkstation: expandCentrifyCSSWorkstationData(d["centrify_css_workstation"]),
			//CentrifyCSSElevation:   expandCentrifyCSSElevationData(d["centrify_css_elevation"]),
			SelfService:      expandSelfServiceData(d["self_service"]),
			PasswordSettings: expandPasswordSettingsData(d["password_settings"]),
			OATHOTP:          expandOATHOTPData(d["oath_otp"]),
			Radius:           expandRadiusData(d["radius"]),
			UserAccount:      expandUserAccountData(d["user_account"]),
			//SystemSet:        expandSystemSetData(d["system_set"]),
			DatabaseSet: expandDatabaseSetData(d["database_set"]),
			DomainSet:   expandDomainSetData(d["domain_set"]),
			//AccountSet:       expandAccountSetData(d["account_set"]),
			//SecretSet:        expandSecretSetData(d["secret_set"]),
			//SSHKeySet:    expandSSHKeySetData(d["sshkey_set"]),
			MobileDevice: expandMobileDeviceData(d["mobile_device"]),
		}

		if data.CentrifyServices, err = expendCentrifyServicesData(d["centrify_services"]); err != nil {
			return nil, err
		}
		logger.Debugf("data.CentrifyServices data: %+v", data.CentrifyServices)
		if data.CentrifyClient, err = expandCentrifyClientData(d["centrify_client"]); err != nil {
			return nil, err
		}
		if data.CentrifyCSSServer, err = expandCentrifyCSSServerData(d["centrify_css_server"]); err != nil {
			return nil, err
		}
		if data.CentrifyCSSWorkstation, err = expandCentrifyCSSWorkstationData(d["centrify_css_workstation"]); err != nil {
			return nil, err
		}
		if data.CentrifyCSSElevation, err = expandCentrifyCSSElevationData(d["centrify_css_elevation"]); err != nil {
			return nil, err
		}
		if data.SystemSet, err = expandSystemSetData(d["system_set"]); err != nil {
			return nil, err
		}
		if data.AccountSet, err = expandAccountSetData(d["account_set"]); err != nil {
			return nil, err
		}
		if data.SecretSet, err = expandSecretSetData(d["secret_set"]); err != nil {
			return nil, err
		}
		if data.SSHKeySet, err = expandSSHKeySetData(d["sshkey_set"]); err != nil {
			return nil, err
		}
		if data.CloudProvidersSet, err = expandCloudProvidersSetData(d["cloudproviders_set"]); err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func expendCentrifyServicesData(v interface{}) (*vault.PolicyCentrifyServices, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCentrifyServices{
			// Session Parameters
			AuthenticationEnabled:  d["authentication_enabled"].(bool),
			DefaultProfileID:       d["default_profile_id"].(string),
			SessionLifespan:        d["session_lifespan"].(int),
			AllowSessionPersist:    d["allow_session_persist"].(bool),
			DefaultSessionPersist:  d["default_session_persist"].(bool),
			PersistSessionLifespan: d["persist_session_lifespan"].(int),
			// Other Settings
			AllowIwa:                   d["allow_iwa"].(bool),
			IwaSetKnownEndpoint:        d["iwa_set_cookie"].(bool),
			IwaSatisfiesAll:            d["iwa_satisfies_all"].(bool),
			UseCertAuth:                d["use_certauth"].(bool),
			CertAuthSkipChallenge:      d["certauth_skip_challenge"].(bool),
			CertAuthSetKnownEndpoint:   d["certauth_set_cookie"].(bool),
			CertAuthSatisfiesAll:       d["certauth_satisfies_all"].(bool),
			NoMfaMechLogin:             d["allow_no_mfa_mech"].(bool),
			FederatedLoginAllowsMfa:    d["auth_rule_federated"].(bool),
			FederatedLoginSatisfiesAll: d["federated_satisfies_all"].(bool),
			BlockMechsOnMobileLogin:    d["block_auth_from_same_device"].(bool),
			ContinueFailedSessions:     d["continue_failed_sessions"].(bool),
			SkipMechsInFalseAdvance:    d["stop_auth_on_prev_failed"].(bool),
			RememberLastAuthFactor:     d["remember_last_factor"].(bool),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandCentrifyClientData(v interface{}) (*vault.PolicyCentrifyClient, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCentrifyClient{
			AuthenticationEnabled: d["authentication_enabled"].(bool),
			DefaultProfileID:      d["default_profile_id"].(string),
			NoMfaMechLogin:        d["allow_no_mfa_mech"].(bool),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandCentrifyCSSServerData(v interface{}) (*vault.PolicyCentrifyCSSServer, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCentrifyCSSServer{
			AuthenticationEnabled: d["authentication_enabled"].(bool),
			DefaultProfileID:      d["default_profile_id"].(string),
			PassThroughMode:       d["pass_through_mode"].(int),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandCentrifyCSSWorkstationData(v interface{}) (*vault.PolicyCentrifyCSSWorkstation, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCentrifyCSSWorkstation{
			AuthenticationEnabled: d["authentication_enabled"].(bool),
			DefaultProfileID:      d["default_profile_id"].(string),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandCentrifyCSSElevationData(v interface{}) (*vault.PolicyCentrifyCSSElevation, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCentrifyCSSElevation{
			AuthenticationEnabled: d["authentication_enabled"].(bool),
			DefaultProfileID:      d["default_profile_id"].(string),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandSelfServiceData(v interface{}) *vault.PolicySelfService {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicySelfService{
			// Password Reset
			AccountSelfServiceEnabled:    d["account_selfservice_enabled"].(bool),
			PasswordResetEnabled:         d["password_reset_enabled"].(bool),
			PasswordResetADEnabled:       d["pwreset_allow_for_aduser"].(bool),
			PasswordResetCookieOnly:      d["pwreset_with_cookie_only"].(bool),
			PasswordResetRequiresRelogin: d["login_after_reset"].(bool),
			PasswordResetAuthProfile:     d["pwreset_auth_profile_id"].(string),
			PasswordResetMaxAttempts:     d["max_reset_attempts"].(int),
			// Account Unlock
			AccountUnlockEnabled:     d["account_unlock_enabled"].(bool),
			AccountUnlockADEnabled:   d["unlock_allow_for_aduser"].(bool),
			AccountUnlockCookieOnly:  d["unlock_with_cookie_only"].(bool),
			ShowAccountLocked:        d["show_locked_message"].(bool),
			AccountUnlockAuthProfile: d["unlock_auth_profile_id"].(string),
			// Active Directory Self Service Settings
			UseADAdmin:  d["use_ad_admin"].(bool),
			ADAdminUser: d["ad_admin_user"].(string),
			ADAdminPass: expandAdminPassword(d["admin_user_password"]),
			// Additional Policy Parameters
			MaxResetAllowed: d["max_reset_allowed"].(int),
			MaxTimeAllowed:  d["max_time_allowed"].(int),
		}
		return data
	}
	return nil
}

func expandAdminPassword(v interface{}) *vault.PolicyADAdminPass {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyADAdminPass{
			Type:  d["type"].(string),
			Value: d["value"].(string),
		}
		return data
	}
	return nil
}

func expandPasswordSettingsData(v interface{}) *vault.PolicyPasswordSettings {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyPasswordSettings{
			// Password Requirements
			MinLength:      d["min_length"].(int),
			MaxLength:      d["max_length"].(int),
			RequireDigit:   d["require_digit"].(bool),
			RequireMixCase: d["require_mix_case"].(bool),
			RequireSymbol:  d["require_symbol"].(bool),
			// Display Requirements
			ShowPasswordComplexity: d["show_password_complexity"].(bool),
			NonCdsComplexityHint:   d["complexity_hint"].(string),
			// Additional Requirements
			AllowRepeatedChar:       d["no_of_repeated_char_allowed"].(int),
			CheckWeakPassword:       d["check_weak_password"].(bool),
			AllowIncludeUsername:    d["allow_include_username"].(bool),
			AllowIncludeDisplayname: d["allow_include_displayname"].(bool),
			RequireUnicode:          d["require_unicode"].(bool),
			// Password Age
			MinAgeInDays:   d["min_age_in_days"].(int),
			MaxAgeInDays:   d["max_age_in_days"].(int),
			History:        d["password_history"].(int),
			NotifySoft:     d["expire_soft_notification"].(int),
			NotifyHard:     d["expire_hard_notification"].(int),
			NotifyOnMobile: d["expire_notification_mobile"].(bool),
			// Capture Settings
			BadAttemptThreshold: d["bad_attempt_threshold"].(int),
			CaptureWindow:       d["capture_window"].(int),
			LockoutDuration:     d["lockout_duration"].(int),
		}
		return data
	}
	return nil
}

func expandOATHOTPData(v interface{}) *vault.PolicyOathOTP {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyOathOTP{
			AllowOTP: d["allow_otp"].(bool),
		}
		return data
	}
	return nil
}

func expandRadiusData(v interface{}) *vault.PolicyRadius {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyRadius{
			AllowRadius:          d["allow_radius"].(bool),
			RadiusUseChallenges:  d["require_challenges"].(bool),
			DefaultProfileID:     d["default_profile_id"].(string),
			SendVendorAttributes: d["send_vendor_attributes"].(bool),
			AllowExternalRadius:  d["allow_external_radius"].(bool),
		}
		return data
	}
	return nil
}

func expandUserAccountData(v interface{}) *vault.PolicyUserAccount {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyUserAccount{
			UserChangePasswordAllow:     d["allow_user_change_password"].(bool),
			PasswordChangeAuthProfileID: d["password_change_auth_profile_id"].(string),
			ShowU2f:                     d["show_fido2"].(bool),
			U2fPrompt:                   d["fido2_prompt"].(string),
			U2fAuthProfileID:            d["fido2_auth_profile_id"].(string),
			ShowQRCode:                  d["show_otp"].(bool),
			OTPPrompt:                   d["otp_prompt"].(string),
			OTPAuthProfileID:            d["otp_auth_profile_id"].(string),
			ConfigureSecurityQuestions:  d["configure_security_questions"].(bool),
			AllowDupAnswers:             d["prevent_dup_answers"].(bool),
			UserDefinedQuestions:        d["user_defined_questions"].(int),
			AdminDefinedQuestions:       d["admin_defined_questions"].(int),
			MinCharInAnswer:             d["min_char_in_answer"].(int),
			QuestionAuthProfileID:       d["question_auth_profile_id"].(string),
			PhonePinChangeAllow:         d["allow_phone_pin_change"].(bool),
			MinPhonePinLength:           d["min_phone_pin_length"].(int),
			PhonePinAuthProfileID:       d["phone_pin_auth_profile_id"].(string),
			AllowUserChangeMFARedirect:  d["allow_mfa_redirect_change"].(bool),
			UserProfileAuthProfileID:    d["user_profile_auth_profile_id"].(string),
			DefaultLanguage:             d["default_language"].(string),
		}
		return data
	}
	return nil
}

func expandSystemSetData(v interface{}) (*vault.PolicySystemSet, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicySystemSet{
			// Account Policy
			DefaultCheckoutTime: d["checkout_lifetime"].(int),
			// System Policy
			AllowRemote:                           d["allow_remote_access"].(bool),
			AllowRdpClipboard:                     d["allow_rdp_clipboard"].(bool),
			AllowAutomaticLocalAccountMaintenance: d["local_account_automatic_maintenance"].(bool),
			AllowManualLocalAccountUnlock:         d["local_account_manual_unlock"].(bool),
			LoginDefaultProfile:                   d["default_profile_id"].(string),
			PrivilegeElevationDefaultProfile:      d["privilege_elevation_default_profile_id"].(string),
			// Security Settings
			RemoveUserOnSessionEnd:            d["remove_user_on_session_end"].(bool),
			AllowMultipleCheckouts:            d["allow_multiple_checkouts"].(bool),
			AllowPasswordRotation:             d["enable_password_rotation"].(bool),
			PasswordRotateDuration:            d["password_rotate_interval"].(int),
			AllowPasswordRotationAfterCheckin: d["enable_password_rotation_after_checkin"].(bool),
			MinimumPasswordAge:                d["minimum_password_age"].(int),
			MinimumSSHKeysAge:                 d["minimum_sshkey_age"].(int),
			AllowSSHKeysRotation:              d["enable_sshkey_rotation"].(bool),
			SSHKeysRotateDuration:             d["sshkey_rotate_interval"].(int),
			SSHKeysGenerationAlgorithm:        d["sshkey_algorithm"].(string),
			// Maintenance Settings
			AllowPasswordHistoryCleanUp:    d["enable_password_history_cleanup"].(bool),
			PasswordHistoryCleanUpDuration: d["password_historycleanup_duration"].(int),
			AllowSSHKeysCleanUp:            d["enable_sshkey_history_cleanup"].(bool),
			SSHKeysCleanUpDuration:         d["sshkey_historycleanup_duration"].(int),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		if v, ok := d["privilege_elevation_rule"]; ok {
			data.PrivilegeElevationRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.PrivilegeElevationRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandDatabaseSetData(v interface{}) *vault.PolicyDatabaseSet {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyDatabaseSet{
			// Account Policy
			DefaultCheckoutTime: d["checkout_lifetime"].(int),
			// Security Settings
			AllowMultipleCheckouts:            d["allow_multiple_checkouts"].(bool),
			AllowPasswordRotation:             d["enable_password_rotation"].(bool),
			PasswordRotateDuration:            d["password_rotate_interval"].(int),
			AllowPasswordRotationAfterCheckin: d["enable_password_rotation_after_checkin"].(bool),
			MinimumPasswordAge:                d["minimum_password_age"].(int),
			// Maintenance Settings
			AllowPasswordHistoryCleanUp:    d["enable_password_history_cleanup"].(bool),
			PasswordHistoryCleanUpDuration: d["password_historycleanup_duration"].(int),
		}
		return data
	}
	return nil
}

func expandDomainSetData(v interface{}) *vault.PolicyDomainSet {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyDomainSet{
			// Domain Policy
			DefaultCheckoutTime: d["checkout_lifetime"].(int),
			// Security Settings
			AllowMultipleCheckouts:            d["allow_multiple_checkouts"].(bool),
			AllowPasswordRotation:             d["enable_password_rotation"].(bool),
			PasswordRotateDuration:            d["password_rotate_interval"].(int),
			AllowPasswordRotationAfterCheckin: d["enable_password_rotation_after_checkin"].(bool),
			MinimumPasswordAge:                d["minimum_password_age"].(int),
			// Maintenance Settings
			AllowPasswordHistoryCleanUp:    d["enable_password_history_cleanup"].(bool),
			PasswordHistoryCleanUpDuration: d["password_historycleanup_duration"].(int),
		}
		return data
	}
	return nil
}

func expandAccountSetData(v interface{}) (*vault.PolicyAccountSet, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyAccountSet{
			DefaultCheckoutTime:                d["checkout_lifetime"].(int),
			PasswordCheckoutDefaultProfile:     d["default_profile_id"].(string),
			AccessSecretCheckoutDefaultProfile: d["access_secret_checkout_dfault_profile_id"].(string),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		if v, ok := d["access_secret_checkout_rule"]; ok {
			data.AccessSecretCheckoutRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.AccessSecretCheckoutRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandSecretSetData(v interface{}) (*vault.PolicySecretSet, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicySecretSet{
			DataVaultDefaultProfile: d["default_profile_id"].(string),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandSSHKeySetData(v interface{}) (*vault.PolicySshKeySet, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicySshKeySet{
			SSHKeysDefaultProfile: d["default_profile_id"].(string),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandCloudProvidersSetData(v interface{}) (*vault.PolicyCloudProvidersSet, error) {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyCloudProvidersSet{
			LoginDefaultProfile:                       d["default_profile_id"].(string),
			EnableUnmanagedPasswordRotation:           d["enable_interactive_password_rotation"].(bool),
			EnableUnmanagedPasswordRotationPrompt:     d["prompt_change_root_password"].(bool),
			EnableUnmanagedPasswordRotationReminder:   d["enable_password_rotation_reminders"].(bool),
			UnmanagedPasswordRotationReminderDuration: d["password_rotation_reminder_duration"].(int),
		}
		if v, ok := d["challenge_rule"]; ok {
			data.ChallengeRules = expandChallengeRules(v.([]interface{}))
			// Perform validations
			if err := validateChallengeRules(data.ChallengeRules); err != nil {
				return nil, fmt.Errorf("Schema setting error: %s", err)
			}
		}
		return data, nil
	}
	return nil, nil
}

func expandMobileDeviceData(v interface{}) *vault.PolicyMobileDevice {
	subsettings := v.([]interface{})
	if len(subsettings) > 0 && subsettings[0] != nil {
		d := subsettings[0].(map[string]interface{})
		data := &vault.PolicyMobileDevice{
			AllowEnrollment:           d["allow_enrollment"].(bool),
			AllowJailBrokenDevices:    d["permit_non_compliant_device"].(bool),
			EnableInviteEnrollment:    d["enable_invite_enrollment"].(bool),
			AllowNotifnMutipleDevices: d["allow_notify_multi_devices"].(bool),
			AllowDebugLogging:         d["enable_debug"].(bool),
			LocationTracking:          d["location_tracking"].(bool),
			ForceFingerprint:          d["force_fingerprint"].(bool),
			AllowFallbackAppPin:       d["allow_fallback_pin"].(bool),
			RequestPasscode:           d["require_passcode"].(bool),
			AutoLockTimeout:           d["auto_lock_timeout"].(int),
			AppLockOnExit:             d["lock_app_on_exit"].(bool),
		}
		return data
	}
	return nil
}
