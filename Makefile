mod:
	@echo "Vendoring..."
	@go mod vendor
test: 
	@echo "Ejecutando tests..."
	@go test ./... -v

coverage: 
	@echo "Coverage..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
