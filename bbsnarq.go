package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "net/http"
)

func main() {
	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			if err != nil {
				log.Fatal(err)
			}
		}
		
        fmt.Printf("stdin = %s\n", stdin)
        url := "http://bbsstore-service:7002/api/dns_store"
        var jsonStr = []byte(`{"domain":"`+stdin+`"}`)
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
	} else {
		fmt.Println("send data via pipe")
    }
}

func mainOLD() {
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
