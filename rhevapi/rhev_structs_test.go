package api


import (
    "testing"
    "fmt"
    "log"
)

func TestVMCreateStruct(t *testing.T) {
    vm := &RHEVVM{
        Name:        "testvm",
        Template:    "mytemplate",
        Memory:      1024 * 1024 * 1024,
        CPUTopology: RHEVAPICpuTopology{Cores:   "1", Sockets: "1",},
        Cluster:     "default",
        BootDevice:  RHEVAPIBootDevice{Dev: "hd",},
        Type:       "server",
    }

    xml, err := CreateVMXML(vm)

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Printf("%v", xml)
    }
}
