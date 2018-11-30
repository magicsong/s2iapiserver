IMG ?= controller:latest

# Generate code
generate:
	apiserver-boot build generated

manager:generate
	CGO_ENABLED=0 go build -o bin/controller-manager cmd/controller-manager/main.go

debug: manager


