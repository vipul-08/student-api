check_install:
	which swagger || GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_install
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models
test:
	GO111MODULE=on go test ./... -v
