# gpu-memory-monitor

`gpu-memory-monitor` is a metrics server for collecting GPU memory usage of kubernetes pods.
If you have a GPU machine, and some pods are using the GPU device, you can run the container by docker or kubernetes when your GPU device belongs to nvidia.
The `gpu-memory-monitor` will collect the GPU memory usage of pods, you can get those metrics by API of `gpu-memory-monitor`.

### Prerequisites
- golang 1.15+
- NVIDIA drivers ~= 361.93
- Nvidia-docker version > 2.0 (see how to [install](https://github.com/NVIDIA/nvidia-docker) and it's [prerequisites](https://github.com/nvidia/nvidia-docker/wiki/Installation-\(version-2.0\)#prerequisites))

### How to build binary?
```
$ git clone https://github.com/lxyzhangqing/gpu-memory-monitor.git
$ cd gpu-memory-monitor
$ go mod tidy
$ go mod vendor
$ make
```

### How to build images?
```
$ git clone https://github.com/lxyzhangqing/gpu-memory-monitor.git
$ cd gpu-memory-monitor
$ go mod tidy
$ go mod vendor
$ docker build -t gpu-memory-monitor:v1 .
```

### How to deploy gpu memory monitor by docker?
You can execute the following command line on your GPU machine.
```
docker run -d --name=gpu-memory-monitor -e NVIDIA_VISIBLE_DEVICES=all -e NVIDIA_DRIVER_CAPABILITIES=utility -v /var/run:/var/run:ro  --net=host gpu-memory-monitor:v1
```

### How to deploy gpu memory monitor by kubernetes?
You can copy `deploy.yaml` to your kubernetes cluster and execute the following command line to deploy `gpu-momory-monitor`. 
Before this, you should to edit `nodeAffinity` for scheduling pods of `gpu-memory-monitor` metrics server to correct GPU machines.
```
kubectl create -f deploy.yaml
```

### How to get the metrics?
You can execute this command line on you machine:
```
curl http://127.0.0.1:5091/metrics
```

Then you may get metrics info like this:
```
# HELP pod gpu memory usage, unit is MiB
# TYPE pod_gpu_memory_usage gauge
pod_gpu_memory_usage{gpu_type="Tesla T4",gpu_uuid="GPU-576ab88b-464f-5903-3ab9-2d25e3ee6c4a",hostname="test-node",name="gpu.test1-85846f7bd4-4ppm9",namespace="default",pid="37691"} 2027
pod_gpu_memory_usage{gpu_type="Tesla T4",gpu_uuid="GPU-6758250c-1793-6349-ba37-332ac77b1d0a",hostname="test-node",name="gpu.test2-57485d95d6-wsngh",namespace="default",pid="54702"} 3449
```