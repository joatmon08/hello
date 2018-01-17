# joatmon08/hello

Just a simple app, serving on two ports. Use for multi-post testing.

| Port      | URI     | Description |
| --------- |-------- | ----------- |
| :8001     | /hello  | Returns Hello World |
| :8002     | /health | Returns I'm healthy! |

## Build
Run `./build.sh` to build the binary and create a Docker image.

## Run
To run as a container, use
`docker run -d -p 8001:8001 -p 8002:8002 joatmon08/hello:latest`.

## Deploy to Kubernetes
Run `kubectl apply -f kubernetes.yml`.
