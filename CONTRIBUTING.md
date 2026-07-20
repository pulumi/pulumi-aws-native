# Contributing to pulumi-aws-native

## Prerequisites
Use the same toolchain contract as CI and prepare local generated artifacts:

```bash
mise exec -- make prepare_local_workspace
```

This installs the tools pinned in `.config/mise.toml`, downloads Go modules, initializes submodules, and creates the
ignored schema/metadata artifacts required by local builds and tests.

## Quick Flow
1. Review `AGENTS.md` (agent instructions) and `ARCHITECTURE.md` (module map and system boundaries).
2. Determine whether target files are hand-written or generated.
3. Implement focused changes.
4. Run verification commands.
5. Open a PR with command evidence.

## Verification Commands
For most provider code changes:

```bash
make lint
make test_provider_fast
```

For broader provider changes:

```bash
make test_provider
```

For slow integration validation (AWS creds required):

```bash
make test
```

## Regeneration Rules
If generation inputs changed, regenerate before PR:

- Schema/codegen changes: `make codegen && make generate_schema`
- SDK/API surface changes: `make local_generate`
- CI config changes (`.ci-mgmt.yaml`): `make ci-mgmt`

## Generated Artifacts Policy
Do not hand-edit generated files:

- `sdk/**`
- generated schema artifacts under `provider/cmd/pulumi-resource-aws-native/`
- `.github/workflows/**` and `.config/mise.toml` (ci-mgmt generated)

If generated outputs change, include the regeneration command in your PR notes.
