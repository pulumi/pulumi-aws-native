PROJECT_NAME := Pulumi Native AWS Resource Provider
include build/common.mk

PACK             := aws-native
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-aws-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION         := $(shell pulumictl get version)

PROVIDER_PKGS    := $(shell cd ./provider && go list ./...)
WORKING_DIR     := $(shell pwd)

VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}"

CFN_SCHEMA_REGION   ?= us-west-2
CFN_SCHEMA_URL      ?= https://cfn-resource-specifications-${CFN_SCHEMA_REGION}-prod.s3.${CFN_SCHEMA_REGION}.amazonaws.com/latest/gzip/CloudFormationResourceSpecification.json
CFN_SCHEMA_DIR      := provider/cmd/pulumi-gen-${PACK}
CFN_SCHEMA_FILE     := ${CFN_SCHEMA_DIR}/cfn-spec-${CFN_SCHEMA_REGION}.json

update_cfm_schema::
	test -f $(CFN_SCHEMA_FILE) || curl -s -L $(CFN_SCHEMA_URL) | gzip -d > $(CFN_SCHEMA_FILE)

ensure::
	@echo "GO111MODULE=on go mod tidy"; cd provider; GO111MODULE=on go mod tidy
	@echo "GO111MODULE=on go mod download"; cd provider; GO111MODULE=on go mod download

local_generate:: clean
	$(WORKING_DIR)/bin/$(CODEGEN) schema,nodejs,dotnet,python $(CFN_SCHEMA_FILE) ${VERSION}
	echo "Finished generating."

generate_schema::
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) schema $(CFN_SCHEMA_FILE) ${VERSION}
	echo "Finished generating schema."

codegen::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -a -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

test_provider::
	(cd provider && go test -v $(PROVIDER_PKGS))

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

generate_nodejs::
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs $(CFN_SCHEMA_FILE) ${VERSION}

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

generate_python::
	$(WORKING_DIR)/bin/$(CODEGEN) python $(CFN_SCHEMA_FILE) ${VERSION}

build_python:: VERSION := $(shell pulumictl get version --language python)
build_python::
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

generate_dotnet::
	$(WORKING_DIR)/bin/$(CODEGEN) dotnet $(CFN_SCHEMA_FILE) ${VERSION}

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
		echo "${PACK}\n${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

generate_go::
	$(WORKING_DIR)/bin/$(CODEGEN) go $(CFN_SCHEMA_FILE) ${VERSION}

build_go::

clean::
	rm -rf sdk/nodejs && mkdir sdk/nodejs && touch sdk/nodejs/go.mod
	rm -rf sdk/python && mkdir sdk/python && touch sdk/python/go.mod
	rm -rf sdk/dotnet && mkdir sdk/dotnet && touch sdk/dotnet/go.mod
	rm -rf sdk/go

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test::
	cd examples && go test -v -tags=all -timeout 2h

build:: clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

.PHONY: ensure generate_schema generate build_provider build
