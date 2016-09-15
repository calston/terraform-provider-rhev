package rhev

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceRHEVVm() *schema.Resource {
    return &schema.Resource{
        Create: resourceRHEVVmCreate,
        Read:   resourceRHEVVmRead,
        Update: nil,
        Delete: resourceRHEVVmDelete,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "template": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "instance_type": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "default_user": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                Default: "terraform",
            },
            "public_key": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
            },
            "network_interface": &schema.Schema{
                Type: schema.TypeList,
                Required: true,
                ForceNew: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "profile": &schema.Schema{Type: schema.TypeString, Required: true},
                    },
                },
            },
            "volumes": &schema.Schema{
                Type: schema.TypeList,
                Optional: true,
                ForceNew: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "size": &schema.Schema{Type: schema.TypeInt, Required: true},
                        "interface": &schema.Schema{Type: schema.TypeString, Optional: true, Default: "VirtIO", ForceNew: true},
                        "domain": &schema.Schema{Type: schema.TypeString, Required: true},
                        "thin": &schema.Schema{Type: schema.TypeBool, Optional: true, Default: false, ForceNew: true},
                        "profile": &schema.Schema{Type: schema.TypeString, Required: true},
                    },
                },
            },
        },
    }
}



func resourceRHEVVmCreate(d *schema.ResourceData, meta interface{}) error {
    //config := meta.(*Config)
    //doAPICall("", config.)
    return nil
}

func resourceRHEVVmRead(d *schema.ResourceData, meta interface{}) error {
    //config := meta.(*Config)
    //doAPICall("", config)
    return nil
}

func resourceRHEVVmDelete(d *schema.ResourceData, meta interface{}) error {
    //config := meta.(*Config)
    //doAPICall("", config)

    return nil
}
