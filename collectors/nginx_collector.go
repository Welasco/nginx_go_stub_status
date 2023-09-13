package collectors

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/Welasco/nginx_go_stub_status/common"
)

func (nlog Nginx_log) CollectLog(bodystr string) error {

	// Regular expression patterns
	activeConnPattern := regexp.MustCompile(`Active connections: (\d+)`)
	serviceConnPattern := regexp.MustCompile(` (\d+) (\d+) (\d+)`)
	rwwPattern := regexp.MustCompile(`Reading: (\d+) Writing: (\d+) Waiting: (\d+)`)

	// Find matches in the text
	activeConnMatches := activeConnPattern.FindStringSubmatch(bodystr)
	serviceConnMatches := serviceConnPattern.FindStringSubmatch(bodystr)
	rwwMatches := rwwPattern.FindStringSubmatch(bodystr)

	nginx_log := &Nginx_log{}
	nginx_log.Active_connections = activeConnMatches[1]
	nginx_log.Server_accepts = serviceConnMatches[1]
	nginx_log.Server_handled = serviceConnMatches[2]
	nginx_log.Server_requests = serviceConnMatches[3]
	nginx_log.Reading = rwwMatches[1]
	nginx_log.Writing = rwwMatches[2]
	nginx_log.Waiting = rwwMatches[3]

	write_to_csv(nginx_log)

	return nil
}

func write_to_csv(nginx_log *Nginx_log) {
	file, err := os.OpenFile("nginx_metrics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	//writer.Write([]string{"Active_connections", "Server_accepts", "Server_handled", "Server_requests", "Reading", "Writing", "Waiting"})
	writer.Write([]string{time.Now().Format("2006-01-02T15:04:05"), nginx_log.Active_connections, nginx_log.Server_accepts, nginx_log.Server_handled, nginx_log.Server_requests, nginx_log.Reading, nginx_log.Writing, nginx_log.Waiting})
}

// Variable to control the status of the collector
// Running or Stopped
var shouldStop bool

func (nlog Nginx_log) Start() error {
	shouldStop = false
	for !shouldStop {
		// need to implement a configuration file to get the url and the specifics of the collector
		req, err := common.Http_client("http://localhost/nginx_status")
		if err != nil {
			fmt.Println(err)
			return err
		}

		nlog.CollectLog(req)
		time.Sleep(5 * time.Second)
	}

	return nil
}

func (nlog Nginx_log) Stop() error {
	shouldStop = true
	return nil
}

func (nlog Nginx_log) GetStatus() string {
	if !shouldStop {
		return "Running"
	} else if shouldStop {
		return "Stopped"
	}
	return "Unknown"
}

func Newnginx_log() *Nginx_log {
	var nlog Nginx_log = Nginx_log{}
	return &nlog
}
