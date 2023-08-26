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
	"context"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPAddressFamilyIPv6VRF struct {
	Device     types.String                  `tfsdk:"device"`
	Id         types.String                  `tfsdk:"id"`
	DeleteMode types.String                  `tfsdk:"delete_mode"`
	Asn        types.String                  `tfsdk:"asn"`
	AfName     types.String                  `tfsdk:"af_name"`
	Vrfs       []BGPAddressFamilyIPv6VRFVrfs `tfsdk:"vrfs"`
}

type BGPAddressFamilyIPv6VRFData struct {
	Device types.String                  `tfsdk:"device"`
	Id     types.String                  `tfsdk:"id"`
	Asn    types.String                  `tfsdk:"asn"`
	AfName types.String                  `tfsdk:"af_name"`
	Vrfs   []BGPAddressFamilyIPv6VRFVrfs `tfsdk:"vrfs"`
}
type BGPAddressFamilyIPv6VRFVrfs struct {
	Name                  types.String                                     `tfsdk:"name"`
	AdvertiseL2vpnEvpn    types.Bool                                       `tfsdk:"advertise_l2vpn_evpn"`
	RedistributeConnected types.Bool                                       `tfsdk:"redistribute_connected"`
	RedistributeStatic    types.Bool                                       `tfsdk:"redistribute_static"`
	Ipv6UnicastNetworks   []BGPAddressFamilyIPv6VRFVrfsIpv6UnicastNetworks `tfsdk:"ipv6_unicast_networks"`
}
type BGPAddressFamilyIPv6VRFVrfsIpv6UnicastNetworks struct {
	Network  types.String `tfsdk:"network"`
	RouteMap types.String `tfsdk:"route_map"`
	Backdoor types.Bool   `tfsdk:"backdoor"`
	Evpn     types.Bool   `tfsdk:"evpn"`
}

func (data BGPAddressFamilyIPv6VRF) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv6=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

func (data BGPAddressFamilyIPv6VRFData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv6=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyIPv6VRF) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyIPv6VRF) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.ValueString())
	}
	if len(data.Vrfs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf", []interface{}{})
		for index, item := range data.Vrfs {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.AdvertiseL2vpnEvpn.IsNull() && !item.AdvertiseL2vpnEvpn.IsUnknown() {
				if item.AdvertiseL2vpnEvpn.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.advertise.l2vpn.evpn", map[string]string{})
				}
			}
			if !item.RedistributeConnected.IsNull() && !item.RedistributeConnected.IsUnknown() {
				if item.RedistributeConnected.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.redistribute-v6.connected", map[string]string{})
				}
			}
			if !item.RedistributeStatic.IsNull() && !item.RedistributeStatic.IsUnknown() {
				if item.RedistributeStatic.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.redistribute-v6.static", map[string]string{})
				}
			}
			if len(item.Ipv6UnicastNetworks) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.network", []interface{}{})
				for cindex, citem := range item.Ipv6UnicastNetworks {
					if !citem.Network.IsNull() && !citem.Network.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.network"+"."+strconv.Itoa(cindex)+"."+"number", citem.Network.ValueString())
					}
					if !citem.RouteMap.IsNull() && !citem.RouteMap.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.network"+"."+strconv.Itoa(cindex)+"."+"route-map", citem.RouteMap.ValueString())
					}
					if !citem.Backdoor.IsNull() && !citem.Backdoor.IsUnknown() {
						if citem.Backdoor.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.network"+"."+strconv.Itoa(cindex)+"."+"backdoor", map[string]string{})
						}
					}
					if !citem.Evpn.IsNull() && !citem.Evpn.IsUnknown() {
						if citem.Evpn.ValueBool() {
							body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"vrf"+"."+strconv.Itoa(index)+"."+"ipv6-unicast.network"+"."+strconv.Itoa(cindex)+"."+"evpn", map[string]string{})
						}
					}
				}
			}
		}
	}
	return body
}

