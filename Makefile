build-dev:
	docker build . --target development --tag openslides-autoupdate-dev

run-tests:
	docker build . --target testing --tag openslides-autoupdate-test
	docker run openslides-autoupdate-test

all: golinter

golinter:
	golint -set_exit_status ./...
