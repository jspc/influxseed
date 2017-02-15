package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
	"github.com/satori/go.uuid"
)

const (
	hostCount         = 3
	metricName        = "request"
	appName           = "generic-content-service"
	region            = "de"
	databaseName      = "requests"
	requestsPerSecond = 11
)

var (
	seconds        int
	datapointHosts []string

	// set on build
	version, built string
)

func init() {
	flag.IntVar(&seconds, "c", 1000, "seconds worth of data to generate")
}

func main() {
	log.Print(info())
	flag.Parse()

	log.Println("Starting seeder")
	rand.Seed(time.Now().Unix())

	datapointHosts = hosts(hostCount)

	c, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	contentLoop(c)

	log.Print("Completed seeder")
}

func hosts(count int) (h []string) {
	for i := 0; i < count; i++ {
		hostname := strings.SplitN(uuid.NewV4().String(), "-", 2)[0]
		h = append(h, hostname)
	}

	return
}

func info() string {
	return fmt.Sprintf(`influxseed:
version: %s
built: %s

build constants:
    hostCount = %d
    metricName = %s
    appName = %s
    region = %s
    databaseName = %s
    requestsPerSecond = %d

`, version, built, hostCount, metricName, appName, region, databaseName, requestsPerSecond)
}