func (data *BGPAddressFamilyIPv6VRF) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "af-name"); value.Exists() && !data.AfName.IsNull() {
		data.AfName = types.StringValue(value.String())
	} else {
		data.AfName = types.StringNull()
	}
	for i := range data.Vrfs {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Vrfs[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "vrf").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("name"); value.Exists() && !data.Vrfs[i].Name.IsNull() {
			data.Vrfs[i].Name = types.StringValue(value.String())
		} else {
			data.Vrfs[i].Name = types.StringNull()
		}
		if value := r.Get("ipv6-unicast.advertise.l2vpn.evpn"); !data.Vrfs[i].AdvertiseL2vpnEvpn.IsNull() {
			if value.Exists() {
				data.Vrfs[i].AdvertiseL2vpnEvpn = types.BoolValue(true)
			} else {
				data.Vrfs[i].AdvertiseL2vpnEvpn = types.BoolValue(false)
			}
		} else {
			data.Vrfs[i].AdvertiseL2vpnEvpn = types.BoolNull()
		}
		if value := r.Get("ipv6-unicast.redistribute-v6.connected"); !data.Vrfs[i].RedistributeConnected.IsNull() {
			if value.Exists() {
				data.Vrfs[i].RedistributeConnected = types.BoolValue(true)
			} else {
				data.Vrfs[i].RedistributeConnected = types.BoolValue(false)
			}
		} else {
			data.Vrfs[i].RedistributeConnected = types.BoolNull()
		}
		if value := r.Get("ipv6-unicast.redistribute-v6.static"); !data.Vrfs[i].RedistributeStatic.IsNull() {
			if value.Exists() {
				data.Vrfs[i].RedistributeStatic = types.BoolValue(true)
			} else {
				data.Vrfs[i].RedistributeStatic = types.BoolValue(false)
			}
		} else {
			data.Vrfs[i].RedistributeStatic = types.BoolNull()
		}
		for ci := range data.Vrfs[i].Ipv6UnicastNetworks {
			keys := [...]string{"number"}
			keyValues := [...]string{data.Vrfs[i].Ipv6UnicastNetworks[ci].Network.ValueString()}

			var cr gjson.Result
			r.Get("ipv6-unicast.network").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			if value := cr.Get("number"); value.Exists() && !data.Vrfs[i].Ipv6UnicastNetworks[ci].Network.IsNull() {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].Network = types.StringValue(value.String())
			} else {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].Network = types.StringNull()
			}
			if value := cr.Get("route-map"); value.Exists() && !data.Vrfs[i].Ipv6UnicastNetworks[ci].RouteMap.IsNull() {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].RouteMap = types.StringValue(value.String())
			} else {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].RouteMap = types.StringNull()
			}
			if value := cr.Get("backdoor"); !data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor.IsNull() {
				if value.Exists() {
					data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor = types.BoolValue(true)
				} else {
					data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor = types.BoolValue(false)
				}
			} else {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor = types.BoolNull()
			}
			if value := cr.Get("evpn"); !data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn.IsNull() {
				if value.Exists() {
					data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn = types.BoolValue(true)
				} else {
					data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn = types.BoolValue(false)
				}
			} else {
				data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn = types.BoolNull()
			}
		}
	}
}

