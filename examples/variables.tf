variable "eyc_token" {
  type      = string
  sensitive = true
}

variable "env_id" {
  type = number
}

variable "env_var_of_env1" {
  type = list(any)
}