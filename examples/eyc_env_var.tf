resource "eyc_env_var" "env1" {
  count = length(var.env_var_of_env1)

  app_id = var.env_var_of_env1[count.index]["app_id"]
  env_id = var.env_var_of_env1[count.index]["env_id"]
  name    = var.env_var_of_env1[count.index]["name"]
  value  = var.env_var_of_env1[count.index]["value"]
}