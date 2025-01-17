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

type AAA struct {
	Device                           types.String                          `tfsdk:"device"`
	Id                               types.String                          `tfsdk:"id"`
	NewModel                         types.Bool                            `tfsdk:"new_model"`
	ServerRadiusDynamicAuthor        types.Bool                            `tfsdk:"server_radius_dynamic_author"`
	SessionId                        types.String                          `tfsdk:"session_id"`
	ServerRadiusDynamicAuthorClients []AAAServerRadiusDynamicAuthorClients `tfsdk:"server_radius_dynamic_author_clients"`
	GroupServerRadius                []AAAGroupServerRadius                `tfsdk:"group_server_radius"`
	GroupTacacsplus                  []AAAGroupTacacsplus                  `tfsdk:"group_tacacsplus"`
}

type AAAData struct {
	Device                           types.String                          `tfsdk:"device"`
	Id                               types.String                          `tfsdk:"id"`
	NewModel                         types.Bool                            `tfsdk:"new_model"`
	ServerRadiusDynamicAuthor        types.Bool                            `tfsdk:"server_radius_dynamic_author"`
	SessionId                        types.String                          `tfsdk:"session_id"`
	ServerRadiusDynamicAuthorClients []AAAServerRadiusDynamicAuthorClients `tfsdk:"server_radius_dynamic_author_clients"`
	GroupServerRadius                []AAAGroupServerRadius                `tfsdk:"group_server_radius"`
	GroupTacacsplus                  []AAAGroupTacacsplus                  `tfsdk:"group_tacacsplus"`
}
type AAAServerRadiusDynamicAuthorClients struct {
	Ip            types.String `tfsdk:"ip"`
	ServerKeyType types.String `tfsdk:"server_key_type"`
	ServerKey     types.String `tfsdk:"server_key"`
}
type AAAGroupServerRadius struct {
	Name        types.String                      `tfsdk:"name"`
	ServerNames []AAAGroupServerRadiusServerNames `tfsdk:"server_names"`
}
type AAAGroupTacacsplus struct {
	Name    types.String                `tfsdk:"name"`
	Servers []AAAGroupTacacsplusServers `tfsdk:"servers"`
}
type AAAGroupServerRadiusServerNames struct {
	Name types.String `tfsdk:"name"`
}
type AAAGroupTacacsplusServers struct {
	Name types.String `tfsdk:"name"`
}

func (data AAA) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa"
}

func (data AAAData) getPath() string {
	return "Cisco-IOS-XE-native:native/aaa"
}

// if last path element has a key -> remove it
func (data AAA) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data AAA) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.NewModel.IsNull() && !data.NewModel.IsUnknown() {
		if data.NewModel.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:new-model", map[string]string{})
		}
	}
	if !data.ServerRadiusDynamicAuthor.IsNull() && !data.ServerRadiusDynamicAuthor.IsUnknown() {
		if data.ServerRadiusDynamicAuthor.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:server.radius.dynamic-author", map[string]string{})
		}
	}
	if !data.SessionId.IsNull() && !data.SessionId.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:session-id", data.SessionId.ValueString())
	}
	if len(data.ServerRadiusDynamicAuthorClients) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:server.radius.dynamic-author.client", []interface{}{})
		for index, item := range data.ServerRadiusDynamicAuthorClients {
			if !item.Ip.IsNull() && !item.Ip.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:server.radius.dynamic-author.client"+"."+strconv.Itoa(index)+"."+"ip", item.Ip.ValueString())
			}
			if !item.ServerKeyType.IsNull() && !item.ServerKeyType.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:server.radius.dynamic-author.client"+"."+strconv.Itoa(index)+"."+"server-key.key", item.ServerKeyType.ValueString())
			}
			if !item.ServerKey.IsNull() && !item.ServerKey.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:server.radius.dynamic-author.client"+"."+strconv.Itoa(index)+"."+"server-key.string", item.ServerKey.ValueString())
			}
		}
	}
	if len(data.GroupServerRadius) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.radius", []interface{}{})
		for index, item := range data.GroupServerRadius {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.radius"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if len(item.ServerNames) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.radius"+"."+strconv.Itoa(index)+"."+"server.name", []interface{}{})
				for cindex, citem := range item.ServerNames {
					if !citem.Name.IsNull() && !citem.Name.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.radius"+"."+strconv.Itoa(index)+"."+"server.name"+"."+strconv.Itoa(cindex)+"."+"name", citem.Name.ValueString())
					}
				}
			}
		}
	}
	if len(data.GroupTacacsplus) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.tacacsplus", []interface{}{})
		for index, item := range data.GroupTacacsplus {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.tacacsplus"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
			if len(item.Servers) > 0 {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.tacacsplus"+"."+strconv.Itoa(index)+"."+"server.name", []interface{}{})
				for cindex, citem := range item.Servers {
					if !citem.Name.IsNull() && !citem.Name.IsUnknown() {
						body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"Cisco-IOS-XE-aaa:group.server.tacacsplus"+"."+strconv.Itoa(index)+"."+"server.name"+"."+strconv.Itoa(cindex)+"."+"name", citem.Name.ValueString())
					}
				}
			}
		}
	}
	return body
}

