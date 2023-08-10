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

type BGPIPv4UnicastVRFNeighbor struct {
	Device                              types.String                         `tfsdk:"device"`
	Id                                  types.String                         `tfsdk:"id"`
	DeleteMode                          types.String                         `tfsdk:"delete_mode"`
	Asn                                 types.String                         `tfsdk:"asn"`
	Vrf                                 types.String                         `tfsdk:"vrf"`
	Ip                                  types.String                         `tfsdk:"ip"`
	RemoteAs                            types.String                         `tfsdk:"remote_as"`
	Description                         types.String                         `tfsdk:"description"`
	Shutdown                            types.Bool                           `tfsdk:"shutdown"`
	ClusterId                           types.String                         `tfsdk:"cluster_id"`
	LogNeighborChangesDisable           types.Bool                           `tfsdk:"log_neighbor_changes_disable"`
	PasswordEnctype                     types.Int64                          `tfsdk:"password_enctype"`
	PasswordText                        types.String                         `tfsdk:"password_text"`
	TimersKeepaliveInterval             types.Int64                          `tfsdk:"timers_keepalive_interval"`
	TimersHoldtime                      types.Int64                          `tfsdk:"timers_holdtime"`
	TimersMinimumNeighborHold           types.Int64                          `tfsdk:"timers_minimum_neighbor_hold"`
	Version                             types.Int64                          `tfsdk:"version"`
	FallOverDefaultRouteMap             types.String                         `tfsdk:"fall_over_default_route_map"`
	FallOverBfdMultiHop                 types.Bool                           `tfsdk:"fall_over_bfd_multi_hop"`
	FallOverBfdSingleHop                types.Bool                           `tfsdk:"fall_over_bfd_single_hop"`
	FallOverBfdCheckControlPlaneFailure types.Bool                           `tfsdk:"fall_over_bfd_check_control_plane_failure"`
	FallOverBfdStrictMode               types.Bool                           `tfsdk:"fall_over_bfd_strict_mode"`
	FallOverMaximumMetricRouteMap       types.String                         `tfsdk:"fall_over_maximum_metric_route_map"`
	DisableConnectedCheck               types.Bool                           `tfsdk:"disable_connected_check"`
	TtlSecurityHops                     types.Int64                          `tfsdk:"ttl_security_hops"`
	LocalAsAsNo                         types.String                         `tfsdk:"local_as_as_no"`
	LocalAsNoPrepend                    types.Bool                           `tfsdk:"local_as_no_prepend"`
	LocalAsReplaceAs                    types.Bool                           `tfsdk:"local_as_replace_as"`
	LocalAsDualAs                       types.Bool                           `tfsdk:"local_as_dual_as"`
	UpdateSourceLoopback                types.String                         `tfsdk:"update_source_loopback"`
	Activate                            types.Bool                           `tfsdk:"activate"`
	SendCommunity                       types.String                         `tfsdk:"send_community"`
	RouteReflectorClient                types.Bool                           `tfsdk:"route_reflector_client"`
	RouteMaps                           []BGPIPv4UnicastVRFNeighborRouteMaps `tfsdk:"route_maps"`
	EbgpMultihop                        types.Bool                           `tfsdk:"ebgp_multihop"`
	EbgpMultihopMaxHop                  types.Int64                          `tfsdk:"ebgp_multihop_max_hop"`
}

