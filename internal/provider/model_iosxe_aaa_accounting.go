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

type AAAAccounting struct {
	Device                         types.String            `tfsdk:"device"`
	Id                             types.String            `tfsdk:"id"`
	DeleteMode                     types.String            `tfsdk:"delete_mode"`
	UpdateNewinfoPeriodic          types.Int64             `tfsdk:"update_newinfo_periodic"`
	IdentityDefaultStartStopGroup1 types.String            `tfsdk:"identity_default_start_stop_group1"`
	IdentityDefaultStartStopGroup2 types.String            `tfsdk:"identity_default_start_stop_group2"`
	IdentityDefaultStartStopGroup3 types.String            `tfsdk:"identity_default_start_stop_group3"`
	IdentityDefaultStartStopGroup4 types.String            `tfsdk:"identity_default_start_stop_group4"`
	Execs                          []AAAAccountingExecs    `tfsdk:"execs"`
	Networks                       []AAAAccountingNetworks `tfsdk:"networks"`
	SystemGuaranteeFirst           types.Bool              `tfsdk:"system_guarantee_first"`
}

type AAAAccountingData struct {
	Device                         types.String            `tfsdk:"device"`
	Id                             types.String            `tfsdk:"id"`
	UpdateNewinfoPeriodic          types.Int64             `tfsdk:"update_newinfo_periodic"`
	IdentityDefaultStartStopGroup1 types.String            `tfsdk:"identity_default_start_stop_group1"`
	IdentityDefaultStartStopGroup2 types.String            `tfsdk:"identity_default_start_stop_group2"`
	IdentityDefaultStartStopGroup3 types.String            `tfsdk:"identity_default_start_stop_group3"`
	IdentityDefaultStartStopGroup4 types.String            `tfsdk:"identity_default_start_stop_group4"`
	Execs                          []AAAAccountingExecs    `tfsdk:"execs"`
	Networks                       []AAAAccountingNetworks `tfsdk:"networks"`
	SystemGuaranteeFirst           types.Bool              `tfsdk:"system_guarantee_first"`
}
type AAAAccountingExecs struct {
	Name            types.String `tfsdk:"name"`
	StartStopGroup1 types.String `tfsdk:"start_stop_group1"`
}
type AAAAccountingNetworks struct {
	Id              types.String `tfsdk:"id"`
	StartStopGroup1 types.String `tfsdk:"start_stop_group1"`
	StartStopGroup2 types.String `tfsdk:"start_stop_group2"`
}

func (data AAAAccounting) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:accounting"
}

func (data AAAAccountingData) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:accounting"
}

// if last path element has a key -> remove it
func (data AAAAccounting) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AAAAccounting) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.UpdateNewinfoPeriodic.IsNull() && !data.UpdateNewinfoPeriodic.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"update.newinfo.periodic", strconv.FormatInt(data.UpdateNewinfoPeriodic.ValueInt64(), 10))
	}
	if !data.IdentityDefaultStartStopGroup1.IsNull() && !data.IdentityDefaultStartStopGroup1.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.default.start-stop.group-config.group1.group", data.IdentityDefaultStartStopGroup1.ValueString())
	}
	if !data.IdentityDefaultStartStopGroup2.IsNull() && !data.IdentityDefaultStartStopGroup2.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.default.start-stop.group-config.group2.group", data.IdentityDefaultStartStopGroup2.ValueString())
	}
	if !data.IdentityDefaultStartStopGroup3.IsNull() && !data.IdentityDefaultStartStopGroup3.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.default.start-stop.group-config.group3.group", data.IdentityDefaultStartStopGroup3.ValueString())
	}
	if !data.IdentityDefaultStartStopGroup4.IsNull() && !data.IdentityDefaultStartStopGroup4.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"identity.default.start-stop.group-config.group4.group", data.IdentityDefaultStartStopGroup4.ValueString())
	}
	if !data.SystemGuaranteeFirst.IsNull() && !data.SystemGuaranteeFirst.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"system.guarantee-first", data.SystemGuaranteeFirst.ValueBool())
	}
	if len(data.Execs) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec", []interface{}{})
		for index, item := range data.Execs {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if !item.StartStopGroup1.IsNull() && !item.StartStopGroup1.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exec"+"."+strconv.Itoa(index)+"."+"start-stop.group1.group", item.StartStopGroup1.ValueString())
			}
		}
	}
	if len(data.Networks) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network", []interface{}{})
		for index, item := range data.Networks {
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"id", item.Id.ValueString())
			}
			if !item.StartStopGroup1.IsNull() && !item.StartStopGroup1.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"start-stop.group-config.group1.group", item.StartStopGroup1.ValueString())
			}
			if !item.StartStopGroup2.IsNull() && !item.StartStopGroup2.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"network"+"."+strconv.Itoa(index)+"."+"start-stop.group-config.group2.group", item.StartStopGroup2.ValueString())
			}
		}
	}
	return body
}

