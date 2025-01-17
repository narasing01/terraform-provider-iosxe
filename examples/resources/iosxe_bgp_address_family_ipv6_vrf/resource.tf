resource "iosxe_bgp_address_family_ipv6_vrf" "example" {
  asn     = "65000"
  af_name = "unicast"
  vrfs = [
    {
      name                   = "VRF1"
      advertise_l2vpn_evpn   = true
      redistribute_connected = true
      redistribute_static    = true
      ipv6_unicast_networks = [
        {
          network   = "2001:1234::0/64"
          route_map = "RM1"
          backdoor  = true
          evpn      = false
        }
      ]
    }
  ]
}
