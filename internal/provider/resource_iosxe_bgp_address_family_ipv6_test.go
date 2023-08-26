// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxeBGPAddressFamilyIPv6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv6.test", "af_name", "unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv6.test", "ipv6_unicast_networks.0.network", "2001:1234::0/64"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv6.test", "ipv6_unicast_networks.0.route_map", "RM1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_address_family_ipv6.test", "ipv6_unicast_networks.0.backdoor", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPAddressFamilyIPv6PrerequisitesConfig + testAccIosxeBGPAddressFamilyIPv6Config_minimum(),
			},
			{
				Config: testAccIosxeBGPAddressFamilyIPv6PrerequisitesConfig + testAccIosxeBGPAddressFamilyIPv6Config_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_bgp_address_family_ipv6.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/address-family/no-vrf/ipv6=unicast",
			},
		},
	})
}

const testAccIosxeBGPAddressFamilyIPv6PrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/ipv6"
	attributes = {
		"unicast-routing" = ""
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

`

func testAccIosxeBGPAddressFamilyIPv6Config_minimum() string {
	config := `resource "iosxe_bgp_address_family_ipv6" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	af_name = "unicast"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPAddressFamilyIPv6Config_all() string {
	config := `resource "iosxe_bgp_address_family_ipv6" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	af_name = "unicast"` + "\n"
	config += `	ipv6_unicast_networks = [{` + "\n"
	config += `		network = "2001:1234::0/64"` + "\n"
	config += `		route_map = "RM1"` + "\n"
	config += `		backdoor = true` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
