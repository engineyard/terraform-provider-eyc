# Terraform Provider EYC


## Deployment

1) Clone this repo
2) cd to the folder & build the provider with `go build -o terraform-provider-eyc`
3) At the root path, install the provider with `make install`
4) (Temporary) Add EYC Token at line 81 of /eyc/data_source_env_var.go
5) To test fetching EYC environment variables with Terraform, run `cd /examples && terraform init && terraform apply`
6) Environment variables would be fetched.