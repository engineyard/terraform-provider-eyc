terraform {
  required_providers {
    eyc = {
      version = "0.2"
      source  = "hashicorp.com/edu/eyc"
    }
  }
}

provider "eyc" {
  token = var.eyc_token
}