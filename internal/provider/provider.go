package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	idatasource "github.com/love2hina-net/terraform-provider-windows/internal/provider/datasource"
)

// インターフェースの確認
var _ provider.Provider = &WindowsProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &WindowsProvider{
			version: version,
		}
	}
}

type WindowsProvider struct {
	// dev - Development(locally)
	// test - Testing
	version string
}

func (p *WindowsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, res *provider.MetadataResponse) {
	res.TypeName = "windows"
	res.Version = p.version
}

func (p *WindowsProvider) Schema(ctx context.Context, req provider.SchemaRequest, res *provider.SchemaResponse) {
	res.Schema = schema.Schema{}
}

func (p *WindowsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {
}

func (p *WindowsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *WindowsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		idatasource.NewConnectionWinRM,
	}
}
