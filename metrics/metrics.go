package metrics

import (
	"github.com/lxyzhangqing/gpu-memory-monitor/flags"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var registry *prometheus.Registry

func init() {
	registry = prometheus.NewRegistry()
	registry.MustRegister(NewGpuMetrics())
}



func HandleMetrics() {
	http.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	if err := http.ListenAndServe(flags.ServerAddr + ":" + flags.ServerPort, nil); err != nil {
		log.Fatalln(err)
	}
}
