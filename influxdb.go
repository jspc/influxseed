package main

import (
	"fmt"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

func contentLoop(c influxdb.Client) {
	bp, _ := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database:  databaseName,
		Precision: "s",
	})

	for i := 0; i < seconds; i++ {
		fmt.Printf("iteration %d/%d\r", i+1, seconds)

		timeOffset, _ := time.ParseDuration(fmt.Sprintf("%ds", (0 - i)))

		for j := 0; j < requestsPerSecond; j++ {
			dp := NewDataPoint()

			tags := dp.Tags()
			fields := dp.Fields()

			pt, err := influxdb.NewPoint(metricName, tags, fields, time.Now().Add(timeOffset))
			if err != nil {
				fmt.Println("Error: ", err.Error())
			}
			bp.AddPoint(pt)
		}
	}

	c.Write(bp)

}
