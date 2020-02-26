# Release the given VERSION

MDCLOUD_GO_VERSION=1.2.0

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
	@docker run -e VERSION=$(VERSION) -v $(shell pwd):/go/src/github.com/OPSWAT/mdcloud-go --rm opswat/mdcloud-go:$(MDCLOUD_GO_VERSION) /go/src/github.com/OPSWAT/mdcloud-go/build.sh
.PHONY: build

# Build docker image
image:
ifndef VERSION
$(error VERSION is not set)
endif
	@docker build --build-arg VERSION=$(VERSION) -t opswat/mdcloud-go:$(VERSION) -f "./image.dockerfile" .
	# @docker push opswat/mdcloud-go:$(VERSION)
.PHONY: image
