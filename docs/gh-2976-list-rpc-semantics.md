# gh-2976 CloudControl List Semantics

## Summary

Issue [#2976](https://github.com/pulumi/pulumi-aws-native/issues/2976) tracks implementing the Pulumi provider `List` RPC for AWS Native resources.

AWS Native should implement `List` only through AWS Cloud Control API `ListResources`. This provider is CloudControl-backed; it should not grow service-native AWS SDK enumeration paths.

The implementation should be schema-driven:

- CloudFormation resource schemas that declare `handlers.list` are listable through CloudControl.
- `handlers.list.handlerSchema`, when present, describes the query input shape that should become Pulumi `ResourceSpec.listInputs`.
- Resources with `handlers.list` but no `handlerSchema` are still listable. They should get an empty `listInputs` object, and runtime should reject non-empty queries because there is no handler schema to validate or map to CloudFormation casing.
- The `List` result `id` must be the CloudControl `ResourceDescription.Identifier`, because the Pulumi `List` protocol expects an importable ID.

The initial implementation should support every generated AWS Native resource whose CloudFormation schema declares a `list` handler and whose resource type is otherwise included in the generated provider.

## Current Flow

AWS Native already carries enough generated metadata to describe resource inputs, outputs, identifiers, and a partial list handler schema.

- `provider/pkg/provider/provider.go` defines `cfnProvider`, embeds `pulumirpc.UnimplementedResourceProviderServer`, and implements the provider RPC surface.
- `provider/pkg/provider/provider.go` implements standard-resource `Read` by passing the Pulumi resource ID to CloudControl `GetResource`.
- `provider/pkg/client/api.go` wraps low-level CloudControl create, update, delete, get, and operation-status calls.
- `provider/pkg/client/client.go` wraps those low-level calls into the higher-level provider client.
- `provider/pkg/schema/gen.go` builds Pulumi `ResourceSpec` values and generated AWS Native metadata from CloudFormation schemas.
- `provider/pkg/schema/gen.go` currently extracts `handlers.list.handlerSchema` into `metadata.CloudAPIResource.ListHandlerSchema`.

The missing pieces are:

- A preserved metadata flag for `handlers.list` independent of `handlerSchema`.
- Generated Pulumi `listInputs` for listable resources.
- CloudControl `ListResources` wrappers.
- A provider `List` RPC implementation.

Local schema inspection shows why `ListHandlerSchema` is not enough as the listability signal. In the current inputs, generated provider metadata has 1372 resources, 1293 of those source schemas declare `handlers.list`, and only 307 declare `handlers.list.handlerSchema`. Using `ListHandlerSchema != nil` as the gate would skip 986 listable generated resources.

Current Pulumi engine support should be treated as protocol and schema plumbing, not as full list query validation. In the current `pulumi/pulumi` checkout, `listInputs` binds and round-trips through package schema, and required properties are preserved on the bound list input object. The `ListRequest.query` field remains an unstructured protobuf property map because the typed query shape is token-specific and comes from package schema, not from the provider protocol message. The AWS Native provider must therefore validate list queries against its generated list metadata before calling CloudControl.

## Desired Semantics

1. AWS Native `List` is implemented only with CloudControl `ListResources`.
2. A generated AWS Native resource is listable when its source CloudFormation schema declares `handlers.list` and the resource is otherwise included in generated provider metadata.
3. `handlers.list.handlerSchema` describes the optional and required query inputs, not listability itself.
4. A listable resource with no `handlerSchema` still gets non-nil empty `listInputs`.
5. `listInputs` includes both required and optional list handler fields when the handler schema is present.
6. `listInputs.required` mirrors `handlerSchema.required`, after converting property names to Pulumi SDK casing.
7. `ListRequest.Query` is converted from Pulumi SDK casing to CloudFormation casing and serialized as CloudControl `ListResourcesInput.ResourceModel`.
8. If the converted query is empty, `ResourceModel` is omitted.
9. `ListRequest.ContinuationToken` is a provider-owned opaque token for an in-memory list session. The session stores the CloudControl `NextToken`, the original logical request identity, and the number of results already returned.
10. `ListRequest.PageSize`, when non-zero and within CloudControl limits, is the requested provider RPC chunk size and maps to CloudControl `MaxResults` for the current streaming response.
11. `ListRequest.Limit`, when non-zero, is the total result cap across the logical list operation, including continuation requests.
12. Each CloudControl `ResourceDescription.Identifier` is streamed as `ListResponse.Result.Id`.
13. `ListResponse.Result.Id` must be importable through the existing provider `Read` path.
14. `ListResponse.Result.Name` is optional. The first implementation should leave it empty unless a clear protocol or product requirement identifies a stable name rule.
15. Custom AWS Native resources that are not CloudControl standard resources are not listable unless they later get an explicit list implementation.
16. The provider validates query requiredness and basic shape. It should not rely on the engine to reject missing required list inputs before the RPC arrives.
17. If the query contains unknown/computed values, the provider returns `ListResponse.Computed` instead of calling CloudControl.
18. Each AWS Native provider RPC returns one bounded page of importable IDs. The provider tracks returned result count in the continuation session so it can stop returning continuations when `Limit` is reached.
19. The provider does not implement arbitrary search by name, tags, or resource properties.
20. A non-empty query for a resource with empty `listInputs` is invalid.

## Design

### Codegen Metadata

Add a metadata field to distinguish listability from list input shape, for example:

```go
HasListHandler bool `json:"hasListHandler,omitempty"`
```

Set this field when the CloudFormation schema contains `handlers.list`, even if `handlerSchema` is absent.

Keep `ListHandlerSchema` for the query input shape. It should remain nil when the source schema has no `handlerSchema` or when the handler schema is empty.

### Pulumi Schema

Populate `pschema.ResourceSpec.ListInputs` for every generated resource with `HasListHandler`.

When `ListHandlerSchema` is present:

- Convert handler schema property names from CloudFormation casing to SDK casing.
- Preserve descriptions and primitive types where available.
- Include required and optional properties.
- Set `ListInputs.Required` from handler schema `Required`, converted to SDK casing.

Some CloudFormation schemas declare required list handler fields through `$ref` chains rather than inline property schemas.
The generator should resolve those references before emitting `listInputs`.
If a required list handler field is missing from `handlerSchema.properties`, the generator may resolve that field from the resource's own property schema.
If a required list handler field cannot be resolved to schema metadata from either source, generation should fail instead of emitting an untyped required `Any` input.

`ListHandlerSchema` should remain stored in CloudFormation casing as the source of truth. Runtime should build a per-resource `sdkName -> cfnName` map from `ListHandlerSchema.Properties`, validate incoming `ListRequest.Query` keys against the SDK names, and serialize accepted values under the mapped CloudFormation names.

If two CloudFormation list handler fields collapse to the same SDK name, the generator should fail. The runtime should not guess which CloudFormation field a colliding SDK name means.

When `ListHandlerSchema` is nil but `HasListHandler` is true:

- Set `ListInputs` to an empty object.

Do not emit `ListInputs` for resources without `handlers.list`.

## Use Cases And Non-Goals

### Import ID Discovery

The primary v1 use case is discovering importable IDs.

For example, a user who wants to import an existing S3 bucket could list `aws-native:s3:Bucket`, inspect returned IDs, and pass the selected ID to `pulumi import` or the resource import option.

This works best when the CloudControl identifier is already human-recognizable, such as a bucket name, log group name, or ARN. It is weaker for opaque IDs where the user cannot identify the target from the ID alone.

A single AWS Native provider RPC should return one bounded page and a continuation token when more results are available and the logical `limit` has not been reached. CLI, SDK, or higher-level tooling should treat `page_size` as an internal provider RPC chunking control and may consume continuation tokens internally. For example, `pulumi resources list <type> --limit 50 --page-size 10` should return up to 50 results to the user while asking the provider for chunks of up to 10 at a time.

### Scoped Child Discovery

The strongest use for `listInputs` is scoped child enumeration.

Many CloudFormation list handlers require parent or scope fields, such as an API Gateway REST API ID when listing stages. Those required list inputs are not general search filters; they are the minimum scope CloudControl needs to perform the list operation.

For these resources, v1 should expose required and optional handler schema fields through `listInputs` and serialize supplied query values as CloudControl `ResourceModel`.

### Bulk Inventory And Bulk Import Tooling

Bulk inventory or bulk import tools can use `List` to enumerate IDs page by page and then decide whether to call `Read`, generate import files, or perform some other follow-up operation.

This is supported, but the cost belongs to the caller-facing tool. The AWS Native provider `List` RPC should not call `GetResource` for every listed item and should not read all pages by default.

### Pulumi Program Inputs

Programmatic use inside Pulumi code is intentionally limited by the protocol response shape.

The response contains an importable `id` and optional `name`, not full resource properties. This can be useful when a program only needs the IDs returned by a scoped list call, but it is not a replacement for typed data sources or get functions when the program needs to select a resource by tag, name, or other property.

Programs that depend on broad list results should also account for normal inventory drift: resources may be added, removed, or paginated differently between deployments.

### Non-Goals

V1 should not implement:

- arbitrary search by substring, name, tag, or resource property;
- provider-side list-plus-read filtering;
- automatic full-account inventory crawling from a single `List` call;
- service-native AWS SDK list operations;
- full resource property projection in `ListResponse`.

### CloudControl Client

Extend the low-level CloudControl API wrapper with `ListResources`.

The wrapper should expose CloudControl concepts without provider protocol details:

- `typeName`
- optional serialized `resourceModel`
- optional `nextToken`
- optional `maxResults`
- returned resource descriptions
- returned next token

Extend the higher-level CloudControl client with a small list method if it simplifies provider tests, but avoid long-running operation machinery. CloudControl `ListResources` is not an asynchronous progress-event operation.

### Provider List RPC

Implement `cfnProvider.List(req, stream)` for standard resources.

The method should:

1. Look up `req.Token` in generated metadata.
2. Return an unimplemented or invalid-argument error if the token is unknown, custom-only, or not listable.
3. If `req.Query` contains unknown/computed values, send a single `ListResponse.Computed` and do not call CloudControl.
4. Reject a non-empty `req.Query` for resources whose list input object is empty.
5. Validate `req.Query` against the generated list input metadata:
   - all required list inputs are present;
   - supplied fields are declared by `ListHandlerSchema`;
   - supplied primitive values match the generated list input type well enough to serialize to CloudControl.
6. Convert `req.Query` from Pulumi SDK names to CloudFormation names using the generated list-handler `sdkName -> cfnName` map.
7. Serialize the converted query to JSON as `ResourceModel` only when non-empty.
8. Resolve or create the provider-owned continuation session, including the stored CloudControl `NextToken` for continuation requests.
9. Compute the remaining logical result budget from `Limit` and the session's returned count.
10. Call CloudControl `ListResources` with `TypeName`, `ResourceModel`, the session's CloudControl `NextToken`, effective `MaxResults`, and configured `RoleArn`.
11. Stream one `ListResponse.Result` per returned `ResourceDescription` with `Id` set to `Identifier`.
12. Update the session's returned count and stored CloudControl `NextToken`.
13. Send a provider-owned `ListResponse.Continuation` only if CloudControl returned `NextToken` and the logical `Limit` has not been reached.

The first implementation should not call `GetResource` for every listed item. That would turn list into list-plus-read and make broad enumeration unexpectedly expensive.

### Continuation Handling

CloudControl already returns `NextToken`, but AWS Native v1 should not expose that token directly as the Pulumi continuation token. The Pulumi continuation token should be provider-owned and should reference a short-lived in-memory list session.

One AWS Native provider `List` call should correspond to one bounded CloudControl `ListResources` call. A continuation request looks up the provider-owned list session, retrieves the stored CloudControl `NextToken`, and performs the next bounded CloudControl call.

The list session only needs to store logical paging state:

- current CloudControl `NextToken`;
- number of results returned so far;
- original `Limit`;
- original resource token;
- a query fingerprint or canonical query representation for continuation validation;
- last-access time for expiry and cleanup.

For each CloudControl call, the effective CloudControl `MaxResults` should be:

- `PageSize`, if `PageSize` is positive;
- otherwise `Limit`, if `Limit` is positive;
- otherwise omitted, so CloudControl can choose its default page size.

The effective value is capped at the CloudControl `MaxResults` limit of 100. If `Limit` is positive, also cap the current call by `Limit - returnedSoFar`. If the remaining budget is zero, delete the session and return no continuation without calling CloudControl.

Continuation requests should normally use the same `Token`, `Query`, and `Limit` as the initial request because the continuation token belongs to that logical query. The provider should reject a continuation request if the token is unknown or expired, if `Token` or `Limit` changed, or if the query does not match the original query.

The provider should not implement partial-page buffering in v1. It should request at most the number of results it is prepared to stream. If CloudControl returns more descriptions than the effective requested maximum, the provider should return an error rather than stream extra items, silently drop items, or invent item-offset state inside an already-fetched CloudControl page.

### Limit And Page Size Semantics

`PageSize` controls provider RPC chunking. `Limit` controls the total number of results the provider should return across the logical list operation.

For a command like:

```sh
pulumi resources list <type> --limit 50 --page-size 10
```

the result should contain up to 50 resources. The CLI or other caller asks the provider for pages of up to 10 resources and consumes provider-owned continuation tokens internally until the provider stops returning a continuation.

AWS Native should use a small in-memory continuation session:

- The first provider RPC calls CloudControl `ListResources` with `MaxResults=10`.
- If CloudControl returns a `NextToken`, the provider stores it in a list session together with `returnedSoFar=10` and returns a provider-owned opaque continuation token.
- The second provider RPC looks up the session, computes `remaining=40`, calls CloudControl with the stored `NextToken` and `MaxResults=10`, updates `returnedSoFar=20`, and returns the same or a replacement provider token if CloudControl has more results.
- The fifth provider RPC reaches `returnedSoFar=50`; the provider deletes the session and returns no continuation even if CloudControl returned another `NextToken`.

This matches the bridge implementation's user-facing contract without adopting the bridge's full buffering design. The bridge needs a background producer and buffered pages because Terraform Plugin Framework `ListResource` is a stream. AWS Native already has CloudControl page tokens, so it only needs provider-owned session state for the current CloudControl token and logical result count.

## Edge Cases

### Empty Query

For resources with empty `listInputs`, `Query` should normally be nil or empty. The provider should omit `ResourceModel`.

If `Query` is non-empty for a resource with empty `listInputs`, the provider should return an invalid-argument error. Empty `listInputs` means the list handler accepts no modeled query shape, not that the provider supports arbitrary search by resource property.

### Scoped List Queries

For resources whose list handler requires parent scope, such as API Gateway child resources, `listInputs.required` should include those parent fields. Runtime should pass them through as `ResourceModel` in CloudFormation casing.

### Unknown or Secret Query Values

Unknown values cannot produce a meaningful CloudControl request. The provider protocol has an explicit `ListResponse.Computed` response for this case. If any query input is unknown/computed, the provider should stream a single computed response and return without calling CloudControl.

This `Computed` response should also take precedence over continuation-token validation. A continuation request with unknown query values cannot prove query identity, but it also cannot perform a CloudControl request or advance session state.

Secret wrappers should be unwrapped for the CloudControl request while preserving normal Pulumi secret handling in logs.

### Missing Identifier

If CloudControl returns a resource description with nil or empty `Identifier`, the provider should skip it with a warning or return an error. The first implementation should prefer an error, because returning non-importable list entries violates the protocol goal.

### Name

`ListResponse.Result.Name` remains unsettled. Possible future rules include using an autonaming property, a single primary identifier component, or a common `name`-like field. None is universally correct for AWS Native today.

The v1 implementation should leave `Name` empty and document that `Id` is the import key.

### Broad Lists In Large Accounts

For unscoped top-level resources, CloudControl may return many IDs. The provider should rely on CloudControl pagination and return a continuation token instead of attempting to collect all results. Caller-facing tools should make unbounded full enumeration explicit, while the provider still enforces any non-zero `Limit` across continuation requests.

## Test Plan

### Schema Tests

Add codegen tests covering:

- `handlers.list` with no `handlerSchema` emits empty `listInputs`.
- `handlers.list.handlerSchema` emits required and optional `listInputs`.
- handler schema properties are converted from CloudFormation casing to SDK casing.
- required handler schema properties that cannot be resolved fail generation.
- CloudFormation list handler field names that collide after SDK casing conversion fail generation.
- resources without `handlers.list` do not emit `listInputs`.
- generated metadata preserves `HasListHandler` independently from `ListHandlerSchema`.

### Client Tests

Add CloudControl client tests covering:

- `ListResources` serializes `TypeName`, `ResourceModel`, `NextToken`, `MaxResults`, and `RoleArn`.
- returned identifiers and next tokens are preserved.
- CloudControl errors are returned unchanged or wrapped with useful operation context.

### Provider Tests

Add provider RPC tests covering:

- unknown or non-listable token is rejected.
- non-empty query against empty `listInputs` is rejected.
- missing required query inputs are rejected by the provider.
- unknown/computed query inputs return `ListResponse.Computed` without calling CloudControl.
- empty-query list call omits `ResourceModel`.
- scoped-query list call sends CloudFormation-cased `ResourceModel`.
- streamed results use CloudControl `Identifier` as `Id`.
- provider-owned continuation token is streamed when CloudControl returns `NextToken` and the logical `Limit` has not been reached.
- continuation request looks up the stored CloudControl `NextToken` and passes it to CloudControl.
- unknown or expired continuation token is rejected.
- continuation request with changed `Token`, `Query`, or `Limit` is rejected.
- `PageSize` maps to effective `MaxResults` for the current CloudControl call.
- `Limit` maps to effective `MaxResults` for the current CloudControl call when `PageSize` is unset or larger than `Limit`.
- combined `PageSize` and `Limit` use the smaller value for `MaxResults`.
- `limit=50` and `pageSize=10` uses `MaxResults=10` for the current provider RPC.
- `limit=50`, `pageSize=20`, and 40 prior returned results uses `MaxResults=10` for the continuation RPC.
- continuation session is deleted and no continuation is returned when `Limit` is reached.
- effective `MaxResults` is capped at the CloudControl limit of 100.
- nil or empty identifier is treated as an error.

### Integration Validation

Prefer a focused live or recorded validation after unit tests prove the generic path:

- one no-query resource whose schema has `handlers.list` but no `handlerSchema`;
- one scoped-query resource whose schema has `handlerSchema.required`.

Do not run the broad provider integration suite locally by default.

## Open Questions

1. Whether `Name` should remain empty permanently or later use a conservative resource-specific rule.
2. Whether list inputs should support only primitive handler schema properties initially, or recursively support object shapes if CloudFormation list handler schemas ever include them.
3. Whether AWS Native should eventually expose richer names or summaries for list results beyond the importable ID.
