package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/calston/terraform-provider-rhev/rhev"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: rhev.Provider,
    })

}
