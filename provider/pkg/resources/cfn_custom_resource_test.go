package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/urn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestCfnCustomResource_Check(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		inputs           resource.PropertyMap
		expectedInputs   resource.PropertyMap
		expectedError    *string
		expectedFailures []ValidationFailure
	}{
		{
			name: "Valid inputs",
			inputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
		},
		{
			name: "Invalid ServiceToken",
			inputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("invalid-arn"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
			expectedFailures: []ValidationFailure{
				{
					Path:   "serviceToken",
					Reason: "serviceToken must be a valid Lambda function ARN.",
				},
			},
		},
		{
			name: "No ServiceToken",
			inputs: resource.PropertyMap{
				"stackId": resource.NewStringProperty("testProject"),
			},
			expectedFailures: []ValidationFailure{
				{
					Path:   "serviceToken",
					Reason: "serviceToken must be a valid Lambda function ARN.",
				},
			},
		},
		{
			name: "Default StackID",
			inputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
		},
		{
			name: "Stringify CustomResourceProperties",
			inputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"stackId":      resource.NewStringProperty("testProject"),
			},
		},
		{
			name: "Unknown inputs",
			inputs: resource.PropertyMap{
				"serviceToken": resource.MakeComputed(resource.NewStringProperty("")),
				"stackId":      resource.MakeComputed(resource.NewStringProperty("")),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken": resource.MakeComputed(resource.NewStringProperty("")),
				"stackId":      resource.MakeComputed(resource.NewStringProperty("")),
			},
		},
		{
			name: "Preserves Secrets",
			inputs: resource.PropertyMap{
				"serviceToken": resource.MakeSecret(resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function")),
				"stackId":      resource.MakeSecret(resource.NewStringProperty("testProject")),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken": resource.MakeSecret(resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function")),
				"stackId":      resource.MakeSecret(resource.NewStringProperty("testProject")),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &cfnCustomResource{}
			ctx := context.Background()
			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")
			engineAutonaming := autonaming.EngineAutoNamingConfig{
				RandomSeed: []byte{},
			}
			state := resource.PropertyMap{}
			defaultTags := map[string]string{}

			newInputs, failures, err := c.Check(ctx, urn, engineAutonaming, tt.inputs, state, defaultTags)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), *tt.expectedError)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedFailures, failures)

			if tt.expectedInputs != nil {
				assert.Equal(t, tt.expectedInputs, newInputs)
			}
		})
	}
}