func (data *AAA) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:new-model"); !data.NewModel.IsNull() {
		if value.Exists() {
			data.NewModel = types.BoolValue(true)
		} else {
			data.NewModel = types.BoolValue(false)
		}
	} else {
		data.NewModel = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:server.radius.dynamic-author"); !data.ServerRadiusDynamicAuthor.IsNull() {
		if value.Exists() {
			data.ServerRadiusDynamicAuthor = types.BoolValue(true)
		} else {
			data.ServerRadiusDynamicAuthor = types.BoolValue(false)
		}
	} else {
		data.ServerRadiusDynamicAuthor = types.BoolNull()
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:session-id"); value.Exists() && !data.SessionId.IsNull() {
		data.SessionId = types.StringValue(value.String())
	} else {
		data.SessionId = types.StringNull()
	}
	for i := range data.ServerRadiusDynamicAuthorClients {
		keys := [...]string{"ip"}
		keyValues := [...]string{data.ServerRadiusDynamicAuthorClients[i].Ip.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-aaa:server.radius.dynamic-author.client").ForEach(
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
		if value := r.Get("ip"); value.Exists() && !data.ServerRadiusDynamicAuthorClients[i].Ip.IsNull() {
			data.ServerRadiusDynamicAuthorClients[i].Ip = types.StringValue(value.String())
		} else {
			data.ServerRadiusDynamicAuthorClients[i].Ip = types.StringNull()
		}
		if value := r.Get("server-key.key"); value.Exists() && !data.ServerRadiusDynamicAuthorClients[i].ServerKeyType.IsNull() {
			data.ServerRadiusDynamicAuthorClients[i].ServerKeyType = types.StringValue(value.String())
		} else {
			data.ServerRadiusDynamicAuthorClients[i].ServerKeyType = types.StringNull()
		}
		if value := r.Get("server-key.string"); value.Exists() && !data.ServerRadiusDynamicAuthorClients[i].ServerKey.IsNull() {
			data.ServerRadiusDynamicAuthorClients[i].ServerKey = types.StringValue(value.String())
		} else {
			data.ServerRadiusDynamicAuthorClients[i].ServerKey = types.StringNull()
		}
	}
	for i := range data.GroupServerRadius {
		keys := [...]string{"name"}
		keyValues := [...]string{data.GroupServerRadius[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-aaa:group.server.radius").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.GroupServerRadius[i].Name.IsNull() {
			data.GroupServerRadius[i].Name = types.StringValue(value.String())
		} else {
			data.GroupServerRadius[i].Name = types.StringNull()
		}
		for ci := range data.GroupServerRadius[i].ServerNames {
			keys := [...]string{"name"}
			keyValues := [...]string{data.GroupServerRadius[i].ServerNames[ci].Name.ValueString()}

			var cr gjson.Result
			r.Get("server.name").ForEach(
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
			if value := cr.Get("name"); value.Exists() && !data.GroupServerRadius[i].ServerNames[ci].Name.IsNull() {
				data.GroupServerRadius[i].ServerNames[ci].Name = types.StringValue(value.String())
			} else {
				data.GroupServerRadius[i].ServerNames[ci].Name = types.StringNull()
			}
		}
	}
	for i := range data.GroupTacacsplus {
		keys := [...]string{"name"}
		keyValues := [...]string{data.GroupTacacsplus[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "Cisco-IOS-XE-aaa:group.server.tacacsplus").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.GroupTacacsplus[i].Name.IsNull() {
			data.GroupTacacsplus[i].Name = types.StringValue(value.String())
		} else {
			data.GroupTacacsplus[i].Name = types.StringNull()
		}
		for ci := range data.GroupTacacsplus[i].Servers {
			keys := [...]string{"name"}
			keyValues := [...]string{data.GroupTacacsplus[i].Servers[ci].Name.ValueString()}

			var cr gjson.Result
			r.Get("server.name").ForEach(
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
			if value := cr.Get("name"); value.Exists() && !data.GroupTacacsplus[i].Servers[ci].Name.IsNull() {
				data.GroupTacacsplus[i].Servers[ci].Name = types.StringValue(value.String())
			} else {
				data.GroupTacacsplus[i].Servers[ci].Name = types.StringNull()
			}
		}
	}
}

func (data *AAAData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:new-model"); value.Exists() {
		data.NewModel = types.BoolValue(true)
	} else {
		data.NewModel = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:server.radius.dynamic-author"); value.Exists() {
		data.ServerRadiusDynamicAuthor = types.BoolValue(true)
	} else {
		data.ServerRadiusDynamicAuthor = types.BoolValue(false)
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:session-id"); value.Exists() {
		data.SessionId = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:server.radius.dynamic-author.client"); value.Exists() {
		data.ServerRadiusDynamicAuthorClients = make([]AAAServerRadiusDynamicAuthorClients, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAServerRadiusDynamicAuthorClients{}
			if cValue := v.Get("ip"); cValue.Exists() {
				item.Ip = types.StringValue(cValue.String())
			}
			if cValue := v.Get("server-key.key"); cValue.Exists() {
				item.ServerKeyType = types.StringValue(cValue.String())
			}
			if cValue := v.Get("server-key.string"); cValue.Exists() {
				item.ServerKey = types.StringValue(cValue.String())
			}
			data.ServerRadiusDynamicAuthorClients = append(data.ServerRadiusDynamicAuthorClients, item)
			return true
		})
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:group.server.radius"); value.Exists() {
		data.GroupServerRadius = make([]AAAGroupServerRadius, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAGroupServerRadius{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("server.name"); cValue.Exists() {
				item.ServerNames = make([]AAAGroupServerRadiusServerNames, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AAAGroupServerRadiusServerNames{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					}
					item.ServerNames = append(item.ServerNames, cItem)
					return true
				})
			}
			data.GroupServerRadius = append(data.GroupServerRadius, item)
			return true
		})
	}
	if value := res.Get(prefix + "Cisco-IOS-XE-aaa:group.server.tacacsplus"); value.Exists() {
		data.GroupTacacsplus = make([]AAAGroupTacacsplus, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AAAGroupTacacsplus{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			if cValue := v.Get("server.name"); cValue.Exists() {
				item.Servers = make([]AAAGroupTacacsplusServers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AAAGroupTacacsplusServers{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					}
					item.Servers = append(item.Servers, cItem)
					return true
				})
			}
			data.GroupTacacsplus = append(data.GroupTacacsplus, item)
			return true
		})
	}
}

