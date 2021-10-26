package metrics

import (
	"fmt"
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu"
	"github.com/prometheus/client_golang/prometheus"
)

func NewGpuMetrics() *GpuMetrics {
	m := new(GpuMetrics)
	m.Init()
	return m
}

type GpuMetrics struct {
	gpuMemoryUsage *prometheus.Desc
}

func (g *GpuMetrics) Init() {
	g.gpuMemoryUsage = prometheus.NewDesc(
		"pod_gpu_memory_usage",
		"pod gpu memory usage, unit is MiB",
		[]string{"namespace","name","hostname","pid","gpu_type","gpu_uuid"},
		nil,
	)
}

// Describe returns all descriptions of the collector.
func (g *GpuMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- g.gpuMemoryUsage
}

// Collect returns the current state of all metrics of the collector.
func (g *GpuMetrics) Collect(ch chan<- prometheus.Metric) {
	pods := gpu.List()
	if len(pods) == 0 {
		return
	}

	for _, pod := range pods {
		gpuUsage := pod.GpuMemoryUsed()
		if gpuUsage == nil {
			continue
		}

		ch <- prometheus.MustNewConstMetric(g.gpuMemoryUsage, prometheus.GaugeValue, float64(gpuUsage.MemoryUsage()/1024/1024),
			pod.PodNamespace(),
			pod.PodName(),
			pod.PodHostName(),
			fmt.Sprintf("%v", pod.Pid()),
			gpuUsage.Type(),
			gpuUsage.UUID(),
		)
	}
}
