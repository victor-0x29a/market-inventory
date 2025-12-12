test-w:
	cd tests && set APP_ENVIRONMENT=test&& go test -v .

test:
	cd tests && APP_ENVIRONMENT=test && go test -v .

dev-w:
	set APP_ENVIRONMENT=dev&& go run main.go

dev:
	APP_ENVIRONMENT=dev && go run main.go

install:
	go mod tidy
