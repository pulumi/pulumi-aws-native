# Agent Instructions

## Purpose
This repository contains the Pulumi AWS Native provider: a Go provider plus generated SDKs (Node.js, Python, .NET, Java, Go).
Read `ARCHITECTURE.md` for a concise module map and system boundaries before making structural changes.

## High-Signal Paths
- `Makefile`: canonical command surface
- `provider/pkg/`: hand-written provider runtime logic
- `provider/cmd/pulumi-gen-aws-native/`: schema/codegen logic
- `provider/cmd/pulumi-resource-aws-native/`: provider entrypoint + generated schema artifacts
- `meta/`: metadata inputs to generation
- `examples/`: integration examples/tests
- `.ci-mgmt.yaml`: source for generated CI config
- `.config/mise.toml`: pinned toolchain versions

## Generated vs Hand-Written
Hand-written (safe to edit directly):
- `provider/pkg/**`
- `provider/cmd/**` (except generated outputs under `provider/cmd/pulumi-resource-aws-native/`)
- `meta/**`
- `examples/**`
- `README.md`, `Makefile`, `.ci-mgmt.yaml`

Generated or externally sourced (do not hand-edit):
- `sdk/nodejs/**`, `sdk/python/**`, `sdk/dotnet/**`, `sdk/java/**`, `sdk/go/**`
- `provider/cmd/pulumi-resource-aws-native/schema.json` and generated metadata artifacts
- `.github/workflows/**` (generated via `make ci-mgmt`)
- `.config/mise.toml` (generated via `make ci-mgmt`)
- `aws-cloudformation-schema/**`, `aws-cloudformation-user-guide/**` submodule content

## Command Canon
Run all `make` commands through `mise exec -- make ...` so the pinned toolchain is used.

- Workspace bootstrap: `make prepare_local_workspace`
- Dependency/setup: `make ensure`
- Build codegen binary: `make codegen`
- Build provider binary: `make provider`
- Regenerate schema: `make generate_schema`
- Regenerate all SDKs: `make local_generate`
- Lint provider code: `make lint`
- Fast provider tests: `make test_provider_fast`
- Full provider tests: `make test_provider`
- Integration suite (slow): `make test`
- Regenerate CI from `.ci-mgmt.yaml`: `make ci-mgmt`

### Local Test Safety
- Do not run package-wide provider tests by default on this machine. In particular, avoid `mise exec -- make test_provider_fast`, `mise exec -- make test_provider`, and broad `go test ./pkg/...` / `go test ./pkg/provider` commands unless the user explicitly asks for them.
- The `provider.test` process can grow until the machine becomes unresponsive when broad provider test targets run locally.
- Prefer targeted tests for the changed code, for example `mise exec -- go test -count=1 ./pkg/resources -run '<TestNamePattern>'` or `mise exec -- go test -count=1 ./pkg/provider -run '<TestNamePattern>'`.
- If broader validation is needed, explain the risk and hand it off to CI instead of running it locally.

## If You Change X, Also Do Y
- `provider/pkg/**`: normally `make lint && make test_provider_fast`, but on this machine follow Local Test Safety above and run targeted `go test -run` filters unless the user explicitly asks for the broad target.
- codegen/schema logic (`provider/cmd/pulumi-gen-aws-native/**`, schema overlays, `meta/**`): `make codegen && make generate_schema` (example: new schema overlay)
- SDK surface: `make local_generate`
- `.ci-mgmt.yaml`: `make ci-mgmt`
- `provider/go.mod` or `examples/go.mod`: `make ensure`

## Forbidden Actions
- Do not hand-edit generated SDK files under `sdk/**`.
- Do not hand-edit `.github/workflows/**` or `.config/mise.toml`; regenerate through ci-mgmt.
- Do not use destructive git commands (`git reset --hard`, force push) unless explicitly approved.
- Do not claim lint/tests passed unless the commands were actually run.

## Escalate When
- Requested edits conflict with generated-file boundaries.
- Public SDK surface changes are requested but regeneration scope is unclear.
- CI/local behavior diverges after following command canon.