type BGPIPv4UnicastVRFNeighborData struct {
	Device                              types.String                         `tfsdk:"device"`
	Id                                  types.String                         `tfsdk:"id"`
	Asn                                 types.String                         `tfsdk:"asn"`
	Vrf                                 types.String                         `tfsdk:"vrf"`
	Ip                                  types.String                         `tfsdk:"ip"`
	RemoteAs                            types.String                         `tfsdk:"remote_as"`
	Description                         types.String                         `tfsdk:"description"`
	Shutdown                            types.Bool                           `tfsdk:"shutdown"`
	ClusterId                           types.String                         `tfsdk:"cluster_id"`
	LogNeighborChangesDisable           types.Bool                           `tfsdk:"log_neighbor_changes_disable"`
	PasswordEnctype                     types.Int64                          `tfsdk:"password_enctype"`
	PasswordText                        types.String                         `tfsdk:"password_text"`
	TimersKeepaliveInterval             types.Int64                          `tfsdk:"timers_keepalive_interval"`
	TimersHoldtime                      types.Int64                          `tfsdk:"timers_holdtime"`
	TimersMinimumNeighborHold           types.Int64                          `tfsdk:"timers_minimum_neighbor_hold"`
	Version                             types.Int64                          `tfsdk:"version"`
	FallOverDefaultRouteMap             types.String                         `tfsdk:"fall_over_default_route_map"`
	FallOverBfdMultiHop                 types.Bool                           `tfsdk:"fall_over_bfd_multi_hop"`
	FallOverBfdSingleHop                types.Bool                           `tfsdk:"fall_over_bfd_single_hop"`
	FallOverBfdCheckControlPlaneFailure types.Bool                           `tfsdk:"fall_over_bfd_check_control_plane_failure"`
	FallOverBfdStrictMode               types.Bool                           `tfsdk:"fall_over_bfd_strict_mode"`
	FallOverMaximumMetricRouteMap       types.String                         `tfsdk:"fall_over_maximum_metric_route_map"`
	DisableConnectedCheck               types.Bool                           `tfsdk:"disable_connected_check"`
	TtlSecurityHops                     types.Int64                          `tfsdk:"ttl_security_hops"`
	LocalAsAsNo                         types.String                         `tfsdk:"local_as_as_no"`
	LocalAsNoPrepend                    types.Bool                           `tfsdk:"local_as_no_prepend"`
	LocalAsReplaceAs                    types.Bool                           `tfsdk:"local_as_replace_as"`
	LocalAsDualAs                       types.Bool                           `tfsdk:"local_as_dual_as"`
	UpdateSourceLoopback                types.String                         `tfsdk:"update_source_loopback"`
	Activate                            types.Bool                           `tfsdk:"activate"`
	SendCommunity                       types.String                         `tfsdk:"send_community"`
	RouteReflectorClient                types.Bool                           `tfsdk:"route_reflector_client"`
	RouteMaps                           []BGPIPv4UnicastVRFNeighborRouteMaps `tfsdk:"route_maps"`
	EbgpMultihop                        types.Bool                           `tfsdk:"ebgp_multihop"`
	EbgpMultihopMaxHop                  types.Int64                          `tfsdk:"ebgp_multihop_max_hop"`
}
type BGPIPv4UnicastVRFNeighborRouteMaps struct {
	InOut        types.String `tfsdk:"in_out"`
	RouteMapName types.String `tfsdk:"route_map_name"`
}

