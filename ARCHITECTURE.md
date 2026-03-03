# Architecture

## Module Map
- `provider/pkg/`: core provider runtime logic
- `provider/cmd/pulumi-resource-aws-native/`: provider entrypoint and generated schema artifacts
- `provider/cmd/pulumi-gen-aws-native/`: schema ingestion and generation logic
- `provider/cmd/cf2pulumi/`: CloudFormation-to-Pulumi conversion CLI
- `provider/tools/ref-parser/`: reference metadata generation helper
- `meta/`: metadata inputs (`ref-db.json`, `regions.json`)
- `examples/`: integration examples/tests
- `sdk/`: generated SDK outputs for each language
- `aws-cloudformation-schema/`, `aws-cloudformation-user-guide/`: upstream submodule inputs

## Ownership Boundaries
- Hand-written behavior should be implemented in `provider/pkg/**` and selected `provider/cmd/**`.
- Generated SDKs in `sdk/**` must be regenerated, not manually edited.
- CI workflow/toolchain files are generated from ci-mgmt inputs.

## Common Change Patterns
Provider runtime bug fix:
- Edit `provider/pkg/**`
- Add/update tests in `provider/pkg/**/*_test.go`
- Run `make lint && make test_provider_fast`

Schema/codegen behavior change:
- Edit `provider/cmd/pulumi-gen-aws-native/**` and/or `meta/**`
- Run `make codegen && make generate_schema`
- If SDK surface changed, run `make local_generate`

CI config change:
- Edit `.ci-mgmt.yaml`
- Run `make ci-mgmt`
