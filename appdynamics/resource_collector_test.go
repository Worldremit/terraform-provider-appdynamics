package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccAppDCollector_Create(t *testing.T) {

	fmt.Println("The string is")

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},

		Steps: []resource.TestStep{
			{
				Config: `

resource collector test {
	name="testAutomation"
	type="MYSQL"
	hostname="test"
	port="3306"
	agentName="test"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
				//					resource.TestCheckResourceAttr(resourceName, "phone_number", phoneNumber),
				//					resource.TestCheckResourceAttr(resourceName, "action_type", "SMS"),
				//					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
				//					resource.TestCheckResourceAttrSet(resourceName, "id"),
				//					CheckActionExists(resourceName),
				),
			},
		},
		CheckDestroy: CheckActionDoesNotExist(resourceName),
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
