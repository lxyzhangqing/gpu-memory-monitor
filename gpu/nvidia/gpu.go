package nvidia

import (
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/types"
)

type GPU struct {
	DeviceName string
	DeviceUUID string
	MemoryUsed uint64
}

func (gpu GPU) Type() string {
	return gpu.DeviceName
}

func (gpu GPU) UUID() string {
	return gpu.DeviceUUID
}

func (gpu GPU) MemoryUsage() uint64 {
	return gpu.MemoryUsed
}

type PodResourceUsage struct {
	Namespace string
	Name      string
	HostName  string
	PodPid    uint32
	GpuMemory *GPU
}

func (pod PodResourceUsage) PodNamespace() string {
	return pod.Namespace
}

func (pod PodResourceUsage) PodName() string {
	return pod.Name
}

func (pod PodResourceUsage) Pid() uint32 {
	return pod.PodPid
}

func (pod PodResourceUsage) PodHostName() string {
	return pod.HostName
}

func (pod PodResourceUsage) GpuMemoryUsed() types.GPU {
	return pod.GpuMemory
}
