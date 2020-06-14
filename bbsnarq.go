package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
    url := "http://bbsstore-service:7002/api/dns_store"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"domain":"`+os.Args[2]+`"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
