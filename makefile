build:
	docker build -t hello:local .
	docker build -t hello:test -f Dockerfile.test .

setup:
	helm version
	helm init
	helm plugin install https://github.com/lrills/helm-unittest

inttest:
	helm install --name testing .
	helm test --cleanup testing
	helm del --purge testing

unittest:
	helm unittest -f 'templates/tests/unit/*.yaml' .