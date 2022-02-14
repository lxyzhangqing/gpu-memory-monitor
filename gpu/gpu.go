package gpu

import (
	"errors"
	"github.com/lxyzhangqing/gpu-memory-monitor/flags"
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/cambricon"
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/nvidia"
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/types"
)

const (
	NVIDIA    = "nvidia"
	CAMBRICON = "cambricon"
)

func Init() error {
	switch flags.GPUType {
	case NVIDIA:
		return nvidia.Init()
	case CAMBRICON:
		return cambricon.Init()
	default:
		return errors.New("not support type")
	}
}

func Finish() {
	switch flags.GPUType {
	case NVIDIA:
		nvidia.Finish()
	case CAMBRICON:
		cambricon.Finish()
	default:
		// do nothing
	}
}

func IsSupported() bool {
	switch flags.GPUType {
	case NVIDIA:
		return true
	case CAMBRICON:
		return true
	default:
		return false
	}
}

func List() []types.PodResourceUsage {
	switch flags.GPUType {
	case NVIDIA:
		return nvidia.List()
	case CAMBRICON:
		return cambricon.List()
	default:
		return nil
	}
}

