package centrify

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func getCentrifyServicesSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Authentication -> Centrify Services",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable authentication policy controls",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"session_lifespan": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Hours until session expires (default 12)",
					ValidateFunc: validation.IntBetween(1, 9999),
				},
				// Session Parameters
				"allow_session_persist": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow 'Keep me signed in' checkbox option at login (session spans browser sessions)",
				},
				"default_session_persist": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Default 'Keep me signed in' checkbox option to enabled",
				},
				"persist_session_lifespan": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Hours until session expires when 'Keep me signed in' option enabled (default 2 weeks)",
					ValidateFunc: validation.IntBetween(1, 9999),
				},
				// Other Settings
				"allow_iwa": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Allow IWA connections (bypasses authentication rules and default profile)",
				},
				"iwa_set_cookie": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Set identity cookie for IWA connections",
				},
				"iwa_satisfies_all": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "IWA connections satisfy all MFA mechanisms",
				},
				"use_certauth": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Use certificates for authentication",
				},
				"certauth_skip_challenge": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Certificate authentication bypasses authentication rules and default profile",
				},
				"certauth_set_cookie": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Set identity cookie for connections using certificate authentication",
				},
				"certauth_satisfies_all": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Connections using certificate authentication satisfy all MFA mechanisms",
				},
				"allow_no_mfa_mech": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow users without a valid authentication factor to log in",
				},
				"auth_rule_federated": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Apply additional authentication rules to federated users",
				},
				"federated_satisfies_all": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Connections via Federation satisfy all MFA mechanisms",
				},
				"block_auth_from_same_device": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Allow additional authentication from same device",
				},
				"continue_failed_sessions": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Continue with additional challenges after failed challenge",
				},
				"stop_auth_on_prev_failed": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Do not send challenge request when previous challenge response failed",
				},
				"remember_last_factor": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Remember and suggest last used authentication factor",
				},
			},
		},
	}
}

func getCentrifyClientSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Authentication -> Centrify Clients -> Login",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable authentication policy controls",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"allow_no_mfa_mech": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow users without a valid authentication factor to log in",
				},
			},
		},
	}
}

func getCentrifyCSSServerSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Authentication -> Centrify Server Suite Agents -> Linux, UNIX and Windows Servers",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable authentication policy controls",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"pass_through_mode": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Apply pass-through duration: Never (Default), If Same Source and Target, If Same Source, If Same Target",
					ValidateFunc: validation.IntBetween(0, 3),
				},
			},
		},
	}
}

func getCentrifyCSSWorkstationSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Authentication -> Centrify Server Suite Agents -> Windows Workstations",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable authentication policy controls",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
			},
		},
	}
}

func getCentrifyCSSElevationSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Authentication -> Centrify Server Suite Agents -> Privilege Elevation",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable authentication policy controls",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
			},
		},
	}
}

func getSelfServiceSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "User Security -> Self Service",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Password Reset
				"account_selfservice_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable account self service controls",
				},
				"password_reset_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Enable password reset",
				},
				"pwreset_allow_for_aduser": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow for Active Directory users",
				},
				"pwreset_with_cookie_only": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Only allow from browsers with identity cookie",
				},
				"login_after_reset": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "User must log in after successful password reset",
				},
				"pwreset_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Password reset authentication profile",
				},
				"max_reset_attempts": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Maximum consecutive password reset attempts per session",
					ValidateFunc: validation.IntBetween(0, 10),
				},
				// Account Unlock
				"account_unlock_enabled": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable account unlock",
				},
				"unlock_allow_for_aduser": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow for Active Directory users",
				},
				"unlock_with_cookie_only": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Only allow from browsers with identity cookie",
				},
				"show_locked_message": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Show a message to end users in desktop login that account is locked (default no)",
				},
				"unlock_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Account unlock authentication profile",
				},
				// Active Directory Self Service Settings
				"use_ad_admin": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Use AD admin for AD self-service",
				},
				"ad_admin_user": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Admin user name",
				},
				"admin_user_password": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"type": {
								Type:        schema.TypeString,
								Optional:    true,
								Default:     "SafeString",
								Description: "Password type",
							},
							"value": {
								Type:        schema.TypeString,
								Optional:    true,
								Sensitive:   true,
								Description: "Actual password",
							},
						},
					},
				},
				// Additional Policy Parameters
				"max_reset_allowed": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Maximum forgotten password resets allowed within window (default 10)",
					ValidateFunc: validation.IntBetween(0, 10),
				},
				"max_time_allowed": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Capture window for forgotten password resets (default 60 minutes)",
					ValidateFunc: validation.IntInSlice([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}),
				},
			},
		},
	}
}

func getPasswordSettingsSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "User Security -> Password Settings",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Password Requirements
				"min_length": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum password length (default 8)",
					ValidateFunc: validation.IntBetween(4, 16),
				},
				"max_length": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Maximum password length (default 64)",
					ValidateFunc: validation.IntBetween(4, 64),
				},
				"require_digit": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require at least one digit (default yes)",
				},
				"require_mix_case": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require at least one upper case and one lower case letter (default yes)",
				},
				"require_symbol": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require at least one symbol (default no)",
				},
				// Display Requirements
				"show_password_complexity": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Show password complexity requirements when entering a new password (default no)",
				},
				"complexity_hint": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Password complexity requirements for directory services other than Centrify Directory",
				},
				// Additional Requirements
				"no_of_repeated_char_allowed": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Limit the number of consecutive repeated characters",
					ValidateFunc: validation.IntBetween(2, 64),
				},
				"check_weak_password": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Check against weak password",
				},
				"allow_include_username": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow username as part of password",
				},
				"allow_include_displayname": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow display name as part of password",
				},
				"require_unicode": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require at least one Unicode characters",
				},
				// Password Age
				"min_age_in_days": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum password age before change is allowed (default 0 days)",
					ValidateFunc: validation.IntBetween(0, 998),
				},
				"max_age_in_days": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Maximum password age (default 365 days)",
					ValidateFunc: validation.IntBetween(0, 3650),
				},
				"password_history": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password history (default 3)",
					ValidateFunc: validation.IntBetween(0, 25),
				},
				"expire_soft_notification": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password Expiration Notification (default 14 days)",
					ValidateFunc: validation.IntInSlice([]int{7, 14, 21, 28, 35, 42, 49}),
				},
				"expire_hard_notification": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Escalated Password Expiration Notification (default 48 hours)",
					ValidateFunc: validation.IntInSlice([]int{24, 48, 72, 96, 120}),
				},
				"expire_notification_mobile": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable password expiration notifications on enrolled mobile devices",
				},
				// Capture Settings
				"bad_attempt_threshold": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Maximum consecutive bad password attempts allowed within window (default Off)",
					ValidateFunc: validation.IntBetween(0, 10),
				},
				"capture_window": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Capture window for consecutive bad password attempts (default 30 minutes)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
				"lockout_duration": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Lockout duration before password re-attempt allowed (default 30 minutes)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
			},
		},
	}
}

func getOATHOTPSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "User Security -> OATH OTP",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_otp": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow OATH OTP integration",
				},
			},
		},
	}
}

func getRadiusSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "User Security -> RADIUS",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_radius": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow RADIUS client connections",
				},
				"require_challenges": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require authentication challenge",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default authentication profile",
				},
				"send_vendor_attributes": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Send vendor specific attributes",
				},
				"allow_external_radius": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow 3rd Party RADIUS Authentication",
				},
			},
		},
	}
}

func getUserAccountSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "User Security -> User Account",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_user_change_password": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to change their passwords",
				},
				"password_change_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to change password",
				},
				"show_fido2": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to enroll FIDO2 Authenticators",
				},
				"fido2_prompt": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "FIDO2 Security Key Display Name",
				},
				"fido2_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to configure FIDO2 Authenticators",
				},
				"show_otp": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to configure an OATH OTP client (requires enabling OATH OTP policy)",
				},
				"otp_prompt": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "OATH OTP Display Name",
				},
				"otp_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to configure OATH OTP client",
				},
				"configure_security_questions": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to configure Security Questions",
				},
				"prevent_dup_answers": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow duplicate security question answers",
				},
				"user_defined_questions": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Required number of user-defined questions",
					ValidateFunc: validation.IntBetween(0, 20),
				},
				"admin_defined_questions": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Required number of admin-defined questions",
					ValidateFunc: validation.IntBetween(0, 20),
				},
				"min_char_in_answer": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum number of characters required in answers",
					ValidateFunc: validation.IntAtLeast(1),
				},
				"question_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to set Security Questions",
				},
				"allow_phone_pin_change": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to configure a Phone PIN for MFA",
				},
				"min_phone_pin_length": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum Phone PIN length",
					ValidateFunc: validation.IntBetween(4, 8),
				},
				"phone_pin_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to configure a Phone PIN",
				},
				"allow_mfa_redirect_change": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable users to redirect multi factor authentication to a different user account",
				},
				"user_profile_auth_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile required to modify Personal Profile",
				},
				"default_language": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Default Language",
					ValidateFunc: validation.StringInSlice([]string{"de", "en", "es", "fr", "it", "ja", "ko", "pt-br", "zh-cn", "zh"}, false),
				},
			},
		},
	}
}

func getSystemSetSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> Systems",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Account Policy
				"checkout_lifetime": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Checkout lifetime (minutes)",
					ValidateFunc: validation.IntBetween(15, 2147483647),
				},
				// System Policy
				"allow_remote_access": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow access from a public network (web client only)",
				},
				"allow_rdp_clipboard": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow RDP client to sync local clipboard with remote session",
				},
				"local_account_automatic_maintenance": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable local account automatic maintenance",
				},
				"local_account_manual_unlock": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable local account manual unlock",
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default System Login Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"privilege_elevation_default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Privilege Elevation Profile (used if no conditions matched)",
				},
				"privilege_elevation_rule": getChallengeRulesSchema(),
				// Security Settings
				"remove_user_on_session_end": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Remove local accounts upon session termination - Windows only ",
				},
				"allow_multiple_checkouts": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow multiple password checkouts for this system",
				},
				"enable_password_rotation": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic password rotation",
				},
				"password_rotate_interval": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password rotation interval (days)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
				"enable_password_rotation_after_checkin": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable password rotation after checkin",
				},
				"minimum_password_age": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum Password Age (days)",
					ValidateFunc: validation.IntBetween(0, 2147483647),
				},
				"minimum_sshkey_age": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum SSH Key Age (days)",
					ValidateFunc: validation.IntBetween(0, 2147483647),
				},
				"enable_sshkey_rotation": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic SSH key rotation",
				},
				"sshkey_rotate_interval": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "SSH key rotation interval (days)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
				"sshkey_algorithm": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "SSH Key Generation Algorithm",
					ValidateFunc: validation.StringInSlice([]string{
						"RSA_1024",
						"RSA_2048",
						"ECDSA_P256",
						"ECDSA_P384",
						"ECDSA_P521",
						"EdDSA_Ed448",
						"EdDSA_Ed25519",
					}, false),
				},
				// Maintenance Settings
				"enable_password_history_cleanup": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic password history cleanup",
				},
				"password_historycleanup_duration": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password history cleanup (days)",
					ValidateFunc: validation.IntBetween(90, 2147483647),
				},
				"enable_sshkey_history_cleanup": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic SSH key history cleanup",
				},
				"sshkey_historycleanup_duration": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "SSH key history cleanup (days)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
			},
		},
	}
}

func getDatabaseAndDomainSetSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> Databases",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Account Security
				"checkout_lifetime": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Checkout lifetime (minutes)",
					ValidateFunc: validation.IntBetween(15, 2147483647),
				},
				// Security Settings
				"allow_multiple_checkouts": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow multiple password checkouts for related accounts",
				},
				"enable_password_rotation": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic password rotation",
				},
				"password_rotate_interval": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password rotation interval (days)",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
				"enable_password_rotation_after_checkin": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable password rotation after checkin",
				},
				"minimum_password_age": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum Password Age (days)",
					ValidateFunc: validation.IntBetween(0, 2147483647),
				},
				// Maintenance Settings
				"enable_password_history_cleanup": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable periodic password history cleanup",
				},
				"password_historycleanup_duration": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Password history cleanup (days)",
					ValidateFunc: validation.IntBetween(90, 2147483647),
				},
			},
		},
	}
}

func getAccountSetSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> Accounts",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"checkout_lifetime": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Checkout lifetime (minutes)",
					ValidateFunc: validation.IntBetween(15, 2147483647),
				},
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Password Checkout Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"access_secret_checkout_dfault_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Password Checkout Profile (used if no conditions matched)",
				},
				"access_secret_checkout_rule": getChallengeRulesSchema(),
			},
		},
	}
}

func getSecretSetSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> Secrets",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Secret Challenge Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
			},
		},
	}
}

func getSSHKeySetSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> SSH Keys",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default SSH Key Challenge Profile",
				},
				"challenge_rule": getChallengeRulesSchema(),
			},
		},
	}
}

func getCloudProvidersSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Resouces -> Cloud Providers",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Default Root Account Login Profile (used if no conditions matched)",
				},
				"challenge_rule": getChallengeRulesSchema(),
				"enable_interactive_password_rotation": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable interactive password rotation",
				},
				"prompt_change_root_password": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Prompt to change root password every login and password checkin",
				},
				"enable_password_rotation_reminders": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable password rotation reminders",
				},
				"password_rotation_reminder_duration": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Minimum number of days since last rotation to trigger a reminder",
					ValidateFunc: validation.IntBetween(1, 2147483647),
				},
			},
		},
	}
}

func getMobileDeviceSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Devices",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_enrollment": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Permit device registration",
				},
				"permit_non_compliant_device": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Permit non-compliant devices to register",
				},
				"enable_invite_enrollment": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable invite based registration",
				},
				"allow_notify_multi_devices": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow user notifications on multiple devices",
				},
				"enable_debug": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enable debug logging",
				},
				"location_tracking": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Report mobile device location",
				},
				"force_fingerprint": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Enforce fingerprint scan for Mobile Authenticator",
				},
				"allow_fallback_pin": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Allow App PIN",
				},
				"require_passcode": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Require client application passcode on device",
				},
				"auto_lock_timeout": {
					Type:         schema.TypeInt,
					Optional:     true,
					Description:  "Auto-Lock (minutes)",
					ValidateFunc: validation.IntInSlice([]int{1, 2, 5, 15, 30}),
				},
				"lock_app_on_exit": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "Lock on exit",
				},
			},
		},
	}
}
