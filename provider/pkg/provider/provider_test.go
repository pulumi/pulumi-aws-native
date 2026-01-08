package provider

import (
	"context"
	"testing"
	"time"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestConfigure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	t.Run("No AutoNaming Config", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:region":                    "us-west-2",
			},
		}

		_, err := provider.Configure(ctx, req)
		assert.NoError(t, err)

		assert.Nil(t, provider.autoNamingConfig)
	})

	t.Run("AutoNaming config", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:autoNaming":                "{\"autoTrim\": true, \"randomSuffixMinLength\": 5}",
				"aws-native:config:region":                    "us-west-2",
			},
		}

		_, err := provider.Configure(ctx, req)
		assert.NoError(t, err)

		assert.NotNil(t, provider.autoNamingConfig)
		assert.Equal(t, autonaming.ProviderAutoNamingConfig{
			AutoTrim:              true,
			RandomSuffixMinLength: 5,
		}, *provider.autoNamingConfig)
	})

	t.Run("AutoNaming empty", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:autoNaming":                "{}",
				"aws-native:config:region":                    "us-west-2",
			},
		}

		_, err := provider.Configure(ctx, req)
		assert.NoError(t, err)

		assert.NotNil(t, provider.autoNamingConfig)
		assert.Equal(t, autonaming.ProviderAutoNamingConfig{}, *provider.autoNamingConfig)
	})

	t.Run("AutoNaming config invalid", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:autoNaming":                "autoTrim: true",
				"aws-native:config:region":                    "us-west-2",
			},
		}

		_, err := provider.Configure(ctx, req)
		assert.ErrorContains(t, err, "failed to unmarshal 'autoNaming' config")
	})

	t.Run("S3ForcePathStyle config", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:s3ForcePathStyle":          "true",
				"aws-native:config:region":                    "us-west-2",
			},
		}

		_, err := provider.Configure(ctx, req)
		assert.NoError(t, err)
	})
}

func TestCheckConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	t.Run("s3ForcePathStyle is accepted", func(t *testing.T) {
		news, err := plugin.MarshalProperties(resource.PropertyMap{
			"s3ForcePathStyle": resource.NewBoolProperty(true),
			"region":           resource.NewStringProperty("us-west-2"),
		}, plugin.MarshalOptions{})
		require.NoError(t, err)

		req := &pulumirpc.CheckRequest{
			News: news,
		}

		resp, err := provider.CheckConfig(ctx, req)
		assert.NoError(t, err)
		assert.Nil(t, resp.Failures)
	})
}

func TestEngineAutonaming(t *testing.T) {
	t.Parallel()

	t.Run("WithoutAutonaming", func(t *testing.T) {
		req := &pulumirpc.CheckRequest{
			RandomSeed: []byte("test-seed"),
		}

		config := engineAutonaming(req)
		assert.Equal(t, []byte("test-seed"), config.RandomSeed)
		assert.Nil(t, config.AutonamingMode)
		assert.Empty(t, config.ProposedName)
	})

	tests := []struct {
		name         string
		mode         pulumirpc.CheckRequest_AutonamingOptions_Mode
		proposedName string
		want         autonaming.EngineAutonamingMode
	}{
		{
			name: "Disable",
			mode: pulumirpc.CheckRequest_AutonamingOptions_DISABLE,
			want: autonaming.EngineAutonamingModeDisable,
		},
		{
			name:         "Enforce",
			mode:         pulumirpc.CheckRequest_AutonamingOptions_ENFORCE,
			proposedName: "test-resource",
			want:         autonaming.EngineAutonamingModeEnforce,
		},
		{
			name:         "Propose",
			mode:         pulumirpc.CheckRequest_AutonamingOptions_PROPOSE,
			proposedName: "test-resource",
			want:         autonaming.EngineAutonamingModePropose,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pulumirpc.CheckRequest{
				RandomSeed: []byte("test-seed"),
				Autonaming: &pulumirpc.CheckRequest_AutonamingOptions{
					Mode:         tt.mode,
					ProposedName: tt.proposedName,
				},
			}

			config := engineAutonaming(req)
			assert.Equal(t, []byte("test-seed"), config.RandomSeed)
			require.NotNil(t, config.AutonamingMode)
			assert.Equal(t, tt.want, *config.AutonamingMode)
			assert.Equal(t, tt.proposedName, config.ProposedName)
		})
	}
}

func TestCreatePreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name: "test-provider",
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws-native:s3:Bucket": metadata.CloudAPIResource{
				CfType: "AWS::S3::Bucket",
				Outputs: map[string]schema.PropertySpec{
					"arn":        {TypeSpec: schema.TypeSpec{Type: "string"}},
					"bucketName": {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
				ReadOnly: []string{"arn", "domainName", "dualStackDomainName", "regionalDomainName", "websiteUrl"},
			},
		}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")

	t.Run("Outputs are computed", func(t *testing.T) {
		req := &pulumirpc.CreateRequest{
			Urn:        string(urn),
			Preview:    true,
			Properties: mustMarshalProperties(t, resource.PropertyMap{"bucketName": resource.NewStringProperty("name")}),
			Timeout:    float64((5 * time.Minute).Seconds()),
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws-native:s3:Bucket", "name"))

		resp, err := provider.Create(ctx, req)
		assert.NoError(t, err)
		assert.Empty(t, resp.Id)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		require.True(t, props.HasValue("arn"), "Expected 'arn' property in response")
		require.True(t, props.HasValue("bucketName"), "Expected 'bucketName' property in response")
		assert.Equal(t, "name", props["bucketName"].StringValue())
		assert.True(t, props["arn"].IsComputed())
	})

	t.Run("Outputs are computed for custom resource", func(t *testing.T) {
		req := &pulumirpc.CreateRequest{
			Urn:        string(urn),
			Preview:    true,
			Properties: mustMarshalProperties(t, resource.PropertyMap{"bucketName": resource.NewStringProperty("name")}),
			Timeout:    float64((5 * time.Minute).Seconds()),
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "custom:resource", "name"))

		mockCustomResource.EXPECT().PreviewCustomResourceOutputs().Return(
			resource.PropertyMap{"outputs": resource.MakeComputed(resource.NewStringProperty(""))},
		)

		resp, err := provider.Create(ctx, req)
		assert.NoError(t, err)
		assert.Empty(t, resp.Id)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		require.True(t, props.HasValue("outputs"), "Expected 'outputs' property in response")
		assert.True(t, props["outputs"].IsComputed())
	})
}

