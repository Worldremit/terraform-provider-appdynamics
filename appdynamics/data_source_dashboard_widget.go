package appdynamics

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/guregu/null.v4"
)

func dataSourceDashboardWidget() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDashboardWidgetRead,
		Schema: map[string]*schema.Schema{
			"json": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"widget_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDashboardWidgetRead(d *schema.ResourceData, meta interface{}) error {
	/*
			Generate nullable json structure to be able to detect changes inside and outside of terraform of the widget.
		    Also provider GUID to fulfill AppD requirements.
	*/
	jsonSource := d.Get("json").(string)
	hash := sha256.Sum224([]byte(jsonSource))
	hashString := "wr-" + hex.EncodeToString(hash[:])
	d.SetId(hashString)
	widget := client.DashboardWidget{}
	json.Unmarshal([]byte(jsonSource), &widget)
	widget.GUID = null.NewString(hashString[0:50], true)
	jsonDoc, err := json.Marshal(widget)
	if err != nil {
		return err
	}
	jsonString := string(jsonDoc)
	d.Set("widget_json", jsonString)
	return nil
}
