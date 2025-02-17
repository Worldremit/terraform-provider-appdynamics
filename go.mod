module github.com/Worldremit/terraform-provider-appdynamics

go 1.16

require (
	github.com/HarryEMartland/terraform-provider-appdynamics v0.1.0
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/imroc/req v0.3.0
	github.com/joho/godotenv v1.4.0
	github.com/stretchr/testify v1.5.1
)

replace github.com/HarryEMartland/terraform-provider-appdynamics => ./
