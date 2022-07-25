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