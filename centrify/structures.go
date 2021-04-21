package centrify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	logger "github.com/marcozj/golang-sdk/logging"
	vault "github.com/marcozj/golang-sdk/platform"
)

// ResourceIDStringInterface - Generic interface for resource ID
type ResourceIDStringInterface interface {
	Id() string
}

// ResourceIDString - Obtain resource ID string
func ResourceIDString(d ResourceIDStringInterface) string {
	id := d.Id()
	if id == "" {
		id = "<new resource>"
	}
	return fmt.Sprintf("(ID = %s)", id)
}

// flattenTypeListToSlice converts simple schema.TypeList to slices
func flattenTypeListToSlice(i interface{}) []string {
	var lstr []string

	for _, v := range i.([]interface{}) {
		lstr = append(lstr, v.(string))
	}

	return lstr
}

// convert string slice to interface slice
func flattenStringSlice(list []string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}

// flattenSchemaSetToString converts ["value1", "value2"] to "value1,value2"
func flattenSchemaSetToString(input *schema.Set) string {
	var str string
	for i, v := range input.List() {
		str = str + v.(string)
		// Append "," if it is not the last element
		if i < input.Len()-1 {
			str = str + ","
		}
	}

	return str
}

// flattenSliceToString converts ["value1", "value2"] to "value1,value2"
func flattenSliceToString(input []string) string {
	var str string
	for i, v := range input {
		str = str + v
		// Append "," if it is not the last element
		if i < len(input)-1 {
			str = str + ","
		}
	}

	return str
}

// flattenSchemaSetToStringSlice used for converting terraform schema set to a string slice
func flattenSchemaSetToStringSlice(s interface{}) []string {
	vL := []string{}

	for _, v := range s.(*schema.Set).List() {
		vL = append(vL, v.(string))
	}

	return vL
}

// flattenSchemaListToStringSlice used for converting terraform attribute of TypeString embedded in TypeList to a string slice.
// it expected interface{} type as []interface{}, usually get the value from `d.Get` of terraform resource data.
func flattenSchemaListToStringSlice(iface interface{}) []string {
	s := []string{}

	for _, i := range iface.([]interface{}) {
		s = append(s, i.(string))
	}

	return s
}

// flattenStringSliceToSet converts slice of string to schema Set
func flattenStringSliceToSet(input []string) *schema.Set {
	var out = make([]interface{}, len(input))
	for i, v := range input {
		out[i] = v
	}
	return schema.NewSet(schema.HashString, out)
}

func expandListToMap(input []interface{}) map[string]interface{} {
	options := make(map[string]interface{})

	for _, option := range input {
		for optKey, optValue := range option.(map[string]interface{}) {
			options[optKey] = optValue
		}

	}

	return options
}

func expendSchemaSetToMap(input *schema.Set) map[string]interface{} {
	options := make(map[string]interface{})
	if input.Len() > 0 {
		options = expandListToMap(input.List())
	}
	return options
}

// Assemble the hash for the system proxy_account
// TypeSet attribute.
func customProxyAccountHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["username"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["managed"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(bool)))
	}
	return hashcode.String(buf.String())
}

func customRoleMemberHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["type"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}

func customPermissionsHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["principal_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["rights"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", v.(*schema.Set).GoString()))
	}
	return hashcode.String(buf.String())
}

func customLoginRuleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["filter"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["condition"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["value"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}

func customCommandParamHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["type"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["target_object_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["value"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}

func customAccessKeyHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["access_key_id"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return hashcode.String(buf.String())
}

func customGroupMappingHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["attribute_value"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["group_name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}

func customZoneRoleHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return hashcode.String(buf.String())
}

func customSamlAttributeHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["value"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return hashcode.String(buf.String())
}

func customOauthScopeHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["description"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["allowed_rest_apis"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(*schema.Set).GoString()))
	}

	return hashcode.String(buf.String())
}

// StringSliceToInterface converts []string to []interface{}
func StringSliceToInterface(s []string) []interface{} {
	i := make([]interface{}, len(s))
	for j, v := range s {
		i[j] = v
	}
	return i
}

func expandRoleMembers(v interface{}) []vault.RoleMember {
	members := []vault.RoleMember{}

	for _, p := range v.(*schema.Set).List() {
		member := vault.RoleMember{}
		member.MemberID = p.(map[string]interface{})["id"].(string)
		member.MemberName = p.(map[string]interface{})["name"].(string)
		member.MemberType = p.(map[string]interface{})["type"].(string)
		members = append(members, member)
	}
	logger.Debugf("Role members: %+v", members)

	return members
}

func expandPermissions(v interface{}, valid map[string]string, validate bool) ([]vault.Permission, error) {
	m := v.(*schema.Set).List()
	var permissions []vault.Permission
	if m != nil {
		for _, v := range m {
			// Validate given list of permissions against a valid map
			existing := v.(map[string]interface{})["rights"].(*schema.Set)
			var converted []string
			for _, r := range existing.List() {
				if valid[r.(string)] != "" {
					converted = append(converted, valid[r.(string)])
				} else {
					if validate {
						return nil, fmt.Errorf("For %s, %v can only contain %v", v.(map[string]interface{})["principal_name"].(string), existing.List(), valid)
					}
				}
			}
			// Convert map to Permission object
			permission := vault.Permission{
				PrincipalID:   v.(map[string]interface{})["principal_id"].(string),
				PrincipalName: v.(map[string]interface{})["principal_name"].(string),
				PrincipalType: v.(map[string]interface{})["principal_type"].(string),
				Rights:        flattenSliceToString(converted),
			}
			permissions = append(permissions, permission)
		}
	}
	return permissions, nil
}

func expandChallengeRules(v []interface{}) *vault.ChallengeRules {
	challengerules := &vault.ChallengeRules{}
	// Deal with root level
	challengerules.Enabled = true
	challengerules.Type = "RowSet"
	challengerules.UniqueKey = "Condition"

	for _, lrv := range v {
		// Deal with "_Value" level
		challengerule := vault.ChallengeRule{}
		challengerule.AuthProfileID = lrv.(map[string]interface{})["authentication_profile_id"].(string)
		rules := lrv.(map[string]interface{})["rule"]

		for _, rv := range rules.(*schema.Set).List() {
			// Deal with "Conditions" level
			rule := vault.ChallengeCondition{}
			rule.Filter = rv.(map[string]interface{})["filter"].(string)
			rule.Condition = rv.(map[string]interface{})["condition"].(string)
			rule.Value = rv.(map[string]interface{})["value"].(string)
			challengerule.ChallengeCondition = append(challengerule.ChallengeCondition, rule)
		}
		challengerules.Rules = append(challengerules.Rules, challengerule)
	}

	return challengerules
}

func expandCommandParams(v interface{}) []vault.DesktopAppParam {
	parms := []vault.DesktopAppParam{}

	for _, p := range v.(*schema.Set).List() {
		parm := vault.DesktopAppParam{}
		parm.ParamName = p.(map[string]interface{})["name"].(string)
		parm.ParamType = p.(map[string]interface{})["type"].(string)
		parm.TargetObjectID = p.(map[string]interface{})["target_object_id"].(string)
		parms = append(parms, parm)
	}
	logger.Debugf("Command params: %+v", parms)

	return parms
}

func getPermissionSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Object permissions",
		Set:         customPermissionsHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"principal_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Uuid of the principal",
				},
				"principal_name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "User name or role name",
				},
				"principal_type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Principal type: User, Role or Group",
					ValidateFunc: validation.StringInSlice([]string{
						"User",
						"Role",
						"Group",
					}, false),
				},
				"rights": {
					Type:     schema.TypeSet,
					Required: true,
					Set:      schema.HashString,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "Permissions: Grant,View,Edit,Delete",
				},
			},
		},
	}
}

func getChallengeRulesSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Authentication rules",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"authentication_profile_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Authentication Profile (if all conditions met)",
				},
				"rule": {
					Type:     schema.TypeSet,
					Optional: true,
					Set:      customLoginRuleHash,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"filter": {
								Type:     schema.TypeString,
								Required: true,
								ValidateFunc: validation.StringInSlice([]string{
									"IpAddress",
									"IdentityCookie",
									"DayOfWeek",
									"Date",
									"DateRange",
									"Time",
									"DeviceOs",
									"Browser",
									"CountryCode",
									"Zso",
								}, false),
							},
							"condition": {
								Type:     schema.TypeString,
								Required: true,
								ValidateFunc: validation.StringInSlice([]string{
									"OpInCorpIpRange",    // IpAddress
									"OpNotInCorpIpRange", // IpAddress
									"OpExists",           // IdentityCookie
									"OpNotExists",        // IdentityCookie
									"OpIsDayOfWeek",      // DayOfWeek
									"OpLessThan",         // Date
									"OpGreaterThan",      // Date
									"OpBetween",          // DateRange, Time
									"OpEqual",            // DeviceOs, Browser, CountryCode
									"OpNotEqual",         // DeviceOs, Browser, CountryCode
									"OpIs",               // Zso
									"OpIsNot",            // Zso
									"OpHeader",           // Header
									"OpArgument",         // Argument
								}, false),
							},
							"value": {
								Type:     schema.TypeString,
								Optional: true,
								/*
									DayOfWeek: "L,0,1,2,3,4,5,6" or "U,1,2"
									Date: "L,08/27/2020" or "U,08/29/2020"
									DateRange: "L,08/26/2020,08/29/2020" or "U,08/26/2020,08/29/2020"
									Time: "L,00:16,15:56" or "U,00:16,15:56"
									DeviceOs: iOS, Android, WindowsMobile, Mac, Windows, Linux
									Browser: Other, Chrome, Firefox, IE, Safari, MicrosoftEdge
								*/
							},
						},
					},
				},
			},
		},
	}
}

func getAccessKeySchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "AWS Access Key",
		Set:         customAccessKeyHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					Description: "AWS access key id",
				},
				"access_key_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "AWS access key id",
				},
				"secret_access_key": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					Description: "AWS secret access key",
				},
			},
		},
	}
}

func expandAccessKeys(v interface{}) []vault.AccessKey {
	accesskeys := []vault.AccessKey{}

	for _, p := range v.(*schema.Set).List() {
		accesskey := vault.AccessKey{}
		accesskey.AccessKeyID = p.(map[string]interface{})["access_key_id"].(string)
		accesskey.SecretAccessKey = p.(map[string]interface{})["secret_access_key"].(string)
		accesskeys = append(accesskeys, accesskey)
	}

	return accesskeys
}

func getZoneRoleSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Zone Role",
		Set:         customZoneRoleHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Zone role name. In the format of <zone role name>/<zone name>",
				},
				"zone_description": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Description of the zone",
				},
				"zone_dn": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Distinguished Name (DN) of the zone",
				},
				"description": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Description of the zone role",
				},
				"zone_canonical_name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Cannoical name of the zone",
				},
				"parent_zone_dn": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "DN of the parent zone",
				},
				"unix": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "The zone role is for Unix",
				},
				"windows": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "The zone role is for Windows",
				},
			},
		},
	}
}

func getWorkflowApproversSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Workflow approvers",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"guid": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "ID of the principal",
				},
				"name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "User or role name of the approver",
				},
				"type": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Type of principal",
					ValidateFunc: validation.StringInSlice([]string{"User", "Role", "Manager"}, false),
				},
				"no_manager_action": {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  "Type of principal",
					ValidateFunc: validation.StringInSlice([]string{"approve", "deny", "useBackup"}, false),
				},
				"backup_approver": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"guid": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "ID of the principal",
							},
							"name": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "User or role name of the approver",
							},
							"type": {
								Type:         schema.TypeString,
								Optional:     true,
								Description:  "Type of principal",
								ValidateFunc: validation.StringInSlice([]string{"User", "Role"}, false),
							},
						},
					},
				},
				"options_selector": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	}
}

