package docker

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/lxyzhangqing/gpu-memory-monitor/flags"
	"log"
	"os"
	"time"
)

const (
	podNamespaceLabelKey = "io.kubernetes.pod.namespace"
	podNameLabelKey      = "io.kubernetes.pod.name"
)

type ContainerMetadata struct {
	Pid          uint32
	PodNamespace string
	PodName      string
	HostName     string
}

type ContainerPid uint32
type ContainerList map[ContainerPid]*ContainerMetadata

func (cl ContainerList) ContainerMetadata(pid uint32) (*ContainerMetadata, error) {
	if v, exist := cl[ContainerPid(pid)]; exist {
		return v, nil
	}
	return nil, errors.New("pid not found")
}

func (cl ContainerList) GetPodInfo(pid uint32) (podNamespace, podName, podHostName string) {
	if cm, err := cl.ContainerMetadata(pid); err == nil {
		podNamespace = cm.PodNamespace
		podName      = cm.PodName
		podHostName  = cm.HostName
	}
	return podNamespace, podName, podHostName
}

func (cl ContainerList) PrintAll() {
	for _, v := range cl {
		fmt.Printf("pod namespace = %v, pod name = %v, pid = %v, pod hostname = %v\n",
			v.PodNamespace, v.PodName, v.Pid, v.HostName)
	}
}

func GetContainerList() (ContainerList, error) {
	cli, err := client.NewClientWithOpts(client.WithHost("unix://" + flags.DockerSock), client.WithVersion(flags.DockerVersion))
	if err != nil {
		log.Printf("new docker client failed: %v", err)
		return nil, err
	}

	defer func() {
		if err := cli.Close(); err != nil {
			log.Printf("docker client Close failed: %v", err)
		}
	}()

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancelFunc()

	containers, err := cli.ContainerList(timeout, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	containerList := ContainerList{}
	for _, container := range containers {
		cj, err := cli.ContainerInspect(timeout, container.ID)
		if err != nil || cj.State == nil {
			continue
		}

		namespace, name := getContainerPodInfo(container.Labels)
		containerList[ContainerPid(cj.State.Pid)] = &ContainerMetadata{
			Pid: uint32(cj.State.Pid),
			PodNamespace: namespace,
			PodName: name,
			HostName: getHostname(),
		}
	}
	return containerList, nil
}

func getContainerPodInfo(labels map[string]string) (podNamespace, podName string) {
	if len(labels) != 0 {
		podNamespace = labels[podNamespaceLabelKey]
		podName      = labels[podNameLabelKey]
	}
	return podNamespace, podName
}

func getHostname() string {
	hostName, _ := os.Hostname()
	return hostName
}
