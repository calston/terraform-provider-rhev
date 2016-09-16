package rhevapi

import (
    "errors"
    "fmt"
    "encoding/xml"
    "encoding/base64"
    "net/http"
    "crypto/tls"
    "strings"
)

func createUrl(path string, url string) string {
    return fmt.Sprintf("%s/%s", url, path)
}

func createAuthorization(config Config) string {
    up := fmt.Sprintf("%s:%s", config.User, config.Password)
    auth_b64 := base64.StdEncoding.EncodeToString([]byte(up))
    return fmt.Sprintf("Basic %s", auth_b64)
}

func doAPICall(path string, method string, data *string, config Config) (*http.Response, error) {
    transport := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: transport}

    url := createUrl(path, config.Url)

    var req *http.Request
    var err error

    if data != nil {
        req, err = http.NewRequest(method, url, strings.NewReader(*data))
    } else {
        req, err = http.NewRequest(method, url, nil)
    }

    req.Header.Add("Authorization", createAuthorization(config))
    req.Header.Add("Content-type", "application/xml")

    resp, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    return resp, nil
}

func Get(path string, config Config) (*http.Response, error) {
    return doAPICall(path, "GET", nil, config)
}

func Post(path string, apiobj interface{}, config Config) (*http.Response, error) {
    payload, err := xml.Marshal(apiobj)

    if err != nil {
        return nil, err
    }

    strPayload := string(payload)

    return doAPICall(path, "POST", &strPayload, config)
}

func getAPIRoot(config Config) (*RHEVRoot, error) {
    res, err := Get("", config)


    if err != nil {
        return nil, err
    }

    if (res.StatusCode == 200) {
        decoder := xml.NewDecoder(res.Body)
        var root RHEVRoot
        err := decoder.Decode(&root)

        if err != nil {
            return nil, err
        }

        return &root, nil
    }

    return nil, errors.New("Bad response code")
}

func httpXMLParse(xmlstruct interface{}, res *http.Response) error {
    if (res.StatusCode == 200) {
        decoder := xml.NewDecoder(res.Body)

        err := decoder.Decode(xmlstruct)

        if err != nil {
            return err
        }

        return nil
    }

    return errors.New(fmt.Sprintf("Bad response code %d != 200", res.StatusCode))
}
