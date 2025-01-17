---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_aaa_authorization Resource - terraform-provider-iosxe"
subcategory: "AAA"
description: |-
  This resource can manage the AAA Authorization configuration.
---

# iosxe_aaa_authorization (Resource)

This resource can manage the AAA Authorization configuration.

## Example Usage

```terraform
resource "iosxe_aaa_authorization" "example" {
  execs = [
    {
      name                = "TEST"
      a1_local            = false
      a1_if_authenticated = true
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `execs` (Attributes List) For starting an exec (shell). (see [below for nested schema](#nestedatt--execs))
- `networks` (Attributes List) For network services. (PPP, SLIP, ARAP) (see [below for nested schema](#nestedatt--networks))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--execs"></a>
### Nested Schema for `execs`

Required:

- `name` (String)

Optional:

- `a1_if_authenticated` (Boolean) Succeed if user has authenticated.
- `a1_local` (Boolean) Use local database


<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Required:

- `id` (String)

Optional:

- `a1_group` (String) Use Server-group

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_aaa_authorization.example "Cisco-IOS-XE-native:native/aaa/Cisco-IOS-XE-aaa:authorization"
```
