# gh-2390 Refresh Input/Output Semantics

## Summary

Issue [#2390](https://github.com/pulumi/pulumi-aws-native/issues/2390) exposes a provider-runtime modeling bug in AWS Native refresh. The old provider path projected live CloudControl outputs into input shape and applied all detected refresh drift to checkpointed `__inputs`.

That is useful for managed drift: if the user set `myField = "abc"` and AWS now has `"xyz"`, refresh can make the next preview compare old input `"xyz"` to program input `"abc"` and repair the drift.

It is wrong for unowned optional-computed-style fields: if the user never set `backupTarget`, AWS defaults/returns `backupTarget = "region"`, and a provider upgrade now exposes that property as an input, refresh must not make the next preview compare old input `"region"` to an absent program input. That would report a user-requested diff that does not exist.

The target model is:

- Outputs remain the source of truth for live cloud state.
- Refresh returns `Inputs` as the engine's next old-input baseline, not as a pure user-intent record.
- That old-input baseline must be ownership-filtered:
  - managed paths are updated to the live input-shaped actual value;
  - unowned optional-computed paths are not added merely because AWS returns them.
- The Pulumi engine still plans standard updates/replacements by diffing old inputs against new program inputs, so the provider must shape refreshed inputs accordingly.
- Patch generation may still need actual output state as the CloudControl patch baseline, especially where input checkpoint shape and output shape diverge.

## Implementation Flow

### Checkpoint State

AWS Native stores output state plus a provider-owned secret `__inputs` field.

- `provider/pkg/resources/checkpoint.go:15` copies outputs into a checkpoint map.
- `provider/pkg/resources/checkpoint.go:23` stores `__inputs` as a secret object.
- `provider/pkg/resources/checkpoint.go:28` reads `__inputs` back out of output state.

This means old inputs are durable even when the provider reconstructs them from output state.

### Read / Refresh

For standard resources, `Read` starts with old output state from `req.Properties` and extracts old inputs from checkpointed `__inputs`.

- `provider/pkg/provider/provider.go:1003` unmarshals old state.
- `provider/pkg/provider/provider.go:1010` extracts `inputs := resources.ParseCheckpointObject(oldState)`.
- `provider/pkg/provider/provider.go:1032` reads live state from CloudControl.
- `provider/pkg/provider/provider.go:1041` converts CloudFormation casing to SDK casing.

If `inputs == nil`, `Read` treats the operation like import and derives inputs from live state:

- `provider/pkg/provider/provider.go:1043` enters the import/no-checkpoint branch.
- `provider/pkg/provider/provider.go:1047` creates a `PathClassifier`.
- `provider/pkg/provider/provider.go:1048` calls `classifier.ProjectActualInputs`.

If `inputs != nil`, `Read` builds ownership-filtered refreshed inputs:

1. Create a `PathClassifier` for the resource schema.
2. Rehydrate non-create-only write-only fields from checkpointed `__inputs` into the live state because CloudControl cannot return them.
3. Project live state into writable input shape.
4. Build an actual input baseline from projected live state plus the checkpointed ownership source.
5. Prune unowned optional-computed fields from that baseline.
6. Suppress AWS-managed and normalized changes before checkpointing refreshed inputs.
7. Store the refreshed inputs in `__inputs` and return them in `ReadResponse.Inputs`.

The recursive helper lives in `provider/pkg/resources/path_classifier.go`. It projects writable input shape, classifies nested schema paths, expands wildcard metadata paths, preserves write-only fallbacks, prunes unowned optional-computed fields, and applies AWS-managed suppression before refresh writes a new `__inputs` baseline.

### Diff

`Diff` continues to reconstruct old inputs from `req.Olds.__inputs`. Protocol `OldInputs` adoption is intentionally deferred to a follow-up PR.

- `provider/pkg/provider/provider.go:824` reads new desired inputs from `req.News`.
- `provider/pkg/provider/provider.go:834` reads old output state from `req.Olds`.
- `provider/pkg/provider/provider.go:845` extracts old inputs from `__inputs`.
- `provider/pkg/provider/provider.go:849` creates a `PathClassifier`.
- `provider/pkg/provider/provider.go:850` projects old output state into writable input shape.
- `provider/pkg/provider/provider.go:851` builds the actual input baseline.
- `provider/pkg/provider/provider.go:857` suppresses AWS-managed and normalized changes.

AWS Native returns `DIFF_UNKNOWN` plus `DeleteBeforeReplace: true` for standard-resource changes. The Pulumi engine normalizes `DIFF_UNKNOWN` by computing its own old-input/new-input diff. Generated SDK `replaceOnChanges` entries only get useful property attribution if the refreshed old inputs differ from new program inputs on the relevant paths.

### Update / Patch

`Update` follows the same old-input model:

- `provider/pkg/provider/provider.go:1121` reads new desired inputs from `req.News`.
- `provider/pkg/provider/provider.go:1131` reads old output state from `req.Olds`.
- `provider/pkg/provider/provider.go:1113` extracts old inputs from `__inputs`.
- `provider/pkg/provider/provider.go:1155` creates a `PathClassifier`.
- `provider/pkg/provider/provider.go:1156` projects old output state into writable input shape.
- `provider/pkg/provider/provider.go:1157` calls `resources.CalcPatchWithActualBaseline(...)`.

`CalcPatchWithActualBaseline` builds CloudControl patches from actual projected state to new desired inputs. It still uses checkpointed `__inputs` for ownership and write-only fallback, then applies AWS-managed/property-transform suppression before converting the diff to a CloudControl JSON patch.

## Problem

The old refresh model applied input-shaped actual drift without checking ownership.

That is correct when the user owns the path:

1. User creates with `myField = "abc"`.
2. AWS is externally changed to `myField = "xyz"`.
3. Refresh returns old inputs with `myField = "xyz"`.
4. Next preview compares old input `"xyz"` to program input `"abc"` and detects the repair diff.

It is wrong when the user does not own the path:

1. User creates with no `backupTarget`.
2. AWS defaults/returns `backupTarget = "region"`.
3. A newer provider version now includes `backupTarget` as an input and output.
4. Refresh returns old inputs with `backupTarget = "region"`.
5. Next preview compares old input `"region"` to absent program input and reports a diff the user did not request.

The bug is not that refresh mutates inputs at all. The bug is that refresh mutates inputs for both managed and unmanaged paths. With Pulumi's current engine, refresh must mutate old inputs for managed drift so that the engine can later diff old inputs against program inputs. It must not adopt unowned optional-computed values.

This design is forward-looking. It prevents future refreshes from adopting newly visible unowned optional-computed values into old inputs. It does not attempt to scrub checkpoints already polluted by prior provider versions that refreshed and saved AWS-returned defaults into `__inputs`.

## Desired Semantics

1. Output state represents live actual cloud state after refresh.
2. Engine inputs after refresh represent an ownership-filtered input-shaped actual baseline for the next engine diff.
3. Refresh must update outputs to the live CloudControl state.
4. Refresh must update returned inputs for paths the user previously managed when CloudControl can read the live value.
5. Refresh must not add newly visible optional input/output overlap fields that the user never managed merely because AWS returns them.
6. Old desired inputs are still needed to decide ownership, especially for optional-computed-style fields, nested objects, maps, arrays, and write-only properties.
7. A user who continues to set `myField = "abc"` after refresh sees AWS drift to `"xyz"` should get a later engine diff from old input `"xyz"` to new input `"abc"`.
8. A user who never set optional-computed `backupTarget`, while AWS returns default `"region"`, should get old inputs that still omit `backupTarget`.
9. A user who later explicitly sets optional-computed `backupTarget` transitions that property into ownership. If the actual value already matches, no cloud operation is required; ownership persistence can wait until the next state write.
10. Import remains distinct: when there is no prior desired-input record, the provider may infer initial inputs from live state.
11. If a path was present in old inputs and absent from new program inputs, the user intentionally removed it. For create-only paths, removal should be replacement-significant.
12. Map and array ownership is at the containing property. If the user manages a map or array input, provider refresh and patch generation should reconcile the actual collection back to the desired collection, except for explicitly suppressed AWS-managed values such as `aws:*` tags.

## Design

### Ownership-Filtered Refresh Inputs

`Read` should shape the `Inputs` it returns to the engine. Those inputs are the baseline the engine will compare against the next program inputs.

For existing standard resources, `Read` should:

1. Read live CloudControl state.
2. Return checkpoint `Properties` containing live outputs.
3. Project live outputs into writable input shape.
4. Build refreshed inputs by applying ownership rules:
   - for managed paths, use the live projected actual value;
   - for unowned optional-computed paths, keep the path absent;
   - for write-only paths, preserve the old input value because CloudControl cannot return live actual state;
   - for maps and arrays, apply ownership at the containing property.
5. Store the same refreshed inputs in `__inputs`.
6. Return those refreshed inputs in `ReadResponse.Inputs`.

This is not a generic actual-state projection. It is the old-input baseline Pulumi needs for the next old-input/new-input diff.

### Engine Diff Contract

AWS Native standard-resource `Diff` currently returns `DIFF_UNKNOWN` for changes. Pulumi handles that by computing an input diff from old inputs to new inputs and then applying `replaceOnChanges`.

That means GH-2390 should preserve the existing SDK/engine replacement path:

- generated SDKs convert CloudFormation `createOnlyProperties` into Pulumi `replaceOnChanges`;
- refresh makes managed create-only drift visible as an old-input/new-input difference;
- the engine marks matching `replaceOnChanges` paths as replacements;
- provider `Update` is not called for normal managed create-only changes.

Do not change standard-resource `Diff` from `DIFF_UNKNOWN` to bare `DIFF_SOME` unless it also returns correct `ChangedKeys`/`DetailedDiff`. Pulumi does not backfill property attribution for bare `DIFF_SOME`, so generated `replaceOnChanges` would have no changed paths to match and the engine could plan an in-place `Update`.

### Update Patch Baseline

Refresh input shaping is about engine planning. CloudControl patch generation still needs to avoid invalid patches.

`Update` / `CalcPatch` should build CloudControl patches from the best available actual baseline to new desired inputs:

- normally use live output state projected into writable input shape;
- use old input values for write-only paths because CloudControl cannot return them;
- preserve existing property transform and AWS-managed suppression behavior.

This patch baseline is not a replacement for ownership-aware refresh inputs. It is needed because CloudControl patches should describe actual-to-desired mutations, not stale-checkpoint-to-desired mutations.

Do not add provider-side create-only patch filtering as a speculative fallback.
Under the GH-2390 model, create-only changes are handled before standard-resource
`Update` by generated SDK `replaceOnChanges` plus engine diff normalization. If
a concrete reachable `Update` path for create-only diffs is identified later,
handle that path explicitly with evidence.

### Ownership Classification

The provider needs a recursive classifier that can answer both projection and ownership questions.

Minimum helper responsibilities:

- project output state into writable input shape;
- read, write, and delete nested property paths such as `code/imageUri`;
- classify paths as concrete schema fields, map keys, array elements, read-only paths, write-only paths, and create-only paths;
- decide whether a concrete schema field is optional-computed-style:
  - present in inputs;
  - present in outputs;
  - not read-only;
  - not required;
- distinguish schema-declared object fields from map keys and array elements;
- match wildcard metadata paths such as `foo/*/bar` against concrete array paths such as `foo/0/bar`.

The helper should use SDK-name paths with `/` separators, matching metadata entries such as `code/imageUri` and wildcard array paths such as `rules/*/statement`. Concrete get/set/delete operations may use `resource.PropertyPath` under a thin slash-path conversion layer. Wildcard expansion still needs a local walk over actual values because `PropertyPath.Get` and `PropertyPath.Set` do not expand `*` over every concrete array element.

Current metadata has a top-level `CloudAPIResource.Required` field, but `CloudAPIType` does not persist nested required properties even though schema generation computes them for Pulumi schema output. Add `Required []string` to `metadata.CloudAPIType` and populate it from the already-computed nested `requiredSpecs` in schema generation. That keeps optional-computed detection consistent between top-level resource properties and nested object fields.

### Checkpoint Ownership Source

Modern Pulumi protocol requests include old inputs, including
`DiffRequest.OldInputs`, `UpdateRequest.OldInputs`, `DeleteRequest.OldInputs`,
and `ReadRequest.Inputs` during refresh.

Using those protocol fields is intentionally deferred to a follow-up PR to keep
this change focused on the refresh ownership model. This PR continues to use the
existing checkpointed `__inputs` field as the ownership source for standard
resources.

Do not remove `__inputs` in this change; existing stacks and current code paths
depend on it. The refresh ownership decision should treat `__inputs` as the
persisted old desired-input baseline, not as live actual state.

### Write-Only Properties

Write-only properties cannot be read from CloudControl, so live outputs cannot provide an actual value. For these properties:

- refresh should preserve the old input value if the path was managed;
- import should warn because the value cannot be inferred;
- patch generation should keep existing behavior where write-only, non-create-only properties are resent when required for CloudControl model validation;
- if a property is both write-only and create-only, do not include it in update patches.

### Create-Only Properties

Create-only input changes are handled primarily by generated SDK `replaceOnChanges` plus the Pulumi engine.

For managed create-only drift:

```text
old inputs before refresh: backupTarget = "local"
live output:               backupTarget = "region"
refreshed inputs:          backupTarget = "region"
next program inputs:       backupTarget = "local"
engine diff:               region -> local
operation:                 replacement via replaceOnChanges
```

For unowned create-only defaults:

```text
old inputs before refresh: no backupTarget
live output:               backupTarget = "region"
refreshed inputs:          no backupTarget
next program inputs:       no backupTarget
engine diff:               none
```

For removal of a previously managed create-only path:

```text
old inputs after refresh: backupTarget = "region"
next program inputs:      no backupTarget
engine diff:              region -> absent
operation:                replacement via replaceOnChanges
```

Provider patch generation does not add a create-only fallback path for GH-2390.
Create-only input changes, managed create-only drift, and managed create-only
removals must be represented as engine-visible old-input/new-input diffs so
generated SDK `replaceOnChanges` can replace the resource before provider
`Update`.

### AWS-Managed Diff Suppression

`SuppressAWSManagedDiffs` receives checkpointed old desired inputs plus actual/desired context. Under ownership-aware refresh, suppression should remain tied to ownership:

- AWS-managed tag keys such as `aws:*` should not become user-owned merely because AWS returns them.
- Non-`aws:*` keys in a user-owned tag map should remain visible drift.
- Suppression should filter AWS-managed tag keys from both the old actual baseline and the new desired value; otherwise an `aws:*` tag present only in refreshed actual inputs becomes false removal drift.
- Property transforms should receive enough context to compare actual and desired values without reusing a tag-ownership argument for unrelated semantics.

## Edge Cases

### External Drift On Managed Input

State before refresh:

```text
old inputs: myField = "abc"
output:     myField = "abc"
```

Cloud changes to:

```text
live output: myField = "xyz"
```

After refresh:

```text
old inputs: myField = "xyz"
output:     myField = "xyz"
```

Next preview with program still setting `"abc"`:

```text
engine diff: xyz -> abc
operation:   update or replacement depending on property semantics
```

### Newly Visible Optional-Computed Default

State before provider upgrade:

```text
old inputs: no backupTarget
output:     no backupTarget in old schema/checkpoint
```

After provider upgrade and refresh:

```text
old inputs: no backupTarget
output:     backupTarget = "region"
```

Next preview with program still omitting `backupTarget`:

```text
engine diff: none
```

### User Starts Managing Existing Default

After refresh:

```text
old inputs: no backupTarget
output:     backupTarget = "region"
```

Program changes to:

```text
new inputs: backupTarget = "region"
```

Expected behavior:

```text
cloud operation: none, because actual already matches desired
ownership:       user now manages backupTarget, recorded on the next state write
```

The provider should not force a CloudControl update solely to rewrite checkpoint ownership.

### User Starts Managing Different Value

After refresh:

```text
old inputs: no backupTarget
output:     backupTarget = "region"
```

Program changes to:

```text
new inputs: backupTarget = "local"
```

Expected behavior:

```text
engine diff: absent -> local
operation:   update or replacement depending on property semantics
```

For create-only properties, generated `replaceOnChanges` should make this a replacement.

### User Removes Previously Managed Optional-Computed Input

After refresh:

```text
old inputs: backupTarget = "region"
output:     backupTarget = "region"
```

Program changes to:

```text
new inputs: no backupTarget
```

Treat removal as intentional user removal. The resulting operation depends on property kind:

- mutable properties should emit the appropriate remove/unset patch when CloudControl supports it;
- create-only properties should be replacement-significant;
- write-only properties should use old input values as the best available baseline because the actual value cannot be read.

### Maps, Tags, And Arrays

Do not apply a simple "path absent from old inputs means ignore it" rule to maps or arrays. For example, if a user manages a tags map and someone adds a tag key externally, that can be real drift, not a newly visible schema field.

Any optional-computed handling for nested paths must distinguish:

- schema-declared object fields;
- freeform map keys;
- array elements;
- AWS-managed tag keys such as `aws:*`, which already have suppression behavior.

For maps and arrays, ownership is at the containing property. If the user desired the map or array property, actual extra keys/elements should generally be reflected into refreshed old inputs so the next engine diff can reconcile back to desired state, except where existing suppression rules identify AWS-managed values.

### Legacy Or Empty `__inputs`

`ParseCheckpointObject` returns non-nil for an empty `__inputs` object. This design does not perform a checkpoint migration:

- use existing `__inputs` exactly as present;
- if `__inputs` is missing, keep the current import/no-checkpoint behavior;
- if `__inputs` exists but is empty, treat it as empty prior ownership;
- if `__inputs` is stale or polluted, do not attempt automatic repair.

Already-polluted checkpoints are out of scope for the forward fix. Users who already refreshed with a provider version that saved AWS defaults into `__inputs` may need a one-time repair workflow; this design does not guarantee automatic cleanup of that historical state.

## Rejected Shortcuts

### Adopt Every Input-Shaped Output During Refresh

This is the current bug. It repairs managed drift, but it also adopts unowned AWS defaults and newly visible optional-computed values.

### Preserve All Old Inputs During Refresh

This avoids adopting unowned defaults, but it breaks managed drift detection under the current engine contract. AWS Native returns `DIFF_UNKNOWN`, and the engine normalizes `DIFF_UNKNOWN` by diffing old inputs against new inputs. If refresh preserves `myField = "abc"` while outputs show live `myField = "xyz"`, a later no-change program still has old input `"abc"` and new input `"abc"`, so the engine sees no diff.

### Bare Provider `DIFF_SOME` Without Attribution

Returning `DIFF_SOME` without `ChangedKeys` or `DetailedDiff` is not enough. Pulumi does not synthesize property attribution for bare `DIFF_SOME`, so generated `replaceOnChanges` has no paths to match.

### Treat Every Input/Output Overlap As User-Owned

Input/output overlap often means "optional input that AWS may also compute or default." It does not prove the user intended to manage the returned value.

### Treat `__inputs` As Live Actual State

`__inputs` is useful as the provider's persisted old-input baseline. It should not be treated as live actual state when output state records the cloud value and CloudControl can read it.

## Test Plan

### Provider Runtime Tests

Unit tests under `provider/pkg/provider` cover standard-resource lifecycle behavior:

1. **External drift on managed input**
   - Old inputs and old output are `"abc"`.
   - Refresh reads live output `"xyz"`.
   - Refresh returns properties with `"xyz"` and inputs with `"xyz"`.
   - Later `Diff`/engine planning with new program input `"abc"` sees old input `"xyz"` vs new input `"abc"`.

2. **Newly visible optional-computed default**
   - Old inputs omit `backupTarget`.
   - Old output omits `backupTarget`.
   - New live output includes `backupTarget = "region"`.
   - Refresh returns properties with `backupTarget`, but returned inputs still omit it.
   - Later preview with program still omitting `backupTarget` reports no diff.

3. **Managed optional-computed drift**
   - Old inputs include `backupTarget = "local"`.
   - New live output includes `backupTarget = "region"`.
   - Refresh returns inputs with `backupTarget = "region"`.
   - Later preview with program setting `backupTarget = "local"` sees a diff.
   - If `backupTarget` is create-only, the diff is replacement-significant through generated `replaceOnChanges`.

4. **User starts managing same value**
   - Old inputs omit `backupTarget`.
   - Old output has `backupTarget = "region"`.
   - New desired sets `backupTarget = "region"`.
   - Provider reports no cloud operation is needed.
   - Ownership is recorded on the next state write; do not force a CloudControl update solely for checkpoint mutation.

5. **User starts managing different value**
   - Old inputs omit `backupTarget`.
   - Old output has `backupTarget = "region"`.
   - New desired sets `backupTarget = "local"`.
   - Provider reports a diff.
   - If `backupTarget` is create-only, generated `replaceOnChanges` makes the operation a replacement.

6. **User removes managed create-only input**
   - Old inputs include `backupTarget`.
   - New desired omits `backupTarget`.
   - The engine-visible diff is replacement-significant for create-only paths.
   - This remains primarily covered by the earlier engine boundary tests because the provider intentionally keeps create-only replacement enforcement in generated SDK `replaceOnChanges` plus engine diff normalization.

7. **Write-only fallback**
   - A write-only property absent from outputs still uses old input values for refresh and patch generation where current behavior requires it.
   - Nested and wildcard write-only paths such as `code/imageUri` and `defaultActions/*/authenticateOidcConfig/clientSecret` are expanded to concrete paths before rehydration/resend.

8. **Map/tag drift**
   - A user-managed tags map gets an externally added non-`aws:*` key.
   - Refresh reflects that key into old inputs so the next engine diff can reconcile it to desired state.
   - An externally added `aws:*` key remains present in output state but is suppressed from refreshed `__inputs` and from CloudControl update patches.

9. **Nested object field**
   - A schema-declared nested input/output field appears after refresh.
   - The test verifies schema-aware optional-computed behavior and proves the implementation is recursive where needed.

### Helper Tests

Helper tests under `provider/pkg/resources` cover:

- deriving input-shaped actual values from outputs;
- identifying optional-computed-style properties;
- applying ownership rules for absent, present, deleted, nested, map, and array paths;
- updating refreshed inputs only for managed paths;
- reading and writing nested metadata paths such as `code/imageUri` without treating them as flat keys.
- preserving array shape when writing concrete wildcard paths such as `defaultActions/0/authenticateOidcConfig/clientSecret`.
- suppressing `aws:*` tags from both old actual and new desired tag values.

### Commands

Follow `AGENTS.md` local test safety. Do not run broad provider test targets locally by default.

Use targeted tests for changed provider/runtime logic, for example:

```bash
mise exec -- go test -count=1 ./pkg/resources -run 'TestPathClassifier|TestCalcPatch|TestSuppressAWSManagedDiffs'
mise exec -- go test -count=1 ./pkg/provider -run 'TestStandardResource(Read|Diff|Update)'
```

If schema metadata or generated artifacts change, use the repo command canon from `AGENTS.md` instead of hand-editing generated files.

## Open Decisions

No semantic open decisions remain. Implementation should validate the helper shape against concrete tests and adjust names/signatures only where the code proves a simpler equivalent API.
