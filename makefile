build:
	docker build -t hello:local .
	docker build -t hello:test -f Dockerfile.test .

binary:
	go build .

unit:
	go test ./test/unit/... -v

setup:
	helm version
	helm init
	helm plugin install https://github.com/lrills/helm-unittest

inttest:
	helm install --name testing --namespace todaystest .
	sleep 30
	helm test --cleanup testing
	helm del --purge testing

unittest:
	helm unittest -f 'templates/tests/unit/*.yaml' .