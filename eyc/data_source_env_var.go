package eyc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	eyc "github.com/engineyard/terraform-eyc-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnvVars() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEnvVarsRead,
		Schema: map[string]*schema.Schema{
			"env_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
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
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEnvVarsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*eyc.Client)

	var envID int
	envID = d.Get("env_id").(int)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var body map[string]interface{}
	var err error
	fmt.Println(strconv.Itoa(5))

	if envID > 0 {
		body, err = c.GetEnvVarsByEnv(envID)
	} else {
		body, err = c.GetEnvVars()
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_variables", body["environment_variables"]); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
