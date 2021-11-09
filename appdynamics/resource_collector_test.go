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
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

provider "appdynamics" {
  secret = "%s"
  controller_base_url = "%s"
}

resource appdynamics_collector test {
	name="testAutomation"
	type="MYSQL"
	hostname="test"
	username="user"
	password="paswd"
	port=3306
	agent_name="test"
}
`, os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL")),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "testAutomation"),
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

/*
func smsAction(phoneNumber string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_action" "test_sms" {
					  application_id = var.application_id
					  action_type = "SMS"
					  phone_number = "%s"
					}
`, configureConfig(), phoneNumber)
}

func configureConfig() string {
	return fmt.Sprintf(`
					provider "appdynamics" {
					  secret = "%s"
					  controller_base_url = "%s"
					}

					variable "scope_id" {
					  type = string
					  default = "%s"
					}

					variable "application_id" {
					  type = number
					  default = %s
					}`, os.Getenv("APPD_SECRET"), os.Getenv("APPD_CONTROLLER_BASE_URL"), os.Getenv("APPD_SCOPE_ID"), os.Getenv("APPD_APPLICATION_ID"))
}

*/
