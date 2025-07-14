override SERVICE=autoupdate
override MAKEFILE_PATH=../dev/scripts/makefile
override DOCKER_COMPOSE_FILE=

# Build images for different contexts

build-prod:
	docker build ./ --tag "openslides-$(SERVICE)" --build-arg CONTEXT="prod" --target "prod"

build-dev:
	docker build ./ --tag "openslides-$(SERVICE)-dev" --build-arg CONTEXT="dev" --target "dev"

build-tests:
	docker build ./ --tag "openslides-$(SERVICE)-tests" --build-arg CONTEXT="tests" --target "tests"

# Development

.PHONY: dev%

dev%:
	bash $(MAKEFILE_PATH)/make-dev.sh "$@" "$(SERVICE)" "$(DOCKER_COMPOSE_FILE)" "$(ARGS)" "$(USED_SHELL)"

# Tests
run-tests:
	bash dev/run-tests.sh

lint:
	bash dev/run-lint.sh -l

gofmt:
	gofmt -l -s -w .

########################## Deprecation List ##########################

deprecation-warning:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh

stop-dev:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh "dev-stop"
	$(DC_DEV) down --volumes --remove-orphans

build-test:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh "build-tests"
	make build-tests

all:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh "run-tests for tests and lints inside a container or run-lint for local linting"
	make gofmt
	make gotest
	make golinter

gotest: | deprecation-warning
	go test ./...

golinter: | deprecation-warning
	golint -set_exit_status ./...

