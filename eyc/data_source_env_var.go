package eyc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnvVars() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEnvVarsRead,
		Schema: map[string]*schema.Schema{
			"environment_variables": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"application_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"environment_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"sensitive": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEnvVarsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	const url = "https://api.engineyard.com"
	// token := m

	const token = "xxx"

	// if token == "" {
	// 	err := errors.New("Missing token")
	// 	return diag.FromErr(err)
	// }

	full_url := fmt.Sprintf("%s/environment_variables", url)

	req, err := http.NewRequest("GET", full_url, nil)
	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"X-EY-TOKEN":   []string{token},
		"Accept":       []string{"application/vnd.engineyard.v3+json"},
	}
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	var env_vars map[string]interface{}
	json.Unmarshal(bodyBytes, &env_vars)

	// env_vars := make([]map[string]interface{}, 0)
	// err = json.NewDecoder(r.Body).Decode(&env_vars)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_variables", env_vars["environment_variables"]); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
