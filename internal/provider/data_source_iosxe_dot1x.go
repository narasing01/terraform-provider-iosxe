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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &Dot1xDataSource{}
	_ datasource.DataSourceWithConfigure = &Dot1xDataSource{}
)

func NewDot1xDataSource() datasource.DataSource {
	return &Dot1xDataSource{}
}

type Dot1xDataSource struct {
	clients map[string]*restconf.Client
}

func (d *Dot1xDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dot1x"
}

func (d *Dot1xDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Dot1x configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"auth_fail_eapol": schema.BoolAttribute{
				MarkdownDescription: "Send EAPOL-Success on successful auth-fail Authorization",
				Computed:            true,
			},
			"credentials": schema.ListNestedAttribute{
				MarkdownDescription: "Configure 802.1X credentials profiles",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"profile_name": schema.StringAttribute{
							MarkdownDescription: "Specify a profile name",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Provide a description for the credentials profile",
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Set the authentication userid",
							Computed:            true,
						},
						"password_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"pki_trustpoint": schema.StringAttribute{
							MarkdownDescription: "Set the default pki trustpoint",
							Computed:            true,
						},
						"anonymous_id": schema.StringAttribute{
							MarkdownDescription: "Set the anonymous userid",
							Computed:            true,
						},
					},
				},
			},
			"critical_eapol_config_block": schema.BoolAttribute{
				MarkdownDescription: "Block all EAPoL transaction on Critical Authentication",
				Computed:            true,
			},
			"critical_recovery_delay": schema.Int64Attribute{
				MarkdownDescription: "Set 802.1x Critical Authentication Recovery Delay period",
				Computed:            true,
			},
			"test_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout for device EAPOL capabilities test in seconds",
				Computed:            true,
			},
			"logging_verbose": schema.BoolAttribute{
				MarkdownDescription: "Show verbose messages in system logs",
				Computed:            true,
			},
			"supplicant_controlled_transient": schema.BoolAttribute{
				MarkdownDescription: "Controlled access is only applied during authentication",
				Computed:            true,
			},
			"supplicant_force_multicast": schema.BoolAttribute{
				MarkdownDescription: "Force 802.1X supplicant to send multicast packets",
				Computed:            true,
			},
			"system_auth_control": schema.BoolAttribute{
				MarkdownDescription: "Enable or Disable SysAuthControl",
				Computed:            true,
			},
		},
	}
}

func (d *Dot1xDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *Dot1xDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Dot1xData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = Dot1xData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}