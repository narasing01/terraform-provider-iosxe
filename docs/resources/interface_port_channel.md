---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_interface_port_channel Resource - terraform-provider-iosxe"
subcategory: "Interface"
description: |-
  This resource can manage the Interface Port Channel configuration.
---

# iosxe_interface_port_channel (Resource)

This resource can manage the Interface Port Channel configuration.

## Example Usage

```terraform
resource "iosxe_interface_port_channel" "example" {
  name                           = 10
  description                    = "My Interface Description"
  shutdown                       = false
  vrf_forwarding                 = "VRF1"
  ipv4_address                   = "192.0.2.1"
  ipv4_address_mask              = "255.255.255.0"
  ip_access_group_in             = "1"
  ip_access_group_in_enable      = true
  ip_access_group_out            = "1"
  ip_access_group_out_enable     = true
  ip_dhcp_relay_source_interface = "Loopback100"
  helper_addresses = [
    {
      address = "10.10.10.10"
      global  = false
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (Number) - Range: `1`-`512`

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `description` (String) Interface specific description
- `device` (String) A device name from the provider configuration.
- `helper_addresses` (Attributes List) Specify a destination address for UDP broadcasts (see [below for nested schema](#nestedatt--helper_addresses))
- `ip_access_group_in` (String)
- `ip_access_group_in_enable` (Boolean) inbound packets
- `ip_access_group_out` (String)
- `ip_access_group_out_enable` (Boolean) outbound packets
- `ip_dhcp_relay_source_interface` (String) Set source interface for relayed messages
- `ip_proxy_arp` (Boolean) Enable proxy ARP
- `ip_redirects` (Boolean) Enable sending ICMP Redirect messages
- `ipv4_address` (String)
- `ipv4_address_mask` (String)
- `shutdown` (Boolean) Shutdown the selected interface
- `switchport` (Boolean)
- `unreachables` (Boolean) Enable sending ICMP Unreachable messages
- `vrf_forwarding` (String) Configure forwarding table

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--helper_addresses"></a>
### Nested Schema for `helper_addresses`

Required:

- `address` (String)

Optional:

- `global` (Boolean) Helper-address is global
- `vrf` (String) VRF name for helper-address (if different from interface VRF)

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_interface_port_channel.example "Cisco-IOS-XE-native:native/interface/Port-channel=10"
```