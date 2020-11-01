build-dev:
	docker build . --target development --tag openslides-permission-dev

run-tests:
	docker build . --target testing --tag openslides-permission-test
	docker run openslides-permission-test
