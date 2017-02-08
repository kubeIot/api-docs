# Kubelot api mock

This mock was generated with go-swagger tool.
To get the documentation install go swagger to you `GOPATH` with
```bash
go get github.com/go-swagger/go-swagger
go install github.com/go-swagger/go-swagger
```

After that run `swagger `

## How to run
This project is built with Godep tool located at [https://github.com/tools/godep](https://github.com/tools/godep). The tool is essential for build,
so follow the instructions found in repository.

After godep installation, the project can be built and started with following commands
```bash
godeps go build ./cmd/api-docs-server
./api-docs-server --port 8080
```
The last commands starts `http` server, which contains implementation of api as well as documentation.

The documentation can be found on `http://localhost:8080/docs` when the mock is running or by running following command
```bash
swagger serve -p 8080 ./swagger.yml
```




