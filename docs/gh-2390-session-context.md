# gh-2390 Session Context

This file is the implementation handoff for
`docs/gh-2390-refresh-input-output-semantics.md`. Keep durable semantics in the
design doc; keep this file focused on current build state, verification, and
handoff notes.

## Source Of Truth

- Design contract: `docs/gh-2390-refresh-input-output-semantics.md`
- Runtime-design skill: `.agents/skills/aws-native-provider-runtime-design/SKILL.md`
- Repo instructions and local test safety: `AGENTS.md`

## Current State

Implemented in this worktree:

- `metadata.CloudAPIType` carries nested `Required` metadata, populated by
  `provider/pkg/schema/gen.go`; `metadata.json` was regenerated.
- Standard-resource `Read` returns ownership-filtered refreshed inputs:
  managed paths use live projected actual values, unowned optional-computed
  paths stay absent, and checkpointed `__inputs` remains the ownership source.
- Standard-resource `Diff` preserves the existing `DIFF_UNKNOWN` engine path.
- `Update` patch generation uses live output state projected into writable input
  shape as the actual baseline for mutable updates, with AWS-managed and
  property-transform suppression applied before CloudControl patch conversion.
- Nested and wildcard write-only rehydration/resend is handled for `Create`,
  `Read`, and `Update`.
- Array element classification and pruning is recursive, so nested
  optional-computed fields inside array elements follow the same ownership rules
  as nested object fields.
- Concrete slash-path get/set/delete helpers use `resource.PropertyPath`
  underneath; AWS Native still keeps slash metadata paths and local wildcard
  expansion because `PropertyPath.Get`/`Set` do not expand `*` over every
  matching concrete array element.
- AWS-managed tag suppression filters `aws:*` tags from both old actual and new
  desired tag values, so refreshed actual baselines do not become false removal
  drift or persisted `__inputs` ownership.
- Provider-level tests assert exact CloudControl update patches for mutable
  actual-baseline updates, write-only resend, AWS-managed tag suppression, and
  array ownership pruning.

## Create-Only Boundary

Do not add provider-side create-only patch filtering as a speculative safety
guard. Create-only input changes, managed create-only drift, and managed
create-only removals are handled before standard-resource `Update` by generated
SDK `replaceOnChanges` plus engine diff normalization.

If CI or review identifies a concrete reachable standard-resource `Update` path
for create-only diffs, handle that path explicitly with evidence. Otherwise keep
create-only behavior owned by the SDK/engine replacement path.

## Deferred Protocol Inputs

This PR intentionally continues to use checkpointed `__inputs` as the ownership
source for standard resources. Using protocol fields such as
`ReadRequest.Inputs`, `DiffRequest.OldInputs`, and `UpdateRequest.OldInputs` is
deferred to a follow-up PR to keep this change scoped to refresh ownership and
actual-baseline patch generation.

## Local Test Safety

Do not run broad provider test targets on this machine unless the user explicitly
asks for them.

Avoid:

```bash
mise exec -- make test_provider_fast
mise exec -- make test_provider
go test ./pkg/...
go test ./pkg/provider
```

The `provider.test` process can grow until the machine becomes unresponsive.
Use targeted tests instead, for example:

```bash
mise exec -- go test -count=1 ./pkg/resources -run 'TestPathClassifier|TestCalcPatch|TestSuppressAWSManagedDiffs'
mise exec -- go test -count=1 ./pkg/provider -run 'TestStandardResource(Read|Diff|Update)'
```

Broad validation should be handed off to CI.

## Verification

Current implementation pass:

```bash
cd provider
mise exec -- go test -count=1 ./pkg/resources -run 'TestPathClassifier|TestPathHelpers|TestExpandMatchingPaths|TestCalcPatch|TestSuppressAWSManagedDiffs|TestSuppressAWSManagedTagAdditions'
mise exec -- go test -count=1 ./pkg/provider -run 'TestStandardResourceDiffUsesActualOutputBaseline|TestStandardResourceReadReturnsOwnershipFilteredInputs|TestRead|TestUpdate'
```

Earlier engine boundary verification:

```bash
cd /Users/chall/work/pulumi/pkg
go test ./resource/deploy -run TestApplyReplaceOnChangesEmptyDetailedDiff -count=1
go test ./engine/lifecycletest -run 'TestReplaceOnChanges/(provider diff|engine diff)' -count=1
```

The broad `make test_provider_fast` target was attempted earlier and caused
machine pressure through `provider.test`; do not repeat it locally.

## Follow-Up Work

1. Add protocol `Inputs`/`OldInputs` usage for standard resources in a follow-up
   PR.
2. Leave broad provider validation to CI unless explicitly requested locally.
