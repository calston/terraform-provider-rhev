package api

import (
    "encoding/xml"
)

type Config struct {
    APIUrl      string
    APIUser     string
    APIPassword string
    APIPort     int
}

type RHEVAPILink struct {
    Href    string      `xml:"href,attr"`
    Name    string      `xml:"rel,attr"`
}

type RHEVAPISpecialObjects struct {
    Links   []RHEVAPILink   `xml:"link"`
}

type RHEVAPIProductInfo struct {
    Name    string      `xml:"name"`
    Vendor  string      `xml:"vendor"`
    Version string      `xml:"full_version"`
}

type RHEVAPIRoot struct {
    Links   []RHEVAPILink   `xml:"link"`
    SpecialObjects  RHEVAPISpecialObjects   `xml:"special_objects"`
    ProductInfo     RHEVAPIProductInfo      `xml:"product_info"`
}

type RHEVAPIBootDevice struct {
    Dev         string      `xml:"dev,attr"`
}

type RHEVAPICpuTopology struct {
    Cores       string      `xml:"cores,attr"`
    Sockets     string      `xml:"sockets,attr"`
}

type RHEVVM struct {
    XMLName     xml.Name            `xml:"vm"`
    Name        string              `xml:"name"`
    Cluster     string              `xml:"cluster>name"`
    Template    string              `xml:"template>name"`
    BootDevice  RHEVAPIBootDevice   `xml:"os>boot>dev"`
    Type        string              `xml:"type"`
    CPUTopology RHEVAPICpuTopology  `xml:"cpu>topology"`
    Memory      int                 `xml:"memory"`
}

type RHEVAPIVMCreateResponse struct {
    UUID          string      `xml:"vm>href,attr"`
}


