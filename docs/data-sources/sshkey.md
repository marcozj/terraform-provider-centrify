---
subcategory: "Resources"
---

# centrifyvault_sshkey (Data Source)

This data source gets information of ssh key.

## Example Usage

```terraform
data "centrifyvault_sshkey" "testkey" {
    name = "testkey"
    key_pair_type = "PrivateKey"
    passphrase = ""
    key_format = "PEM"
    checkout = true
}

output "testkey_sshkey" {
  value = data.centrifyvault_sshkey.testkey.ssh_key
}
```

More examples can be found [here](https://github.com/marcozj/terraform-provider-centrifyvault/tree/main/examples/centrifyvault_sshkey)

## Search Attributes

### Required

- `name` - (String) Name of the ssh key.
- `key_pair_type` - (String) Which key to retrieve from the pair. Can be set to `PublicKey`, `PrivateKey`, or `PPK`.

### Optional

- `passphrase` - (String, Sensitive) Passphrase to use for decrypting the PrivateKey.
- `key_format` - (String) KeyFormat to retrieve the key in. Default is `PEM`. For PublicKey, it can be set to `PEM` or `OpenSSH`.
- `checkout` - (Boolean) Whether to retrieve SSH Key. Default is `false`. If `true`, `ssh_key` will be populated.

## Attributes Reference

- `id` - id of the ssh key.
- `name` - (String) Name of the ssh key.
- `description` - (String) Description of the SSH Key.
- `key_pair_type` - (String) Which key to retrieve from the pair.
- `passphrase` - (String, Sensitive) Passphrase to use for decrypting the PrivateKey.
- `key_format` - (String) KeyFormat to retrieve the key in.
- `key_type` - (String) Key type.
- `challenge_rule` - (Block List) Authentication rules. Refer to [challenge_rule](./attribute_challengerule.md) attribute for details.
- `default_profile_id` - (String) Default SSH Key Challenge Profile ID (used if no conditions matched).
- `ssh_key` - (String, Sensitive) SSH private key. This attribute value is available only if `checkout` is set to `true`.
- `passphrase` - (String, Sensitive) Passphrase to use for encrypting the PrivateKey.
- `sets` (Set of String) List of Set IDs the resource belongs to. Refer to [sets](./attribute_sets.md) attribute for details.
