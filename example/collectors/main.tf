
provider "appdynamics" {
  secret = var.secret
  controller_base_url = var.controller_url
}



resource appdynamics_collector test {
	name="example2"
	type="MYSQL"
	hostname="test2"
	username="user"
	password="paswd3"
	port=17
	agent_name="test"
}

resource appdynamics_collector test2 {
	name="example1"
	type="MYSQL"
	hostname="test2"
	username="u"
	password="password2"
	port=17
	agent_name="test"
}

variable secret{}

variable controller_url {}