func (data BGPIPv4UnicastVRFNeighbor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv4=unicast/vrf=%s/ipv4-unicast/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

func (data BGPIPv4UnicastVRFNeighborData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/with-vrf/ipv4=unicast/vrf=%s/ipv4-unicast/neighbor=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Vrf.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.Ip.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPIPv4UnicastVRFNeighbor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPIPv4UnicastVRFNeighbor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"id", data.Ip.ValueString())
	}
	if !data.RemoteAs.IsNull() && !data.RemoteAs.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"remote-as", data.RemoteAs.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.Shutdown.IsNull() && !data.Shutdown.IsUnknown() {
		if data.Shutdown.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"shutdown", map[string]string{})
		}
	}
	if !data.ClusterId.IsNull() && !data.ClusterId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"cluster-id", data.ClusterId.ValueString())
	}
	if !data.LogNeighborChangesDisable.IsNull() && !data.LogNeighborChangesDisable.IsUnknown() {
		if data.LogNeighborChangesDisable.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"log-neighbor-changes.disable", map[string]string{})
		}
	}
	if !data.PasswordEnctype.IsNull() && !data.PasswordEnctype.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.enctype", strconv.FormatInt(data.PasswordEnctype.ValueInt64(), 10))
	}
	if !data.PasswordText.IsNull() && !data.PasswordText.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"password.text", data.PasswordText.ValueString())
	}
	if !data.TimersKeepaliveInterval.IsNull() && !data.TimersKeepaliveInterval.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timers.keepalive-interval", strconv.FormatInt(data.TimersKeepaliveInterval.ValueInt64(), 10))
	}
	if !data.TimersHoldtime.IsNull() && !data.TimersHoldtime.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timers.holdtime", strconv.FormatInt(data.TimersHoldtime.ValueInt64(), 10))
	}
	if !data.TimersMinimumNeighborHold.IsNull() && !data.TimersMinimumNeighborHold.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"timers.minimum-neighbor-hold", strconv.FormatInt(data.TimersMinimumNeighborHold.ValueInt64(), 10))
	}
	if !data.Version.IsNull() && !data.Version.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"version", strconv.FormatInt(data.Version.ValueInt64(), 10))
	}
	if !data.FallOverDefaultRouteMap.IsNull() && !data.FallOverDefaultRouteMap.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.default.route-map", data.FallOverDefaultRouteMap.ValueString())
	}
	if !data.FallOverBfdMultiHop.IsNull() && !data.FallOverBfdMultiHop.IsUnknown() {
		if data.FallOverBfdMultiHop.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.bfd.multi-hop", map[string]string{})
		}
	}
	if !data.FallOverBfdSingleHop.IsNull() && !data.FallOverBfdSingleHop.IsUnknown() {
		if data.FallOverBfdSingleHop.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.bfd.single-hop", map[string]string{})
		}
	}
	if !data.FallOverBfdCheckControlPlaneFailure.IsNull() && !data.FallOverBfdCheckControlPlaneFailure.IsUnknown() {
		if data.FallOverBfdCheckControlPlaneFailure.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.bfd.check-control-plane-failure", map[string]string{})
		}
	}
	if !data.FallOverBfdStrictMode.IsNull() && !data.FallOverBfdStrictMode.IsUnknown() {
		if data.FallOverBfdStrictMode.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.bfd.strict-mode", map[string]string{})
		}
	}
	if !data.FallOverMaximumMetricRouteMap.IsNull() && !data.FallOverMaximumMetricRouteMap.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"fall-over.maximum-metric.route-map", data.FallOverMaximumMetricRouteMap.ValueString())
	}
	if !data.DisableConnectedCheck.IsNull() && !data.DisableConnectedCheck.IsUnknown() {
		if data.DisableConnectedCheck.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"disable-connected-check", map[string]string{})
		}
	}
	if !data.TtlSecurityHops.IsNull() && !data.TtlSecurityHops.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ttl-security.hops", strconv.FormatInt(data.TtlSecurityHops.ValueInt64(), 10))
	}
	if !data.LocalAsAsNo.IsNull() && !data.LocalAsAsNo.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"local-as.as-no", data.LocalAsAsNo.ValueString())
	}
	if !data.LocalAsNoPrepend.IsNull() && !data.LocalAsNoPrepend.IsUnknown() {
		if data.LocalAsNoPrepend.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"local-as.no-prepend", map[string]string{})
		}
	}
	if !data.LocalAsReplaceAs.IsNull() && !data.LocalAsReplaceAs.IsUnknown() {
		if data.LocalAsReplaceAs.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"local-as.replace-as", map[string]string{})
		}
	}
	if !data.LocalAsDualAs.IsNull() && !data.LocalAsDualAs.IsUnknown() {
		if data.LocalAsDualAs.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"local-as.dual-as", map[string]string{})
		}
	}
	if !data.UpdateSourceLoopback.IsNull() && !data.UpdateSourceLoopback.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"update-source.interface.Loopback", data.UpdateSourceLoopback.ValueString())
	}
	if !data.Activate.IsNull() && !data.Activate.IsUnknown() {
		if data.Activate.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"activate", map[string]string{})
		}
	}
	if !data.SendCommunity.IsNull() && !data.SendCommunity.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"send-community.send-community-where", data.SendCommunity.ValueString())
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.IsUnknown() {
		if data.RouteReflectorClient.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-reflector-client", map[string]string{})
		}
	}
	if !data.EbgpMultihop.IsNull() && !data.EbgpMultihop.IsUnknown() {
		if data.EbgpMultihop.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ebgp-multihop", map[string]string{})
		}
	}
	if !data.EbgpMultihopMaxHop.IsNull() && !data.EbgpMultihopMaxHop.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ebgp-multihop.max-hop", strconv.FormatInt(data.EbgpMultihopMaxHop.ValueInt64(), 10))
	}
	if len(data.RouteMaps) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map", []interface{}{})
		for index, item := range data.RouteMaps {
			if !item.InOut.IsNull() && !item.InOut.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map"+"."+strconv.Itoa(index)+"."+"inout", item.InOut.ValueString())
			}
			if !item.RouteMapName.IsNull() && !item.RouteMapName.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"route-map"+"."+strconv.Itoa(index)+"."+"route-map-name", item.RouteMapName.ValueString())
			}
		}
	}
	return body
}

