package eyc

import (
	"context"
	"encoding/json"
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
			"env_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"app_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_variable": &schema.Schema{
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
	value := d.Get("value").(string)
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

	body, err := c.CreateEnvVar(param)
	if err != nil {
		return diag.FromErr(err)
	}
	// Convert map to json string
	jsonStr, _ := json.Marshal(body)

	// Convert struct
	var mapData map[string]map[string]interface{}
	if err := json.Unmarshal(jsonStr, &mapData); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_variable", mapData["environment_variable"]); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(body["environment_variable"].ID))

	return diags
}

func resourceEnvVarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	evID, _ := strconv.Atoi(d.Id())

	// envID, hasEnvID := d.Get("env_id").(int)
	// appID := strconv.Itoa(d.Get("app_id").(int))

	var body map[string]eyc.EnvVar
	var err error

	body, err = c.GetEnvVarByID(evID)

	// fmt.Printf("appID: %v\n", appID)

	if err != nil {
		return diag.FromErr(err)
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(body)

	// Convert struct
	var mapData map[string]map[string]interface{}
	if err := json.Unmarshal(jsonStr, &mapData); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("environment_variable", mapData["environment_variable"]); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEnvVarUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	evID, _ := strconv.Atoi(d.Id())

	if d.HasChange("environment_variable") {
		key := d.Get("key").(string)
		value := d.Get("value").(string)
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

		body, err := c.UpdateEnvVar(param, evID)
		if err != nil {
			return diag.FromErr(err)
		}

		// Convert map to json string
		jsonStr, err := json.Marshal(body)

		// Convert struct
		var mapData map[string]map[string]interface{}
		if err := json.Unmarshal(jsonStr, &mapData); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("environment_variable", mapData["environment_variable"]); err != nil {
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

	evID, _ := strconv.Atoi(d.Id())

	_, err := c.DeleteEnvVar(evID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
