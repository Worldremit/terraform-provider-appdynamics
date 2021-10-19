package appdynamics

import (
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"appdynamics_health_rule":                resourceHealthRule(),
			"appdynamics_action":                     resourceAction(),
			"appdynamics_policy":                     resourcePolicy(),
			"appdynamics_transaction_detection_rule": resourceTransactionDetectionRule(),
			"appdynamics_dashboard":                  resourceDashboard(),
			"appdynamics_import_export_dashboard":    resourceImportExportDashboard(),
			"appdynamics_tier_template_association":  resourceTierTemplateAssociation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"appdynamics_dashboard_widget": dataSourceDashboardWidget(),
			"appdynamics_appd_service":     dataSourceAppdService(),
		},
		Schema: map[string]*schema.Schema{
			"secret":              {Type: schema.TypeString, Sensitive: true, Required: true},
			"controller_base_url": {Type: schema.TypeString, Required: true},
		},
		ConfigureFunc: func(data *schema.ResourceData) (interface{}, error) {
			return &client.AppDClient{
				BaseUrl: data.Get("controller_base_url").(string),
				Secret:  data.Get("secret").(string),
			}, nil
		},
	}
}
