package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi-go-provider/resourcex"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/urn"
)

// This is the default timeout for custom resource operations in CloudFormation
const DefaultCustomResourceTimeout = 60 * time.Minute

var lambdaFunctionArnRegex = regexp.MustCompile(`^arn:[^:]+:lambda:[^:]+:[^:]+:function:[a-zA-Z0-9-_]+(:[a-zA-Z0-9-_]+)?$`)

type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

func (realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

type cfnCustomResource struct {
	providerName string
	lambdaClient client.LambdaClient
	s3Client     client.S3Client
	clock        Clock
}

// Check CfnCustomResource implements CustomResource
var _ CustomResource = (*cfnCustomResource)(nil)

// The CloudFormation Custom Resource Emulator provides a bridge between Pulumi's resource model
// and AWS CloudFormation Custom Resources. It enables Pulumi programs to interact with CloudFormation
// Custom Resources while maintaining Pulumi's strong typing and error handling capabilities.
//
// When are Cloudformation Custom Resources used?
// - Managing resources outside of AWS (e.g., GitHub repositories, external APIs)
// - Implementing complex provisioning logic not possible with standard CloudFormation
// - Performing custom validations or transformations during resource lifecycle events
// - Integrating with legacy systems or third-party services
// - Implementing organization-specific infrastructure patterns
//
// The Custom Resource implementation (i.e. the Lambda function) is responsible for:
//   - Processing the request based on the RequestType
//   - Sending a response to the pre-signed URL. The response should include:
//   - success/failure status and a reason
//   - a PhysicalResourceId for the resource
//   - optionally return data that can be referenced by other resources
//
// Example CloudFormation Custom Resource:
//
//	Resources:
//	  MyCustomResource:
//	    Type: Custom::MyResource
//	    Properties:
//	      ServiceToken: arn:aws:lambda:region:account:function:name
//	      Property1: value1
//	      Property2: value2
//
// This emulator implements this lifecycle in Pulumi by:
// - Translating Pulumi resource operations to CloudFormation custom resource requests
// - Managing the asynchronous response collection via S3 with pre-signed URLs
// - Handling timeout and error scenarios according to CloudFormation specifications
//
// Architecture:
//
//	┌──────────────┐         ┌─────────────┐         ┌──────────────┐
//	│   Pulumi     │  CRUD   │   Custom    │  AWS    │    Lambda    │
//	│   Engine     ├────────►│   Resource  ├────────►│   Function   │
//	│              │         │   Emulator  │         │              │
//	└──────────────┘         └─────────────┘         └──────────────┘
//	                              │                         │
//	                              │                         │
//	                              │      ┌──────────┐       │
//	                              └─────►│    S3    │◄──────┘
//	                           Poll for  │  Bucket  │ Response
//	                           Response  └──────────┘
func NewCfnCustomResource(providerName string, s3Client client.S3Client, lambdaClient client.LambdaClient) *cfnCustomResource {
	return &cfnCustomResource{
		providerName: providerName,
		s3Client:     s3Client,
		lambdaClient: lambdaClient,
		clock:        &realClock{},
	}
}

type CfnCustomResourceInputs struct {
	// The name of the S3 bucket to use for storing the response from the custom resource
	BucketName string
	// The prefix to use for the bucket key when storing the response from the custom resource
	BucketKeyPrefix string
	// The service token, such as a Lambda function ARN. The service token must be in the same Region as this resource
	ServiceToken string
	// The custom resource properties to pass as an input to the Lambda function
	CustomResourceProperties map[string]interface{}
	// The CloudFormation type of the custom resource (e.g. "Custom::MyCustomResource")
	ResourceType string
	// A stand-in value for the CloudFormation stack ID required by the custom resource. If not provided, the project ID is used.
	StackID *string
}

type CfnCustomResourceState struct {
	// The physical resource ID of the custom resource
	PhysicalResourceID string `json:"physicalResourceId"`
	// The response data returned by invoking the custom resource lambda
	Data map[string]interface{} `json:"data"`
	// The stand-in value for the CloudFormation stack ID required by the custom resource
	StackID string `json:"stackId"`
	// The service token, such as a Lambda function ARN. The service token must be in the same Region as this resource
	ServiceToken string `json:"serviceToken"`
	// The name of the S3 bucket to use for storing the response from the custom resource
	Bucket string `json:"bucket"`
	// The CloudFormation type of the custom resource (e.g. "Custom::MyCustomResource")
	ResourceType string `json:"resourceType"`
	// Whether the response data contains sensitive information that should be marked as secret and not logged
	NoEcho bool `json:"noEcho"`
}

func (s CfnCustomResourceState) ToPropertyMap() resource.PropertyMap {
	return resource.NewPropertyMap(s)
}

func CfnCustomResourceSpec(description string) pschema.ResourceSpec {
	return pschema.ResourceSpec{
		ObjectTypeSpec: pschema.ObjectTypeSpec{
			Description: description,
			Properties: map[string]pschema.PropertySpec{
				"physicalResourceId": {
					Description: "The name or unique identifier that corresponds to the `PhysicalResourceId` included in the Custom Resource response. If no `PhysicalResourceId` is provided in the response, a random ID will be generated.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"data": {
					Description: "The response data returned by invoking the Custom Resource.",
					TypeSpec: pschema.TypeSpec{
						Type: "object",
						AdditionalProperties: &pschema.TypeSpec{
							Ref: "pulumi.json#/Any",
						},
					},
				},
				"stackId": {
					Description: "A stand-in value for the CloudFormation stack ID.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"serviceToken": {
					Description: "The service token, such as a Lambda function ARN, that is invoked when the resource is created, updated, or deleted.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"bucket": {
					Description: "The name of the S3 bucket to use for storing the response from the Custom Resource.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"resourceType": {
					Description: "The CloudFormation type of the Custom Resource provider. For example, `Custom::MyCustomResource`.",
					TypeSpec:    pschema.TypeSpec{Type: "string"},
				},
				"noEcho": {
					Description: "Whether the response data contains sensitive information that should be marked as secret and not logged.",
					TypeSpec:    pschema.TypeSpec{Type: "boolean"},
				},
			},
			Required: []string{"physicalResourceId", "data", "stackId", "serviceToken", "bucket", "resourceType", "noEcho"},
		},
		InputProperties: map[string]pschema.PropertySpec{
			"bucketName": {
				Description: "The name of the S3 bucket to use for storing the response from the Custom Resource.\n\n" +
					"The IAM principal configured for the provider must have `s3:PutObject`, `s3:HeadObject` and `s3:GetObject` permissions on this bucket.",
				TypeSpec: pschema.TypeSpec{Type: "string"},
			},
			"bucketKeyPrefix": {
				Description: "The prefix to use for the bucket key when storing the response from the Custom Resource provider.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"serviceToken": {
				Description: "The service token to use for the Custom Resource. The service token is invoked when the resource is created, updated, or deleted.\n" +
					"This can be a Lambda Function ARN with optional version or alias identifiers.\n\n" +
					"The IAM principal configured for the provider must have `lambda:InvokeFunction` permissions on this service token.",
				TypeSpec: pschema.TypeSpec{Type: "string"},
			},
			"customResourceProperties": {
				Description: "The properties to pass as an input to the Custom Resource.\nThe properties are passed as a map of key-value pairs whereas all " +
					"primitive values (number, boolean) are converted to strings for CloudFormation interoperability.",
				TypeSpec: pschema.TypeSpec{
					Type: "object",
					AdditionalProperties: &pschema.TypeSpec{
						Ref: "pulumi.json#/Any",
					},
				},
			},
			"resourceType": {
				Description: "The CloudFormation type of the Custom Resource. For example, `Custom::MyCustomResource`.\n" +
					"This is required for CloudFormation interoperability.",
				TypeSpec: pschema.TypeSpec{Type: "string"},
			},
			"stackId": {
				Description: "A stand-in value for the CloudFormation stack ID. This is required for CloudFormation interoperability.\n" +
					"If not provided, the Pulumi Stack ID is used.",
				TypeSpec: pschema.TypeSpec{Type: "string"},
			},
		},
		RequiredInputs: []string{"bucketName", "bucketKeyPrefix", "serviceToken", "customResourceProperties", "resourceType"},
	}
}

type customResourceInvokeData struct {
	event           cfn.Event
	bucket          string
	bucketKeyPrefix string
	timeout         time.Duration
	loggingLabel    string
	serviceToken    string
}

// Check validates the inputs of the resource and applies default values if necessary.
// It returns the inputs, validation failures, and an error if the inputs cannot be unmarshalled.
func (c *cfnCustomResource) Check(ctx context.Context, urn urn.URN, _ autonaming.EngineAutoNamingConfig, inputs resource.PropertyMap, state resource.PropertyMap, defaultTags map[string]string) (resource.PropertyMap, []ValidationFailure, error) {
	var typedInputs CfnCustomResourceInputs
	_, err := resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	var failures []ValidationFailure

	serviceTokenInput, hasServiceToken := inputs[resource.PropertyKey("serviceToken")]
	if (!hasServiceToken || !serviceTokenInput.ContainsUnknowns()) && !lambdaFunctionArnRegex.MatchString(typedInputs.ServiceToken) {
		failures = append(failures, ValidationFailure{
			Path:   "serviceToken",
			Reason: "serviceToken must be a valid Lambda function ARN.",
		})
	}

	_, hasStackID := inputs[resource.PropertyKey("stackId")]
	if !hasStackID && typedInputs.StackID == nil {
		// if the stack ID is not provided, we use the pulumi stack ID as the stack ID
		inputs[resource.PropertyKey("stackId")] = resource.NewStringProperty(urn.Stack().String())
	}

	return inputs, failures, nil
}

// Create creates the Custom Resource by invoking the Lambda function with the CREATE request type.
// Returns the physical resource ID and outputs if the creation is successful, otherwise an error.
//
// Creation Flow:
//
//	┌─────────┐   ┌─────────────┐   ┌───────────────┐   ┌──────────┐
//	│ Create  │   │Generate S3  │   │ Invoke Lambda │   │ Wait for │
//	│ Request ├──►│Presigned URL├──►│ with CREATE   ├──►│ Response │
//	└─────────┘   └─────────────┘   │ RequestType   │   └────┬─────┘
//	                                └───────────────┘        │
//	                                                         ▼
//	                                                 ┌───────────────┐
//	                                                 │Return Physical│
//	                                                 │Resource ID &  │
//	                                                 │Outputs        │
//	                                                 └───────────────┘
//
// The Create operation:
// 1. Generates a presigned S3 URL for response collection
// 2. Constructs a CloudFormation CREATE event with:
//   - Unique RequestID (UUID)
//   - ResponseURL (presigned S3 URL)
//   - ResourceType from input
//   - LogicalResourceId from Pulumi URN
//   - Custom properties from input
//
// 3. Invokes the Lambda function asynchronously
// 4. Waits for response in S3 bucket
// 5. Processes response:
//   - On success: Returns PhysicalResourceId and properties
//   - On failure: Returns error with reason
//   - Handles `NoEcho` for sensitive data
func (c *cfnCustomResource) Create(ctx context.Context, urn urn.URN, inputs resource.PropertyMap, timeout time.Duration) (*string, resource.PropertyMap, error) {
	var typedInputs CfnCustomResourceInputs
	_, err := resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	label := fmt.Sprintf("%s.Create(%s)", c.providerName, urn)

	event := cfn.Event{
		RequestType:        cfn.RequestCreate,
		ResourceType:       typedInputs.ResourceType,
		LogicalResourceID:  urn.Name(),
		StackID:            *typedInputs.StackID,
		ResourceProperties: naming.ToStringifiedMap(typedInputs.CustomResourceProperties),
	}

	response, err := c.invokeCustomResource(ctx, customResourceInvokeData{
		event:           event,
		bucket:          typedInputs.BucketName,
		bucketKeyPrefix: typedInputs.BucketKeyPrefix,
		timeout:         timeout,
		loggingLabel:    label,
		serviceToken:    typedInputs.ServiceToken,
	})

	if err != nil {
		return nil, nil, fmt.Errorf("failed to create custom resource: %w", err)
	}

	outputs := typedInputs.makeOutputs(inputs, response)

	if response.Status == cfn.StatusFailed {
		glog.V(9).Infof("%s custom resource creation failed: %s", label, response.Reason)

		var partialID *string

		// if the custom resource creation partially succeeded, we should return the physical resource ID
		// this could happen if parts of the custom resource creation succeeded but the overall operation failed
		if response.PhysicalResourceID != "" {
			partialID = &response.PhysicalResourceID
			glog.V(9).Infof("%s custom resource creation partially succeeded, physical resource ID: %s", label, *partialID)
		}

		return partialID, outputs, fmt.Errorf("failed to create custom resource: %s", response.Reason)
	}

	glog.V(9).Infof("%s custom resource creation succeeded", label)
	return &response.PhysicalResourceID, outputs, nil
}

// Delete deletes the Custom Resource by invoking the Lambda function with the DELETE request type.
// Returns an error if the delete operation fails, otherwise nil.
//
// Delete Flow:
//
//	┌─────────┐   ┌─────────────┐   ┌───────────────┐   ┌──────────┐
//	│ Delete  │   │Generate S3  │   │ Invoke Lambda │   │ Wait for │
//	│ Request ├──►│Presigned URL├──►│ with DELETE   ├──►│Response  │
//	└─────────┘   └─────────────┘   │ RequestType   │   └────┬─────┘
//	                                └───────────────┘        │
//	                                                         ▼
//	                                                 ┌───────────────┐
//	                                                 │Verify Delete  │
//	                                                 │Success        │
//	                                                 └───────────────┘
//
// The Delete operation:
// 1. Generates a presigned S3 URL for response collection
// 2. Constructs CloudFormation DELETE event with:
//   - Existing PhysicalResourceId
//   - Current ResourceProperties
//   - All standard CloudFormation fields
//
// 3. Invokes Lambda and waits for response
// 4. Handles response:
//   - Success: Completes deletion
//   - Failure: Returns error with reason from Lambda
func (c *cfnCustomResource) Delete(ctx context.Context, urn urn.URN, id string, inputs, state resource.PropertyMap, timeout time.Duration) error {
	var typedInputs CfnCustomResourceInputs
	_, err := resourcex.Unmarshal(&typedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("failed to unmarshal inputs: %w", err)
	}

	var typedState CfnCustomResourceState
	_, err = resourcex.Unmarshal(&typedState, state, resourcex.UnmarshalOptions{})
	if err != nil {
		return fmt.Errorf("failed to unmarshal state: %w", err)
	}

	label := fmt.Sprintf("%s.Delete(%s)", c.providerName, urn)

	event := cfn.Event{
		PhysicalResourceID: typedState.PhysicalResourceID,
		RequestType:        cfn.RequestDelete,
		ResourceType:       typedInputs.ResourceType,
		LogicalResourceID:  urn.Name(),
		StackID:            *typedInputs.StackID,
		ResourceProperties: naming.ToStringifiedMap(typedInputs.CustomResourceProperties),
	}

	response, err := c.invokeCustomResource(ctx, customResourceInvokeData{
		event:           event,
		bucket:          typedInputs.BucketName,
		bucketKeyPrefix: typedInputs.BucketKeyPrefix,
		timeout:         timeout,
		loggingLabel:    label,
		serviceToken:    typedInputs.ServiceToken,
	})

	if err != nil {
		return fmt.Errorf("failed to delete custom resource: %w", err)
	}

	if response.Status == cfn.StatusFailed {
		glog.V(9).Infof("%s custom resource deletion failed: %s", label, response.Reason)
		return fmt.Errorf("failed to delete custom resource: %s", response.Reason)
	}

	glog.V(9).Infof("%s custom resource deletion succeeded", label)
	return nil
}

// Update updates the custom resource with the given inputs and state by invoking the Lambda function with the UPDATE request type.
// If the update is successful and the physical resource ID has changed,
// it deletes the old resource. The function returns the updated resource properties or an error.
//
// Update Flow:
//
//	┌─────────┐   ┌─────────────┐   ┌───────────────┐   ┌──────────┐
//	│ Update  │   │Generate S3  │   │ Invoke Lambda │   │ Wait for │
//	│ Request ├──►│Presigned URL├──►│ with UPDATE   ├──►│ Response │
//	└─────────┘   └─────────────┘   │ RequestType   │   └────┬─────┘
//	                                └───────────────┘        │
//	                                                         ▼
//	                                                 ┌───────────────┐
//	                                                 │Check Physical │
//	                                                 │Resource ID    │
//	                                                 └───────┬───────┘
//	                                                         │
//	                                                         ▼
//	                                                 ┌───────────────┐
//	                                                 │Delete Old     │
//	                                                 │Resource       │
//	                                                 │(if ID changed)│
//	                                                 └───────┬───────┘
//	                                                         │
//	                                                         ▼
//	                                                 ┌───────────────┐
//	                                                 │Return updated │
//	                                                 │Physical       │
//	                                                 │Resource ID &  │
//	                                                 │Outputs        │
//	                                                 └───────────────┘
//
// The Update operation:
// 1. Generates a presigned S3 URL for response collection
// 2. Constructs a CloudFormation UPDATE event with:
//   - Existing PhysicalResourceId
//   - Old and new ResourceProperties
//   - All standard CloudFormation fields
//
// 3. Invokes Lambda and collects response
// 4. If PhysicalResourceId changes:
//   - Initiates cleanup of old resource
//   - Sends DELETE event for old PhysicalResourceId
//
// 5. Returns updated properties and new PhysicalResourceId
func (c *cfnCustomResource) Update(ctx context.Context, urn urn.URN, id string, inputs, oldInputs, state resource.PropertyMap, timeout time.Duration) (resource.PropertyMap, error) {
	var oldTypedInputs CfnCustomResourceInputs
	_, err := resourcex.Unmarshal(&oldTypedInputs, oldInputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal old inputs: %w", err)
	}

	var newTypedInputs CfnCustomResourceInputs
	_, err = resourcex.Unmarshal(&newTypedInputs, inputs, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal new inputs: %w", err)
	}

	var typedState CfnCustomResourceState
	_, err = resourcex.Unmarshal(&typedState, state, resourcex.UnmarshalOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal state: %w", err)
	}

	label := fmt.Sprintf("%s.Update(%s)", c.providerName, urn)

	event := cfn.Event{
		PhysicalResourceID:    typedState.PhysicalResourceID,
		RequestType:           cfn.RequestUpdate,
		ResourceType:          newTypedInputs.ResourceType,
		LogicalResourceID:     urn.Name(),
		StackID:               *newTypedInputs.StackID,
		ResourceProperties:    naming.ToStringifiedMap(newTypedInputs.CustomResourceProperties),
		OldResourceProperties: naming.ToStringifiedMap(oldTypedInputs.CustomResourceProperties),
	}

	startTime := c.clock.Now()
	response, err := c.invokeCustomResource(ctx, customResourceInvokeData{
		event:           event,
		bucket:          newTypedInputs.BucketName,
		bucketKeyPrefix: newTypedInputs.BucketKeyPrefix,
		timeout:         timeout,
		loggingLabel:    label,
		serviceToken:    newTypedInputs.ServiceToken,
	})
	updateDuration := c.clock.Since(startTime)
	glog.V(9).Infof("%s custom resource update took %s", label, updateDuration)

	if err != nil {
		return nil, fmt.Errorf("failed to update custom resource: %w", err)
	}

	if response.Status == cfn.StatusFailed {
		glog.V(9).Infof("%s custom resource update failed: %s", label, response.Reason)
		return nil, fmt.Errorf("failed to update custom resource: %s", response.Reason)
	}

	glog.V(9).Infof("%s custom resource update succeeded", label)

	// if the physical resource ID has changed, we need to delete the old resource
	if response.PhysicalResourceID != typedState.PhysicalResourceID {
		glog.V(9).Infof("%s physical resource ID changed from %q to %q, deleting old resource", label, typedState.PhysicalResourceID, response.PhysicalResourceID)

		deleteEvent := cfn.Event{
			PhysicalResourceID: typedState.PhysicalResourceID,
			RequestType:        cfn.RequestDelete,
			ResourceType:       typedState.ResourceType,
			LogicalResourceID:  urn.Name(),
			StackID:            typedState.StackID,
			ResourceProperties: naming.ToStringifiedMap(oldTypedInputs.CustomResourceProperties),
		}

		deleteTimeout := DefaultCustomResourceTimeout
		// if a custom timeout is set, the delete operation is allowed to take the remaining time
		// otherwise we allow it to take the default timeout
		if timeout != 0 {
			deleteTimeout = timeout - updateDuration
			glog.V(9).Infof("%s custom resource update took %s, clean up of the old resource is allowed to take %s", label, updateDuration, deleteTimeout)
			if deleteTimeout <= 0 {
				return nil, fmt.Errorf("failed to clean up old custom resource: not enough time left to delete the old resource. Consider increasing the timeout")
			}
		}

		deleteResponse, err := c.invokeCustomResource(ctx, customResourceInvokeData{
			event:           deleteEvent,
			bucket:          newTypedInputs.BucketName,
			bucketKeyPrefix: newTypedInputs.BucketKeyPrefix,
			timeout:         deleteTimeout,
			loggingLabel:    label,
			serviceToken:    newTypedInputs.ServiceToken,
		})

		if err != nil {
			return nil, fmt.Errorf("failed to clean up old custom resource: %w", err)
		}
		if deleteResponse.Status == cfn.StatusFailed {
			return nil, fmt.Errorf("failed to clean up old custom resource %q: %s", typedState.PhysicalResourceID, deleteResponse.Reason)
		}
		glog.V(9).Infof("%s old custom resource %q successfully cleaned up", label, typedState.PhysicalResourceID)
	}

	outputs := newTypedInputs.makeOutputs(inputs, response)
	return outputs, nil
}

// Read returns the current inputs and outputs of the custom resource because CFN custom resources do not store state.
// They are just a stateless wrapper around a Lambda function or SNS topic.
func (c *cfnCustomResource) Read(ctx context.Context, urn urn.URN, id string, oldInputs resource.PropertyMap, oldState resource.PropertyMap) (resource.PropertyMap, resource.PropertyMap, bool, error) {
	// Assuming that Read without old state is an import operation
	if len(oldState) == 0 {
		// We can't support import because CustomResources do not store any state
		return nil, nil, false, fmt.Errorf("CustomResourceEmulator import not implemented")
	}

	return oldState, oldInputs, true, nil
}

func (c *cfnCustomResource) invokeCustomResource(ctx context.Context, invokeData customResourceInvokeData) (*cfn.Response, error) {
	timeout := invokeData.timeout
	if timeout == 0 {
		timeout = DefaultCustomResourceTimeout
	}

	requestID := uuid.New().String()
	bucketKey := fmt.Sprintf("%s/%s.json", invokeData.bucketKeyPrefix, requestID)

	responseUrl, err := c.s3Client.PresignPutObject(ctx, invokeData.bucket, bucketKey, timeout)
	if err != nil {
		// presigning is not an API call, it should not fail unless there's issues with the SDK or crypto libs
		return nil, fmt.Errorf("failed to generate response URL: %w", err)
	}
	glog.V(9).Infof("%s created presigned response URL for s3://%s/%s", invokeData.loggingLabel, invokeData.bucket, bucketKey)

	event := invokeData.event
	event.RequestID = requestID
	event.ResponseURL = responseUrl

	eventPayload, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event to JSON: %w", err)
	}

	err = c.lambdaClient.InvokeAsync(ctx, invokeData.serviceToken, eventPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke lambda for custom resource: %w", err)
	}
	glog.V(9).Infof("%s invoked custom resource with request ID %q", invokeData.loggingLabel, requestID)

	glog.V(9).Infof("%s custom resource invocation succeeded, waiting for response to be sent", invokeData.loggingLabel)
	body, err := c.s3Client.WaitForObject(ctx, invokeData.bucket, bucketKey, timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch custom resource response: %w", err)
	}
	defer body.Close()

	var response cfn.Response
	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode custom resource response: %w", err)
	}

	if glog.V(9) {
		logCustomResourceResponse(invokeData.loggingLabel, &response)
	}

	return sanitizeCustomResourceResponse(&event, &response), nil
}

