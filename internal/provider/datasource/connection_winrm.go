package datasource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = &ConnectionWinRM{}
)

func NewConnectionWinRM() datasource.DataSource {
	return &ConnectionWinRM{}
}

type ConnectionWinRM struct {
}

type ConnectionWinRMModel struct {
	Hostname types.String `tfsdk:"hostname"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (d *ConnectionWinRM) Metadata(ctx context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_connection_winrm"
}

func (d *ConnectionWinRM) Schema(ctx context.Context, req datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Required: true,
			},
			"username": schema.StringAttribute{
				Required: true,
			},
			"password": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (d *ConnectionWinRM) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	var data ConnectionWinRMModel

	// 設定値を読み込む
	res.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if res.Diagnostics.HasError() {
		return
	}

	// TODO: dummy read
	_ = data.Hostname
	_ = data.Username
	_ = data.Password

	// 反映する
	res.Diagnostics.Append(res.State.Set(ctx, &data)...)
}
