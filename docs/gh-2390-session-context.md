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

`mise exec -- make test_provider_fast` is acceptable for broad fast provider
validation on this machine.

Do not run the full provider suite locally by default. Avoid:

```bash
mise exec -- make test_provider
go test ./pkg/...
go test ./pkg/provider
```

The full provider suite can run provider 2e2 tests that require live AWS
credentials and may spend several minutes retrying EC2 IMDS when credentials are
unavailable. Use targeted tests for narrow changes, for example:

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

The earlier broad-test memory pressure was fixed. A later local run of
`mise exec -- make test_provider_fast` passed without visible memory pressure.
`mise exec -- make test_provider` still failed locally because live AWS
credential resolution fell back to EC2 IMDS retries.

## Review Follow-Up Required

The current branch has review feedback from the user plus a Claude review. Do
not work the lower-priority Claude items unless the user asks. In particular,
leave polluted legacy `__inputs`, type-graph cycle hardening, and map-of-object
optional-computed pruning alone for this PR.

Latest implementation pass addressed items 1, 2, 3, 5, and 6 below. Item 4 is
explicitly deferred to #2944 because precise nested CloudControl patch
generation would expand this PR beyond the GH-2390 refresh/input semantics.

Status of review items:

1. **Done: make `Diff` explicit about custom resources vs missing specs.**
   `provider/pkg/provider/provider.go` currently falls back to
   `oldInputs.Diff(newInputs)` when `resourceMap.Resources[resourceToken]` has
   no spec. That fallback exists for custom resources, because
   `CustomResource` has no `Diff` method, but it also silently covers impossible
   missing-standard-spec cases. The control flow now separates:
   - standard resource with spec: ownership-aware baseline diff;
   - custom resource: old-input diff fallback;
   - neither: hard error, matching `Check`/`Create`/`Read`/`Update`/`Delete`.

2. **Done: fix nested wildcard write-only fallback when AWS omits the parent array.**
   `SetPath` / `slashPathForValue` only treats numeric segments as array
   indices when the destination already contains an array. If AWS omits the
   parent array and old desired inputs contain a wildcard write-only path such
   as `defaultActions/0/authenticateOidcConfig/clientSecret`, rehydration can
   create `defaultActions: {"0": ...}` instead of an array. The setter now uses
   old desired input shape as a guide, with focused tests for the omitted-parent
   array case.

3. **Done: extract shared write-only output rehydration.**
   The same nested write-only copy loop exists in `Create` and `Update`, and
   `Read` used the same concept through `AddWriteOnlyFallbacks`. A single
   helper now copies write-only values from inputs/old desired into output state
   when AWS did not return them, and it handles nested paths, wildcard
   expansion, and the array-shape fix above.

4. **Deferred: resolve nested write-only patch resend semantics under #2944.**
   `provider/pkg/resources/patching.go` currently truncates nested write-only
   paths such as `code/imageUri` to the top-level `code` key before forcing an
   add patch. This is because `resource.ObjectDiff` and `naming.DiffToPatch`
   operate at top-level input keys, but it means the provider resends the whole
   parent object. This is related to #2944 and is not required for this PR.
   Keep this PR focused on preserving write-only state correctly and avoiding
   create-only write-only update resends.

5. **Done: split create-only-write-only state rehydration from update resend.**
   `addWriteOnlyFallbacks` skips write-only paths that are also create-only,
   which is correct for update patch resend, but incorrect for
   refresh/create output rehydration. The implementation uses two modes:
   refresh/create output state carries create-only-write-only values forward for
   checkpoint continuity, while update patch baselines exclude them so
   CloudControl does not receive create-only data during update.

6. **Done: keep `ProjectWritableOutputState` / `ActualInputBaseline` hard to misuse.**
   Review feedback asked whether these should be private. The concern is valid:
   `ProjectWritableOutputState` is only a projection and is not safe to return as
   refreshed inputs for existing resources. `ActualInputBaseline` is now
   private; provider lifecycle callers use `ActualInputBaselineFromOutputs`.
   `ProjectWritableOutputState` remains public for the import/no-checkpoint
   branch.

PR notes to include when opening the PR:

- The `Create`/`Update` write-only output change primarily adds nested and
  wildcard write-only support; the old `inputs.Mappable()[writeOnlyProp]` path
  only worked for flat keys.
- Import changed from shallow `schema.GetInputsFromState` to recursive writable
  output projection. That means import still infers inputs from live state when
  there is no prior ownership record, but now excludes nested read-only and
  write-only values. This is related to #2941 and is an acceptable direction
  for this PR, but it may only partially address #2941 because import can still
  infer provider-observed writable defaults. Call this implication out in the PR
  because it is a behavior change.
- Update-time AWS-managed/property-transform suppression is intentional because
  patch generation now compares live actual output baseline against desired
  inputs. Without suppression, AWS-managed tags or normalized values can become
  false removal/replace patches. Consider adding a property-transform update
  test if current coverage only proves tag suppression.

## Follow-Up Work

1. Add protocol `Inputs`/`OldInputs` usage for standard resources in a follow-up
   PR.
2. Track precise nested CloudControl patch generation under #2944.
3. Leave full provider validation to CI unless explicitly requested locally with
   suitable AWS credentials.
