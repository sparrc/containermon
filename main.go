package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	containerName = flag.String("container", "", "Name or ID of the container to monitor")
	outputFormat  = flag.String("output-format", "json", "Output format: json | csv")
	interval      = flag.Int("interval", 10, "collection interval (in seconds)")

	previousTotalUsage uint64
)

func main() {
	flag.Parse()
	if *containerName == "" {
		fmt.Println("--container flag is required")
		flag.Usage()
		return
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(context.TODO())

	start := time.Now()
	stats := getStats(cli)
	startUsage := stats.CPUStats.CPUUsage.TotalUsage
	previousTotalUsage = stats.CPUStats.CPUUsage.TotalUsage
	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	if *outputFormat == "csv" {
		fmt.Println("timeElapsed,cpuTimeElapsed,percentCPUSinceStart,percentCPUThisInterval")
	}
	for {
		select {
		case t := <-ticker.C:
			stats = getStats(cli)
			elapsed := t.Sub(start)
			printStats(stats, elapsed, startUsage)
			previousTotalUsage = stats.CPUStats.CPUUsage.TotalUsage
		}
	}
}

func getStats(cli *client.Client) *types.StatsJSON {
	resp, err := cli.ContainerStats(context.TODO(), *containerName, false)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	data := new(types.StatsJSON)
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}
	return data
}

func printStats(stats *types.StatsJSON, elapsed time.Duration, startUsage uint64) {
	if *outputFormat == "csv" {
		// csv
		// timeElapsed,cpuTimeElapsed,percentCPUSinceStart,percentCPUThisInterval
		fmt.Printf("%.2f,%.2f,%.2f,%.2f\n",
			elapsed.Seconds(),
			float64(stats.CPUStats.CPUUsage.TotalUsage-startUsage)/1000000000,
			float64(stats.CPUStats.CPUUsage.TotalUsage-startUsage)/float64(elapsed.Nanoseconds())*100,
			float64(stats.CPUStats.CPUUsage.TotalUsage-previousTotalUsage)/float64(*interval*1000000000)*100)
	} else {
		// json
		fmt.Printf(`{"timeElapsed":%.2f,"cpuTimeElapsed":%.2f,"percentCPUSinceStart":%.2f,"percentCPUThisInterval":%.2f}`,
			elapsed.Seconds(),
			float64(stats.CPUStats.CPUUsage.TotalUsage-startUsage)/1000000000,
			float64(stats.CPUStats.CPUUsage.TotalUsage-startUsage)/float64(elapsed.Nanoseconds())*100,
			float64(stats.CPUStats.CPUUsage.TotalUsage-previousTotalUsage)/float64(*interval*1000000000)*100)
		fmt.Println()
	}
}
