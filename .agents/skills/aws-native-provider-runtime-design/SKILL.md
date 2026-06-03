---
name: aws-native-provider-runtime-design
description: Use when planning or designing substantial Pulumi AWS Native provider runtime changes, especially lifecycle semantics around Check, Diff, Read, refresh, Update, imports, CloudControl patches, schema input/output mapping, read-only/write-only/create-only properties, optional-computed behavior, or provider upgrade regressions. Do not use for routine generated SDK churn, small local bug fixes, or straightforward schema/codegen edits unless provider lifecycle semantics are being debated.
---

# AWS Native Provider Runtime Design

Use this skill to turn ambiguous provider-runtime behavior into a repo-grounded design or implementation plan. The goal is not a generic planning ritual; it is to preserve the AWS Native distinctions that are easy to blur.

## Start With The Current Flow

Before drafting a plan or spec, inspect the relevant code and summarize the actual lifecycle:

- The issue, repro, PR comment, or failing test that motivated the work.
- The Pulumi RPCs involved: `Check`, `Diff`, `Create`, `Read`, `Update`, `Delete`, import, or refresh.
- What each RPC receives and returns: old desired inputs, new desired inputs, old output state, live CloudControl state, and checkpointed provider state.
- The relevant metadata in `metadata.CloudAPIResource`: `Inputs`, `Outputs`, `ReadOnly`, `WriteOnly`, `CreateOnly`, `PrimaryIdentifier`, `PropertyTransforms`, and `TagsStyle`.
- The local code path, with file/line references. For provider lifecycle work, start in `provider/pkg/provider/provider.go`, then follow helpers in `provider/pkg/resources`, `provider/pkg/schema`, `provider/pkg/naming`, and `provider/pkg/outputs`.

Do not ask the user to explain code facts that can be read from the repo.

## Decide The Artifact

Use the smallest artifact that reduces real risk:

- **Chat-only plan**: narrow behavior change with obvious implementation and low migration risk.
- **Checked-in design spec**: cross-cutting lifecycle semantics, schema-evolution behavior, migration risk, or multiple plausible models.
- **Handoff/status doc**: multi-session work where current decisions, blockers, and verification state must survive context loss.

Prefer one provider-runtime behavior/design spec over separate product and tech specs unless the user explicitly wants that split.

## Spec Shape

When a checked-in spec is warranted, use a compact Markdown document with these sections:

```markdown
# <issue-or-topic> Provider Runtime Semantics

## Summary
What behavior is wrong, what semantic model should replace it, and which issue/PR this serves.

## Current Flow
Lifecycle walkthrough grounded in code. Distinguish old desired inputs, new desired inputs, old actual outputs, live AWS state, checkpoint state, and CloudControl patch payloads.

## Desired Semantics
Numbered, testable invariants from the consumer's perspective. The consumer may be the Pulumi engine, provider runtime, SDK user, CloudControl, or a future maintainer.

## Design
Implementation approach that fits the existing provider. Name the key modules and explain ownership boundaries, schema metadata used, and rejected shortcuts.

## Edge Cases
Import, refresh with and without running the program, write-only properties, create-only properties, read-only paths, optional-computed input/output overlap, nested objects, maps/tags, arrays, legacy checkpoints, and provider/schema upgrades.

## Test Plan
Provider unit tests, helper tests, recorded provider-upgrade fixtures, or examples needed to prove the invariants.
```

Omit sections only when they truly add no information.

## Domain Rules

- Treat **desired inputs** and **actual outputs** as separate concepts. Do not use `__inputs`, `OldInputs`, or output projection interchangeably without explaining which role is needed.
- For refresh issues, trace `Read -> returned Inputs/Properties -> later Diff -> later Update`.
- For update issues, trace `Diff` separately from `CalcPatch`; preview behavior and CloudControl patch generation can diverge.
- For schema-evolution bugs, compare old desired inputs, old output checkpoint, new live outputs, and new schema visibility.
- For optional-computed-style fields, look for input/output overlap (`Inputs` and `Outputs`) and exclude explicit read-only paths before deciding ownership semantics.
- For write-only fields, verify how the old value is preserved because CloudControl cannot return it.
- For create-only fields, call out replacement semantics and whether the provider or engine actually enforces them.
- For nested paths, maps, tags, or arrays, inspect whether the helper is shallow or recursive before relying on a path-based rule.
- Do not hand-edit generated SDKs, generated schema artifacts, `.github/workflows`, `.config/mise.toml`, or submodule content.

## Validation

Map every important invariant to concrete verification:

- Use provider runtime unit tests under `provider/pkg/provider` for RPC-level lifecycle behavior.
- Use helper tests under `provider/pkg/resources` or `provider/pkg/schema` only when behavior is factored there.
- Prefer tests that model old checkpoint/new metadata or old provider/new provider shape when the bug is schema-evolution-related.
- Include cases for input/output overlap, nested object fields, maps/tags, arrays, write-only fields, and create-only fields when the design depends on them.
- Follow `AGENTS.md` local test safety. For `provider/pkg/**` changes on this machine, targeted `mise exec -- go test -count=1 ./pkg/<package> -run '<TestNamePattern>'` commands are appropriate for narrow changes, and `mise exec -- make test_provider_fast` is acceptable for broad fast provider validation. Do not run the full provider suite locally by default because it can run provider 2e2 tests that require live AWS credentials and may spend several minutes retrying EC2 IMDS when credentials are unavailable.

## Stop Conditions

Stop and return to the user when the semantic model is not settled, when two plausible ownership models have different user-visible outcomes, or when the implementation would require changing public SDK/schema behavior beyond the requested issue.