func expandZoneRoles(v interface{}) []vault.ZoneRole {
	m := v.(*schema.Set).List()
	var zoneroles []vault.ZoneRole
	for _, lrv := range m {
		zonerole := vault.ZoneRole{}
		zonerole.Name = lrv.(map[string]interface{})["name"].(string)
		zonerole.ZoneDescription = lrv.(map[string]interface{})["zone_description"].(string)
		zonerole.ZoneDn = lrv.(map[string]interface{})["zone_dn"].(string)
		zonerole.Description = lrv.(map[string]interface{})["description"].(string)
		zonerole.ZoneCanonicalName = lrv.(map[string]interface{})["zone_canonical_name"].(string)
		zonerole.ParentZoneDn = lrv.(map[string]interface{})["parent_zone_dn"].(string)
		zonerole.Unix = lrv.(map[string]interface{})["unix"].(bool)
		zonerole.Windows = lrv.(map[string]interface{})["windows"].(bool)
		zoneroles = append(zoneroles, zonerole)
	}

	return zoneroles
}

func expandWorkflowApprovers(v []interface{}) []vault.WorkflowApprover {
	var approvers []vault.WorkflowApprover

	for _, lrv := range v {
		approver := vault.WorkflowApprover{}
		approver.Guid = lrv.(map[string]interface{})["guid"].(string)
		approver.Name = lrv.(map[string]interface{})["name"].(string)
		approver.Type = lrv.(map[string]interface{})["type"].(string)
		approver.NoManagerAction = lrv.(map[string]interface{})["no_manager_action"].(string)
		something := lrv.(map[string]interface{})["backup_approver"].([]interface{})
		if len(something) > 0 && something[0] != nil {
			d := something[0].(map[string]interface{})
			approver.BackupApprover = &vault.BackupApprover{
				Guid: d["guid"].(string),
				Name: d["name"].(string),
				Type: d["type"].(string),
			}
		}
		approver.OptionsSelector = lrv.(map[string]interface{})["options_selector"].(bool)
		approvers = append(approvers, approver)
	}

	return approvers
}

func getWorkflowDefaultOptionsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"grant_minute": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
}

func expandWorkflowDefaultOptions(v interface{}) *vault.WorkflowDefaultOptions {
	options := v.([]interface{})
	if len(options) > 0 && options[0] != nil {
		d := options[0].(map[string]interface{})
		data := &vault.WorkflowDefaultOptions{
			GrantMin: d["grant_minute"].(int),
		}
		return data
	}
	return nil
}

func convertWorkflowSchema(v string) (interface{}, error) {
	str := strings.Replace(v, "\\", "", -1)
	var wfapprovers vault.ProxyWorkflowApprover
	err := json.Unmarshal([]byte(str), &wfapprovers)
	if err != nil {
		logger.ErrorTracef(err.Error())
		return nil, fmt.Errorf("failed to unmarshal ProxyWorkflowApprover: %v", err)
	}
	approvers, err := vault.GenerateSchemaMap(wfapprovers)
	if err != nil {
		return nil, err
	}
	return processBackupApproverSchema(approvers["proxy_approver"]), nil
}

// processBackupApproverSchema converts "backup_approver" from map to list of map so that it conforms to schema
func processBackupApproverSchema(v interface{}) []interface{} {
	approvers := v.([]interface{})
	for i, v := range approvers {
		backupapprover := v.(map[string]interface{})["backup_approver"]
		if backupapprover != nil {
			// backup_approver is schema.TypeList but BackupApprover struct is not, so convert it
			v.(map[string]interface{})["backup_approver"] = []interface{}{backupapprover}
			approvers[i] = v
		}
	}
	return approvers
}

func convertZoneRoleSchema(v string) (interface{}, error) {
	str := strings.Replace(v, "\\", "", -1)
	var zoneroles vault.ProxyZoneRole
	err := json.Unmarshal([]byte(str), &zoneroles)
	if err != nil {
		logger.ErrorTracef(err.Error())
		return nil, fmt.Errorf("failed to unmarshal ProxyZoneRole: %v", err)
	}
	zroles, err := vault.GenerateSchemaMap(zoneroles)
	if err != nil {
		return nil, err
	}
	return processBackupApproverSchema(zroles["proxy_zonerole"]), nil
}
