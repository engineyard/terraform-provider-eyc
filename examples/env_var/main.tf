terraform {
  required_providers {
    eyc = {
      version = "0.2"
      source  = "hashicorp.com/edu/eyc"
    }
  }
}

data "eyc_env_vars" "all" {}

# Returns all env_varss
output "all_env_vars" {
  value = data.eyc_env_vars.all.environment_variables
}