func logCustomResourceResponse(label string, response *cfn.Response) {
	if response == nil {
		return
	}

	if response.NoEcho {
		redactedResponse := *response
		redactedResponse.Data = map[string]interface{}{}
		responseJSON, err := json.Marshal(&redactedResponse)
		if err != nil {
			glog.V(9).Infof("%s failed to marshal custom resource response for logging: %v", label, err)
			return
		}
		glog.V(9).Infof("%s received custom resource response with redacted secret data: %s", label, responseJSON)
	} else {
		responseJSON, err := json.Marshal(response)
		if err != nil {
			glog.V(9).Infof("%s failed to marshal custom resource response for logging: %v", label, err)
			return
		}
		glog.V(9).Infof("%s received custom resource response: %s", label, responseJSON)
	}
}

func (i CfnCustomResourceInputs) makeOutputs(inputs resource.PropertyMap, response *cfn.Response) resource.PropertyMap {
	state := CfnCustomResourceState{
		PhysicalResourceID: response.PhysicalResourceID,
		Data:               response.Data,
		StackID:            *i.StackID,
		ServiceToken:       i.ServiceToken,
		Bucket:             i.BucketName,
		ResourceType:       i.ResourceType,
	}
	checkpoint := CheckpointPropertyMap(inputs, state.ToPropertyMap())

	// if NoEcho is set to true, it means the response contains sensitive data and we should mark it as a secret
	if data, ok := checkpoint["data"]; ok && response.NoEcho {
		checkpoint["data"] = resource.MakeSecret(data)
	}

	return checkpoint
}

func sanitizeCustomResourceResponse(event *cfn.Event, response *cfn.Response) *cfn.Response {
	if response == nil || response.PhysicalResourceID == "" {
		return response
	}

	// ensure PhysicalResourceID is set. For Create requests we fall back to the RequestID,
	// for Update and Delete requests we fall back to the PhysicalResourceID from state
	if response.PhysicalResourceID == "" && (event.RequestType == cfn.RequestDelete || event.RequestType == cfn.RequestUpdate) {
		response.PhysicalResourceID = event.PhysicalResourceID
	} else if response.PhysicalResourceID == "" && event.RequestType == cfn.RequestCreate {
		response.PhysicalResourceID = event.RequestID
	}

	// ensure Data is not nil
	if response.Data == nil {
		response.Data = map[string]interface{}{}
	}

	return response
}

// cfn custom resources have outputs returned in a "data" property
// since it can be any arbitrary data, we mark the entire thing as computed
func (c *cfnCustomResource) PreviewCustomResourceOutputs() resource.PropertyMap {
	return resource.PropertyMap{
		"data": resource.MakeComputed(resource.NewStringProperty("")),
	}
}
