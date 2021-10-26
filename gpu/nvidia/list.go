package nvidia

import (
	"github.com/lxyzhangqing/gpu-memory-monitor/gpu/types"
)

func List() []types.PodResourceUsage {
	if podResourceUsageList, err := GetProcessInfo(); err != nil {
		return nil
	}else{
		return podResourceUsageList
	}
}
