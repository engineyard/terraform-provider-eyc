# Terraform Provider EYC


## Deployment

1) Clone this repo
2) At the root path, install the provider with `make install`, which would install the plugin locally at `~/.terraform.d/plugins/` path which you could then setup provider in Terraform with 

```
terraform {
  required_providers {
    eyc = {
      version = "0.1"
      source  = "engineyard/terraform/eyc"
    }
  }
}

provider "eyc" {
  token = var.eyc_token
}
```

note: `eyc_token` could be retrieved from https://cloud.engineyard.com/cli

3) Change directory to `examples` folder, and change `terraform.tfvars.example` to `terraform.tfvars`
4) (Under `examples` folder) Input corresponding variables under `terraform.tfvars` file
5) (Under `examples` folder) Run `terraform init && terraform apply`
6) You could then visit the corresponding EYC environment to observe the changes.