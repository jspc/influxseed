package main

import (
	"fmt"
	"math/rand"
)

var (
	routes = []string{"getContentTypes", "getContentByID", "getContentBySlug", "getContentByType"}
)

type DataPoint struct {
	Host, Region, App, Route, SiteID string
	Size, TimeTaken, StatusCode      int
}

func NewDataPoint() (dp DataPoint) {
	dp.Host = datapointHosts[rand.Intn(hostCount)]
	dp.Region = region
	dp.App = appName
	dp.Route = routes[rand.Intn(4)]
	dp.SiteID = fmt.Sprintf("%d", rand.Intn(3))

	dp.Size = endpointSizer(dp.Route)
	dp.TimeTaken = (rand.Intn(800) + dp.Size/60) - 100
	dp.StatusCode = []int{200, 200, 200, 200, 200, 403, 404, 500}[rand.Intn(8)]

	return
}

func (dp DataPoint) Tags() map[string]string {
	return map[string]string{"host": dp.Host, "region": dp.Region, "app": dp.App, "route": dp.Route, "status": fmt.Sprintf("%d", dp.StatusCode), "siteID": dp.SiteID}
}

func (dp DataPoint) Fields() map[string]interface{} {
	return map[string]interface{}{
		"size":     dp.Size,
		"duration": dp.TimeTaken,
	}
}

func endpointSizer(e string) int {
	switch e {
	case "getContentTypes":
		return 5000 + rand.Intn(500)
	case "getContentByID":
		return 24000 + rand.Intn(1200)
	case "getContentByType":
		return 50000 + rand.Intn(5000)
	case "getContentBySlug":
		return 18000 + rand.Intn(1800)
	}
	return 0
}