func TestCfnCustomResource_Create(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		timeout              time.Duration
		noEcho               bool
		expectedError        string
		customResourceData   map[string]interface{}
		customResourceInputs map[string]interface{}
	}{
		{
			name: "Success",
			customResourceData: map[string]interface{}{
				"prop1": "value1",
				"prop2": true,
				"prop3": []interface{}{"a", "b", "c"},
				"prop4": map[string]interface{}{
					"nestedProp1": "nestedValue1",
					"nestedProp2": 42,
				},
			},
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name:   "SecretResponse",
			noEcho: true,
			customResourceData: map[string]interface{}{
				"prop1": "value1",
				"prop2": true,
				"prop3": []interface{}{"a", "b", "c"},
				"prop4": map[string]interface{}{
					"nestedProp1": "nestedValue1",
					"nestedProp2": 42,
				},
			},
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name:    "CustomTimeout",
			timeout: 10 * time.Minute,
			customResourceData: map[string]interface{}{
				"prop1": "value1",
				"prop2": true,
				"prop3": []interface{}{"a", "b", "c"},
				"prop4": map[string]interface{}{
					"nestedProp1": "nestedValue1",
					"nestedProp2": 42,
				},
			},
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name: "EnsurePhysicalResourceID",
			customResourceData: map[string]interface{}{
				"prop1": "value1",
				"prop2": true,
				"prop3": []interface{}{"a", "b", "c"},
				"prop4": map[string]interface{}{
					"nestedProp1": "nestedValue1",
					"nestedProp2": 42,
				},
			},
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name: "Stringify CustomResourceInputs",
			customResourceData: map[string]interface{}{
				"prop1": "value1",
				"prop2": true,
				"prop3": []interface{}{"a", "b", "c"},
				"prop4": map[string]interface{}{
					"nestedProp1": "nestedValue1",
					"nestedProp2": 42,
				},
			},
			customResourceInputs: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
				"key3": true,
				"key4": map[string]interface{}{
					"nestedKey1": "nestedValue1",
					"nestedKey2": 100,
				},
			},
		},
		{
			name:               "CustomResource without response data",
			customResourceData: nil,
			customResourceInputs: map[string]interface{}{
				"key1": "value1",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLambdaClient := client.NewMockLambdaClient(ctrl)
			mockS3Client := client.NewMockS3Client(ctrl)

			stackID := "stack-id"
			serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
			bucketKeyPrefix := "bucket-key-prefix"
			bucketName := "bucket-name"
			physicalResourceID := "physical-resource-id"
			resourceType := "Custom::MyResource"

			expectedTimeout := DefaultCustomResourceTimeout
			if tt.timeout != 0 {
				expectedTimeout = tt.timeout
			}

			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")
			responseUrl := "https://example.com"
			expectedPayload := cfn.Event{
				RequestType:        cfn.RequestCreate,
				ResponseURL:        responseUrl,
				ResourceType:       resourceType,
				LogicalResourceID:  urn.Name(),
				StackID:            stackID,
				ResourceProperties: naming.ToStringifiedMap(tt.customResourceInputs),
			}

			mockLambdaClient.EXPECT().InvokeAsync(
				gomock.Any(),
				serviceToken,
				gomock.Any(),
			).DoAndReturn(func(ctx context.Context, functionName string, payload []byte) error {
				var event cfn.Event
				err := json.Unmarshal(payload, &event)
				require.NoError(t, err)
				// ignore the RequestID field as it is randomly generated
				expectedPayload.RequestID = event.RequestID
				assert.Equal(t, expectedPayload, event)
				return nil
			})

			mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return(responseUrl, nil)

			response := cfn.Response{
				Status:             cfn.StatusSuccess,
				RequestID:          "request-id",
				LogicalResourceID:  "logical-resource-id",
				StackID:            stackID,
				PhysicalResourceID: physicalResourceID,
				Data:               tt.customResourceData,
				NoEcho:             tt.noEcho,
			}

			responseMessage, err := json.Marshal(response)
			require.NoError(t, err)

			mockS3Client.EXPECT().WaitForObject(
				gomock.Any(),
				bucketName,
				matchesBucketKeyPrefix(bucketKeyPrefix),
				expectedTimeout,
			).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

			c := &cfnCustomResource{
				providerName: "testProvider",
				lambdaClient: mockLambdaClient,
				s3Client:     mockS3Client,
			}
			ctx := context.Background()

			inputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(tt.customResourceInputs)),
			}

			id, outputs, err := c.Create(ctx, urn, inputs, tt.timeout)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, outputs)
			} else {
				require.NoError(t, err)

				expectedID := physicalResourceID
				expectedCustomResourceData := tt.customResourceData
				if expectedCustomResourceData == nil {
					expectedCustomResourceData = map[string]interface{}{}
				}

				expectedState := CfnCustomResourceState{
					PhysicalResourceID: physicalResourceID,
					Data:               expectedCustomResourceData,
					StackID:            stackID,
					ServiceToken:       serviceToken,
					Bucket:             bucketName,
					ResourceType:       resourceType,
				}

				assert.Equal(t, &expectedID, id)
				expectedOutputs := CheckpointPropertyMap(inputs, expectedState.ToPropertyMap())
				if tt.noEcho {
					expectedOutputs["data"] = resource.MakeSecret(expectedOutputs["data"])
				}
				assert.Equal(t, expectedOutputs, outputs)
			}
		})
	}
}

