# joatmon08/hello

Just a simple app, serving on two ports. Use for multi-post testing.

| Port      | URI     | Description |
| --------- |-------- | ----------- |
| :8001     | /version  | Returns version |
| :8001     | /hello  | Returns Hello World |
| :8001     | /phone?targetService=nginx  | Expects to call a service denoted by targetService and get a return. Default: nginx. |
| :8001     | /cpu?testTime=1m  | Generates CPU load for the specified testTime. Default: 1m. |
| :8002     | /health | Returns I'm healthy! |

## Build
Run `go build main.go` to build the binary.

## Run
To run as a container, use
`docker run -d -p 8001:8001 -p 8002:8002 joatmon08/hello:latest`.

## Deploy to Kubernetes
Run `kubectl apply -f kubernetes.yml`.
