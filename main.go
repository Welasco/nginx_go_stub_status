package main

import (
	"fmt"
	"time"

	"github.com/Welasco/nginx_go_stub_status/collectors"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))

	var logcollector []collectors.LogCollector

	logcollector = append(logcollector, collectors.Newnginx_log())

	for _, collector := range logcollector {
		go collector.Start()
	}

	fmt.Println("Sleeping for 50 seconds before stopping the collector")
	time.Sleep(50 * time.Second)

	for _, collector := range logcollector {
		collector.Stop()
	}
	fmt.Println("Sleeping for 50 seconds after stop the collector")
	time.Sleep(50 * time.Second)

	// Add config file for all collectors
	// Config file must support the specifics of each collector, URL, File, etc
	// Error handling everywhere
	// Add a test module
	// Add a log module
	// check if we must need a type.go file for each collector

	////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////////////////////////////////////////////
	// Temp code for reference only
	// req, err := http.NewRequest("GET", "http://localhost/nginx_status", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)
	// bodystr := string(body)
	// fmt.Println(bodystr)

	// scanner := bufio.NewScanner(strings.NewReader(bodystr))
	// var sahr bool = false
	// for scanner.Scan() {
	// 	// Grab Active connections
	// 	if strings.Contains(scanner.Text(), "Active connections") {
	// 		fmt.Println(scanner.Text())
	// 		txt := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
	// 		fmt.Println(txt)
	// 	}

	// 	// Grab Server accepts handled requests
	// 	if sahr {
	// 		fmt.Println(scanner.Text())
	// 		txt := strings.Split(scanner.Text(), " ")

	// 		fmt.Println(len(txt))
	// 		for i := 0; i < len(txt); i++ {
	// 			fmt.Println("i=", i, " ", txt[i])
	// 		}
	// 		fmt.Println(txt)
	// 		sahr = false
	// 	}
	// 	if strings.Contains(scanner.Text(), "server accepts handled requests") {
	// 		sahr = true
	// 		fmt.Println(scanner.Text())
	// 	}

	// 	// Grab Reading: Writing: Waiting:
	// 	if strings.Contains(scanner.Text(), "Reading") {
	// 		txt := strings.Split(scanner.Text(), " ")
	// 		for i := 0; i < len(txt); i++ {
	// 			fmt.Println("i=", i, " ", txt[i])
	// 		}
	// 	}
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }
	////////////////////////////////////////////////////////////////////////////////
}
