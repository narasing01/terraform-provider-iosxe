---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_cdp Resource - terraform-provider-iosxe"
subcategory: "System"
description: |-
  This resource can manage the CDP configuration.
---

# iosxe_cdp (Resource)

This resource can manage the CDP configuration.

## Example Usage

```terraform
resource "iosxe_cdp" "example" {
  holdtime        = 15
  timer           = 5
  run             = true
  filter_tlv_list = "TLIST"
  tlv_lists = [
    {
      name            = "TLIST"
      vtp_mgmt_domain = true
      cos             = true
      duplex          = true
      trust           = true
      version         = true
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.
- `filter_tlv_list` (String) Apply tlv-list globally
- `holdtime` (Number) Specify the holdtime (in sec) to be sent in packets
  - Range: `10`-`255`
- `run` (Boolean) Enable CDP
- `timer` (Number) Specify the rate at which CDP packets are sent (in sec)
  - Range: `5`-`254`
- `tlv_lists` (Attributes List) Configure tlv-list (see [below for nested schema](#nestedatt--tlv_lists))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--tlv_lists"></a>
### Nested Schema for `tlv_lists`

Required:

- `name` (String) Tlv-list

Optional:

- `cos` (Boolean) Select cos TLV
- `duplex` (Boolean) Select duplex TLV
- `trust` (Boolean) Select trust TLV
- `version` (Boolean) Select version TLV
- `vtp_mgmt_domain` (Boolean) Select vtp mgmt domain TLV

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_cdp.example "Cisco-IOS-XE-native:native/cdp"
```