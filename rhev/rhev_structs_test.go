package rhev


import (
    "testing"
    "os"
    "log"
    "encoding/xml"
)

func TestVMCreateStruct(t *testing.T) {
    vm := &RHEVAPIVMCreate{
        Name:        "testvm",
        Template:    "mytemplate",
        Memory:      1024 * 1024 * 1024,
        CPUTopology: RHEVAPICpuTopology{Cores:   "1", Sockets: "1",},
        Cluster:     "default",
        BootDevice:  RHEVAPIBootDevice{Dev: "hd",},
        Type:       "server",
    }

    blob, err := xml.MarshalIndent(vm, "  ", "    ")
    if err != nil {
        log.Fatal(err)
    }
    os.Stdout.Write(blob)
}
