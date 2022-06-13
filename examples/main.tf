terraform {
  required_providers {
    eyc = {
      version = "0.2"
      source  = "hashicorp.com/edu/eyc"
    }
  }
}


provider "eyc" {
  # token = "xxx"
}

data "eyc_env_vars" "all" {}



output "out" {
  value = data.eyc_env_vars.all
}


# module "env_var" {
#   source = "./env_var"
# } 
