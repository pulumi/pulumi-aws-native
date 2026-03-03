## Summary
<!-- What changed and why? -->

## Scope
<!-- Files/areas touched and what was intentionally not changed. -->

## Validation
<!-- Paste exact commands run -->
- [ ] `make lint`
- [ ] `make test_provider_fast`
- [ ] `make test_provider` (if provider behavior changed)
- [ ] `make test` (only when integration validation is required)

## Generation / Drift Checks
- [ ] No hand-edits to generated files (`sdk/**`, generated schema artifacts)
- [ ] Ran `make generate_schema` (if schema/codegen inputs changed)
- [ ] Ran `make local_generate` (if SDK/API surface changed)
- [ ] Ran `make ci-mgmt` (if `.ci-mgmt.yaml` changed)

## Risk
<!-- Blast radius and likely failure modes -->

## Rollback
<!-- Exact rollback approach -->
