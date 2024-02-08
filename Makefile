run:## Run personal-finance project 
	go run cmd/main.go

clean: ## Erase the database file
	rm test.db

test: ## Run all unit tests in the project
	go test -v ./...

test-cover: ## Run all unit tests in the porject with test coverage
	go test -v ./... -covermode=count -coverprofile=coverage.out
