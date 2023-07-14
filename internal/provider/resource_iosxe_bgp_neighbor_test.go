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

func TestAccIosxeBGPNeighbor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_neighbor.test", "ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_neighbor.test", "remote_as", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_neighbor.test", "description", "BGP Neighbor Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_neighbor.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxe_bgp_neighbor.test", "update_source_loopback", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeBGPNeighborPrerequisitesConfig + testAccIosxeBGPNeighborConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxe_bgp_neighbor.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000/neighbor=3.3.3.3",
			},
		},
	})
}

const testAccIosxeBGPNeighborPrerequisitesConfig = `
resource "iosxe_restconf" "PreReq0" {
	path = "Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=65000"
	attributes = {
		"id" = "65000"
	}
}

resource "iosxe_restconf" "PreReq1" {
	path = "Cisco-IOS-XE-native:native/interface/Loopback=100"
	attributes = {
		"name" = "100"
	}
}

`

func testAccIosxeBGPNeighborConfig_minimum() string {
	config := `resource "iosxe_bgp_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxeBGPNeighborConfig_all() string {
	config := `resource "iosxe_bgp_neighbor" "test" {` + "\n"
	config += `	asn = "65000"` + "\n"
	config += `	ip = "3.3.3.3"` + "\n"
	config += `	remote_as = "65000"` + "\n"
	config += `	description = "BGP Neighbor Description"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	update_source_loopback = "100"` + "\n"
	config += `	depends_on = [iosxe_restconf.PreReq0, iosxe_restconf.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}