func TestCfnCustomResource_Create_PartialError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		physicalResourceID string
		expectedError      string
		customResourceData map[string]interface{}
	}{
		{
			name:               "With PhysicalResourceID",
			physicalResourceID: "physical-resource-id",
			expectedError:      "some error occurred",
			customResourceData: map[string]interface{}{"key": "value"},
		},
		{
			name:               "Without PhysicalResourceID",
			physicalResourceID: "",
			expectedError:      "some error occurred",
			customResourceData: map[string]interface{}{"key": "value"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLambdaClient := client.NewMockLambdaClient(ctrl)
			mockS3Client := client.NewMockS3Client(ctrl)

			stackID := "stack-id"
			serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
			bucketKeyPrefix := "bucket-key-prefix"
			bucketName := "bucket-name"
			resourceType := "Custom::MyResource"

			mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(nil)
			mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)

			response := cfn.Response{
				Status:             cfn.StatusFailed,
				RequestID:          "request-id",
				LogicalResourceID:  "logical-resource-id",
				StackID:            stackID,
				PhysicalResourceID: tt.physicalResourceID,
				Data:               tt.customResourceData,
				Reason:             tt.expectedError,
			}

			responseMessage, err := json.Marshal(response)
			require.NoError(t, err)

			mockS3Client.EXPECT().WaitForObject(
				gomock.Any(),
				bucketName,
				matchesBucketKeyPrefix(bucketKeyPrefix),
				DefaultCustomResourceTimeout,
			).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

			c := &cfnCustomResource{
				providerName: "testProvider",
				lambdaClient: mockLambdaClient,
				s3Client:     mockS3Client,
			}
			ctx := context.Background()
			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

			inputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{"key": "value"})),
			}

			id, outputs, err := c.Create(ctx, urn, inputs, 0)

			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedError)
			if tt.physicalResourceID != "" {
				assert.Equal(t, &tt.physicalResourceID, id)
			} else {
				assert.Nil(t, id)
			}

			expectedState := CfnCustomResourceState{
				PhysicalResourceID: tt.physicalResourceID,
				Data:               tt.customResourceData,
				StackID:            stackID,
				ServiceToken:       serviceToken,
				Bucket:             bucketName,
				ResourceType:       resourceType,
			}

			expectedOutputs := CheckpointPropertyMap(inputs, expectedState.ToPropertyMap())
			assert.Equal(t, expectedOutputs, outputs)
		})
	}
}

func TestCfnCustomResource_Create_PresignPutObjectFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to presign put object")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"

	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("", expectedError)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	id, outputs, err := c.Create(ctx, urn, inputs, 0)

	require.Error(t, err)
	assert.Nil(t, id)
	assert.Nil(t, outputs)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Create_LambdaInvokeFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to invoke lambda")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(expectedError)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	id, outputs, err := c.Create(ctx, urn, inputs, 0)

	require.Error(t, err)
	assert.Nil(t, id)
	assert.Nil(t, outputs)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Create_S3WaitForObjectFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to fetch custom resource response")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(nil)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)
	mockS3Client.EXPECT().WaitForObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), DefaultCustomResourceTimeout).Return(nil, expectedError)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	id, outputs, err := c.Create(ctx, urn, inputs, 0)

	require.Error(t, err)
	assert.Nil(t, id)
	assert.Nil(t, outputs)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		timeout               time.Duration
		noEcho                bool
		newCustomResourceData map[string]interface{}
	}{
		{
			name:                  "Success",
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "SecretResponse",
			noEcho:                true,
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "CustomTimeout",
			timeout:               10 * time.Minute,
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "CustomResource without response data",
			newCustomResourceData: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLambdaClient := client.NewMockLambdaClient(ctrl)
			mockS3Client := client.NewMockS3Client(ctrl)

			stackID := "stack-id"
			serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
			bucketKeyPrefix := "bucket-key-prefix"
			bucketName := "bucket-name"
			physicalResourceID := "physical-resource-id"
			resourceType := "Custom::MyResource"
			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")
			expectedTimeout := DefaultCustomResourceTimeout
			if tt.timeout != 0 {
				expectedTimeout = tt.timeout
			}

			oldResourceProperties := map[string]interface{}{
				"inputs": "old",
				"key":    42,
			}
			newResourceProperties := map[string]interface{}{
				"inputs": "new",
				"key":    42,
			}

			responseUrl := "https://example.com"
			expectedPayload := cfn.Event{
				RequestType:           cfn.RequestUpdate,
				ResponseURL:           responseUrl,
				ResourceType:          resourceType,
				PhysicalResourceID:    physicalResourceID,
				LogicalResourceID:     urn.Name(),
				StackID:               stackID,
				ResourceProperties:    naming.ToStringifiedMap(newResourceProperties),
				OldResourceProperties: naming.ToStringifiedMap(oldResourceProperties),
			}

			mockLambdaClient.EXPECT().InvokeAsync(
				gomock.Any(),
				serviceToken,
				gomock.Any(),
			).DoAndReturn(func(ctx context.Context, functionName string, payload []byte) error {
				var event cfn.Event
				err := json.Unmarshal(payload, &event)
				require.NoError(t, err)
				// ignore the RequestID field as it is randomly generated
				expectedPayload.RequestID = event.RequestID
				assert.Equal(t, expectedPayload, event)
				return nil
			})

			mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), expectedTimeout).Return("https://example.com", nil)

			response := cfn.Response{
				Status:             cfn.StatusSuccess,
				RequestID:          "request-id",
				LogicalResourceID:  "logical-resource-id",
				StackID:            stackID,
				PhysicalResourceID: physicalResourceID,
				Data:               tt.newCustomResourceData,
				NoEcho:             tt.noEcho,
			}

			responseMessage, err := json.Marshal(response)
			require.NoError(t, err)

			mockS3Client.EXPECT().WaitForObject(
				gomock.Any(),
				bucketName,
				matchesBucketKeyPrefix(bucketKeyPrefix),
				expectedTimeout,
			).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

			c := &cfnCustomResource{
				providerName: "testProvider",
				lambdaClient: mockLambdaClient,
				s3Client:     mockS3Client,
				clock:        &MockClock{},
			}
			ctx := context.Background()

			oldInputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(oldResourceProperties)),
			}

			inputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(newResourceProperties)),
			}

			oldState := CfnCustomResourceState{
				PhysicalResourceID: physicalResourceID,
				Data: map[string]interface{}{
					"old": "value",
				},
				StackID:      stackID,
				ServiceToken: serviceToken,
				Bucket:       bucketName,
				ResourceType: resourceType,
			}

			outputs, err := c.Update(ctx, urn, physicalResourceID, inputs, oldInputs, CheckpointPropertyMap(oldInputs, oldState.ToPropertyMap()), tt.timeout)

			require.NoError(t, err)

			expectedCustomResourceData := tt.newCustomResourceData
			if expectedCustomResourceData == nil {
				expectedCustomResourceData = map[string]interface{}{}
			}

			expectedState := CfnCustomResourceState{
				PhysicalResourceID: physicalResourceID,
				Data:               expectedCustomResourceData,
				StackID:            stackID,
				ServiceToken:       serviceToken,
				Bucket:             bucketName,
				ResourceType:       resourceType,
			}

			expectedOutputs := CheckpointPropertyMap(inputs, expectedState.ToPropertyMap())
			if tt.noEcho {
				expectedOutputs["data"] = resource.MakeSecret(expectedOutputs["data"])
			}

			assert.Equal(t, expectedOutputs, outputs)
		})
	}
}

