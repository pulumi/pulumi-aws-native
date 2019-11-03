PROJECT_NAME := Pulumi CLoudformation Resource Provider
include build/common.mk

PACK             := cloudformation
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-cloudformation
NODE_MODULE_NAME := @pulumi/cloudformation

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         ?= $(shell scripts/get-version)
PYPI_VERSION    := $(shell scripts/get-py-version)

SCHEMA_REGION   ?= us-west-2
SCHEMA_URL      ?= https://cfn-resource-specifications-${SCHEMA_REGION}-prod.s3.${SCHEMA_REGION}.amazonaws.com/latest/gzip/CloudFormationResourceSpecification.json
SCHEMA_DIR      := pkg/gen/
SCHEMA_FILE     := ${SCHEMA_DIR}/cfn-spec-${SCHEMA_REGION}.json

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-cloudformation/pkg/version.Version=${VERSION}"

GO              ?= go
CURL            ?= curl
PYTHON          ?= python3

TESTPARALLELISM := 10
TESTABLE_PKGS   := ./pkg/... ./examples/...

# Set NOPROXY to true to skip GOPROXY on 'ensure'
NOPROXY := false

$(SCHEMA_FILE)::
	@mkdir -p $(SCHEMA_DIR)
	test -f $(SCHEMA_FILE) || $(CURL) -s -L $(SCHEMA_URL) | gzip -d > $(SCHEMA_FILE)

build:: $(SCHEMA_FILE)
	go generate ./pkg/gen
	$(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(PROVIDER)
	$(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(CODEGEN)
	$(CODEGEN) <$(SCHEMA_FILE) >sdk/nodejs/cloudformation.ts || exit 3 ; \
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc
	cp README.md LICENSE ${PACKDIR}/nodejs/package.json ${PACKDIR}/nodejs/yarn.lock ${PACKDIR}/nodejs/bin/
	sed -i.bak 's/$${VERSION}/$(VERSION)/g' ${PACKDIR}/nodejs/bin/package.json

lint::
	golangci-lint run

install::
	GOBIN=$(PULUMI_BIN) $(GO) install $(VERSION_FLAGS) $(PROJECT)/cmd/$(PROVIDER)
	[ ! -e "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" ] || rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	mkdir -p "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	cp -r sdk/nodejs/bin/. "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/node_modules"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/tests"
	cd "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" && \
		yarn install --offline --production && \
		(yarn unlink > /dev/null 2>&1 || true) && \
		yarn link

test_fast::
	$(GO_TEST_FAST) $(TESTABLE_PKGS)

test_all::
	$(GO_TEST) $(TESTABLE_PKGS)

.PHONY: publish_tgz
publish_tgz:
	$(call STEP_MESSAGE)
	./scripts/publish_tgz.sh

# While pulumi-cloudformation is not built using tfgen/tfbridge, the layout of the source tree is the same as these
# style of repositories, so we can re-use the common publishing scripts.
.PHONY: publish_packages
publish_packages:
	$(call STEP_MESSAGE)
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/publish-tfgen-package .
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/build-package-docs.sh cloudformation

.PHONY: check_clean_worktree
check_clean_worktree:
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/check-worktree-is-clean.sh

# The travis_* targets are entrypoints for CI.
.PHONY: travis_cron travis_push travis_pull_request travis_api
travis_cron: all
travis_push: only_build check_clean_worktree publish_tgz only_test publish_packages
travis_pull_request: only_build check_clean_worktree only_test
travis_api: all
