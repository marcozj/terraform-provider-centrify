package centrify

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/biter777/countries"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	vault "github.com/marcozj/golang-sdk/platform"
)

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

//Validate the incoming set only contains values from the specified set
func validateSetValues(valid *schema.Set) schema.SchemaValidateFunc {
	return func(value interface{}, field string) (ws []string, errors []error) {
		if valid.Intersection(value.(*schema.Set)).Len() != value.(*schema.Set).Len() {
			errors = append(errors, fmt.Errorf("%q can only contain %v", field, value.(*schema.Set).List()))
		}
		return
	}
}

// hashStringCaseInsensitive hashes strings in a case insensitive manner.
// If you want a Set of strings and are case inensitive, this is the SchemaSetFunc you want.
func hashStringCaseInsensitive(v interface{}) int {
	return hashcode.String(strings.ToLower(v.(string)))
}

func flattenCaseInsensitiveStringSet(list []*string) *schema.Set {
	return schema.NewSet(hashStringCaseInsensitive, flattenStringList(list))
}

// Takes list of pointers to strings. Expand to an array
// of raw strings and returns a []interface{}
// to keep compatibility w/ schema.NewSetschema.NewSet
func flattenStringList(list []*string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, *v)
	}
	return vs
}

// StringSlice converts []string to []*string
func StringSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

func isEmptyValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func validateChallengeRules(input *vault.ChallengeRules) error {
	if input != nil && input.Rules != nil {
		for _, rule := range input.Rules {
			for _, v := range rule.ChallengeCondition {
				// Validate Filter and Condition pair
				switch v.Filter {
				case "IpAddress":
					if v.Condition != "OpInCorpIpRange" && v.Condition != "OpNotInCorpIpRange" {
						return fmt.Errorf("In %+v: IpAddress must have condition: OpInCorpIpRange or OpNotInCorpIpRange", v)
					}
				case "IdentityCookie":
					if v.Condition != "OpExists" && v.Condition != "OpNotExists" {
						return fmt.Errorf("In %+v: IdentityCookie must have condition: OpExists or OpNotExists", v)
					}
				case "DayOfWeek":
					if v.Condition != "OpIsDayOfWeek" {
						return fmt.Errorf("In %+v: DayOfWeek must have condition: OpIsDayOfWeek", v)
					}
				case "Date":
					if v.Condition != "OpLessThan" && v.Condition != "OpGreaterThan" {
						return fmt.Errorf("In %+v: Date must have condition: OpLessThan or OpGreaterThan", v)
					}
				case "DateRange":
					if v.Condition != "OpBetween" {
						return fmt.Errorf("In %+v: DateRange must have condition: OpBetween", v)
					}
				case "Time":
					if v.Condition != "OpBetween" {
						return fmt.Errorf("In %+v: Time must have condition: OpBetween", v)
					}
				case "DeviceOs":
					if v.Condition != "OpEqual" && v.Condition != "OpNotEqual" {
						return fmt.Errorf("In %+v: DeviceOs must have condition: OpEqual or OpNotEqual", v)
					}
					// Validate device value
					devices := []string{"iOS", "Android", "WindowsMobile", "Mac", "Windows", "Linux"}
					if !contains(devices, v.Value) {
						return fmt.Errorf("In %+v: DeviceOs must have value: %+v", v, devices)
					}
				case "Browser":
					if v.Condition != "OpEqual" && v.Condition != "OpNotEqual" {
						return fmt.Errorf("In %+v: Browser must have condition: OpEqual or OpNotEqual", v)
					}
					// Validate browser value
					browser := []string{"Other", "Chrome", "Firefox", "IE", "Safari", "MicrosoftEdge"}
					if !contains(browser, v.Value) {
						return fmt.Errorf("In %+v: Browser must have value: %+v", v, browser)
					}
				case "CountryCode":
					if v.Condition != "OpEqual" && v.Condition != "OpNotEqual" {
						return fmt.Errorf("In %+v: CountryCode must have condition: OpEqual or OpNotEqual", v)
					}
					if len(v.Value) != 2 {
						return fmt.Errorf("In %+v: CountryCode must be valid 2 digit country code", v)
					}
					country := countries.ByName(v.Value)
					if country == countries.Unknown {
						return fmt.Errorf("In %+v: %s is not a valid country code", v, v.Value)
					}
				case "Zso":
					if v.Condition != "OpIs" && v.Condition != "OpIsNot" {
						return fmt.Errorf("In %+v: Zso must have condition: OpIs or OpIsNot", v)
					}
				} // end of switch

			}
		}
	}
	return nil
}
