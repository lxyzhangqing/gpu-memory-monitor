package cambricon

import (
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/types"
)

type GPU struct {

}

func (gpu *GPU) Type() string {
	// TODO:
	return ""
}

func (gpu *GPU) UUID() string {
	// TODO:
	return ""
}

func (gpu *GPU) MemoryUsage() uint64 {
	// TODO:
	return 0
}

type PodResourceUsage struct {

}

func (pod PodResourceUsage) PodNamespace() string {
	// TODO:
	return "default"
}

func (pod PodResourceUsage) PodName() string {
	// TODO:
	return "unknown"
}

func (pod PodResourceUsage) Pid() uint32 {
	// TODO:
	return 0
}

func (pod PodResourceUsage) PodHostName() string {
	// TODO:
	return ""
}

func (pod PodResourceUsage) GpuMemoryUsed() types.GPU {
	// TODO:
	return nil
}
