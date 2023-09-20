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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeVLANAccessMap(t *testing.T) {
	if os.Getenv("C9000V") == "" {
		t.Skip("skipping test, set environment variable C9000V")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_vlan_access_map.test", "match_ipv6_address.0", "ipv6_address1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_vlan_access_map.test", "match_ip_address.0", "ip_address1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_vlan_access_map.test", "action", "forward"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeVLANAccessMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeVLANAccessMapConfig() string {
	config := `resource "iosxe_vlan_access_map" "test" {` + "\n"
	config += `	name = "accessmap1"` + "\n"
	config += `	value = 1000` + "\n"
	config += `	match_ipv6_address = ["ipv6_address1"]` + "\n"
	config += `	match_ip_address = ["ip_address1"]` + "\n"
	config += `	action = "forward"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_vlan_access_map" "test" {
			name = "accessmap1"
			value = 1000
			depends_on = [iosxe_vlan_access_map.test]
		}
	`
	return config
}