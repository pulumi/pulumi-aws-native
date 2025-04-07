PROJECT_NAME := Pulumi Native AWS Resource Provider

PACK             := aws-native
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-aws-native
PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}

WORKING_DIR		:= $(shell pwd)

JAVA_GEN		 := pulumi-java-gen
JAVA_GEN_VERSION := v0.9.7

# Override during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local & branch builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 1.0.0-alpha.0+dev
# Use this normalised version everywhere rather than the raw input to ensure consistency.
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")
VERSION_FLAGS   := -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION_GENERIC}"

CFN_SCHEMA_DIR  := aws-cloudformation-schema
METADATA_DIR    := meta

# Only use explicitly installed plugins - this is to avoid using any ambient plugins from the PATH
export PULUMI_IGNORE_AMBIENT_PLUGINS = true

init_submodules::
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		if [ ! -f "$$submodule/.git" ]; then \
			echo "Initializing submodule $$submodule" ; \
			(cd $$submodule && git submodule update --init); \
		fi; \
	done

update_submodules:: init_submodules
	@for submodule in $$(git submodule status | awk {'print $$2'}); do \
		echo "Updating submodule $$submodule" ; \
		(cd $$submodule && git checkout main && git pull origin main); \
	done

docs:: DOCS_VERSION := $(shell cat .docs.version)
docs:: .docs.version
	$(WORKING_DIR)/bin/$(CODEGEN) -schema-folder $(CFN_SCHEMA_DIR) -version ${VERSION_GENERIC} -docs-url https://github.com/cdklabs/awscdk-service-spec/raw/${DOCS_VERSION}/sources/CloudFormationDocumentation/CloudFormationDocumentation.json docs

discovery:: update_submodules codegen
	git ls-remote https://github.com/cdklabs/awscdk-service-spec refs/heads/main | awk '{print $$1}' > .docs.version
	$(WORKING_DIR)/bin/$(CODEGEN) -schema-folder $(CFN_SCHEMA_DIR) -version ${VERSION_GENERIC} -schema-urls https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip,https://schema.cloudformation.us-west-2.amazonaws.com/CloudformationSchema.zip discovery

	# Fetch the latest version of botocore
	git -c 'versionsort.suffix=-' ls-remote --exit-code --refs --sort='version:refname' --tags https://github.com/boto/botocore '*.*.*' \
		| tail -n 1 \
		| cut -d '/' -f 3 \
		> "$(METADATA_DIR)/.botocore.version"

	# Extract the regions from botocore
	cat $(METADATA_DIR)/.botocore.version | xargs -I{} curl -fsSL https://raw.githubusercontent.com/boto/botocore/{}/botocore/data/endpoints.json \
		| jq '[.partitions[].regions | to_entries[] | { name: .key, description: .value.description}]' \
		> "$(METADATA_DIR)/regions.json"

ensure:: init_submodules
	@echo "GO111MODULE=on go mod tidy"
	cd provider && GO111MODULE=on go mod tidy
	cd sdk/go && GO111MODULE=on go mod tidy
	cd examples && GO111MODULE=on go mod tidy

local_generate:: generate_schema generate_nodejs generate_python generate_dotnet generate_java generate_go

generate_schema:: docs
	echo "Generating Pulumi schema..."
	$(WORKING_DIR)/bin/$(CODEGEN) --schema-folder $(CFN_SCHEMA_DIR) --version ${VERSION_GENERIC} --metadata-folder $(METADATA_DIR) schema
	echo "Finished generating schema."

codegen::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(CODEGEN))

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/$(PROVIDER) $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

install_provider::
	(cd provider && go install $(VERSION_FLAGS) $(PROJECT)/provider/cmd/$(PROVIDER))

cf2pulumi::
	(cd provider && go build -o $(WORKING_DIR)/bin/cf2pulumi $(VERSION_FLAGS) $(PROJECT)/provider/cmd/cf2pulumi)

test_provider::
	(cd provider && go test -v -coverpkg=./... -coverprofile=coverage.txt ./...)

