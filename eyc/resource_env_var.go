package eyc

import (
	"context"
	"strconv"
	"time"

	eyc "github.com/engineyard/terraform-eyc-sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnvVar() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnvVarCreate,
		ReadContext:   resourceEnvVarRead,
		UpdateContext: resourceEnvVarUpdate,
		DeleteContext: resourceEnvVarDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"environment_variable": &schema.Schema{
				Type:     schema.TypeMap,
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
			}},
	}
}

func resourceEnvVarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	value := d.Get("key").(string)
	EnvID := d.Get("env_id").(int)
	AppID := d.Get("app_id").(int)

	param := eyc.EnvVarParam{
		Environment_variable: eyc.EnvVarNameValue{
			Name:  key,
			Value: value,
		},
		Application_id: AppID,
		Environment_id: EnvID,
	}

	o, err := c.CreateEnvVar(param)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(o.ID))

	return diags
}

func resourceEnvVarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	evID := d.Id()

	// envID, hasEnvID := d.Get("env_id").(int)
	// appID := strconv.Itoa(d.Get("app_id").(int))

	var body map[string]interface{}
	var err error

	body, err = c.GetEnvVarByID(evID)

	// fmt.Printf("appID: %v\n", appID)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_variable", body["environment_variable"]); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEnvVarUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	evID := d.Id()

	if d.HasChange("environment_variable") {
		environment_variable := d.Get("environment_variable").(interface{})

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Inspecting returned env var",
			Detail:   string(environment_variable),
		})

		return nil, diags

		_, err := c.UpdateEnvVar(environment_variable, evID)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceEnvVarRead(ctx, d, m)
}

func resourceEnvVarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	evID := d.Id()

	err := c.DeleteEnvVar(evID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
