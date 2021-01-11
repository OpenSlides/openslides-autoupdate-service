build-dev:
	docker build . --target development --tag openslides-permission-dev

run-tests:
	docker build . --target testing --tag openslides-permission-test
	docker run openslides-permission-test

cover:
	go test ./...  -coverprofile=cover.out -coverpkg=./internal/collection
	go tool cover -html cover.out