func TestCfnCustomResource_Update_LambdaInvokeFailure(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to invoke lambda")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	physicalResourceID := "physical-resource-id"
	resourceType := "Custom::MyResource"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(expectedError)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
		clock:        &MockClock{},
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	oldInputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"inputs": "old",
		})),
	}

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"inputs": "new",
		})),
	}

	oldState := CfnCustomResourceState{
		PhysicalResourceID: physicalResourceID,
		Data: map[string]interface{}{
			"old": "value",
		},
		StackID:      stackID,
		ServiceToken: serviceToken,
		Bucket:       bucketName,
		ResourceType: resourceType,
	}

	outputs, err := c.Update(ctx, urn, physicalResourceID, inputs, oldInputs, CheckpointPropertyMap(oldInputs, oldState.ToPropertyMap()), 0)

	require.Error(t, err)
	assert.Nil(t, outputs)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Update_PhysicalResourceIDChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		timeout               time.Duration
		noEcho                bool
		newPhysicalResourceID string
		newCustomResourceData map[string]interface{}
		expectedError         string
		deleteError           string
	}{
		{
			name:                  "Success",
			newPhysicalResourceID: "new-physical-resource-id",
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "SecretResponse",
			noEcho:                true,
			newPhysicalResourceID: "new-physical-resource-id",
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "CustomTimeout",
			timeout:               10 * time.Minute,
			newPhysicalResourceID: "new-physical-resource-id",
			newCustomResourceData: map[string]interface{}{"new": "value"},
		},
		{
			name:                  "DeleteOldResourceFail",
			newPhysicalResourceID: "new-physical-resource-id",
			newCustomResourceData: map[string]interface{}{"new": "value"},
			expectedError:         "failed to clean up old custom resource",
			deleteError:           "failed to delete old resource",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClock := NewMockClock()
			mockLambdaClient := client.NewMockLambdaClient(ctrl)
			mockS3Client := client.NewMockS3Client(ctrl)

			stackID := "stack-id"
			serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
			bucketKeyPrefix := "bucket-key-prefix"
			bucketName := "bucket-name"
			physicalResourceID := "physical-resource-id"
			resourceType := "Custom::MyResource"
			expectedTimeout := DefaultCustomResourceTimeout
			if tt.timeout != 0 {
				expectedTimeout = tt.timeout
			}

			oldResourceProperties := map[string]interface{}{
				"inputs": "old",
			}
			newResourceProperties := map[string]interface{}{
				"inputs": "new",
			}

			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

			responseUrl := "https://example.com"
			expectedUpdatePayload := cfn.Event{
				RequestType:           cfn.RequestUpdate,
				ResponseURL:           responseUrl,
				ResourceType:          resourceType,
				PhysicalResourceID:    physicalResourceID,
				LogicalResourceID:     urn.Name(),
				StackID:               stackID,
				ResourceProperties:    newResourceProperties,
				OldResourceProperties: oldResourceProperties,
			}

			mockLambdaClient.EXPECT().InvokeAsync(
				gomock.Any(),
				serviceToken,
				gomock.Any(),
			).DoAndReturn(func(ctx context.Context, functionName string, payload []byte) error {
				var event cfn.Event
				err := json.Unmarshal(payload, &event)
				require.NoError(t, err)
				// ignore the RequestID field as it is randomly generated
				expectedUpdatePayload.RequestID = event.RequestID
				assert.Equal(t, expectedUpdatePayload, event)
				return nil
			})

			expectedDeletePayload := cfn.Event{
				RequestType:        cfn.RequestDelete,
				ResponseURL:        responseUrl,
				ResourceType:       resourceType,
				PhysicalResourceID: physicalResourceID,
				LogicalResourceID:  urn.Name(),
				StackID:            stackID,
				ResourceProperties: oldResourceProperties,
			}

			mockLambdaClient.EXPECT().InvokeAsync(
				gomock.Any(),
				serviceToken,
				gomock.Any(),
			).DoAndReturn(func(ctx context.Context, functionName string, payload []byte) error {
				var event cfn.Event
				err := json.Unmarshal(payload, &event)
				require.NoError(t, err)
				// ignore the RequestID field as it is randomly generated
				expectedDeletePayload.RequestID = event.RequestID
				assert.Equal(t, expectedDeletePayload, event)
				return nil
			})

			mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), expectedTimeout).Return("https://example.com", nil).Times(2)

			response := cfn.Response{
				Status:             cfn.StatusSuccess,
				RequestID:          "request-id",
				LogicalResourceID:  "logical-resource-id",
				StackID:            stackID,
				PhysicalResourceID: tt.newPhysicalResourceID,
				Data:               tt.newCustomResourceData,
				NoEcho:             tt.noEcho,
			}

			responseMessage, err := json.Marshal(response)
			require.NoError(t, err)

			mockS3Client.EXPECT().WaitForObject(
				gomock.Any(),
				bucketName,
				matchesBucketKeyPrefix(bucketKeyPrefix),
				expectedTimeout,
			).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

			// Whether the deletion of the old resource succeeds or fails
			if tt.deleteError != "" {
				mockS3Client.EXPECT().WaitForObject(
					gomock.Any(),
					bucketName,
					matchesBucketKeyPrefix(bucketKeyPrefix),
					expectedTimeout,
				).Return(nil, errors.New(tt.deleteError))
			} else {
				mockS3Client.EXPECT().WaitForObject(
					gomock.Any(),
					bucketName,
					matchesBucketKeyPrefix(bucketKeyPrefix),
					expectedTimeout,
				).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)
			}

			c := &cfnCustomResource{
				providerName: "testProvider",
				lambdaClient: mockLambdaClient,
				s3Client:     mockS3Client,
				clock:        mockClock,
			}
			ctx := context.Background()

			oldInputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(oldResourceProperties)),
			}

			inputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(newResourceProperties)),
			}

			oldState := CfnCustomResourceState{
				PhysicalResourceID: physicalResourceID,
				Data: map[string]interface{}{
					"old": "value",
				},
				StackID:      stackID,
				ServiceToken: serviceToken,
				Bucket:       bucketName,
				ResourceType: resourceType,
			}

			outputs, err := c.Update(ctx, urn, physicalResourceID, inputs, oldInputs, CheckpointPropertyMap(oldInputs, oldState.ToPropertyMap()), tt.timeout)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, outputs)
			} else {
				require.NoError(t, err)

				expectedState := CfnCustomResourceState{
					PhysicalResourceID: tt.newPhysicalResourceID,
					Data:               tt.newCustomResourceData,
					StackID:            stackID,
					ServiceToken:       serviceToken,
					Bucket:             bucketName,
					ResourceType:       resourceType,
				}

				expectedOutputs := CheckpointPropertyMap(inputs, expectedState.ToPropertyMap())
				if tt.noEcho {
					expectedOutputs["data"] = resource.MakeSecret(expectedOutputs["data"])
				}

				assert.Equal(t, expectedOutputs, outputs)
			}
		})
	}
}

