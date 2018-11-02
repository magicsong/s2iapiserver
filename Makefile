IMG ?= controller:latest

# Generate code
generate:
	apiserver-boot build generated

manager:
	apiserver-boot build executables

debug: 
	apiserver-boot run local

