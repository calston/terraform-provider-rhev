package rhev

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
    "github.com/calston/terraform-provider-rhev/rhevapi"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "api_url": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "api_user": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "api_password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "rhev_vm": resourceRHEVVm(),
        },
        ConfigureFunc: configureProvider,
    }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
    config := rhevapi.Config{
        Url:      d.Get("api_url").(string),
        User:     d.Get("api_user").(string),
        Password: d.Get("api_password").(string),
    }

    return &config, nil
}