func TestCfnCustomResource_Update_PhysicalResourceIDChangeDeleteTimeout(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClock := NewMockClock()
	expectedTimeout := 10 * time.Minute
	// update consumed the whole timeout, nothing left for delete
	mockClock.customDuration = expectedTimeout + 10*time.Second

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	physicalResourceID := "physical-resource-id"
	resourceType := "Custom::MyResource"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(nil)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), expectedTimeout).Return("https://example.com", nil)

	response := cfn.Response{
		Status:             cfn.StatusSuccess,
		RequestID:          "request-id",
		LogicalResourceID:  "logical-resource-id",
		StackID:            stackID,
		PhysicalResourceID: "new-physical-resource-id",
		Data:               map[string]interface{}{"new": "value"},
		NoEcho:             false,
	}

	responseMessage, err := json.Marshal(response)
	require.NoError(t, err)

	mockS3Client.EXPECT().WaitForObject(
		gomock.Any(),
		bucketName,
		matchesBucketKeyPrefix(bucketKeyPrefix),
		expectedTimeout,
	).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
		clock:        mockClock,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	oldInputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"inputs": "old",
		})),
	}

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
			"inputs": "new",
		})),
	}

	oldState := CfnCustomResourceState{
		PhysicalResourceID: physicalResourceID,
		Data: map[string]interface{}{
			"old": "value",
		},
		StackID:      stackID,
		ServiceToken: serviceToken,
		Bucket:       bucketName,
		ResourceType: resourceType,
	}

	outputs, err := c.Update(ctx, urn, physicalResourceID, inputs, oldInputs, CheckpointPropertyMap(oldInputs, oldState.ToPropertyMap()), expectedTimeout)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to clean up old custom resource: not enough time left to delete the old resource. Consider increasing the timeout")
	assert.Nil(t, outputs)
}

