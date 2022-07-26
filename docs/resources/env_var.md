---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "eyc_env_var Resource - terraform-provider-eyc"
subcategory: ""
description: |-
  
---

# eyc_env_var (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (Number)
- `env_id` (Number)
- `name` (String)
- `value` (String)

### Optional

- `last_updated` (String)

### Read-Only

- `environment_variable` (List of Object) (see [below for nested schema](#nestedatt--environment_variable))
- `id` (String) The ID of this resource.

<a id="nestedatt--environment_variable"></a>
### Nested Schema for `environment_variable`

Read-Only:

- `application` (String)
- `application_id` (Number)
- `application_name` (String)
- `environment` (String)
- `environment_id` (Number)
- `environment_name` (String)
- `id` (Number)
- `name` (String)
- `value` (String)