func (data *AAAAccounting) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "update.newinfo.periodic"); value.Exists() && !data.UpdateNewinfoPeriodic.IsNull() {
		data.UpdateNewinfoPeriodic = types.Int64Value(value.Int())
	} else {
		data.UpdateNewinfoPeriodic = types.Int64Null()
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group1.group"); value.Exists() && !data.IdentityDefaultStartStopGroup1.IsNull() {
		data.IdentityDefaultStartStopGroup1 = types.StringValue(value.String())
	} else {
		data.IdentityDefaultStartStopGroup1 = types.StringNull()
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group2.group"); value.Exists() && !data.IdentityDefaultStartStopGroup2.IsNull() {
		data.IdentityDefaultStartStopGroup2 = types.StringValue(value.String())
	} else {
		data.IdentityDefaultStartStopGroup2 = types.StringNull()
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group3.group"); value.Exists() && !data.IdentityDefaultStartStopGroup3.IsNull() {
		data.IdentityDefaultStartStopGroup3 = types.StringValue(value.String())
	} else {
		data.IdentityDefaultStartStopGroup3 = types.StringNull()
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group4.group"); value.Exists() && !data.IdentityDefaultStartStopGroup4.IsNull() {
		data.IdentityDefaultStartStopGroup4 = types.StringValue(value.String())
	} else {
		data.IdentityDefaultStartStopGroup4 = types.StringNull()
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
		if value := r.Get("start-stop.group1.group"); value.Exists() && !data.Execs[i].StartStopGroup1.IsNull() {
			data.Execs[i].StartStopGroup1 = types.StringValue(value.String())
		} else {
			data.Execs[i].StartStopGroup1 = types.StringNull()
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
		if value := r.Get("start-stop.group-config.group1.group"); value.Exists() && !data.Networks[i].StartStopGroup1.IsNull() {
			data.Networks[i].StartStopGroup1 = types.StringValue(value.String())
		} else {
			data.Networks[i].StartStopGroup1 = types.StringNull()
		}
		if value := r.Get("start-stop.group-config.group2.group"); value.Exists() && !data.Networks[i].StartStopGroup2.IsNull() {
			data.Networks[i].StartStopGroup2 = types.StringValue(value.String())
		} else {
			data.Networks[i].StartStopGroup2 = types.StringNull()
		}
	}
	if value := res.Get(prefix + "system.guarantee-first"); !data.SystemGuaranteeFirst.IsNull() {
		if value.Exists() {
			data.SystemGuaranteeFirst = types.BoolValue(value.Bool())
		}
	} else {
		data.SystemGuaranteeFirst = types.BoolNull()
	}
}

func (data *AAAAccountingData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "update.newinfo.periodic"); value.Exists() {
		data.UpdateNewinfoPeriodic = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group1.group"); value.Exists() {
		data.IdentityDefaultStartStopGroup1 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group2.group"); value.Exists() {
		data.IdentityDefaultStartStopGroup2 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group3.group"); value.Exists() {
		data.IdentityDefaultStartStopGroup3 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "identity.default.start-stop.group-config.group4.group"); value.Exists() {
		data.IdentityDefaultStartStopGroup4 = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "exec"); value.Exists() {
		data.Execs = make([]AAAAccountingExecs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAccountingExecs{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group1.group"); cValue.Exists() {
				item.StartStopGroup1 = types.StringValue(cValue.String())
			}
			data.Execs = append(data.Execs, item)
			return true
		})
	}
	if value := res.Get(prefix + "network"); value.Exists() {
		data.Networks = make([]AAAAccountingNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAAccountingNetworks{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group-config.group1.group"); cValue.Exists() {
				item.StartStopGroup1 = types.StringValue(cValue.String())
			}
			if cValue := v.Get("start-stop.group-config.group2.group"); cValue.Exists() {
				item.StartStopGroup2 = types.StringValue(cValue.String())
			}
			data.Networks = append(data.Networks, item)
			return true
		})
	}
	if value := res.Get(prefix + "system.guarantee-first"); value.Exists() {
		data.SystemGuaranteeFirst = types.BoolValue(value.Bool())
	} else {
		data.SystemGuaranteeFirst = types.BoolValue(false)
	}
}

func (data *AAAAccounting) getDeletedListItems(ctx context.Context, state AAAAccounting) []string {
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

func (data *AAAAccounting) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *AAAAccounting) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.UpdateNewinfoPeriodic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/update/newinfo/periodic", data.getPath()))
	}
	if !data.IdentityDefaultStartStopGroup1.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/default/start-stop/group-config/group1/group", data.getPath()))
	}
	if !data.IdentityDefaultStartStopGroup2.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/default/start-stop/group-config/group2/group", data.getPath()))
	}
	if !data.IdentityDefaultStartStopGroup3.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/default/start-stop/group-config/group3/group", data.getPath()))
	}
	if !data.IdentityDefaultStartStopGroup4.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/identity/default/start-stop/group-config/group4/group", data.getPath()))
	}
	for i := range data.Execs {
		keyValues := [...]string{data.Execs[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/exec=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Networks {
		keyValues := [...]string{data.Networks[i].Id.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/network=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.SystemGuaranteeFirst.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/system/guarantee-first", data.getPath()))
	}
	return deletePaths
}