func TestCfnCustomResource_Read(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		oldState       resource.PropertyMap
		oldInputs      resource.PropertyMap
		expectedState  resource.PropertyMap
		expectedInputs resource.PropertyMap
		expectedError  string
	}{
		{
			name: "Success",
			oldState: resource.PropertyMap{
				"physicalResourceId": resource.NewStringProperty("physical-resource-id"),
				"data": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
					"key": "value",
				})),
				"stackId":      resource.NewStringProperty("stack-id"),
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"bucket":       resource.NewStringProperty("bucket-name"),
				"resourceType": resource.NewStringProperty("Custom::MyResource"),
			},
			oldInputs: resource.PropertyMap{
				"serviceToken":    resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"resourceType":    resource.NewStringProperty("Custom::MyResource"),
				"stackID":         resource.NewStringProperty("stack-id"),
				"bucketName":      resource.NewStringProperty("bucket-name"),
				"bucketKeyPrefix": resource.NewStringProperty("bucket-key-prefix"),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
					"key": "value",
				})),
			},
			expectedState: resource.PropertyMap{
				"physicalResourceId": resource.NewStringProperty("physical-resource-id"),
				"data": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
					"key": "value",
				})),
				"stackId":      resource.NewStringProperty("stack-id"),
				"serviceToken": resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"bucket":       resource.NewStringProperty("bucket-name"),
				"resourceType": resource.NewStringProperty("Custom::MyResource"),
			},
			expectedInputs: resource.PropertyMap{
				"serviceToken":    resource.NewStringProperty("arn:aws:lambda:us-west-2:123456789012:function:my-function"),
				"resourceType":    resource.NewStringProperty("Custom::MyResource"),
				"stackID":         resource.NewStringProperty("stack-id"),
				"bucketName":      resource.NewStringProperty("bucket-name"),
				"bucketKeyPrefix": resource.NewStringProperty("bucket-key-prefix"),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{
					"key": "value",
				})),
			},
		},
		{
			name:          "No State",
			expectedError: "CustomResourceEmulator import not implemented",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &cfnCustomResource{}
			ctx := context.Background()
			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

			var oldState resource.PropertyMap
			if tt.oldState != nil {
				oldState = CheckpointPropertyMap(tt.oldInputs, tt.oldState)
			}
			expectedState := CheckpointPropertyMap(tt.expectedInputs, tt.expectedState)

			state, inputs, supported, err := c.Read(ctx, urn, "physical-resource-id", tt.oldInputs, oldState)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				assert.Nil(t, state)
				assert.Nil(t, inputs)
				assert.False(t, supported)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expectedState, state)
				assert.Equal(t, tt.expectedInputs, inputs)
				assert.True(t, supported)
			}
		})
	}
}

