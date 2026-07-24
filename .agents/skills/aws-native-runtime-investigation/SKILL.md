---
name: aws-native-runtime-investigation
description: Use after triage or repository evidence establishes that an issue involves Pulumi AWS Native runtime behavior across the Pulumi provider protocol, generated CloudFormation metadata, and AWS Cloud Control API. Do not use for initial issue triage, routine generated SDK churn, or isolated schema/codegen edits with no runtime semantics.
---

# AWS Native Runtime Investigation

Use this skill to find the relevant AWS Native lifecycle boundary and keep the
change small. It is a routing aid, not a required planning framework or an
edge-case checklist.

## Start With The Repository Guide

Read [`docs/provider-runtime.md`](../../../docs/provider-runtime.md), following
only the Pulumi engine and CloudControl references relevant to the reported
behavior. Read [`ARCHITECTURE.md`](../../../ARCHITECTURE.md) for repository
boundaries and [`AGENTS.md`](../../../AGENTS.md) for generated-file and command
rules.

Trace the motivating issue, repro, or failing test through the current code. Do
not ask the user to explain facts that can be read from the repository.

## Keep The Change Narrow

- Identify the concrete resource and lifecycle transition that is wrong.
- Name the boundary that owns it: Pulumi engine behavior, AWS Native
  translation/runtime behavior, CloudFormation schema metadata, or
  CloudControl behavior.
- Propose the smallest end-to-end change that fixes the reported path.
- Treat adjacent metadata and lifecycle cases as context to disposition, not
  requirements to implement.
- Preserve existing behavior when stronger semantics cannot be established
  safely.
- Before adding machinery for an uncommon case, check whether it occurs in the
  current schemas, tests, or reported user behavior. Report its incidence,
  fallback, and follow-up option when support is costly.
- Broaden across `Read`, `Diff`, or patch generation only when those boundaries
  must agree for the motivating behavior to be correct.
- Prefer a focused test that proves the user-visible regression. Add lower-level
  tests when they provide distinct proof or useful failure localization rather
  than duplicating the lifecycle test.

Do not create a design document merely because this skill was loaded. Use the
conversation for narrow changes. Create a durable plan only when the user asks
for one or when unresolved cross-boundary decisions need to survive a handoff.

## Stop And Ask

Return to the maintainer with concrete options when:

- Pulumi, AWS Native, and CloudControl evidence point to different owners;
- two plausible semantics have different user-visible results;
- the safe fallback is unclear;
- the implementation is becoming substantially broader than the reported
  problem; or
- a rare case requires disproportionate complexity and cannot be deferred.

Follow `AGENTS.md` for validation and generated artifacts. Do not hand-edit
generated SDKs, generated schema or metadata, CI workflows, toolchain files, or
upstream submodule content.