func (data *BGPAddressFamilyIPv6VRFData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "vrf"); value.Exists() {
		data.Vrfs = make([]BGPAddressFamilyIPv6VRFVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv6VRFVrfs{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv6-unicast.advertise.l2vpn.evpn"); cValue.Exists() {
				item.AdvertiseL2vpnEvpn = types.BoolValue(true)
			} else {
				item.AdvertiseL2vpnEvpn = types.BoolValue(false)
			}
			if cValue := v.Get("ipv6-unicast.redistribute-v6.connected"); cValue.Exists() {
				item.RedistributeConnected = types.BoolValue(true)
			} else {
				item.RedistributeConnected = types.BoolValue(false)
			}
			if cValue := v.Get("ipv6-unicast.redistribute-v6.static"); cValue.Exists() {
				item.RedistributeStatic = types.BoolValue(true)
			} else {
				item.RedistributeStatic = types.BoolValue(false)
			}
			if cValue := v.Get("ipv6-unicast.network"); cValue.Exists() {
				item.Ipv6UnicastNetworks = make([]BGPAddressFamilyIPv6VRFVrfsIpv6UnicastNetworks, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := BGPAddressFamilyIPv6VRFVrfsIpv6UnicastNetworks{}
					if ccValue := cv.Get("number"); ccValue.Exists() {
						cItem.Network = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("route-map"); ccValue.Exists() {
						cItem.RouteMap = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("backdoor"); ccValue.Exists() {
						cItem.Backdoor = types.BoolValue(true)
					} else {
						cItem.Backdoor = types.BoolValue(false)
					}
					if ccValue := cv.Get("evpn"); ccValue.Exists() {
						cItem.Evpn = types.BoolValue(true)
					} else {
						cItem.Evpn = types.BoolValue(false)
					}
					item.Ipv6UnicastNetworks = append(item.Ipv6UnicastNetworks, cItem)
					return true
				})
			}
			data.Vrfs = append(data.Vrfs, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv6VRF) getDeletedListItems(ctx context.Context, state BGPAddressFamilyIPv6VRF) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Vrfs {
		stateKeyValues := [...]string{state.Vrfs[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Vrfs[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vrfs {
			found = true
			if state.Vrfs[i].Name.ValueString() != data.Vrfs[j].Name.ValueString() {
				found = false
			}
			if found {
				for ci := range state.Vrfs[i].Ipv6UnicastNetworks {
					cstateKeyValues := [...]string{state.Vrfs[i].Ipv6UnicastNetworks[ci].Network.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.Vrfs[i].Ipv6UnicastNetworks[ci].Network.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.Vrfs[j].Ipv6UnicastNetworks {
						found = true
						if state.Vrfs[i].Ipv6UnicastNetworks[ci].Network.ValueString() != data.Vrfs[j].Ipv6UnicastNetworks[cj].Network.ValueString() {
							found = false
						}
						if found {
							break
						}
					}
					if !found {
						deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/network=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vrf=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *BGPAddressFamilyIPv6VRF) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.Vrfs {
		keyValues := [...]string{data.Vrfs[i].Name.ValueString()}
		if !data.Vrfs[i].AdvertiseL2vpnEvpn.IsNull() && !data.Vrfs[i].AdvertiseL2vpnEvpn.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/advertise/l2vpn/evpn", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Vrfs[i].RedistributeConnected.IsNull() && !data.Vrfs[i].RedistributeConnected.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/redistribute-v6/connected", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Vrfs[i].RedistributeStatic.IsNull() && !data.Vrfs[i].RedistributeStatic.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/redistribute-v6/static", data.getPath(), strings.Join(keyValues[:], ",")))
		}

		for ci := range data.Vrfs[i].Ipv6UnicastNetworks {
			ckeyValues := [...]string{data.Vrfs[i].Ipv6UnicastNetworks[ci].Network.ValueString()}
			if !data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor.IsNull() && !data.Vrfs[i].Ipv6UnicastNetworks[ci].Backdoor.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/network=%v/backdoor", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
			if !data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn.IsNull() && !data.Vrfs[i].Ipv6UnicastNetworks[ci].Evpn.ValueBool() {
				emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/vrf=%v/ipv6-unicast/network=%v/evpn", data.getPath(), strings.Join(keyValues[:], ","), strings.Join(ckeyValues[:], ",")))
			}
		}
	}
	return emptyLeafsDelete
}

func (data *BGPAddressFamilyIPv6VRF) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Vrfs {
		keyValues := [...]string{data.Vrfs[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/vrf=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
