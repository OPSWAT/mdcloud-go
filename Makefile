# Release the given VERSION
release:
ifndef VERSION
$(error VERSION is not set)
endif
	@echo "releasing $(VERSION)"
	@echo "building"
	# TODO: git release
	@$(MAKE) build
.PHONY: release

# Test all packages.
test:
	@go test -cover ./...
.PHONY: test

# Build release binaries
# TODO: versioning based on tags
build:
ifndef VERSION
$(error VERSION is not set)
endif
	# VERSION=$(git describe --abbrev=0 --tags --dirty="-$(git rev-parse --abbrev-ref HEAD)")
	@docker build --tag opswat/mdcloud-go . --build-arg VERSION=$(VERSION)
	@docker run --name mdcloud-go opswat/mdcloud-go
	@sleep 1
	@docker cp mdcloud-go:/app/mdcloud-go_* .
	@docker rm mdcloud-go
.PHONY: build

# Build docker image
image:
ifndef VERSION
$(error VERSION is not set)
endif
	@docker build --build-arg VERSION=$(VERSION) -t opswat/mdcloud-go:$(VERSION) -f "./image.dockerfile" .
	@docker push opswat/mdcloud-go:$(VERSION)
.PHONY: image