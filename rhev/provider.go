package rhev

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
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
                Default:  "admin",
            },
            "api_password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
                Default:  "admin",
            },
            "api_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default:  443,
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "rhev_vm": resourceRHEVVm(),
        },
        ConfigureFunc: configureProvider,
    }
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        APIUrl:      d.Get("api_url").(string),
        APIUser:     d.Get("api_user").(string),
        APIPassword: d.Get("api_password").(string),
        APIPort:     d.Get("api_port").(string),
    }

    return &config, nil
}