func (data *BGPIPv4UnicastVRFNeighbor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "id"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() && !data.RemoteAs.IsNull() {
		data.RemoteAs = types.StringValue(value.String())
	} else {
		data.RemoteAs = types.StringNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(prefix + "shutdown"); !data.Shutdown.IsNull() {
		if value.Exists() {
			data.Shutdown = types.BoolValue(true)
		} else {
			data.Shutdown = types.BoolValue(false)
		}
	} else {
		data.Shutdown = types.BoolNull()
	}
	if value := res.Get(prefix + "cluster-id"); value.Exists() && !data.ClusterId.IsNull() {
		data.ClusterId = types.StringValue(value.String())
	} else {
		data.ClusterId = types.StringNull()
	}
	if value := res.Get(prefix + "log-neighbor-changes.disable"); !data.LogNeighborChangesDisable.IsNull() {
		if value.Exists() {
			data.LogNeighborChangesDisable = types.BoolValue(true)
		} else {
			data.LogNeighborChangesDisable = types.BoolValue(false)
		}
	} else {
		data.LogNeighborChangesDisable = types.BoolNull()
	}
	if value := res.Get(prefix + "password.enctype"); value.Exists() && !data.PasswordEnctype.IsNull() {
		data.PasswordEnctype = types.Int64Value(value.Int())
	} else {
		data.PasswordEnctype = types.Int64Null()
	}
	if value := res.Get(prefix + "password.text"); value.Exists() && !data.PasswordText.IsNull() {
		data.PasswordText = types.StringValue(value.String())
	} else {
		data.PasswordText = types.StringNull()
	}
	if value := res.Get(prefix + "timers.keepalive-interval"); value.Exists() && !data.TimersKeepaliveInterval.IsNull() {
		data.TimersKeepaliveInterval = types.Int64Value(value.Int())
	} else {
		data.TimersKeepaliveInterval = types.Int64Null()
	}
	if value := res.Get(prefix + "timers.holdtime"); value.Exists() && !data.TimersHoldtime.IsNull() {
		data.TimersHoldtime = types.Int64Value(value.Int())
	} else {
		data.TimersHoldtime = types.Int64Null()
	}
	if value := res.Get(prefix + "timers.minimum-neighbor-hold"); value.Exists() && !data.TimersMinimumNeighborHold.IsNull() {
		data.TimersMinimumNeighborHold = types.Int64Value(value.Int())
	} else {
		data.TimersMinimumNeighborHold = types.Int64Null()
	}
	if value := res.Get(prefix + "version"); value.Exists() && !data.Version.IsNull() {
		data.Version = types.Int64Value(value.Int())
	} else {
		data.Version = types.Int64Null()
	}
	if value := res.Get(prefix + "fall-over.default.route-map"); value.Exists() && !data.FallOverDefaultRouteMap.IsNull() {
		data.FallOverDefaultRouteMap = types.StringValue(value.String())
	} else {
		data.FallOverDefaultRouteMap = types.StringNull()
	}
	if value := res.Get(prefix + "fall-over.bfd.multi-hop"); !data.FallOverBfdMultiHop.IsNull() {
		if value.Exists() {
			data.FallOverBfdMultiHop = types.BoolValue(true)
		} else {
			data.FallOverBfdMultiHop = types.BoolValue(false)
		}
	} else {
		data.FallOverBfdMultiHop = types.BoolNull()
	}
	if value := res.Get(prefix + "fall-over.bfd.single-hop"); !data.FallOverBfdSingleHop.IsNull() {
		if value.Exists() {
			data.FallOverBfdSingleHop = types.BoolValue(true)
		} else {
			data.FallOverBfdSingleHop = types.BoolValue(false)
		}
	} else {
		data.FallOverBfdSingleHop = types.BoolNull()
	}
	if value := res.Get(prefix + "fall-over.bfd.check-control-plane-failure"); !data.FallOverBfdCheckControlPlaneFailure.IsNull() {
		if value.Exists() {
			data.FallOverBfdCheckControlPlaneFailure = types.BoolValue(true)
		} else {
			data.FallOverBfdCheckControlPlaneFailure = types.BoolValue(false)
		}
	} else {
		data.FallOverBfdCheckControlPlaneFailure = types.BoolNull()
	}
	if value := res.Get(prefix + "fall-over.bfd.strict-mode"); !data.FallOverBfdStrictMode.IsNull() {
		if value.Exists() {
			data.FallOverBfdStrictMode = types.BoolValue(true)
		} else {
			data.FallOverBfdStrictMode = types.BoolValue(false)
		}
	} else {
		data.FallOverBfdStrictMode = types.BoolNull()
	}
	if value := res.Get(prefix + "fall-over.maximum-metric.route-map"); value.Exists() && !data.FallOverMaximumMetricRouteMap.IsNull() {
		data.FallOverMaximumMetricRouteMap = types.StringValue(value.String())
	} else {
		data.FallOverMaximumMetricRouteMap = types.StringNull()
	}
	if value := res.Get(prefix + "disable-connected-check"); !data.DisableConnectedCheck.IsNull() {
		if value.Exists() {
			data.DisableConnectedCheck = types.BoolValue(true)
		} else {
			data.DisableConnectedCheck = types.BoolValue(false)
		}
	} else {
		data.DisableConnectedCheck = types.BoolNull()
	}
	if value := res.Get(prefix + "ttl-security.hops"); value.Exists() && !data.TtlSecurityHops.IsNull() {
		data.TtlSecurityHops = types.Int64Value(value.Int())
	} else {
		data.TtlSecurityHops = types.Int64Null()
	}
	if value := res.Get(prefix + "local-as.as-no"); value.Exists() && !data.LocalAsAsNo.IsNull() {
		data.LocalAsAsNo = types.StringValue(value.String())
	} else {
		data.LocalAsAsNo = types.StringNull()
	}
	if value := res.Get(prefix + "local-as.no-prepend"); !data.LocalAsNoPrepend.IsNull() {
		if value.Exists() {
			data.LocalAsNoPrepend = types.BoolValue(true)
		} else {
			data.LocalAsNoPrepend = types.BoolValue(false)
		}
	} else {
		data.LocalAsNoPrepend = types.BoolNull()
	}
	if value := res.Get(prefix + "local-as.replace-as"); !data.LocalAsReplaceAs.IsNull() {
		if value.Exists() {
			data.LocalAsReplaceAs = types.BoolValue(true)
		} else {
			data.LocalAsReplaceAs = types.BoolValue(false)
		}
	} else {
		data.LocalAsReplaceAs = types.BoolNull()
	}
	if value := res.Get(prefix + "local-as.dual-as"); !data.LocalAsDualAs.IsNull() {
		if value.Exists() {
			data.LocalAsDualAs = types.BoolValue(true)
		} else {
			data.LocalAsDualAs = types.BoolValue(false)
		}
	} else {
		data.LocalAsDualAs = types.BoolNull()
	}
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() && !data.UpdateSourceLoopback.IsNull() {
		data.UpdateSourceLoopback = types.StringValue(value.String())
	} else {
		data.UpdateSourceLoopback = types.StringNull()
	}
	if value := res.Get(prefix + "activate"); !data.Activate.IsNull() {
		if value.Exists() {
			data.Activate = types.BoolValue(true)
		} else {
			data.Activate = types.BoolValue(false)
		}
	} else {
		data.Activate = types.BoolNull()
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() && !data.SendCommunity.IsNull() {
		data.SendCommunity = types.StringValue(value.String())
	} else {
		data.SendCommunity = types.StringNull()
	}
	if value := res.Get(prefix + "route-reflector-client"); !data.RouteReflectorClient.IsNull() {
		if value.Exists() {
			data.RouteReflectorClient = types.BoolValue(true)
		} else {
			data.RouteReflectorClient = types.BoolValue(false)
		}
	} else {
		data.RouteReflectorClient = types.BoolNull()
	}
	for i := range data.RouteMaps {
		keys := [...]string{"inout"}
		keyValues := [...]string{data.RouteMaps[i].InOut.ValueString()}

		var r gjson.Result
		res.Get(prefix + "route-map").ForEach(
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
		if value := r.Get("inout"); value.Exists() && !data.RouteMaps[i].InOut.IsNull() {
			data.RouteMaps[i].InOut = types.StringValue(value.String())
		} else {
			data.RouteMaps[i].InOut = types.StringNull()
		}
		if value := r.Get("route-map-name"); value.Exists() && !data.RouteMaps[i].RouteMapName.IsNull() {
			data.RouteMaps[i].RouteMapName = types.StringValue(value.String())
		} else {
			data.RouteMaps[i].RouteMapName = types.StringNull()
		}
	}
	if value := res.Get(prefix + "ebgp-multihop"); !data.EbgpMultihop.IsNull() {
		if value.Exists() {
			data.EbgpMultihop = types.BoolValue(true)
		} else {
			data.EbgpMultihop = types.BoolValue(false)
		}
	} else {
		data.EbgpMultihop = types.BoolNull()
	}
	if value := res.Get(prefix + "ebgp-multihop.max-hop"); value.Exists() && !data.EbgpMultihopMaxHop.IsNull() {
		data.EbgpMultihopMaxHop = types.Int64Value(value.Int())
	} else {
		data.EbgpMultihopMaxHop = types.Int64Null()
	}
}

func (data *BGPIPv4UnicastVRFNeighborData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "remote-as"); value.Exists() {
		data.RemoteAs = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "shutdown"); value.Exists() {
		data.Shutdown = types.BoolValue(true)
	} else {
		data.Shutdown = types.BoolValue(false)
	}
	if value := res.Get(prefix + "cluster-id"); value.Exists() {
		data.ClusterId = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "log-neighbor-changes.disable"); value.Exists() {
		data.LogNeighborChangesDisable = types.BoolValue(true)
	} else {
		data.LogNeighborChangesDisable = types.BoolValue(false)
	}
	if value := res.Get(prefix + "password.enctype"); value.Exists() {
		data.PasswordEnctype = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "password.text"); value.Exists() {
		data.PasswordText = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "timers.keepalive-interval"); value.Exists() {
		data.TimersKeepaliveInterval = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "timers.holdtime"); value.Exists() {
		data.TimersHoldtime = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "timers.minimum-neighbor-hold"); value.Exists() {
		data.TimersMinimumNeighborHold = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "version"); value.Exists() {
		data.Version = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "fall-over.default.route-map"); value.Exists() {
		data.FallOverDefaultRouteMap = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "fall-over.bfd.multi-hop"); value.Exists() {
		data.FallOverBfdMultiHop = types.BoolValue(true)
	} else {
		data.FallOverBfdMultiHop = types.BoolValue(false)
	}
	if value := res.Get(prefix + "fall-over.bfd.single-hop"); value.Exists() {
		data.FallOverBfdSingleHop = types.BoolValue(true)
	} else {
		data.FallOverBfdSingleHop = types.BoolValue(false)
	}
	if value := res.Get(prefix + "fall-over.bfd.check-control-plane-failure"); value.Exists() {
		data.FallOverBfdCheckControlPlaneFailure = types.BoolValue(true)
	} else {
		data.FallOverBfdCheckControlPlaneFailure = types.BoolValue(false)
	}
	if value := res.Get(prefix + "fall-over.bfd.strict-mode"); value.Exists() {
		data.FallOverBfdStrictMode = types.BoolValue(true)
	} else {
		data.FallOverBfdStrictMode = types.BoolValue(false)
	}
	if value := res.Get(prefix + "fall-over.maximum-metric.route-map"); value.Exists() {
		data.FallOverMaximumMetricRouteMap = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "disable-connected-check"); value.Exists() {
		data.DisableConnectedCheck = types.BoolValue(true)
	} else {
		data.DisableConnectedCheck = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ttl-security.hops"); value.Exists() {
		data.TtlSecurityHops = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "local-as.as-no"); value.Exists() {
		data.LocalAsAsNo = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "local-as.no-prepend"); value.Exists() {
		data.LocalAsNoPrepend = types.BoolValue(true)
	} else {
		data.LocalAsNoPrepend = types.BoolValue(false)
	}
	if value := res.Get(prefix + "local-as.replace-as"); value.Exists() {
		data.LocalAsReplaceAs = types.BoolValue(true)
	} else {
		data.LocalAsReplaceAs = types.BoolValue(false)
	}
	if value := res.Get(prefix + "local-as.dual-as"); value.Exists() {
		data.LocalAsDualAs = types.BoolValue(true)
	} else {
		data.LocalAsDualAs = types.BoolValue(false)
	}
	if value := res.Get(prefix + "update-source.interface.Loopback"); value.Exists() {
		data.UpdateSourceLoopback = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "activate"); value.Exists() {
		data.Activate = types.BoolValue(true)
	} else {
		data.Activate = types.BoolValue(false)
	}
	if value := res.Get(prefix + "send-community.send-community-where"); value.Exists() {
		data.SendCommunity = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "route-reflector-client"); value.Exists() {
		data.RouteReflectorClient = types.BoolValue(true)
	} else {
		data.RouteReflectorClient = types.BoolValue(false)
	}
	if value := res.Get(prefix + "route-map"); value.Exists() {
		data.RouteMaps = make([]BGPIPv4UnicastVRFNeighborRouteMaps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPIPv4UnicastVRFNeighborRouteMaps{}
			if cValue := v.Get("inout"); cValue.Exists() {
				item.InOut = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map-name"); cValue.Exists() {
				item.RouteMapName = types.StringValue(cValue.String())
			}
			data.RouteMaps = append(data.RouteMaps, item)
			return true
		})
	}
	if value := res.Get(prefix + "ebgp-multihop"); value.Exists() {
		data.EbgpMultihop = types.BoolValue(true)
	} else {
		data.EbgpMultihop = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ebgp-multihop.max-hop"); value.Exists() {
		data.EbgpMultihopMaxHop = types.Int64Value(value.Int())
	}
}

func (data *BGPIPv4UnicastVRFNeighbor) getDeletedListItems(ctx context.Context, state BGPIPv4UnicastVRFNeighbor) []string {
	deletedListItems := make([]string, 0)
	for i := range state.RouteMaps {
		stateKeyValues := [...]string{state.RouteMaps[i].InOut.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.RouteMaps[i].InOut.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.RouteMaps {
			found = true
			if state.RouteMaps[i].InOut.ValueString() != data.RouteMaps[j].InOut.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/route-map=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *BGPIPv4UnicastVRFNeighbor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Shutdown.IsNull() && !data.Shutdown.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.LogNeighborChangesDisable.IsNull() && !data.LogNeighborChangesDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/log-neighbor-changes/disable", data.getPath()))
	}
	if !data.FallOverBfdMultiHop.IsNull() && !data.FallOverBfdMultiHop.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fall-over/bfd/multi-hop", data.getPath()))
	}
	if !data.FallOverBfdSingleHop.IsNull() && !data.FallOverBfdSingleHop.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fall-over/bfd/single-hop", data.getPath()))
	}
	if !data.FallOverBfdCheckControlPlaneFailure.IsNull() && !data.FallOverBfdCheckControlPlaneFailure.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fall-over/bfd/check-control-plane-failure", data.getPath()))
	}
	if !data.FallOverBfdStrictMode.IsNull() && !data.FallOverBfdStrictMode.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/fall-over/bfd/strict-mode", data.getPath()))
	}
	if !data.DisableConnectedCheck.IsNull() && !data.DisableConnectedCheck.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/disable-connected-check", data.getPath()))
	}
	if !data.LocalAsNoPrepend.IsNull() && !data.LocalAsNoPrepend.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/local-as/no-prepend", data.getPath()))
	}
	if !data.LocalAsReplaceAs.IsNull() && !data.LocalAsReplaceAs.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/local-as/replace-as", data.getPath()))
	}
	if !data.LocalAsDualAs.IsNull() && !data.LocalAsDualAs.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/local-as/dual-as", data.getPath()))
	}
	if !data.Activate.IsNull() && !data.Activate.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/activate", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() && !data.RouteReflectorClient.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}

	if !data.EbgpMultihop.IsNull() && !data.EbgpMultihop.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ebgp-multihop", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *BGPIPv4UnicastVRFNeighbor) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	if !data.Shutdown.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/shutdown", data.getPath()))
	}
	if !data.ClusterId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/cluster-id", data.getPath()))
	}
	if !data.LogNeighborChangesDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/log-neighbor-changes/disable", data.getPath()))
	}
	if !data.PasswordEnctype.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/password/enctype", data.getPath()))
	}
	if !data.PasswordText.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/password/text", data.getPath()))
	}
	if !data.TimersKeepaliveInterval.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/keepalive-interval", data.getPath()))
	}
	if !data.TimersHoldtime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/holdtime", data.getPath()))
	}
	if !data.TimersMinimumNeighborHold.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timers/minimum-neighbor-hold", data.getPath()))
	}
	if !data.Version.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/version", data.getPath()))
	}
	if !data.FallOverDefaultRouteMap.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/default/route-map", data.getPath()))
	}
	if !data.FallOverBfdMultiHop.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/bfd/multi-hop", data.getPath()))
	}
	if !data.FallOverBfdSingleHop.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/bfd/single-hop", data.getPath()))
	}
	if !data.FallOverBfdCheckControlPlaneFailure.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/bfd/check-control-plane-failure", data.getPath()))
	}
	if !data.FallOverBfdStrictMode.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/bfd/strict-mode", data.getPath()))
	}
	if !data.FallOverMaximumMetricRouteMap.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/fall-over/maximum-metric/route-map", data.getPath()))
	}
	if !data.DisableConnectedCheck.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/disable-connected-check", data.getPath()))
	}
	if !data.TtlSecurityHops.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ttl-security/hops", data.getPath()))
	}
	if !data.LocalAsAsNo.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/local-as/as-no", data.getPath()))
	}
	if !data.LocalAsNoPrepend.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/local-as/no-prepend", data.getPath()))
	}
	if !data.LocalAsReplaceAs.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/local-as/replace-as", data.getPath()))
	}
	if !data.LocalAsDualAs.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/local-as/dual-as", data.getPath()))
	}
	if !data.UpdateSourceLoopback.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/update-source/interface/Loopback", data.getPath()))
	}
	if !data.SendCommunity.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/send-community/send-community-where", data.getPath()))
	}
	if !data.RouteReflectorClient.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-reflector-client", data.getPath()))
	}
	for i := range data.RouteMaps {
		keyValues := [...]string{data.RouteMaps[i].InOut.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-map=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.EbgpMultihop.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ebgp-multihop", data.getPath()))
	}
	if !data.EbgpMultihopMaxHop.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ebgp-multihop/max-hop", data.getPath()))
	}
	return deletePaths
}
