package eyc

import (
	"context"

	eyc "github.com/engineyard/terraform-eyc-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"EYC_TOKEN",
					"EYC_ACCESS_TOKEN",
				}, nil),
				Description: "The token key for API operations.",
			},
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EYC_API_URL", "https://api.engineyard.com"),
				Description: "The URL to use for the EYC API.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"eyc_env_var": resourceEnvVar(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"eyc_env_vars": dataSourceEnvVars(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	token := d.Get("token").(string)
	APIEndpoint := d.Get("api_endpoint").(string)

	c, err := eyc.NewClient(&APIEndpoint, &token)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return c, diags
}
