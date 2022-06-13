package eyc

import (
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
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"eyc_env_vars": dataSourceEnvVars(),
		},
		// ConfigureFunc: providerConfigure,
	}
}

// type Config struct {
// 	Token       string
// 	APIEndpoint string
// }

// func providerConfigure(d *schema.ResourceData) (interface{}, error) {
// 	config := Config{
// 		Token:       d.Get("token").(string),
// 		APIEndpoint: d.Get("api_endpoint").(string),
// 	}

// 	return config
// }
