.ONESHELL:

PACKAGE_DIRS=$(shell go list ./... | grep -v /vendor/)

tidy:
	@go mod tidy

test:
	@echo "crossed fingers emoji running tests"
	@go test -v $(PACKAGE_DIRS)

secrets:
	yq . ${HOME}/.sux/secrets.yml > config/secrets.yml
	@echo "local secrets file is now tainted, use \"make rmsecrets\" to remove before committing"

rmsecrets:
	@echo "removing local secrets"
	cp config/secrets_example.yml config/secrets.yml

version:
	@echo "Updating version data"
	@echo "version:" > config/version.yml
	@echo "  build_date: \"`date`\"" >> config/version.yml
	@echo "  build: \"`git describe --tags --always`\"" >> config/version.yml
	@echo "  branch: \"`git branch | grep '^*' | cut -d' ' -f 2`\"" >> config/version.yml

build: test version secrets 