lint_provider:: provider # lint the provider code
	cd provider && GOGC=20 golangci-lint run -c ../.golangci.yml

.pulumi/bin/pulumi: PULUMI_VERSION := $(shell cat .pulumi.version)
.pulumi/bin/pulumi: HOME := $(WORKING_DIR)
.pulumi/bin/pulumi: .pulumi.version
	curl -fsSL https://get.pulumi.com | sh -s -- --version "$(PULUMI_VERSION)"

generate_nodejs: .pulumi/bin/pulumi
	rm -rf sdk/nodejs
	mkdir sdk/nodejs
	echo "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/nodejs/go.mod'
	.pulumi/bin/pulumi package gen-sdk provider/cmd/pulumi-resource-aws-native/schema.json --language nodejs --version "$(VERSION_GENERIC)"

build_nodejs::
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run build && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/

generate_python: .pulumi/bin/pulumi
	rm -rf sdk/python
	mkdir sdk/python
	echo "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/python/go.mod'
	.pulumi/bin/pulumi package gen-sdk provider/cmd/pulumi-resource-aws-native/schema.json --language python --version "$(VERSION_GENERIC)"

build_python::
	cd sdk/python/ && \
        cp ../../README.md . && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
	python3 -m venv venv && \
	./venv/bin/python -m pip install build && \
        cd ./bin && \
        ../venv/bin/python -m build .

generate_dotnet: .pulumi/bin/pulumi
	rm -rf sdk/dotnet
	mkdir sdk/dotnet
	echo "module fake_dotnet_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/dotnet/go.mod'
	.pulumi/bin/pulumi package gen-sdk provider/cmd/pulumi-resource-aws-native/schema.json --language dotnet --version "$(VERSION_GENERIC)"

build_dotnet::
	cd ${PACKDIR}/dotnet/ && \
		echo "${PACK}\n${VERSION_GENERIC}" > version.txt && \
		dotnet build

generate_java:: bin/pulumi-java-gen
	rm -rf sdk/java
	mkdir sdk/java
	echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/java/go.mod'
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema $(WORKING_DIR)/provider/cmd/$(PROVIDER)/schema.json --out sdk/java --build gradle-nexus

build_java:: PACKAGE_VERSION := $(VERSION_GENERIC)
build_java::
	cd ${PACKDIR}/java/ && \
		gradle --console=plain build

generate_go: .pulumi/bin/pulumi
	rm -rf sdk/go && mkdir sdk/go
	.pulumi/bin/pulumi package gen-sdk provider/cmd/pulumi-resource-aws-native/schema.json --language go --version "$(VERSION_GENERIC)"

build_go::
	cd sdk/ && go build github.com/pulumi/pulumi-aws-native/sdk/go/aws/...

bin/pulumi-java-gen::
	pulumictl download-binary -n pulumi-language-java -v $(JAVA_GEN_VERSION) -r pulumi/pulumi-java

clean::
	rm -rf sdk/nodejs && mkdir sdk/nodejs && echo "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/nodejs/go.mod'
	rm -rf sdk/python && mkdir sdk/python && echo "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/python/go.mod'
	rm -rf sdk/dotnet && mkdir sdk/dotnet && echo "module fake_dotnet_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/dotnet/go.mod'
	rm -rf sdk/java && mkdir sdk/java && echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > 'sdk/java/go.mod'
	rm -rf sdk/go

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_java_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

test:: PATH := $(WORKING_DIR)/bin:$(PATH)
test::
	cd examples && go test -v -tags=all -timeout 2h

build:: clean codegen local_generate provider build_sdks install_sdks
build_sdks: build_nodejs build_dotnet build_python build_go build_java
install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

ref-db::
	mkdir -p bin
	(cd provider/tools/ref-parser && go build -o ../../../bin/ref-parser)
	./bin/ref-parser -guide ./aws-cloudformation-user-guide -schema ./aws-cloudformation-schema -db ./meta/ref-db.json

