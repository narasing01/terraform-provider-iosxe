---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_bgp_address_family_ipv6 Data Source - terraform-provider-iosxe"
subcategory: "BGP"
description: |-
  This data source can read the BGP Address Family IPv6 configuration.
---

# iosxe_bgp_address_family_ipv6 (Data Source)

This data source can read the BGP Address Family IPv6 configuration.

## Example Usage

```terraform
data "iosxe_bgp_address_family_ipv6" "example" {
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
- `ipv6_unicast_networks` (Attributes List) Specify a network to announce via BGP (see [below for nested schema](#nestedatt--ipv6_unicast_networks))

<a id="nestedatt--ipv6_unicast_networks"></a>
### Nested Schema for `ipv6_unicast_networks`

Read-Only:

- `backdoor` (Boolean) Specify a BGP backdoor route
- `network` (String)
- `route_map` (String) Route-map to modify the attributes
