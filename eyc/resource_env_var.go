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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceEnvVarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	value := d.Get("value").(string)
	EnvID := d.Get("env_id").(int)
	AppID := d.Get("app_id").(int)

	param := eyc.EnvVarParam{
		Environment_variable: eyc.EnvVarNameValue{
			Name:  name,
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

	convertedMapData := make([]interface{}, 0)
	convertedMapData = append(convertedMapData, mapData["environment_variable"])

	d.SetId(strconv.Itoa(body["environment_variable"].ID))

	return diags
}

func resourceEnvVarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	evID, _ := strconv.Atoi(d.Id())

	body, err := c.GetEnvVarByID(evID)

	if body == nil {
		d.SetId("")
		return diags
	}

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

	convertedMapData := make([]interface{}, 0)
	convertedMapData = append(convertedMapData, mapData["environment_variable"])

	d.Set("value", mapData["environment_variable"]["value"])
	d.Set("name", mapData["environment_variable"]["name"])

	return diags
}

func resourceEnvVarUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	evID, _ := strconv.Atoi(d.Id())

	if d.HasChange("value") || d.HasChange("name") {
		name := d.Get("name").(string)
		value := d.Get("value").(string)
		EnvID := d.Get("env_id").(int)
		AppID := d.Get("app_id").(int)

		param := eyc.EnvVarParam{
			Environment_variable: eyc.EnvVarNameValue{
				Name:  name,
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

		convertedMapData := make([]interface{}, 0)
		convertedMapData = append(convertedMapData, mapData["environment_variable"])

		d.Set("last_updated", time.Now().Format(time.RFC850))
		d.Set("value", mapData["environment_variable"]["value"])
		d.Set("name", mapData["environment_variable"]["name"])
	}

	return resourceEnvVarRead(ctx, d, m)
}

func resourceEnvVarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*eyc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	evID, _ := strconv.Atoi(d.Id())

	c.DeleteEnvVar(evID)

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
