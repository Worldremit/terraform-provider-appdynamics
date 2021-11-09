
provider "appdynamics" {
  secret = var.secret
  controller_base_url = var.controller_url
}

resource appdynamics_collector test {
	name="example3"
	type="MYSQL"
	hostname="test2"
	username="user"
	password="paswd3"
	port=3306
	agent_name="test"
}

variable secret{}

variable controller_url {}