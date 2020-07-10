PROJECT_NAME := Pulumi CloudFormation Resource Provider
include build/common.mk

PACK             := cloudformation
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-cloudformation
NODE_MODULE_NAME := @pulumi/cloudformation
NUGET_PKG_NAME   := Pulumi.CloudFormation

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         ?= $(shell scripts/get-version)
PYPI_VERSION    := $(shell cd scripts && ./get-py-version)

CFN_SCHEMA_REGION   ?= us-west-2
CFN_SCHEMA_URL      ?= https://cfn-resource-specifications-${CFN_SCHEMA_REGION}-prod.s3.${CFN_SCHEMA_REGION}.amazonaws.com/latest/gzip/CloudFormationResourceSpecification.json
CFN_SCHEMA_DIR      := provider/cmd/pulumi-gen-cloudformation
CFN_SCHEMA_FILE     := ${CFN_SCHEMA_DIR}/cfn-spec-${CFN_SCHEMA_REGION}.json

SCHEMA_FILE := provider/cmd/pulumi-resource-cloudformation/schema.json

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-cloudformation/provider/pkg/version.Version=${VERSION}"

GO              ?= go
CURL            ?= curl
PYTHON          ?= python3

DOTNET_PREFIX  := $(firstword $(subst -, ,${VERSION:v%=%})) # e.g. 1.5.0
DOTNET_SUFFIX  := $(word 2,$(subst -, ,${VERSION:v%=%}))    # e.g. alpha.1

ifeq ($(strip ${DOTNET_SUFFIX}),)
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})
else
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})-$(strip ${DOTNET_SUFFIX})
endif

TESTPARALLELISM := 10
TESTABLE_PKGS   := ./provider/pkg/... ./examples/...

# Set NOPROXY to true to skip GOPROXY on 'ensure'
NOPROXY := false

$(CFN_SCHEMA_FILE)::
	test -f $(CFN_SCHEMA_FILE) || $(CURL) -s -L $(CFN_SCHEMA_URL) | gzip -d > $(CFN_SCHEMA_FILE)

cfngen::
	cd provider && $(GO) install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN)

$(SCHEMA_FILE):: cfngen $(CFN_SCHEMA_FILE)
	$(CODEGEN) schema $(CFN_SCHEMA_FILE) $(CURDIR)

cfnprovider:: $(SCHEMA_FILE)
	#$(CODEGEN) provider $(CFN_SCHEMA_FILE) $(CURDIR)
	cd provider && VERSION=${VERSION} $(GO) generate cmd/${PROVIDER}/main.go
	cd provider && $(GO) install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)

dotnet_sdk:: cfngen $(SCHEMA_FILE)
	$(CODEGEN) -version=${VERSION} dotnet $(SCHEMA_FILE) $(CURDIR)
	cd ${PACKDIR}/dotnet/&& \
		echo "${VERSION:v%=%}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

go_sdk:: cfngen $(SCHEMA_FILE)
	$(CODEGEN) -version=${VERSION} go $(SCHEMA_FILE) $(CURDIR)

nodejs_sdk:: cfngen $(SCHEMA_FILE)
	$(CODEGEN) -version=${VERSION} nodejs $(SCHEMA_FILE) $(CURDIR)
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc
	cp README.md LICENSE ${PACKDIR}/nodejs/package.json ${PACKDIR}/nodejs/yarn.lock ${PACKDIR}/nodejs/bin/
	sed -i.bak 's/$${VERSION}/$(VERSION)/g' ${PACKDIR}/nodejs/bin/package.json

python_sdk:: cfngen $(SCHEMA_FILE)
	# Delete only files and folders that are generated.
	rm -r sdk/python/pulumi_cloudformation/*/ sdk/python/pulumi_cloudformation/__init__.py
	$(CODEGEN) -version=${VERSION} python $(SCHEMA_FILE) $(CURDIR)
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		$(PYTHON) setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && $(PYTHON) setup.py build sdist

lint::
	for DIR in "provider" "sdk" "examples" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

install::
	cd provider && GOBIN=$(PULUMI_BIN) $(GO) install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER)
	[ ! -e "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" ] || rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	mkdir -p "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	cp -r sdk/nodejs/bin/. "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/node_modules"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/tests"
	cd "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" && \
		yarn install --offline --production && \
		(yarn unlink > /dev/null 2>&1 || true) && \
		yarn link
	echo "Copying ${NUGET_PKG_NAME} NuGet packages to ${PULUMI_NUGET}"
	mkdir -p $(PULUMI_NUGET)
	rm -rf "$(PULUMI_NUGET)/$(NUGET_PKG_NAME).*.nupkg"
	find . -name '$(NUGET_PKG_NAME).*.nupkg' -exec cp -p {} ${PULUMI_NUGET} \;

test_fast::
	cd provider/pkg && $(GO_TEST_FAST) ./...
	cd examples && $(GO_TEST_FAST) ./...

test_all::
	cd provider/pkg && $(GO_TEST) ./...
	cd examples && $(GO_TEST) ./...

generate_schema:: $(SCHEMA_FILE)

.PHONY: build
build:: cfngen cfnprovider dotnet_sdk go_sdk nodejs_sdk python_sdk

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
