package api

import (
    "errors"
    "log"
    "fmt"
    "encoding/xml"
    "encoding/base64"
    "net/http"
    "crypto/tls"
    "strings"
)

func createUrl(path string, url string, port int) string {
    return fmt.Sprintf("https://%s:%d/api/%s", url, port, path)
}

func createAuthorization(config Config) string {
    up := fmt.Sprintf("%s:%s", config.APIUser, config.APIPassword)
    auth_b64 := base64.StdEncoding.EncodeToString([]byte(up))
    return fmt.Sprintf("Basic %s", auth_b64)
}

func doAPICall(path string, method string, data *string, config Config) *http.Response {
    transport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: transport}

    url := createUrl(path, config.APIUrl, config.APIPort)

    var req *http.Request
    var err error

    if data != nil {
        req, err = http.NewRequest(method, url, strings.NewReader(*data))
    } else {
        req, err = http.NewRequest(method, url, nil)
    }

    req.Header.Add("Authorization", createAuthorization(config))

    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    return resp
}

func RHEVAPIGet(path string, config Config) *http.Response {
    return doAPICall(path, "GET", nil, config)
}

func RHEVAPIPost(path string, payload *string, config Config) *http.Response {
    return doAPICall(path, "POST", payload, config)
}

func getAPIRoot(config Config) (RHEVAPIRoot, error) {
    res := RHEVAPIGet("", config)

    if (res.StatusCode == 200) {
        decoder := xml.NewDecoder(res.Body)
        var root RHEVAPIRoot
        err := decoder.Decode(&root)

        if err != nil {
            log.Fatal(err)
        }

        return root, nil
    }
    return RHEVAPIRoot{}, errors.New("Bad response code")
}

func CreateVMXML(vm *RHEVVM) (string, error) {
    // Creates a VM and returns the unique ID

    blob, err := xml.MarshalIndent(*vm, "  ", "    ")

    if err != nil {
        log.Fatal(err)
        return "", err
    }

    return string(blob), nil
}
