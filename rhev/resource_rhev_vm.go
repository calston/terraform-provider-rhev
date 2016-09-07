package rhev

import (
    "fmt"
    "io"
    "encoding/base64"
    "net/http"
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

func createUrl(path string) string {
    return Sprintf("https://%s:%d/api/%s", config.APIUrl, config.APIPort, path)
}

func doAPICall(path string) string {
    client := &http.Client{}

    resp, err := http.NewRequest("GET", createUrl(path), nil)

    auth_b64 := base64.StdEncoding.EncodeToString(
        []byte(Sprintf("%s:%s", config.APIUser, config.APIPassword))
    )

    req.Header.Add("Authorization", Sprintf("Basic %s", auth_b64))

    resp, err := client.Do(req)

    return resp
}

func resourceRHEVVmCreate(d *schema.ResourceData, meta interface[]) error {
    http.get(createUrl("")
}
