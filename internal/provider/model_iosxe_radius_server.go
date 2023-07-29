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
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RadiusServer struct {
	Device     types.String             `tfsdk:"device"`
	Id         types.String             `tfsdk:"id"`
	Attributes []RadiusServerAttributes `tfsdk:"attributes"`
}

type RadiusServerData struct {
	Device     types.String             `tfsdk:"device"`
	Id         types.String             `tfsdk:"id"`
	Attributes []RadiusServerAttributes `tfsdk:"attributes"`
}
type RadiusServerAttributes struct {
	Number  types.String                    `tfsdk:"number"`
	Attri31 []RadiusServerAttributesAttri31 `tfsdk:"attri31"`
}
type RadiusServerAttributesAttri31 struct {
	CallingStationId types.String `tfsdk:"calling_station_id"`
	Attri31Format    types.String `tfsdk:"attri31_format"`
	Attri31LuCase    types.String `tfsdk:"attri31_lu_case"`
}

func (data RadiusServer) getPath() string {
	return "Cisco-IOS-XE-native:native/radius-server"
}

func (data RadiusServerData) getPath() string {
	return "Cisco-IOS-XE-native:native/radius-server"
}

// if last path element has a key -> remove it
func (data RadiusServer) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data RadiusServer) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if len(data.Attributes) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute", []interface{}{})
		for index, item := range data.Attributes {
			if !item.Number.IsNull() && !item.Number.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute"+"."+strconv.Itoa(index)+"."+"number", item.Number.ValueString())
			}
			if len(item.Attri31) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute"+"."+strconv.Itoa(index)+"."+"attri31.attri31-list", []interface{}{})
				for cindex, citem := range item.Attri31 {
					if !citem.CallingStationId.IsNull() && !citem.CallingStationId.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute"+"."+strconv.Itoa(index)+"."+"attri31.attri31-list"+"."+strconv.Itoa(cindex)+"."+"calling-station-id", citem.CallingStationId.ValueString())
					}
					if !citem.Attri31Format.IsNull() && !citem.Attri31Format.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute"+"."+strconv.Itoa(index)+"."+"attri31.attri31-list"+"."+strconv.Itoa(cindex)+"."+"id-mac.format", citem.Attri31Format.ValueString())
					}
					if !citem.Attri31LuCase.IsNull() && !citem.Attri31LuCase.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:attribute"+"."+strconv.Itoa(index)+"."+"attri31.attri31-list"+"."+strconv.Itoa(cindex)+"."+"id-mac.lu-case", citem.Attri31LuCase.ValueString())
					}
				}
			}
		}
	}
	return body
}

func (data *RadiusServer) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	for i := range data.Attributes {
		keys := [...]string{"number"}
		keyValues := [...]string{data.Attributes[i].Number.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-aaa:attribute").ForEach(
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
		if value := r.Get("number"); value.Exists() && !data.Attributes[i].Number.IsNull() {
			data.Attributes[i].Number = types.StringValue(value.String())
		} else {
			data.Attributes[i].Number = types.StringNull()
		}
		for ci := range data.Attributes[i].Attri31 {
			keys := [...]string{"calling-station-id"}
			keyValues := [...]string{data.Attributes[i].Attri31[ci].CallingStationId.ValueString()}

			var cr gjson.Result
			r.Get("attri31.attri31-list").ForEach(
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
			if value := cr.Get("calling-station-id"); value.Exists() && !data.Attributes[i].Attri31[ci].CallingStationId.IsNull() {
				data.Attributes[i].Attri31[ci].CallingStationId = types.StringValue(value.String())
			} else {
				data.Attributes[i].Attri31[ci].CallingStationId = types.StringNull()
			}
			if value := cr.Get("id-mac.format"); value.Exists() && !data.Attributes[i].Attri31[ci].Attri31Format.IsNull() {
				data.Attributes[i].Attri31[ci].Attri31Format = types.StringValue(value.String())
			} else {
				data.Attributes[i].Attri31[ci].Attri31Format = types.StringNull()
			}
			if value := cr.Get("id-mac.lu-case"); value.Exists() && !data.Attributes[i].Attri31[ci].Attri31LuCase.IsNull() {
				data.Attributes[i].Attri31[ci].Attri31LuCase = types.StringValue(value.String())
			} else {
				data.Attributes[i].Attri31[ci].Attri31LuCase = types.StringNull()
			}
		}
	}
}

func (data *RadiusServerData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:attribute"); value.Exists() {
		data.Attributes = make([]RadiusServerAttributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RadiusServerAttributes{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Number = types.StringValue(cValue.String())
			}
			if cValue := v.Get("attri31.attri31-list"); cValue.Exists() {
				item.Attri31 = make([]RadiusServerAttributesAttri31, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := RadiusServerAttributesAttri31{}
					if ccValue := cv.Get("calling-station-id"); ccValue.Exists() {
						cItem.CallingStationId = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("id-mac.format"); ccValue.Exists() {
						cItem.Attri31Format = types.StringValue(ccValue.String())
					}
					if ccValue := cv.Get("id-mac.lu-case"); ccValue.Exists() {
						cItem.Attri31LuCase = types.StringValue(ccValue.String())
					}
					item.Attri31 = append(item.Attri31, cItem)
					return true
				})
			}
			data.Attributes = append(data.Attributes, item)
			return true
		})
	}
}

func (data *RadiusServer) getDeletedListItems(ctx context.Context, state RadiusServer) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Attributes {
		stateKeyValues := [...]string{state.Attributes[i].Number.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Attributes[i].Number.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Attributes {
			found = true
			if state.Attributes[i].Number.ValueString() != data.Attributes[j].Number.ValueString() {
				found = false
			}
			if found {
				for ci := range state.Attributes[i].Attri31 {
					cstateKeyValues := [...]string{state.Attributes[i].Attri31[ci].CallingStationId.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.Attributes[i].Attri31[ci].CallingStationId.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.Attributes[j].Attri31 {
						found = true
						if state.Attributes[i].Attri31[ci].CallingStationId.ValueString() != data.Attributes[j].Attri31[cj].CallingStationId.ValueString() {
							found = false
						}
						if found {
							break
						}
					}
					if !found {
						deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:attribute=%v/attri31/attri31-list=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:attribute=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *RadiusServer) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *RadiusServer) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Attributes {
		keyValues := [...]string{data.Attributes[i].Number.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:attribute=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}