IMG ?= controller:latest

# Generate code
generate:
	apiserver-boot build generated

manager:generate
	CGO_ENABLED=0 go build -o bin/apiserver cmd/apiserver/main.go
	CGO_ENABLED=0 go build -o bin/controller-manager cmd/controller-manager/main.go

.PHONY: build
build: 
	apiserver-boot build executables --generate=false
debug: manager
	apiserver-boot run local --run=apiserver --run=controller-manager --etcd "http://192.168.98.8:2379" --generate=false --build=false

