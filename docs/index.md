---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cribl-terraform Provider"
subcategory: ""
description: |-
  Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at docs.cribl.io http://docs.cribl.io.
---

# cribl-terraform Provider

Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at [docs.cribl.io](http://docs.cribl.io).

## Example Usage

```terraform
terraform {
  required_providers {
    cribl-terraform = {
      source  = "speakeasy/cribl-terraform"
      version = "0.0.12"
    }
  }
}

provider "cribl-terraform" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `server_url` (String) Server URL

### Optional

- `bearer_auth` (String, Sensitive)
