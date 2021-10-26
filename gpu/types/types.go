package types

type GPU interface {
	Type() string
	UUID() string
	MemoryUsage() uint64
}

type PodResourceUsage interface {
	PodNamespace() string
	PodName() string
	PodHostName() string
	Pid() uint32
	GpuMemoryUsed() GPU
}