ref-db-auto::
	mkdir -p bin
	(cd provider/tools/ref-parser && go build -o ../../../bin/ref-parser)
	./bin/ref-parser -guide ./aws-cloudformation-user-guide -schema ./aws-cloudformation-schema -db ./meta/ref-db.json -auto

ref-db-report::
	mkdir -p bin
	(cd provider/tools/ref-parser && go build -o ../../../bin/ref-parser)
	./bin/ref-parser -guide ./aws-cloudformation-user-guide -schema ./aws-cloudformation-schema -db ./meta/ref-db.json -report

.PHONY: ensure generate_schema generate build_provider build

# Set these variables to enable signing of the windows binary
AZURE_SIGNING_CLIENT_ID ?=
AZURE_SIGNING_CLIENT_SECRET ?=
AZURE_SIGNING_TENANT_ID ?=
AZURE_SIGNING_KEY_VAULT_URI ?=
SKIP_SIGNING ?=

bin/jsign-6.0.jar:
	wget https://github.com/ebourg/jsign/releases/download/6.0/jsign-6.0.jar --output-document=bin/jsign-6.0.jar

sign-goreleaser-exe-amd64: GORELEASER_ARCH := amd64_v1
sign-goreleaser-exe-arm64: GORELEASER_ARCH := arm64

# Set the shell to bash to allow for the use of bash syntax.
sign-goreleaser-exe-%: SHELL:=/bin/bash
sign-goreleaser-exe-%: bin/jsign-6.0.jar
	@# Only sign windows binary if fully configured.
	@# Test variables set by joining with | between and looking for || showing at least one variable is empty.
	@# Move the binary to a temporary location and sign it there to avoid the target being up-to-date if signing fails.
	@set -e; \
	if [[ "${SKIP_SIGNING}" != "true" ]]; then \
		if [[ "|${AZURE_SIGNING_CLIENT_ID}|${AZURE_SIGNING_CLIENT_SECRET}|${AZURE_SIGNING_TENANT_ID}|${AZURE_SIGNING_KEY_VAULT_URI}|" == *"||"* ]]; then \
			echo "Can't sign windows binaries as required configuration not set: AZURE_SIGNING_CLIENT_ID, AZURE_SIGNING_CLIENT_SECRET, AZURE_SIGNING_TENANT_ID, AZURE_SIGNING_KEY_VAULT_URI"; \
			echo "To rebuild with signing delete the unsigned windows exe file and rebuild with the fixed configuration"; \
			if [[ "${CI}" == "true" ]]; then exit 1; fi; \
		else \
			file=dist/build-provider-sign-windows_windows_${GORELEASER_ARCH}/pulumi-resource-aws-native.exe; \
			mv $${file} $${file}.unsigned; \
			az login --service-principal \
				--username "${AZURE_SIGNING_CLIENT_ID}" \
				--password "${AZURE_SIGNING_CLIENT_SECRET}" \
				--tenant "${AZURE_SIGNING_TENANT_ID}" \
				--output none; \
			ACCESS_TOKEN=$$(az account get-access-token --resource "https://vault.azure.net" | jq -r .accessToken); \
			java -jar bin/jsign-6.0.jar \
				--storetype AZUREKEYVAULT \
				--keystore "PulumiCodeSigning" \
				--url "${AZURE_SIGNING_KEY_VAULT_URI}" \
				--storepass "$${ACCESS_TOKEN}" \
				$${file}.unsigned; \
			mv $${file}.unsigned $${file}; \
			az logout; \
		fi; \
	fi

# To make an immediately observable change to .ci-mgmt.yaml:
#
# - Edit .ci-mgmt.yaml
# - Run make ci-mgmt to apply the change locally.
#
ci-mgmt: .ci-mgmt.yaml
	go run github.com/pulumi/ci-mgmt/provider-ci@b6bfde4bf3d1f9e539671e20aad7801e4ba5d300 generate
.PHONY: ci-mgmt
	fi
