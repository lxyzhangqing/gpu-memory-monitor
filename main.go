package main

import (
	"flag"
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu"
	"github.com/lxyzhangqing/gpu-memory-monitor/metrics"
	"github.com/lxyzhangqing/gpu-memory-monitor/version"
	"log"
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	version.PrintVersionOrContinue()

	if !gpu.IsSupported() {
		log.Fatalf("gpu-type is invalid, please check")
	}

	if err := gpu.Init(); err != nil {
		log.Fatalf("gpu device init failed: %v", err)
	}
	defer gpu.Finish()

	metrics.HandleMetrics()
}
