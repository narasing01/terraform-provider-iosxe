---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_bgp_address_family_ipv4_vrf Data Source - terraform-provider-iosxe"
subcategory: "BGP"
description: |-
  This data source can read the BGP Address Family IPv4 VRF configuration.
---

# iosxe_bgp_address_family_ipv4_vrf (Data Source)

This data source can read the BGP Address Family IPv4 VRF configuration.

## Example Usage

```terraform
data "iosxe_bgp_address_family_ipv4_vrf" "example" {
  asn     = "65000"
  af_name = "unicast"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `af_name` (String)
- `asn` (String)

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `vrfs` (Attributes List) (see [below for nested schema](#nestedatt--vrfs))

<a id="nestedatt--vrfs"></a>
### Nested Schema for `vrfs`

Read-Only:

- `advertise_l2vpn_evpn` (Boolean) Advertise/export prefixes to l2vpn evpn table
- `ipv4_unicast_networks` (Attributes List) Specify a network to announce via BGP (see [below for nested schema](#nestedatt--vrfs--ipv4_unicast_networks))
- `ipv4_unicast_networks_mask` (Attributes List) Specify a network to announce via BGP (see [below for nested schema](#nestedatt--vrfs--ipv4_unicast_networks_mask))
- `name` (String)
- `redistribute_connected` (Boolean) Connected
- `redistribute_static` (Boolean) Static routes

<a id="nestedatt--vrfs--ipv4_unicast_networks"></a>
### Nested Schema for `vrfs.ipv4_unicast_networks`

Read-Only:

- `backdoor` (Boolean) Specify a BGP backdoor route
- `evpn` (Boolean) Advertise or export to EVPN address-family
- `network` (String)
- `route_map` (String) Route-map to modify the attributes


<a id="nestedatt--vrfs--ipv4_unicast_networks_mask"></a>
### Nested Schema for `vrfs.ipv4_unicast_networks_mask`

Read-Only:

- `backdoor` (Boolean) Specify a BGP backdoor route
- `evpn` (Boolean) Advertise or export to EVPN address-family
- `mask` (String) Network mask
- `network` (String)
- `route_map` (String) Route-map to modify the attributes
