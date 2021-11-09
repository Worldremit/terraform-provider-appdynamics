package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccAppDCollector_Create(t *testing.T) {

	fmt.Println("TestAccAppDCollector_Create test started")

	resourceName := "appdynamics_collector.test"
	collectorName := "testAutomationCreate"
	tf := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), collectorName)
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: tf,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", collectorName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckCollectorDoesNotExist(resourceName),
	})

}

func TestAccAppDCollector_Update(t *testing.T) {

	fmt.Println("TestAccAppDCollector_Update test started")

	resourceName := "appdynamics_collector.test"
	collectorName := "testAutomationUpdate1"
	updatedName := "testAutomationUpdate2"
	tf := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), collectorName)
	tfUpdated := configureCollectorTest(os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), updatedName)
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: tf,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", collectorName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
			{
				Config: tfUpdated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "type", "MYSQL"),
					resource.TestCheckResourceAttr(resourceName, "hostname", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					CheckCollectorExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckCollectorDoesNotExist(resourceName),
	})
}

func configureCollectorTest(appdSecret string, appdControllerUrl string, collectorName string) string {
	tf := fmt.Sprintf(`
provider "appdynamics" {
  secret = "%s"
  controller_base_url = "%s"
}

resource appdynamics_collector test {
	name="%s"
	type="MYSQL"
	hostname="test"
	username="user"
	password="paswd"
	port=3306
	agent_name="test"
}`, appdSecret, appdControllerUrl, collectorName)

	return tf
}
