package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    values := url.Values{}
    values.Add("city", "400040")
    resp, err := http.Get("http://weather.livedoor.com/forecast/webservice/json/v1" + "?" + values.Encode())

    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()

    execute(resp)
}

func execute(response *http.Response) {
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}
