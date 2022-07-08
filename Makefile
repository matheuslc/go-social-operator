swag:
	@swag init -d cmd --parseDependency --parseDepth 10 -o docs/

run:
	@go run cmd/main.go