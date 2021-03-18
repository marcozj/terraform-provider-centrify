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
- `description` - description property.
- `key_type` - key_type property.
- `ssh_key` - Retrieved ssh key.
