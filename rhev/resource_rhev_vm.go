package rhev

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/calston/terraform-provider-rhev/rhevapi"
    "log"
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
            "default_user": &schema.Schema{
                Type: schema.TypeString,
                Optional: true,
                ForceNew: true,
                Default: "terraform",
            },
            "memory": &schema.Schema{
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
                Default: 1024,
            },
            "cpu_cores": &schema.Schema{
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
                Default: 2,
            },
            "cpu_sockets": &schema.Schema{
                Type: schema.TypeInt,
                Optional: true,
                ForceNew: true,
                Default: 1,
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
    config := meta.(*rhevapi.Config)

    vm := &rhevapi.RHEVVM{
        Name:           d.Get("name").(string),
        Template:       d.Get("template").(string),
        Memory:         d.Get("memory").(int) * 1024 * 1024,
        CPUTopology:    rhevapi.RHEVCpuTopology{
            Cores:      string(d.Get("cpu_cores").(int)),
            Sockets:    string(d.Get("cpu_sockets").(int)),
        },
        BootDevice:     rhevapi.RHEVBootDevice{Dev: "hd",},
        Cluster:        "default",
        Type:           "server",
        CloudInit:      rhevapi.RHEVCloudInit{
            Hostname:   d.Get("name").(string),
            PublicKey:  d.Get("public_key").(string),
            Username:   d.Get("default_user").(string),
        },
    }

    response, err := rhevapi.Post("/vms", vm, *config)

    if err != nil {
        log.Fatal(err)
    }

    if response != nil {}

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
