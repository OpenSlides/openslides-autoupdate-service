build-dev:
	docker build . -f docker/Dockerfile.dev --tag openslides-autoupdate-dev

run-tests:
	docker build . -f docker/Dockerfile.test --tag openslides-autoupdate-test
	docker run openslides-autoupdate-test
