package rhevapi

import (
    "encoding/xml"
)

type Config struct {
    Url      string
    User     string
    Password string
}

type RHEVLink struct {
    Href    string      `xml:"href,attr"`
    Name    string      `xml:"rel,attr"`
}

type RHEVSpecialObjects struct {
    Links   []RHEVLink  `xml:"link"`
}

type RHEVProductInfo struct {
    Name    string      `xml:"name"`
    Vendor  string      `xml:"vendor"`
    Version string      `xml:"full_version"`
}

type RHEVRoot struct {
    Links           []RHEVLink          `xml:"link"`
    SpecialObjects  RHEVSpecialObjects  `xml:"special_objects"`
    ProductInfo     RHEVProductInfo     `xml:"product_info"`
}

type RHEVBootDevice struct {
    Dev         string      `xml:"dev,attr"`
}

type RHEVCpuTopology struct {
    Cores       string      `xml:"cores,attr"`
    Sockets     string      `xml:"sockets,attr"`
}

type RHEVCloudInit struct {
    Hostname    string      `xml:"host>address"`
    PublicKey   string      `xml:"authorized_ssh_keys"`
    Username    string      `xml:"user_name"`
    Password    string      `xml:"root_password"`
}

type RHEVVM struct {
    XMLName     xml.Name            `xml:"vm"`
    Name        string              `xml:"name"`
    Cluster     string              `xml:"cluster>name"`
    Template    string              `xml:"template>name"`
    BootDevice  RHEVBootDevice      `xml:"os>boot>dev"`
    Type        string              `xml:"type"`
    CPUTopology RHEVCpuTopology     `xml:"cpu>topology"`
    Memory      int                 `xml:"memory"`
    CloudInit   RHEVCloudInit       `xml:"initialization>cloud_init"`
}

type RHEVVMResponse struct {
    UUID        string      `xml:"id,attr"`
    Status      string      `xml:"status>state"`
    Name        string      `xml:"name"`
}