func TestUpdatePreview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name: "test-provider",
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws-native:s3:Bucket": metadata.CloudAPIResource{
				CfType: "AWS::S3::Bucket",
				Outputs: map[string]schema.PropertySpec{
					"arn":        {TypeSpec: schema.TypeSpec{Type: "string"}},
					"bucketName": {TypeSpec: schema.TypeSpec{Type: "string"}},
				},
				ReadOnly: []string{"arn", "domainName", "dualStackDomainName", "regionalDomainName", "websiteUrl"},
			},
		}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")

	t.Run("Stable outputs appear in preview", func(t *testing.T) {
		req := &pulumirpc.UpdateRequest{
			Urn:     string(urn),
			Preview: true,
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"bucketName": resource.NewStringProperty("name"),
				"arn":        resource.NewStringProperty("bucketArn"),
			}),
			News:    mustMarshalProperties(t, resource.PropertyMap{"bucketName": resource.NewStringProperty("name")}),
			Timeout: float64((5 * time.Minute).Seconds()),
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws-native:s3:Bucket", "name"))

		resp, err := provider.Update(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		require.True(t, props.HasValue("arn"), "Expected 'arn' property in response")
		require.True(t, props.HasValue("bucketName"), "Expected 'bucketName' property in response")
		assert.Equal(t, "name", props["bucketName"].StringValue())
		assert.Equal(t, "bucketArn", props["arn"].StringValue())
	})

	t.Run("custom resource", func(t *testing.T) {
		req := &pulumirpc.UpdateRequest{
			Urn:     string(urn),
			Preview: true,
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"bucketName": resource.NewStringProperty("name"),
				"arn":        resource.NewStringProperty("bucketArn"),
			}),
			News:    mustMarshalProperties(t, resource.PropertyMap{"bucketName": resource.NewStringProperty("name")}),
			Timeout: float64((5 * time.Minute).Seconds()),
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "custom:resource", "name"))

		mockCustomResource.EXPECT().PreviewCustomResourceOutputs().Return(
			resource.PropertyMap{"data": resource.MakeComputed(resource.NewStringProperty(""))},
		)

		resp, err := provider.Update(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		require.True(t, props.HasValue("data"), "Expected 'data' property in response")
		assert.True(t, props["data"].IsComputed())
	})
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")
	req := &pulumirpc.CreateRequest{
		Urn:        string(urn),
		Properties: mustMarshalProperties(t, resource.PropertyMap{"my": resource.NewStringProperty("input value")}),
		Timeout:    float64((5 * time.Minute).Seconds()),
	}

	t.Run("CustomResource", func(t *testing.T) {
		mockCustomResource.EXPECT().Create(ctx, urn, gomock.Any(), 5*time.Minute).Return(
			stringPtr("resource-id"), resource.PropertyMap{"foo": resource.NewStringProperty("bar")}, nil,
		)

		resp, err := provider.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, "resource-id", resp.Id)
		assert.NotNil(t, resp.Properties)
	})

	t.Run("CustomResource/Error", func(t *testing.T) {
		mockCustomResource.EXPECT().Create(ctx, urn, gomock.Any(), 5*time.Minute).Return(
			nil, nil, assert.AnError,
		)

		resp, err := provider.Create(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("StandardResource", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Create(ctx, gomock.Any(), "AWS::S3::Bucket", gomock.Any()).Return(
			stringPtr("bucket-id"), map[string]interface{}{"foo": "bar"}, nil,
		)

		resp, err := provider.Create(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, "bucket-id", resp.Id)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		require.True(t, props.HasValue("foo"), "Expected 'foo' property in response")
		assert.Equal(t, "bar", props["foo"].StringValue())
		require.True(t, props.HasValue("__inputs"), "Expected '__inputs' property in response")
		require.True(t, props["__inputs"].IsSecret(), "Expected '__inputs' to be a secret")
		inputs := props["__inputs"].SecretValue().Element.ObjectValue()
		require.True(t, inputs.HasValue("my"), "Expected 'my' property in '__inputs'")
		assert.Equal(t, "input value", inputs["my"].StringValue())
	})

	t.Run("StandardResource/NotFound", func(t *testing.T) {
		req.Urn = string(resource.NewURN("stack", "project", "parent", "unknown:resource", "name"))

		resp, err := provider.Create(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("StandardResource/Error", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Create(ctx, gomock.Any(), "AWS::S3::Bucket", gomock.Any()).Return(
			nil, nil, assert.AnError,
		)

		resp, err := provider.Create(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("PartialError", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Create(ctx, gomock.Any(), "AWS::S3::Bucket", gomock.Any()).Return(
			stringPtr("bucket-id"), map[string]interface{}{"foo": "bar"}, assert.AnError,
		)

		resp, err := provider.Create(ctx, req)
		assert.Error(t, err)
		resourceStatus, id, ins, outs, resourceErr := parsePluginError(err)

		assert.Nil(t, resp)
		assert.Equal(t, resource.StatusPartialFailure, resourceStatus)
		assert.Equal(t, "bucket-id", id.String())
		assert.Error(t, resourceErr)

		require.NotNil(t, ins)
		require.NotNil(t, outs)
		props := mustUnmarshalProperties(t, outs)

		require.True(t, props.HasValue("foo"), "Expected 'foo' property in response")
		assert.Equal(t, "bar", props["foo"].StringValue())
		require.True(t, props.HasValue("__inputs"), "Expected '__inputs' property in response")
		require.True(t, props["__inputs"].IsSecret(), "Expected '__inputs' to be a secret")
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		require.True(t, checkpoint.HasValue("my"), "Expected 'my' property in '__inputs'")
		assert.Equal(t, "input value", checkpoint["my"].StringValue())

		inputs := mustUnmarshalProperties(t, ins)
		assert.True(t, inputs.DeepEquals(checkpoint), "Expected inputs and checkpoint to be equal")
	})
}

func TestRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")
	req := &pulumirpc.ReadRequest{
		Urn:        string(urn),
		Id:         "resource-id",
		Properties: mustMarshalProperties(t, resource.PropertyMap{"foo": resource.NewStringProperty("bar")}),
	}

	t.Run("CustomResource", func(t *testing.T) {
		mockCustomResource.EXPECT().Read(ctx, urn, "resource-id", gomock.Any(), gomock.Any()).Return(
			resource.PropertyMap{"foo": resource.NewStringProperty("bar")},
			resource.PropertyMap{"foo": resource.NewStringProperty("bar")},
			true,
			nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, "resource-id", resp.Id)
		assert.NotNil(t, resp.Properties)
	})

	t.Run("CustomResource/NotFound", func(t *testing.T) {
		mockCustomResource.EXPECT().Read(ctx, urn, "resource-id", gomock.Any(), gomock.Any()).Return(
			nil,
			nil,
			false,
			nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		assert.Empty(t, resp.Id, "Expected empty ID for non-existent resource")
	})

	t.Run("CustomResource/Error", func(t *testing.T) {
		mockCustomResource.EXPECT().Read(ctx, urn, "resource-id", gomock.Any(), gomock.Any()).Return(
			nil,
			nil,
			false,
			assert.AnError,
		)

		resp, err := provider.Read(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	// In the import case there is no `__inputs` field in the old state.
	t.Run("StandardResource/Import", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName":        {TypeSpec: schema.TypeSpec{Type: "string"}},
				"objectLockEnabled": {TypeSpec: schema.TypeSpec{Type: "boolean"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Read(ctx, "AWS::S3::Bucket", "resource-id").Return(
			map[string]interface{}{
				"bucketName":        "my-bucket",
				"objectLockEnabled": true,
			}, true, nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, "resource-id", resp.Id)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.True(t, props.HasValue("bucketName"), "Expected 'bucketName' property in response")
		assert.Equal(t, "my-bucket", props["bucketName"].StringValue())
		assert.True(t, props.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in response")
		assert.True(t, props["objectLockEnabled"].BoolValue())

		require.NotNil(t, resp.Inputs)
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.True(t, inputs.HasValue("bucketName"), "Expected 'bucketName' property in inputs")
		assert.Equal(t, "my-bucket", inputs["bucketName"].StringValue())
		assert.True(t, inputs.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in inputs")
		assert.True(t, inputs["objectLockEnabled"].BoolValue())
	})

	t.Run("StandardResource/NotFound", func(t *testing.T) {
		req.Urn = string(resource.NewURN("stack", "project", "parent", "unknown:resource", "name"))

		resp, err := provider.Read(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("StandardResource/Error", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Read(ctx, "AWS::S3::Bucket", "resource-id").Return(
			nil, false, assert.AnError,
		)

		resp, err := provider.Read(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("StandardResource/NoDiff", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName":        {TypeSpec: schema.TypeSpec{Type: "string"}},
				"objectLockEnabled": {TypeSpec: schema.TypeSpec{Type: "boolean"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		inputs := resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
		}

		req.Properties = mustMarshalProperties(t, resource.PropertyMap{
			"foo":      resource.NewStringProperty("bar"),
			"__inputs": resource.MakeSecret(resource.NewObjectProperty(inputs)),
		})

		mockCCC.EXPECT().Read(ctx, "AWS::S3::Bucket", "resource-id").Return(
			map[string]interface{}{
				"foo":               "bar",
				"bucketName":        "my-bucket",
				"objectLockEnabled": true,
			}, true, nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.True(t, props.HasValue("foo"), "Expected 'foo' property in response")
		assert.Equal(t, "bar", props["foo"].StringValue())

		require.NotNil(t, resp.Inputs)
		inputs = mustUnmarshalProperties(t, resp.Inputs)
		assert.True(t, inputs.HasValue("bucketName"), "Expected 'bucketName' property in inputs")
		assert.Equal(t, "my-bucket", inputs["bucketName"].StringValue())
		assert.True(t, inputs.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in inputs")
		assert.True(t, inputs["objectLockEnabled"].BoolValue())
	})

	t.Run("StandardResource/Diff", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName":        {TypeSpec: schema.TypeSpec{Type: "string"}},
				"objectLockEnabled": {TypeSpec: schema.TypeSpec{Type: "boolean"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		inputs := resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
		}

		req.Properties = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
			"__inputs":          resource.MakeSecret(resource.NewObjectProperty(inputs)),
		})

		mockCCC.EXPECT().Read(ctx, "AWS::S3::Bucket", "resource-id").Return(
			map[string]interface{}{
				"foo": "bar",
				// Change the bucket name and object lock status
				"bucketName":        "other-bucket",
				"objectLockEnabled": false,
			}, true, nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.True(t, props.HasValue("bucketName"), "Expected 'bucketName' property in response")
		assert.Equal(t, "other-bucket", props["bucketName"].StringValue())
		assert.True(t, props.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in response")
		assert.False(t, props["objectLockEnabled"].BoolValue())

		require.NotNil(t, resp.Inputs)
		// TODO: Enable those assertions after fixing https://github.com/pulumi/pulumi-aws-native/issues/1796
		// inputs = mustUnmarshalProperties(t, resp.Inputs)
		// assert.True(t, inputs.HasValue("bucketName"), "Expected 'bucketName' property in inputs")
		// assert.Equal(t, "other-bucket", inputs["bucketName"].StringValue())
		// assert.True(t, inputs.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in inputs")
		// assert.False(t, inputs["objectLockEnabled"].BoolValue())
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")
	req := &pulumirpc.UpdateRequest{
		Urn:     string(urn),
		Id:      "resource-id",
		Olds:    mustMarshalProperties(t, resource.PropertyMap{"my": resource.NewStringProperty("old value")}),
		News:    mustMarshalProperties(t, resource.PropertyMap{"my": resource.NewStringProperty("new value")}),
		Timeout: float64((5 * time.Minute).Seconds()),
	}

	t.Run("CustomResource", func(t *testing.T) {
		mockCustomResource.EXPECT().Update(ctx, urn, "resource-id", gomock.Any(), gomock.Any(), gomock.Any(), 5*time.Minute).Return(
			resource.PropertyMap{"foo": resource.NewStringProperty("bar")}, nil,
		)

		resp, err := provider.Update(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.True(t, props.HasValue("foo"), "Expected 'foo' property in response")
		assert.Equal(t, "bar", props["foo"].StringValue())
	})

	t.Run("CustomResource/Error", func(t *testing.T) {
		mockCustomResource.EXPECT().Update(ctx, urn, "resource-id", gomock.Any(), gomock.Any(), gomock.Any(), 5*time.Minute).Return(
			nil, assert.AnError,
		)

		resp, err := provider.Update(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("StandardResource", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName":        {TypeSpec: schema.TypeSpec{Type: "string"}},
				"objectLockEnabled": {TypeSpec: schema.TypeSpec{Type: "boolean"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		inputs := resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("new-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(false),
		}
		req.News = mustMarshalProperties(t, inputs)

		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
			"__inputs":          resource.MakeSecret(resource.NewObjectProperty(inputs)),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id", gomock.Any()).Return(
			map[string]interface{}{
				// Change the bucket name and object lock status according to the inputs
				"bucketName":        resource.NewStringProperty("new-bucket"),
				"objectLockEnabled": resource.NewBoolProperty(false),
			}, nil,
		)

		resp, err := provider.Update(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.True(t, props.HasValue("bucketName"), "Expected 'bucketName' property in response")
		assert.Equal(t, "new-bucket", props["bucketName"].StringValue())
		assert.True(t, props.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in response")
		assert.False(t, props["objectLockEnabled"].BoolValue())

		require.True(t, props.HasValue("__inputs"), "Expected '__inputs' property in response")
		require.True(t, props["__inputs"].IsSecret(), "Expected '__inputs' to be a secret")
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		require.True(t, checkpoint.HasValue("bucketName"), "Expected 'bucketName' property in '__inputs'")
		assert.Equal(t, "new-bucket", checkpoint["bucketName"].StringValue())
		require.True(t, checkpoint.HasValue("objectLockEnabled"), "Expected 'objectLockEnabled' property in '__inputs'")
		assert.False(t, checkpoint["objectLockEnabled"].BoolValue())
	})

	t.Run("StandardResource/Error", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id", gomock.Any()).Return(
			nil, assert.AnError,
		)

		resp, err := provider.Update(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	mockCustomResource := resources.NewMockCustomResource(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name:            "test-provider",
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{"custom:resource": mockCustomResource},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}

	urn := resource.NewURN("stack", "project", "parent", "custom:resource", "name")
	req := &pulumirpc.DeleteRequest{
		Urn: string(urn),
		Id:  "resource-id",
	}

	t.Run("CustomResource", func(t *testing.T) {
		mockCustomResource.EXPECT().Delete(ctx, urn, "resource-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		_, err := provider.Delete(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("CustomResource/Error", func(t *testing.T) {
		mockCustomResource.EXPECT().Delete(ctx, urn, "resource-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(assert.AnError)

		_, err := provider.Delete(ctx, req)
		assert.Error(t, err)
	})

	t.Run("StandardResource", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Delete(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id").Return(nil)

		_, err := provider.Delete(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("StandardResource/Error", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))

		mockCCC.EXPECT().Delete(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id").Return(assert.AnError)

		_, err := provider.Delete(ctx, req)
		assert.Error(t, err)
	})

	t.Run("StandardResource/NotFound", func(t *testing.T) {
		req.Urn = string(resource.NewURN("stack", "project", "parent", "unknown:resource", "name"))

		_, err := provider.Delete(ctx, req)
		assert.Error(t, err)
	})
}

func mustMarshalProperties(t *testing.T, props resource.PropertyMap) *structpb.Struct {
	marshaled, err := plugin.MarshalProperties(props, plugin.MarshalOptions{
		Label:        "test",
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		t.Fatalf("failed to marshal properties: %v", err)
	}
	return marshaled
}

func mustUnmarshalProperties(t *testing.T, props *structpb.Struct) resource.PropertyMap {
	unmarshaled, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{
		Label:        "test",
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		t.Fatalf("failed to unmarshal properties: %v", err)
	}
	return unmarshaled
}

func parsePluginError(err error) (
	resourceStatus resource.Status, id resource.ID, inputs, props *structpb.Struct, resourceErr error,
) {
	responseErr := rpcerror.Convert(err)

	// If resource was successfully created but failed to initialize, the error will be packed
	// with the live properties of the object.
	resourceErr = responseErr
	for _, detail := range responseErr.Details() {
		if initErr, ok := detail.(*pulumirpc.ErrorResourceInitFailed); ok {
			id = resource.ID(initErr.GetId())
			props = initErr.GetProperties()
			inputs = initErr.GetInputs()
			resourceStatus = resource.StatusPartialFailure
			resourceErr = &plugin.InitError{Reasons: initErr.Reasons}
			break
		}
	}

	return resourceStatus, id, inputs, props, resourceErr
}

func stringPtr(s string) *string {
	return &s
}
