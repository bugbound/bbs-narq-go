package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "net/http"
    "bytes"
    "io/ioutil"
)


func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        var linetext = s.Text()
        log.Println("line", linetext)
        
        if os.Args[1] == "domain" {
            adddomain(linetext)
        }
        if os.Args[1] == "url" {
            addurl(linetext)
        }        
    }
}


func addurl(domain string) {
    url := "http://bbsstore-service:7002/api/url_store"
    var jsonStrStart = []byte(`{"url":"`)
    var jsonStrEnd = []byte(`"}`)
    var part1 = append(jsonStrStart, domain...)
    var completeValue = append(part1, jsonStrEnd...)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(completeValue))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}

func adddomain(domain string) {
    url := "http://bbsstore-service:7002/api/dns_store"
    var jsonStrStart = []byte(`{"domain":"`)
    var jsonStrEnd = []byte(`"}`)
    var part1 = append(jsonStrStart, domain...)
    var completeValue = append(part1, jsonStrEnd...)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(completeValue))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func mainOLD() {
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
        var jsonStrStart = []byte(`{"domain":"`)
        var jsonStrEnd = []byte(`"}`)
        var part1 = append(jsonStrStart, stdin...)
        var completeValue = append(part1, jsonStrEnd...)
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(completeValue))
        req.Header.Set("X-Custom-Header", "myvalue")
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        //fmt.Println("response Headers:", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))
	} else {
		fmt.Println("send data via pipe")
    }
}
