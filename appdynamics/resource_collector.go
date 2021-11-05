package appdynamics

import (
	"github.com/HarryEMartland/terraform-provider-appdynamics/appdynamics/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceCollectorCreate,
		Read:   resourceCollectorRead,
		//Update: resourceCollectorUpdate,
		//Delete: resourceCollectorDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateList([]string{
					"MYSQL",
					"MONGO",
				}),
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
			},
			"agentName": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCollectorCreate(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)

	collector := createCollector(d)

	id, err := appdClient.CreateCollector(&collector)
	if err != nil {
		return err
	}
	//TODO change it
	d.SetId(id)

	return nil
}

func createCollector(d *schema.ResourceData) client.Collector {

	collector := client.Collector{
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
		Hostname:  d.Get("hostname").(string),
		Port:      d.Get("port").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		AgentName: d.Get("agentName").(string),
	}
	return collector
}

func resourceCollectorRead(d *schema.ResourceData, m interface{}) error {
	appdClient := m.(*client.AppDClient)
	applicationId := d.Get("application_id").(int)
	id := d.Id()

	actionId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	action, err := appdClient.GetAction(actionId, applicationId)
	if err != nil {
		return err
	}

	updateAction(d, *action)

	return nil
}
