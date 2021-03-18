---
subcategory: "Access"
---

# centrifyvault_policy (Data Source)

This data source gets information of policy.

## Example Usage

```terraform
data "centrifyvault_policy" "Default_Policy" {
    name = "Default Policy"
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/blob/main/examples/centrifyvault_policy/policyorder.tf)

## Search Attributes

### Required

- `name` - (String) The name of the policy.

## Attributes Reference

- `id` - id of the policy.
- `description` - description property.
- `link_type` - link_type property.
