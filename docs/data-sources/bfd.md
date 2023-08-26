---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_bfd Data Source - terraform-provider-iosxe"
subcategory: "BFD"
description: |-
  This data source can read the BFD configuration.
---

# iosxe_bfd (Data Source)

This data source can read the BFD configuration.

## Example Usage

```terraform
data "iosxe_bfd" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `ipv4_both_vrfs` (Attributes List) IPv4 Address Family with vrf (see [below for nested schema](#nestedatt--ipv4_both_vrfs))
- `ipv4_with_dst_vrfs` (Attributes List) IPv4 Address Family with vrf (see [below for nested schema](#nestedatt--ipv4_with_dst_vrfs))
- `ipv4_with_src_vrfs` (Attributes List) IPv4 Address Family with vrf (see [below for nested schema](#nestedatt--ipv4_with_src_vrfs))
- `ipv4_without_vrfs` (Attributes List) IPv4 Address Family with vrf (see [below for nested schema](#nestedatt--ipv4_without_vrfs))
- `ipv6_with_both_vrfs` (Attributes List) IPv6 Address Family with vrf (see [below for nested schema](#nestedatt--ipv6_with_both_vrfs))
- `ipv6_with_dst_vrfs` (Attributes List) IPv6 Address Family with vrf (see [below for nested schema](#nestedatt--ipv6_with_dst_vrfs))
- `ipv6_with_src_vrfs` (Attributes List) IPv6 Address Family with vrf (see [below for nested schema](#nestedatt--ipv6_with_src_vrfs))
- `ipv6_without_vrfs` (Attributes List) IPv6 Address Family with vrf (see [below for nested schema](#nestedatt--ipv6_without_vrfs))
- `slow_timers` (Number) Value in ms to use for slow timers

<a id="nestedatt--ipv4_both_vrfs"></a>
### Nested Schema for `ipv4_both_vrfs`

Read-Only:

- `dest_ip` (String) Destination IP prefix/len
- `dst_vrf` (String) Destination VRF instance name
- `src_ip` (String) Source IP prefix/len
- `src_vrf` (String) source VRF instance name
- `template_name` (String) BFD template name


<a id="nestedatt--ipv4_with_dst_vrfs"></a>
### Nested Schema for `ipv4_with_dst_vrfs`

Read-Only:

- `dest_ip` (String) Destination IP prefix/len
- `dst_vrf` (String) Destination VRF instance name
- `src_ip` (String) Source IP prefix/len
- `template_name` (String) BFD template name


<a id="nestedatt--ipv4_with_src_vrfs"></a>
### Nested Schema for `ipv4_with_src_vrfs`

Read-Only:

- `dest_ip` (String) Destination IP prefix/len
- `src_ip` (String) Source IP prefix/len
- `src_vrf` (String) source VRF instance name
- `template_name` (String) BFD template name


<a id="nestedatt--ipv4_without_vrfs"></a>
### Nested Schema for `ipv4_without_vrfs`

Read-Only:

- `dest_ip` (String) Destination IP prefix/len
- `src_ip` (String) Source IP prefix/len
- `template_name` (String) BFD template name


<a id="nestedatt--ipv6_with_both_vrfs"></a>
### Nested Schema for `ipv6_with_both_vrfs`

Read-Only:

- `dest_ipv6` (String) Destination IPv6 prefix/len
- `dst_vrf` (String) Destination VRF instance name
- `src_ipv6` (String) Source IPv6 prefix/len
- `src_vrf` (String) source VRF instance name
- `template_name` (String) BFD template name


<a id="nestedatt--ipv6_with_dst_vrfs"></a>
### Nested Schema for `ipv6_with_dst_vrfs`

Read-Only:

- `dest_ipv6` (String) Destination IPv6 prefix/len
- `dst_vrf` (String) Destination VRF instance name
- `src_ipv6` (String) Source IPv6 prefix/len
- `template_name` (String) BFD template name


<a id="nestedatt--ipv6_with_src_vrfs"></a>
### Nested Schema for `ipv6_with_src_vrfs`

Read-Only:

- `dest_ipv6` (String) Destination IPv6 prefix/len
- `src_ipv6` (String) Source IPv6 prefix/len
- `src_vrf` (String) source VRF instance name
- `template_name` (String) BFD template name


<a id="nestedatt--ipv6_without_vrfs"></a>
### Nested Schema for `ipv6_without_vrfs`

Read-Only:

- `dest_ipv6` (String) Destination IPv6 prefix/len
- `src_ipv6` (String) Source IPv6 prefix/len
- `template_name` (String) BFD template name