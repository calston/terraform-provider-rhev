package rhevapi


import (
    "testing"
    "log"
    "io/ioutil"
    "strings"
    "encoding/xml"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
)

func TestVMCreateStruct(t *testing.T) {
    vm := &RHEVVM{
        Name:        "testvm",
        Template:    "mytemplate",
        Memory:      1024 * 1024 * 1024,
        CPUTopology: RHEVCpuTopology{Cores:   "1", Sockets: "1",},
        Cluster:     "default",
        BootDevice:  RHEVBootDevice{Dev: "hd",},
        Type:       "server",
        CloudInit:   RHEVCloudInit{
            Hostname:   "testbox.acme.com",
            PublicKey:  "ssh-rsa BLAHBLAH",
            Username:   "steve",
            Password:   "steveshouldhaveabetterpassword",
        },
    }

    xml, err := xml.Marshal(*vm)

    if err != nil {
        log.Fatal(err)
    } else {
        var vm2 RHEVVM
        fakeResponse := http.Response{
            Status: "200 OK",
            StatusCode: 200,
            Body: ioutil.NopCloser(strings.NewReader(string(xml))),
        }
        err = httpXMLParse(&vm2, &fakeResponse)
        if err != nil { log.Fatal(err) }

        assert.Equal(t, vm2.Name, "testvm")
    }
}

func TestVMCreateAPI(t *testing.T) {
    ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(test_api_vm_createresponse))
    }))
    defer ts.Close()

    config := Config{
        Url: ts.URL,
        User: "test",
        Password: "testpass",
    }

    vm := RHEVVM{
        Name:        "testvm",
        Template:    "mytemplate",
        Memory:      1024 * 1024 * 1024,
        CPUTopology: RHEVCpuTopology{Cores:   "1", Sockets: "1",},
        Cluster:     "default",
        BootDevice:  RHEVBootDevice{Dev: "hd",},
        Type:       "server",
        CloudInit:   RHEVCloudInit{
            Hostname:   "testbox.acme.com",
            PublicKey:  "ssh-rsa BLAHBLAH",
            Username:   "steve",
            Password:   "steveshouldhaveabetterpassword",
        },
    }

    response, err := Post("/vms", &vm, config)

    if err != nil { log.Fatal(err) }

    var newvm RHEVVMResponse
    err = httpXMLParse(&newvm, response)

    if err != nil { log.Fatal(err) }

    assert.Equal(t, newvm.UUID, "6efc0cfa-8495-4a96-93e5-ee490328cf48")
}