func TestCfnCustomResource_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		timeout              time.Duration
		expectedError        string
		customResourceInputs map[string]interface{}
	}{
		{
			name: "Success",
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name:    "CustomTimeout",
			timeout: 10 * time.Minute,
			customResourceInputs: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name: "Stringify CustomResourceInputs",
			customResourceInputs: map[string]interface{}{
				"key1": "value1",
				"key2": 42,
				"key3": true,
				"key4": map[string]interface{}{
					"nestedKey1": "nestedValue1",
					"nestedKey2": 100,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLambdaClient := client.NewMockLambdaClient(ctrl)
			mockS3Client := client.NewMockS3Client(ctrl)

			stackID := "stack-id"
			serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
			bucketKeyPrefix := "bucket-key-prefix"
			bucketName := "bucket-name"
			physicalResourceID := "physical-resource-id"
			resourceType := "Custom::MyResource"

			expectedTimeout := DefaultCustomResourceTimeout
			if tt.timeout != 0 {
				expectedTimeout = tt.timeout
			}

			urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")
			responseUrl := "https://example.com"
			expectedPayload := cfn.Event{
				RequestType:        cfn.RequestDelete,
				ResponseURL:        responseUrl,
				ResourceType:       resourceType,
				LogicalResourceID:  urn.Name(),
				StackID:            stackID,
				PhysicalResourceID: physicalResourceID,
				ResourceProperties: naming.ToStringifiedMap(tt.customResourceInputs),
			}

			mockLambdaClient.EXPECT().InvokeAsync(
				gomock.Any(),
				serviceToken,
				gomock.Any(),
			).DoAndReturn(func(ctx context.Context, functionName string, payload []byte) error {
				var event cfn.Event
				err := json.Unmarshal(payload, &event)
				require.NoError(t, err)
				// ignore the RequestID field as it is randomly generated
				expectedPayload.RequestID = event.RequestID
				assert.Equal(t, expectedPayload, event)
				return nil
			})

			mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return(responseUrl, nil)

			response := cfn.Response{
				Status:             cfn.StatusSuccess,
				RequestID:          "request-id",
				LogicalResourceID:  "logical-resource-id",
				StackID:            stackID,
				PhysicalResourceID: physicalResourceID,
				Data:               map[string]interface{}{},
			}

			responseMessage, err := json.Marshal(response)
			require.NoError(t, err)

			mockS3Client.EXPECT().WaitForObject(
				gomock.Any(),
				bucketName,
				matchesBucketKeyPrefix(bucketKeyPrefix),
				expectedTimeout,
			).Return(io.NopCloser(bytes.NewReader(responseMessage)), nil)

			c := &cfnCustomResource{
				providerName: "testProvider",
				lambdaClient: mockLambdaClient,
				s3Client:     mockS3Client,
			}
			ctx := context.Background()

			inputs := resource.PropertyMap{
				"serviceToken":             resource.NewStringProperty(serviceToken),
				"resourceType":             resource.NewStringProperty(resourceType),
				"stackID":                  resource.NewStringProperty(stackID),
				"bucketName":               resource.NewStringProperty(bucketName),
				"bucketKeyPrefix":          resource.NewStringProperty(bucketKeyPrefix),
				"customResourceProperties": resource.NewObjectProperty(resource.NewPropertyMapFromMap(tt.customResourceInputs)),
			}

			state := resource.PropertyMap{
				"physicalResourceId": resource.NewStringProperty(physicalResourceID),
				"data":               resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{})),
				"stackId":            resource.NewStringProperty(stackID),
				"serviceToken":       resource.NewStringProperty(serviceToken),
				"bucket":             resource.NewStringProperty(bucketName),
				"resourceType":       resource.NewStringProperty(resourceType),
			}

			err = c.Delete(ctx, urn, physicalResourceID, inputs, state, tt.timeout)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCfnCustomResource_Delete_PresignPutObjectFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to presign put object")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"
	physicalResourceID := "physical-resource-id"

	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("", expectedError)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	state := resource.PropertyMap{
		"physicalResourceId": resource.NewStringProperty(physicalResourceID),
		"data":               resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{})),
		"stackId":            resource.NewStringProperty(stackID),
		"serviceToken":       resource.NewStringProperty(serviceToken),
		"bucket":             resource.NewStringProperty(bucketName),
		"resourceType":       resource.NewStringProperty(resourceType),
	}

	err := c.Delete(ctx, urn, physicalResourceID, inputs, state, 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Delete_LambdaInvokeFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to invoke lambda")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"
	physicalResourceID := "physical-resource-id"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(expectedError)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	state := resource.PropertyMap{
		"physicalResourceId": resource.NewStringProperty(physicalResourceID),
		"data":               resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{})),
		"stackId":            resource.NewStringProperty(stackID),
		"serviceToken":       resource.NewStringProperty(serviceToken),
		"bucket":             resource.NewStringProperty(bucketName),
		"resourceType":       resource.NewStringProperty(resourceType),
	}

	err := c.Delete(ctx, urn, physicalResourceID, inputs, state, 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func TestCfnCustomResource_Delete_S3WaitForObjectFail(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLambdaClient := client.NewMockLambdaClient(ctrl)
	mockS3Client := client.NewMockS3Client(ctrl)

	expectedError := fmt.Errorf("failed to fetch custom resource response")
	stackID := "stack-id"
	serviceToken := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	bucketKeyPrefix := "bucket-key-prefix"
	bucketName := "bucket-name"
	resourceType := "Custom::MyResource"
	physicalResourceID := "physical-resource-id"

	mockLambdaClient.EXPECT().InvokeAsync(gomock.Any(), serviceToken, gomock.Any()).Return(nil)
	mockS3Client.EXPECT().PresignPutObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), gomock.Any()).Return("https://example.com", nil)
	mockS3Client.EXPECT().WaitForObject(gomock.Any(), bucketName, matchesBucketKeyPrefix(bucketKeyPrefix), DefaultCustomResourceTimeout).Return(nil, expectedError)

	c := &cfnCustomResource{
		providerName: "testProvider",
		lambdaClient: mockLambdaClient,
		s3Client:     mockS3Client,
	}
	ctx := context.Background()
	urn := urn.URN("urn:pulumi:testProject::test::aws-native:cloudformation:CfnCustomResource::dummy")

	inputs := resource.PropertyMap{
		"serviceToken":    resource.NewStringProperty(serviceToken),
		"resourceType":    resource.NewStringProperty(resourceType),
		"stackID":         resource.NewStringProperty(stackID),
		"bucketName":      resource.NewStringProperty(bucketName),
		"bucketKeyPrefix": resource.NewStringProperty(bucketKeyPrefix),
		"customResourceProperties": resource.NewObjectProperty(resource.PropertyMap{
			"key": resource.NewStringProperty("value"),
		}),
	}

	state := resource.PropertyMap{
		"physicalResourceId": resource.NewStringProperty(physicalResourceID),
		"data":               resource.NewObjectProperty(resource.NewPropertyMapFromMap(map[string]interface{}{})),
		"stackId":            resource.NewStringProperty(stackID),
		"serviceToken":       resource.NewStringProperty(serviceToken),
		"bucket":             resource.NewStringProperty(bucketName),
		"resourceType":       resource.NewStringProperty(resourceType),
	}

	err := c.Delete(ctx, urn, physicalResourceID, inputs, state, 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError.Error())
}

func matchesBucketKeyPrefix(prefix string) gomock.Matcher {
	return gomock.Cond(func(x any) bool { return strings.HasPrefix(x.(string), prefix+"/") })
}

type MockClock struct {
	freezeTime     time.Time
	customDuration time.Duration
}

func NewMockClock() *MockClock {
	return &MockClock{freezeTime: time.Now()}
}

func (c *MockClock) Now() time.Time {
	return c.freezeTime
}

func (c *MockClock) Since(t time.Time) time.Duration {
	return c.customDuration
}
