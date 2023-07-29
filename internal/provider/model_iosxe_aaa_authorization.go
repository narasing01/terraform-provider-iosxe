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

type AAAAuthorization struct {
	Device     types.String               `tfsdk:"device"`
	Id         types.String               `tfsdk:"id"`
	DeleteMode types.String               `tfsdk:"delete_mode"`
	Execs      []AAAAuthorizationExecs    `tfsdk:"execs"`
	Networks   []AAAAuthorizationNetworks `tfsdk:"networks"`
}

type AAAAuthorizationData struct {
	Device   types.String               `tfsdk:"device"`
	Id       types.String               `tfsdk:"id"`
	Execs    []AAAAuthorizationExecs    `tfsdk:"execs"`
	Networks []AAAAuthorizationNetworks `tfsdk:"networks"`
}
type AAAAuthorizationExecs struct {
	Name              types.String `tfsdk:"name"`
	A1Local           types.Bool   `tfsdk:"a1_local"`
	A1IfAuthenticated types.Bool   `tfsdk:"a1_if_authenticated"`
}
type AAAAuthorizationNetworks struct {
	Id      types.String `tfsdk:"id"`
	A1Group types.String `tfsdk:"a1_group"`
}

func (data AAAAuthorization) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:authorization"
}

func (data AAAAuthorizationData) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:authorization"
}

// if last path element has a key -> remove it
func (data AAAAuthorization) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AAAAuthorization) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if len(data.Execs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec", []interface{}{})
		for index, item := range data.Execs {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.A1Local.IsNull() && !item.A1Local.IsUnknown() {
				if item.A1Local.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"a1.local", map[string]string{})
				}
			}
			if !item.A1IfAuthenticated.IsNull() && !item.A1IfAuthenticated.IsUnknown() {
				if item.A1IfAuthenticated.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"a1.if-authenticated", map[string]string{})
				}
			}
		}
	}
	if len(data.Networks) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network", []interface{}{})
		for index, item := range data.Networks {
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"id", item.Id.ValueString())
			}
			if !item.A1Group.IsNull() && !item.A1Group.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"a1.group", item.A1Group.ValueString())
			}
		}
	}
	return body
}

func (data *AAAAuthorization) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	for i := range data.Execs {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Execs[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "exec").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Execs[i].Name.IsNull() {
			data.Execs[i].Name = types.StringValue(value.String())
		} else {
			data.Execs[i].Name = types.StringNull()
		}
		if value := r.Get("a1.local"); !data.Execs[i].A1Local.IsNull() {
			if value.Exists() {
				data.Execs[i].A1Local = types.BoolValue(true)
			} else {
				data.Execs[i].A1Local = types.BoolValue(false)
			}
		} else {
			data.Execs[i].A1Local = types.BoolNull()
		}
		if value := r.Get("a1.if-authenticated"); !data.Execs[i].A1IfAuthenticated.IsNull() {
			if value.Exists() {
				data.Execs[i].A1IfAuthenticated = types.BoolValue(true)
			} else {
				data.Execs[i].A1IfAuthenticated = types.BoolValue(false)
			}
		} else {
			data.Execs[i].A1IfAuthenticated = types.BoolNull()
		}
	}
	for i := range data.Networks {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Networks[i].Id.ValueString()}

		var r gjson.Result
		res.Get(prefix + "network").ForEach(
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
		if value := r.Get("id"); value.Exists() && !data.Networks[i].Id.IsNull() {
			data.Networks[i].Id = types.StringValue(value.String())
		} else {
			data.Networks[i].Id = types.StringNull()
		}
		if value := r.Get("a1.group"); value.Exists() && !data.Networks[i].A1Group.IsNull() {
			data.Networks[i].A1Group = types.StringValue(value.String())
		} else {
			data.Networks[i].A1Group = types.StringNull()
		}
	}
}

func (data *AAAAuthorizationData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "exec"); value.Exists() {
		data.Execs = make([]AAAAuthorizationExecs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAuthorizationExecs{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("a1.local"); cValue.Exists() {
				item.A1Local = types.BoolValue(true)
			} else {
				item.A1Local = types.BoolValue(false)
			}
			if cValue := v.Get("a1.if-authenticated"); cValue.Exists() {
				item.A1IfAuthenticated = types.BoolValue(true)
			} else {
				item.A1IfAuthenticated = types.BoolValue(false)
			}
			data.Execs = append(data.Execs, item)
			return true
		})
	}
	if value := res.Get(prefix + "network"); value.Exists() {
		data.Networks = make([]AAAAuthorizationNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAuthorizationNetworks{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			}
			if cValue := v.Get("a1.group"); cValue.Exists() {
				item.A1Group = types.StringValue(cValue.String())
			}
			data.Networks = append(data.Networks, item)
			return true
		})
	}
}

func (data *AAAAuthorization) getDeletedListItems(ctx context.Context, state AAAAuthorization) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Execs {
		stateKeyValues := [...]string{state.Execs[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Execs[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Execs {
			found = true
			if state.Execs[i].Name.ValueString() != data.Execs[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/exec=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Networks {
		stateKeyValues := [...]string{state.Networks[i].Id.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Networks[i].Id.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Networks {
			found = true
			if state.Networks[i].Id.ValueString() != data.Networks[j].Id.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/network=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *AAAAuthorization) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	for i := range data.Execs {
		keyValues := [...]string{data.Execs[i].Name.ValueString()}
		if !data.Execs[i].A1Local.IsNull() && !data.Execs[i].A1Local.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/exec=%v/a1/local", data.getPath(), strings.Join(keyValues[:], ",")))
		}
		if !data.Execs[i].A1IfAuthenticated.IsNull() && !data.Execs[i].A1IfAuthenticated.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/exec=%v/a1/if-authenticated", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	return emptyLeafsDelete
}

func (data *AAAAuthorization) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	for i := range data.Execs {
		keyValues := [...]string{data.Execs[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/exec=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Networks {
		keyValues := [...]string{data.Networks[i].Id.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/network=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}