package rhev

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
