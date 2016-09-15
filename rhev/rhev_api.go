package rhev

import (
    "errors"
    "log"
    "fmt"
    "encoding/xml"
    "encoding/base64"
    "net/http"
    "crypto/tls"
)

func createUrl(path string, url string, port int) string {
    return fmt.Sprintf("https://%s:%d/api/%s", url, port, path)
}

func doAPICall(path string, config Config) *http.Response {
    transport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: transport}

    req, err := http.NewRequest("GET", createUrl(path, config.APIUrl, config.APIPort), nil)

    up := fmt.Sprintf("%s:%s", config.APIUser, config.APIPassword)

    auth_b64 := base64.StdEncoding.EncodeToString([]byte(up))

    req.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth_b64))

    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    return resp
}

func getAPIRoot(config Config) (RHEVAPIRoot, error) {
    res := doAPICall("", config)

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