func (data *AAA) getDeletedListItems(ctx context.Context, state AAA) []string {
	deletedListItems := make([]string, 0)
	for i := range state.ServerRadiusDynamicAuthorClients {
		stateKeyValues := [...]string{state.ServerRadiusDynamicAuthorClients[i].Ip.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.ServerRadiusDynamicAuthorClients[i].Ip.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.ServerRadiusDynamicAuthorClients {
			found = true
			if state.ServerRadiusDynamicAuthorClients[i].Ip.ValueString() != data.ServerRadiusDynamicAuthorClients[j].Ip.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:server/radius/dynamic-author/client=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.GroupServerRadius {
		stateKeyValues := [...]string{state.GroupServerRadius[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.GroupServerRadius[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.GroupServerRadius {
			found = true
			if state.GroupServerRadius[i].Name.ValueString() != data.GroupServerRadius[j].Name.ValueString() {
				found = false
			}
			if found {
				for ci := range state.GroupServerRadius[i].ServerNames {
					cstateKeyValues := [...]string{state.GroupServerRadius[i].ServerNames[ci].Name.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.GroupServerRadius[i].ServerNames[ci].Name.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.GroupServerRadius[j].ServerNames {
						found = true
						if state.GroupServerRadius[i].ServerNames[ci].Name.ValueString() != data.GroupServerRadius[j].ServerNames[cj].Name.ValueString() {
							found = false
						}
						if found {
							break
						}
					}
					if !found {
						deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/radius=%v/server/name=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/radius=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.GroupTacacsplus {
		stateKeyValues := [...]string{state.GroupTacacsplus[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.GroupTacacsplus[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.GroupTacacsplus {
			found = true
			if state.GroupTacacsplus[i].Name.ValueString() != data.GroupTacacsplus[j].Name.ValueString() {
				found = false
			}
			if found {
				for ci := range state.GroupTacacsplus[i].Servers {
					cstateKeyValues := [...]string{state.GroupTacacsplus[i].Servers[ci].Name.ValueString()}

					cemptyKeys := true
					if !reflect.ValueOf(state.GroupTacacsplus[i].Servers[ci].Name.ValueString()).IsZero() {
						cemptyKeys = false
					}
					if cemptyKeys {
						continue
					}

					found := false
					for cj := range data.GroupTacacsplus[j].Servers {
						found = true
						if state.GroupTacacsplus[i].Servers[ci].Name.ValueString() != data.GroupTacacsplus[j].Servers[cj].Name.ValueString() {
							found = false
						}
						if found {
							break
						}
					}
					if !found {
						deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/tacacsplus=%v/server/name=%v", state.getPath(), strings.Join(stateKeyValues[:], ","), strings.Join(cstateKeyValues[:], ",")))
					}
				}
				break
			}
		}
		if !found {
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/tacacsplus=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedListItems
}

func (data *AAA) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.NewModel.IsNull() && !data.NewModel.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:new-model", data.getPath()))
	}
	if !data.ServerRadiusDynamicAuthor.IsNull() && !data.ServerRadiusDynamicAuthor.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:server/radius/dynamic-author", data.getPath()))
	}

	return emptyLeafsDelete
}

func (data *AAA) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.NewModel.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:new-model", data.getPath()))
	}
	if !data.ServerRadiusDynamicAuthor.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:server/radius/dynamic-author", data.getPath()))
	}
	if !data.SessionId.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:session-id", data.getPath()))
	}
	for i := range data.ServerRadiusDynamicAuthorClients {
		keyValues := [...]string{data.ServerRadiusDynamicAuthorClients[i].Ip.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:server/radius/dynamic-author/client=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.GroupServerRadius {
		keyValues := [...]string{data.GroupServerRadius[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/radius=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.GroupTacacsplus {
		keyValues := [...]string{data.GroupTacacsplus[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/Cisco-IOS-XE-aaa:group/server/tacacsplus=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
