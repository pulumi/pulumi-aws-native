# AWS Native Provider Runtime

This document explains the boundaries that are easy to lose when tracing AWS
Native lifecycle behavior. It documents how AWS Native connects Pulumi's
resource provider protocol, generated CloudFormation metadata, and AWS Cloud
Control API. The upstream projects own their respective contracts.

## Upstream contracts

Use these references for behavior outside this repository:

### Pulumi engine and provider protocol

- [Resource providers](https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/providers/README.html)
- [ResourceProvider gRPC protocol](https://pulumi-developer-docs.readthedocs.io/latest/docs/_generated/proto/provider.html)
- [Resource registration and diffing](https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/deployment-execution/resource-registration.html)
- [Deployment steps, including refresh and import](https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/deployment-execution/steps.html)
- [Pulumi values, unknowns, secrets, and property paths](https://pulumi-developer-docs.readthedocs.io/latest/docs/architecture/types/README.html)

These documents determine when the engine calls provider RPCs, how responses
select deployment steps, and how Pulumi values are represented.

### AWS Cloud Control API and resource schemas

- [How Cloud Control API works](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/how-it-works.html)
- [Cloud Control API resource operations](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html)
- [Identifying resources](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html)
- [Managing asynchronous resource requests](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html)
- [CloudFormation resource type schema](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html)

These documents define CloudControl CRUD-L operations, asynchronous requests,
resource identifiers, and CloudFormation resource schema semantics.

## System boundary

For standard generated resources, AWS Native is a translation layer:

```text
Pulumi program
    |
    | SDK-shaped inputs
    v
Pulumi engine
    |
    | ResourceProvider gRPC
    v
AWS Native provider
    |-- validates and defaults inputs
    |-- uses generated metadata to interpret properties
    |-- translates between SDK and CloudFormation shapes
    |-- checkpoints an input baseline with output state
    |
    | CloudControl CRUD-L requests
    v
AWS Cloud Control API
    |
    v
Underlying AWS service resource
```

The engine owns deployment planning and RPC orchestration. CloudControl owns the
cloud-side resource handlers. AWS Native owns the translation between those
systems.

Two resources use hand-written lifecycle implementations rather than the
standard generated-resource path:

- `aws-native:index:ExtensionResource`
- `aws-native:cloudformation:CustomResourceEmulator`

They implement [`resources.CustomResource`](../provider/pkg/resources/custom.go).
Do not assume standard-resource behavior applies to them.

## Runtime values

Avoid using "state" without identifying the concrete value and its owner.

| Value | Meaning in AWS Native |
| --- | --- |
| Program inputs | Values authored by the Pulumi program before provider validation and defaults. |
| Checked inputs | SDK-shaped desired values returned by `Check`. |
| Checkpointed input baseline | The provider-owned value stored under the secret `__inputs` property. It starts as checked desired inputs after create or update, but refresh may replace managed readable values with observed values. |
| Provider output state | SDK-shaped resource properties returned by AWS Native and recorded by the engine. |
| Live resource model | CloudFormation-shaped JSON returned by CloudControl `GetResource`. It excludes write-only properties. |
| Actual input baseline | An input-shaped projection of observable output state, filtered by ownership and supplemented with checkpointed values that CloudControl cannot return. |
| CloudControl desired state | CloudFormation-shaped JSON sent to `CreateResource`. |
| CloudControl patch | JSON Patch operations sent to `UpdateResource`, calculated from an actual input baseline to new desired inputs. |

The checkpointed input baseline is not always a pure record of user intent.
During refresh, managed readable paths can be updated to their live values so a
later engine diff can detect and repair drift. Unowned optional input/output
paths are not adopted merely because AWS returns them.

## Generated artifacts

AWS Native generates two artifacts with different consumers from the
CloudFormation resource schemas:

```text
CloudFormation resource schemas
    |
    v
provider/pkg/schema and pulumi-gen-aws-native
    |-- Pulumi package schema: SDK shape and engine-facing behavior
    `-- CloudAPIMetadata: CloudControl-facing provider behavior
```

The package schema can affect generated SDK types and engine replacement
behavior. Runtime metadata supports validation, naming conversion, state
projection, semantic comparison, and CloudControl requests. Before changing
generation, determine which artifact the behavior's consumer reads.

The runtime representation is
[`metadata.CloudAPIResource`](../provider/pkg/metadata/metadata.go), generation
lives under [`provider/pkg/schema`](../provider/pkg/schema), and generated
artifacts are embedded by
[`provider/cmd/pulumi-resource-aws-native`](../provider/cmd/pulumi-resource-aws-native).
CloudFormation schemas and generated artifacts must not be edited directly.

## Lifecycle boundaries

The provider RPC entry points are in
[`provider/pkg/provider/provider.go`](../provider/pkg/provider/provider.go).
Follow the calls from the relevant RPC rather than treating every lifecycle as a
required checklist.

### Check and planning

`Check` validates and defaults SDK-shaped inputs without calling CloudControl.

For standard resources, `Diff` compares new checked inputs with an actual input
baseline built from recorded outputs and the checkpointed input baseline. It
returns `DIFF_NONE` when it can establish that no relevant change remains after
semantic suppression. Otherwise it returns `DIFF_UNKNOWN`, allowing the engine
to compare old and new inputs and apply replacement information from the Pulumi
package schema. `DeleteBeforeReplace` matters only if the engine determines that
replacement is required.

CloudFormation `createOnlyProperties` are marked as replacement-sensitive in the
Pulumi package schema. The normal lifecycle relies on the engine to select a
replacement before `Update`; patch generation is not a second general
create-only enforcement layer.

### Create and output checkpointing

A real `Create` converts checked inputs to CloudFormation desired state, waits
for the asynchronous CloudControl operation, reads the resulting resource, and
checkpoints converted outputs with `__inputs`.

After a successful create, the CloudControl client retries an initially missing
`GetResource` result for a targeted eventual-consistency window. If a failed
create still yields an identifier and readable state, AWS Native can return
Pulumi partial-initialization state.

Preview does not call CloudControl. It synthesizes available output shape and
marks unavailable read-only values unknown.

### Read, refresh, and import

`Read` converts the live CloudControl model into SDK-shaped output state. A
not-found result becomes an empty response ID, indicating that the resource no
longer exists.

For an existing managed resource, `Read` also creates the input baseline that
the engine will use after refresh. It projects observable output values into
input shape, preserves ownership and values that cannot be read, suppresses
known AWS-managed representation differences, and returns and checkpoints that
baseline. Live outputs and refreshed inputs therefore have different roles.

For import, no prior `__inputs` value exists. AWS Native projects writable
values from the live model into initial inputs. Write-only values cannot be
recovered and produce a warning.

### Update and patching

`Update` calculates a CloudFormation-named JSON Patch from an actual input
baseline to new checked inputs. Non-create-only write-only values may need to be
resent because CloudControl cannot read them while validating the updated
model. After CloudControl completes, AWS Native reads and checkpoints the new
output state with the new desired inputs.

Provider `Diff` and CloudControl patch calculation are separate correctness
boundaries. A preview can report the right change while update computes the
wrong patch, or the reverse, if only one boundary is fixed.

### Delete and list

`Delete` waits for CloudControl deletion and treats CloudControl `NotFound` as
success.

For generated resources with a list handler, `List` validates and translates a
generated query shape before calling CloudControl `ListResources`. It streams
CloudControl identifiers as Pulumi IDs and uses provider-owned continuation
sessions to mediate Pulumi and CloudControl pagination.

## Key implementation helpers

Most runtime investigations should start at the relevant RPC in
[`provider/pkg/provider/provider.go`](../provider/pkg/provider/provider.go).
The main cross-cutting helpers are:

- [`provider/pkg/resources/path_classifier.go`](../provider/pkg/resources/path_classifier.go)
  for output projection, ownership, and write-only fallback;
- [`provider/pkg/resources/patching.go`](../provider/pkg/resources/patching.go)
  for actual-to-desired patch calculation;
- [`provider/pkg/resources/checkpoint.go`](../provider/pkg/resources/checkpoint.go)
  for `__inputs` storage; and
- [`provider/pkg/client`](../provider/pkg/client) for CloudControl operations and
  asynchronous completion.

Inspect only the helpers reached by the reported lifecycle path.
