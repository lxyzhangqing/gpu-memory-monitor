FROM golang:1.15-stretch as build

WORKDIR /go/src/github.com/lxyzhangqing/gpu-memory-monitor
COPY . .
RUN make

FROM debian:bullseye-slim

COPY --from=build /go/src/github.com/lxyzhangqing/gpu-memory-monitor/gpu-memory-monitor /usr/bin/gpu-memory-monitor

CMD ["gpu-memory-monitor"]
