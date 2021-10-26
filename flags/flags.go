package flags

import (
	"flag"
)

var DockerSock string
var DockerVersion string
var ServerAddr string
var ServerPort string
var GPUType string

func init() {
	flag.StringVar(&DockerSock, "docker-sock", "/var/run/docker.sock", "set docker sock path")
	flag.StringVar(&DockerVersion, "docker-version", "v1.38", "set docker version")
	flag.StringVar(&ServerAddr,"addr", "0.0.0.0", "server lister address")
	flag.StringVar(&ServerPort, "port", "5091", "set server listen port")
	flag.StringVar(&GPUType, "gpu-type", "nvidia", "set GPU type")
}
