override SERVICE=autoupdate
override MAKEFILE_PATH=../dev/scripts/makefile
override DOCKER_COMPOSE_FILE=

# Build images for different contexts

build build-prod build-dev build-tests:
	bash $(MAKEFILE_PATH)/make-build-service.sh $@ $(SERVICE)

# Development

run-dev run-dev-standalone run-dev-attached run-dev-detached run-dev-help run-dev-stop run-dev-clean run-dev-exec run-dev-enter:
	bash $(MAKEFILE_PATH)/make-run-dev.sh "$@" "$(SERVICE)" "$(DOCKER_COMPOSE_FILE)" "$(ARGS)"

# Tests
run-tests:
	bash dev/run-tests.sh

run-lint:
	bash dev/run-lint.sh -l


########################## Deprecation List ##########################

deprecation-warning:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh

stop-dev:
	bash $(MAKEFILE_PATH)/make-deprecation-warning.sh "run-dev-stop"
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

gofmt: | deprecation-warning
	gofmt -l -s -w .
