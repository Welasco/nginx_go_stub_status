package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	req, err := http.NewRequest("GET", "http://localhost/nginx_status", nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)

	scanner := bufio.NewScanner(strings.NewReader(bodystr))
	var sahr bool = false
	for scanner.Scan() {
		// Grab Active connections
		if strings.Contains(scanner.Text(), "Active connections") {
			fmt.Println(scanner.Text())
			txt := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])
			fmt.Println(txt)
		}

		// Grab Server accepts handled requests
		if sahr {
			fmt.Println(scanner.Text())
			txt := strings.Split(scanner.Text(), " ")

			fmt.Println(len(txt))
			for i := 0; i < len(txt); i++ {
				fmt.Println("i=", i, " ", txt[i])
			}
			fmt.Println(txt)
			sahr = false
		}
		if strings.Contains(scanner.Text(), "server accepts handled requests") {
			sahr = true
			fmt.Println(scanner.Text())
		}

		// Grab Reading: Writing: Waiting:
		if strings.Contains(scanner.Text(), "Reading") {
			txt := strings.Split(scanner.Text(), " ")
			for i := 0; i < len(txt); i++ {
				fmt.Println("i=", i, " ", txt[i])
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
