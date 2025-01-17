## 0.4.0

- Fix issue when using `tunnel_destination_ipv4` or `tunnel_mode_ipsec_ipv4` attributes of `iosxe_interface_tunnel` resource
- Add `iosxe_static_route_vrf` resource and data source
- Make `lists.key` attribute of `iosxe_restconf` resource mandatory
- Fix issue with nested `lists.items` attributes of `iosxe_restconf` resource
- Add `iosxe_radius` resource and data source
- Add `iosxe_bfd_template_single_hop` resource and data source
- Add `iosxe_cdp` resource and data source
- Add `cluster_id` ,`fall_over` ,`disable_connected_check`, `local_as`, `log_neighbor_changes`, `password`, `timers`, `ttl_security` attributes to `iosxe_bgp_ipv4_unicast_vrf_neighbor` resource and data source
- Add `iosxe_tacacs_server` resource and data source
- Add `iosxe_bfd` resource and data source
- Add BFD attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_port_channel_subinterface`, `iosxe_interface_tunnel` and `iosxe_interface_vlan` resources and data sources
- BREAKING CHANGE: Rename `unreachables` attribute to `ip_unreachables` of `iosxe_interface_ethernet`, `iosxe_interface_port_channel`, `iosxe_interface_port_channel_subinterface`, `iosxe_interface_tunnel`, `iosxe_interface_loopback` and `iosxe_interface_vlan` resources and data sources
- BREAKING CHANGE: Rename `multicast_routing` attribute to `ip_multicast_routing` of `iosxe_system` resource and data source
- BREAKING CHANGE: Rename `multicast_routing_distributed` attribute to `ip_multicast_routing_distributed` of `iosxe_system` resource and data source
- Add `ipv6` attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` `iosxe_interface_vlan` `iosxe_interface_loopback` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `iosxe_bfd_template_multi_hop` resource and data source
- Add `disable_connected_check`, `fall_over`, `local_as`, `log_neighbor_changes`, `password`, `timers` and `ttl_security` attributes to `iosxe_bgp_neighbor` resource and data source
- Add `ttl_security` and `process_ids` attributes to `iosxe_interface_ospf` resource and data source
- BREAKING CHANGE: Remove `iosxe_interface_ospf_process` resource and data source, functionality moved to `iosxe_interface_ospf` resource and data source
- Add `iosxe_arp` resource and data source
- Add `iosxe_class_map` resource and data source
- BREAKING CHANGE: Rename `ipv6_address_prefix_list` attribute to `ipv6_addresses` of `iosxe_interface_tunnel` resource and data source
- Add `snooping_information_option_format_remote_id_hostname` attribute to `iosxe_dhcp` resource and data source
- Add `iosxe_dot1x` resource and data source
- Add `arp_timeout`, `spanning_tree_link_type` and `spanning_tree_portfast_trunk` attribute to `iosxe_interface_ethernet` resource and data source
- Add `message_digest_keys` attribute to `iosxe_interface_ospf` resource and data source
- Add `areas` and `passive_interface_default` attributes to `iosxe_ospf_vrf` resource and data source
- Add `iosxe_policy_map` resource and data source
- Add `iosxe_policy_map_event` resource and data source
- Add `accounting_port`, `pac_key` and `automate_tester` attributes to `iosxe_radius` resource and data source
- Add `arp_timeout` attribute to `iosxe_interface_loopback` resource and data source
- Add `ip_arp_inspection` and `ip_dhcp_snooping` attributes to `iosxe_interface_ethernet` resource and data source
- Add `arp_timeout` and `ip_arp_inspection` attributes to `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `areas` attribute to `iosxe_ospf` resource and data source
- Add `iosxe_udld` resource and data source
- Add `iosxe_vtp` resource and data source
- BREAKING CHANGE: Rename `neighbor` attribute to `neighbors` of `iosxe_ospf` resource and data source
- BREAKING CHANGE: Rename `network` attribute to `networks` of `iosxe_ospf` resource and data source
- BREAKING CHANGE: Rename `summary_address` attribute to `summary_addresses` of `iosxe_ospf` resource and data source
- Add `ipv4_unicast_networks_mask` and `ipv4_unicast_networks` attribute to `iosxe_bgp_address_family_ipv4` and `iosxe_bgp_address_family_ipv4_vrf` resources and data sources
- Add `ipv6_unicast_networks` attribute to `iosxe_bgp_address_family_ipv6` and `iosxe_bgp_address_family_ipv6_vrf` resources and data sources

## 0.3.3

- Add `iosxe_aaa` resource and data source
- Add `iosxe_aaa_accounting` resource and data source
- Add `iosxe_aaa_authorization` resource and data source
- Add `iosxe_interface_mpls` resource and data source
- Add `iosxe_interface_ospfv3` resource and data source
- Add `iosxe_interface_tunnel` resource and data source
- Add `iosxe_crypto_ipsec_transform_set` resource and data source
- Add `iosxe_aaa_authentication` resource and data source
- Add `iosxe_crypto_ikev2` resource and data source
- Add `iosxe_crypto_ikev2_proposal` resource and data source
- Add `iosxe_crypto_ipsec_profile` resource and data source
- Add `iosxe_radius_server` resource and data source

## 0.3.2

- Add `auto_qos` attributes to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `spanning_tree_guard` attribute to `iosxe_interface_ethernet` and `iosxe_interface_port_channel` resources and data sources
- Add `trust_device` attribute to `iosxe_interface_ethernet`, `iosxe_interface_port_channel` and `iosxe_interface_port_channel_subinterface` resources and data sources
- Add `trunk_allowed_vlans_none` attribute to `iosxe_interface_switchport` and `iosxe_template` resources and data sources
- Add `trunk_allowed_vlans_all` attribute to `iosxe_template` resource and data source
- Add `ebgp_multihop` attributes to `iosxe_bgp_neighbor` and `iosxe_bgp_ipv4_unicast_vrf_neighbor` resources and data sources

## 0.3.1

- Fix issue with deletion of servers and peers of `iosxe_ntp` resource

## 0.3.0

- BREAKING CHANGE: Completely revamped the provider based on `github.com/netascode/terraform-provider-iosxe` codebase, replacing all existing resources and data sources
- BREAKING CHANGE: Remove `attributes` map of list items in `iosxe_restconf` resource

## 0.1.1

## 0.1.0

- Initial release
