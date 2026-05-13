package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/mattbaird/jsonpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
	grpcmetadata "google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/autonaming"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/client"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/default_tags"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/resources"
)

type listTestStream struct {
	ctx       context.Context
	sendErr   error
	responses []*pulumirpc.ListResponse
}

const testBucketA = "bucket-a"

func (s *listTestStream) Send(response *pulumirpc.ListResponse) error {
	if s.sendErr != nil {
		return s.sendErr
	}
	s.responses = append(s.responses, response)
	return nil
}

func (s *listTestStream) SetHeader(grpcmetadata.MD) error {
	return nil
}

func (s *listTestStream) SendHeader(grpcmetadata.MD) error {
	return nil
}

func (s *listTestStream) SetTrailer(grpcmetadata.MD) {
}

func (s *listTestStream) Context() context.Context {
	if s.ctx != nil {
		return s.ctx
	}
	return context.Background()
}

func (s *listTestStream) SendMsg(any) error {
	return nil
}

func (s *listTestStream) RecvMsg(any) error {
	return nil
}

func listQueryStruct(t *testing.T, props resource.PropertyMap) *structpb.Struct {
	query, err := plugin.MarshalProperties(props, plugin.MarshalOptions{
		Label:        "list.query",
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	require.NoError(t, err)
	return query
}

func TestListRejectsNonListableResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType: "AWS::S3::Bucket",
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             client.NewMockCloudControlClient(ctrl),
	}

	err := provider.List(&pulumirpc.ListRequest{Token: "aws:s3/bucket:Bucket"}, &listTestStream{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not support List")
}

func TestListRejectsUnknownResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := &cfnProvider{
		resourceMap:     &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{}},
		customResources: map[string]resources.CustomResource{},
		ccc:             client.NewMockCloudControlClient(ctrl),
	}

	err := provider.List(&pulumirpc.ListRequest{Token: "aws:missing/resource:Resource"}, &listTestStream{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestListReturnsComputedForUnknownQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:apigateway/stage:Stage": {
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Properties: map[string]metadata.ListHandlerProperty{
						"RestApiId": {Type: "string"},
					},
				},
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             client.NewMockCloudControlClient(ctrl),
	}

	stream := &listTestStream{}
	err := provider.List(&pulumirpc.ListRequest{
		Token: "aws:apigateway/stage:Stage",
		Query: listQueryStruct(t, resource.PropertyMap{
			"restApiId": resource.MakeComputed(resource.NewStringProperty("")),
		}),
	}, stream)

	require.NoError(t, err)
	require.Len(t, stream.responses, 1)
	assert.NotNil(t, stream.responses[0].GetComputed())
}

func TestListRejectsQueryWhenListInputsAreEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:            "AWS::S3::Bucket",
				HasListHandler:    true,
				ListHandlerSchema: &metadata.ListHandlerSchema{},
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             client.NewMockCloudControlClient(ctrl),
	}

	err := provider.List(&pulumirpc.ListRequest{
		Token: "aws:s3/bucket:Bucket",
		Query: listQueryStruct(t, resource.PropertyMap{
			"name": resource.NewStringProperty(testBucketA),
		}),
	}, &listTestStream{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "List query is not supported")
}

func TestListRejectsInvalidQueryBeforeCloudControl(t *testing.T) {
	tests := []struct {
		name      string
		spec      metadata.CloudAPIResource
		query     resource.PropertyMap
		wantError string
	}{
		{
			name: "nil list handler schema with query",
			spec: metadata.CloudAPIResource{
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
			query: resource.PropertyMap{
				"name": resource.NewStringProperty(testBucketA),
			},
			wantError: "List query is not supported",
		},
		{
			name: "missing required",
			spec: metadata.CloudAPIResource{
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Properties: map[string]metadata.ListHandlerProperty{"RestApiId": {Type: "string"}},
					Required:   []string{"RestApiId"},
				},
			},
			query:     resource.PropertyMap{},
			wantError: "missing required List query property",
		},
		{
			name: "unknown property",
			spec: metadata.CloudAPIResource{
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Properties: map[string]metadata.ListHandlerProperty{"RestApiId": {Type: "string"}},
				},
			},
			query: resource.PropertyMap{
				"restApiId": resource.NewStringProperty("api-123"),
				"unknown":   resource.NewStringProperty("value"),
			},
			wantError: "unknown List query property",
		},
		{
			name: "wrong type",
			spec: metadata.CloudAPIResource{
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Properties: map[string]metadata.ListHandlerProperty{"RestApiId": {Type: "string"}},
				},
			},
			query: resource.PropertyMap{
				"restApiId": resource.NewNumberProperty(42),
			},
			wantError: "expected string",
		},
		{
			name: "required without schema",
			spec: metadata.CloudAPIResource{
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Required: []string{"RestApiId"},
				},
			},
			query:     resource.PropertyMap{},
			wantError: "required properties without schemas",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			provider := &cfnProvider{
				resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
					"aws:apigateway/stage:Stage": tt.spec,
				}},
				customResources: map[string]resources.CustomResource{},
				ccc:             client.NewMockCloudControlClient(ctrl),
			}

			err := provider.List(&pulumirpc.ListRequest{
				Token: "aws:apigateway/stage:Stage",
				Query: listQueryStruct(t, tt.query),
			}, &listTestStream{})

			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantError)
		})
	}
}

func TestListEmptyQueryAndContinuation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
	}

	id1, id2, continuation := testBucketA, "bucket-b", "next-page"
	expectedMaxResults := int32(50)
	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{MaxResults: &expectedMaxResults}).
		Return([]string{id1, id2}, &continuation, nil)

	stream := &listTestStream{}
	err := provider.List(&pulumirpc.ListRequest{
		Token:    "aws:s3/bucket:Bucket",
		PageSize: 75,
		Limit:    50,
	}, stream)

	require.NoError(t, err)
	require.Len(t, stream.responses, 3)
	assert.Equal(t, id1, stream.responses[0].GetResult().GetId())
	assert.Equal(t, id2, stream.responses[1].GetResult().GetId())
	assert.Equal(t, continuation, stream.responses[2].GetContinuation().GetContinuationToken())
}

func TestListMaxResults(t *testing.T) {
	tests := []struct {
		name     string
		pageSize int64
		limit    int64
		expected *int32
	}{
		{name: "page size only", pageSize: 25, expected: int32Ptr(25)},
		{name: "limit only", limit: 15, expected: int32Ptr(15)},
		{name: "omitted", expected: nil},
		{name: "capped", pageSize: 250, expected: int32Ptr(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCCC := client.NewMockCloudControlClient(ctrl)
			provider := &cfnProvider{
				resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
					"aws:s3/bucket:Bucket": {
						CfType:         "AWS::S3::Bucket",
						HasListHandler: true,
					},
				}},
				customResources: map[string]resources.CustomResource{},
				ccc:             mockCCC,
			}

			mockCCC.EXPECT().
				List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{MaxResults: tt.expected}).
				Return(nil, nil, nil)

			err := provider.List(&pulumirpc.ListRequest{
				Token:    "aws:s3/bucket:Bucket",
				PageSize: tt.pageSize,
				Limit:    tt.limit,
			}, &listTestStream{})

			require.NoError(t, err)
		})
	}
}

func TestListErrorsWhenCloudControlReturnsMoreThanRequested(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
	}

	id1, id2 := testBucketA, "bucket-b"
	expectedMaxResults := int32(1)
	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{MaxResults: &expectedMaxResults}).
		Return([]string{id1, id2}, nil, nil)

	err := provider.List(&pulumirpc.ListRequest{
		Token: "aws:s3/bucket:Bucket",
		Limit: 1,
	}, &listTestStream{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "more than requested maximum")
}

func TestListWrapsCloudControlErrorsWithTokenContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
	}

	expectedErr := errors.New("access denied")
	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{}).
		Return(nil, nil, expectedErr)

	err := provider.List(&pulumirpc.ListRequest{Token: "aws:s3/bucket:Bucket"}, &listTestStream{})

	require.ErrorIs(t, err, expectedErr)
	assert.Contains(t, err.Error(), "aws:s3/bucket:Bucket")
	assert.Contains(t, err.Error(), "AWS::S3::Bucket")
}

func TestListPropagatesProviderCancellation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	providerCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: providerCtx,
			cancel:  cancel,
		},
	}

	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{}).
		DoAndReturn(func(
			ctx context.Context,
			_ string,
			_ client.ListRequest,
		) ([]string, *string, error) {
			cancel()
			select {
			case <-ctx.Done():
			case <-time.After(time.Second):
				t.Fatal("timed out waiting for List context cancellation")
			}
			return nil, nil, ctx.Err()
		})

	err := provider.List(&pulumirpc.ListRequest{Token: "aws:s3/bucket:Bucket"}, &listTestStream{})

	require.ErrorIs(t, err, context.Canceled)
}

func TestListReturnsStreamSendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:s3/bucket:Bucket": {
				CfType:         "AWS::S3::Bucket",
				HasListHandler: true,
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
	}

	id := testBucketA
	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::S3::Bucket", client.ListRequest{}).
		Return([]string{id}, nil, nil)

	sendErr := errors.New("send failed")
	err := provider.List(&pulumirpc.ListRequest{Token: "aws:s3/bucket:Bucket"}, &listTestStream{sendErr: sendErr})

	require.ErrorIs(t, err, sendErr)
}

func TestListScopedQueryConvertsToCloudFormationCasing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	provider := &cfnProvider{
		resourceMap: &metadata.CloudAPIMetadata{Resources: map[string]metadata.CloudAPIResource{
			"aws:apigateway/stage:Stage": {
				CfType:         "AWS::ApiGateway::Stage",
				HasListHandler: true,
				ListHandlerSchema: &metadata.ListHandlerSchema{
					Properties: map[string]metadata.ListHandlerProperty{
						"RestApiId":      {Type: "string"},
						"IncludeValues":  {Type: "boolean"},
						"MinimumVersion": {Type: "integer"},
					},
					Required: []string{"RestApiId"},
				},
			},
		}},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
	}

	nextToken := "next"
	mockCCC.EXPECT().
		List(gomock.Any(), "AWS::ApiGateway::Stage", gomock.Any()).
		DoAndReturn(func(
			_ context.Context,
			_ string,
			request client.ListRequest,
		) ([]string, *string, error) {
			require.NotNil(t, request.ResourceModel)
			assert.JSONEq(t, `{"RestApiId":"api-123","IncludeValues":true,"MinimumVersion":2}`, *request.ResourceModel)
			assert.Equal(t, &nextToken, request.NextToken)
			assert.Nil(t, request.MaxResults)
			return nil, nil, nil
		})

	err := provider.List(&pulumirpc.ListRequest{
		Token:             "aws:apigateway/stage:Stage",
		ContinuationToken: nextToken,
		Query: listQueryStruct(t, resource.PropertyMap{
			"restApiId": resource.NewOutputProperty(resource.Output{
				Element: resource.MakeSecret(resource.NewStringProperty("api-123")),
				Known:   true,
			}),
			"includeValues":  resource.NewBoolProperty(true),
			"minimumVersion": resource.NewNumberProperty(2),
		}),
	}, &listTestStream{})

	require.NoError(t, err)
}

func TestListRejectsInvalidQuery(t *testing.T) {
	spec := metadata.CloudAPIResource{
		HasListHandler: true,
		ListHandlerSchema: &metadata.ListHandlerSchema{
			Properties: map[string]metadata.ListHandlerProperty{
				"RestApiId": {Type: "string"},
			},
			Required: []string{"RestApiId"},
		},
	}

	_, err := convertListQueryToCfn(&spec, resource.PropertyMap{})
	require.ErrorContains(t, err, "missing required List query property")

	_, err = convertListQueryToCfn(&spec, resource.PropertyMap{
		"restApiId": resource.NewStringProperty("api-123"),
		"unknown":   resource.NewStringProperty("value"),
	})
	require.ErrorContains(t, err, "unknown List query property")

	_, err = convertListQueryToCfn(&spec, resource.PropertyMap{
		"restApiId": resource.NewNumberProperty(42),
	})
	require.ErrorContains(t, err, "expected string")

	_, err = convertListQueryToCfn(&metadata.CloudAPIResource{
		HasListHandler: true,
		ListHandlerSchema: &metadata.ListHandlerSchema{
			Required: []string{"RestApiId"},
		},
	}, resource.PropertyMap{})
	require.ErrorContains(t, err, "required properties without schemas")
}

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

	t.Run("S3UsePathStyle config", func(t *testing.T) {
		req := &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:skipCredentialsValidation": "true",
				"aws-native:config:s3UsePathStyle":            "true",
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

	t.Run("s3UsePathStyle is accepted", func(t *testing.T) {
		news, err := plugin.MarshalProperties(resource.PropertyMap{
			"s3UsePathStyle": resource.NewBoolProperty(true),
			"region":         resource.NewStringProperty("us-west-2"),
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
			"aws-native:s3:Bucket": {
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
			"aws-native:s3:Bucket": {
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

	t.Run("StandardResource/PreservesAnyPropertyKeys", func(t *testing.T) {
		provider.resourceMap.Resources["aws:iam/role:Role"] = metadata.CloudAPIResource{
			CfType: "AWS::IAM::Role",
			Inputs: map[string]schema.PropertySpec{
				"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
			},
			Outputs: map[string]schema.PropertySpec{
				"awsId": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			IrreversibleNames: map[string]string{
				"awsId": "Id",
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:iam/role:Role", "name"))
		policyDocument := resource.NewObjectProperty(resource.PropertyMap{
			"Version": resource.NewStringProperty("2012-10-17"),
			"Statement": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"Action": resource.NewStringProperty("sts:AssumeRole"),
				}),
			}),
		})
		req.Properties = mustMarshalProperties(t, resource.PropertyMap{
			"assumeRolePolicyDocument": policyDocument,
		})

		mockCCC.EXPECT().Create(ctx, gomock.Any(), "AWS::IAM::Role", gomock.Any()).Return(
			stringPtr("role-id"),
			map[string]interface{}{
				"Id": "role-output-id",
				"AssumeRolePolicyDocument": map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []interface{}{
						map[string]interface{}{"Action": "sts:AssumeRole"},
					},
				},
			}, nil,
		)

		resp, err := provider.Create(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "role-output-id", props["awsId"].StringValue())
		doc := props["assumeRolePolicyDocument"].ObjectValue()
		assert.Equal(t, "2012-10-17", doc["Version"].StringValue())
		statement := doc["Statement"].ArrayValue()[0].ObjectValue()
		assert.Equal(t, "sts:AssumeRole", statement["Action"].StringValue())
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
		req.Properties = mustMarshalProperties(t, resource.PropertyMap{"my": resource.NewStringProperty("input value")})

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

	t.Run("StandardResource/ImportPreservesAnyPropertyKeys", func(t *testing.T) {
		provider.resourceMap.Resources["aws:iam/role:Role"] = metadata.CloudAPIResource{
			CfType: "AWS::IAM::Role",
			Inputs: map[string]schema.PropertySpec{
				"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:iam/role:Role", "name"))
		req.Properties = mustMarshalProperties(t, resource.PropertyMap{"foo": resource.NewStringProperty("bar")})

		mockCCC.EXPECT().Read(ctx, "AWS::IAM::Role", "resource-id").Return(
			map[string]interface{}{
				"AssumeRolePolicyDocument": map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []interface{}{
						map[string]interface{}{"Action": "sts:AssumeRole"},
					},
				},
			}, true, nil,
		)

		resp, err := provider.Read(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		doc := props["assumeRolePolicyDocument"].ObjectValue()
		assert.True(t, doc.HasValue("Version"), "Expected Any property key to be preserved")
		assert.Equal(t, "2012-10-17", doc["Version"].StringValue())
		statement := doc["Statement"].ArrayValue()[0].ObjectValue()
		assert.True(t, statement.HasValue("Action"), "Expected nested Any property key to be preserved")
		assert.Equal(t, "sts:AssumeRole", statement["Action"].StringValue())

		require.NotNil(t, resp.Inputs)
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		inputDoc := inputs["assumeRolePolicyDocument"].ObjectValue()
		assert.True(t, inputDoc.HasValue("Version"), "Expected imported inputs to preserve Any property key")
		assert.Equal(t, "2012-10-17", inputDoc["Version"].StringValue())
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
		oldInputs := resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
		}
		req.News = mustMarshalProperties(t, inputs)

		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName":        resource.NewStringProperty("my-bucket"),
			"objectLockEnabled": resource.NewBoolProperty(true),
			"__inputs":          resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id", jsonPatchEquals([]jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/BucketName", Value: "new-bucket"},
			{Operation: "replace", Path: "/ObjectLockEnabled", Value: false},
		})).Return(
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

	t.Run("StandardResource/PreservesAnyPropertyKeys", func(t *testing.T) {
		provider.resourceMap.Resources["aws:iam/role:Role"] = metadata.CloudAPIResource{
			CfType: "AWS::IAM::Role",
			Inputs: map[string]schema.PropertySpec{
				"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
			},
			Outputs: map[string]schema.PropertySpec{
				"awsId": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			IrreversibleNames: map[string]string{
				"awsId": "Id",
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:iam/role:Role", "name"))
		oldPolicyDocument := resource.NewObjectProperty(resource.PropertyMap{
			"Version": resource.NewStringProperty("2012-10-17"),
		})
		newPolicyDocument := resource.NewObjectProperty(resource.PropertyMap{
			"Version": resource.NewStringProperty("2012-10-17"),
			"Statement": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"Action": resource.NewStringProperty("sts:AssumeRole"),
				}),
			}),
		})
		req.News = mustMarshalProperties(t, resource.PropertyMap{
			"assumeRolePolicyDocument": newPolicyDocument,
		})
		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"assumeRolePolicyDocument": oldPolicyDocument,
			"__inputs": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
				"assumeRolePolicyDocument": oldPolicyDocument,
			})),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::IAM::Role", "resource-id", gomock.Any()).Return(
			map[string]interface{}{
				"Id": "role-output-id",
				"AssumeRolePolicyDocument": map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []interface{}{
						map[string]interface{}{"Action": "sts:AssumeRole"},
					},
				},
			}, nil,
		)

		resp, err := provider.Update(ctx, req)
		assert.NoError(t, err)
		require.NotNil(t, resp.Properties)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "role-output-id", props["awsId"].StringValue())
		doc := props["assumeRolePolicyDocument"].ObjectValue()
		assert.Equal(t, "2012-10-17", doc["Version"].StringValue())
		statement := doc["Statement"].ArrayValue()[0].ObjectValue()
		assert.Equal(t, "sts:AssumeRole", statement["Action"].StringValue())
	})

	t.Run("StandardResource/ActualBaselineRepairDrift", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			Outputs: map[string]schema.PropertySpec{
				"bucketName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))
		req.News = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName": resource.NewStringProperty("desired-bucket"),
		})
		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName": resource.NewStringProperty("actual-drifted-bucket"),
			"__inputs": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
				"bucketName": resource.NewStringProperty("desired-bucket"),
			})),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id", jsonPatchEquals([]jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/BucketName", Value: "desired-bucket"},
		})).Return(
			map[string]interface{}{
				"bucketName": "desired-bucket",
			}, nil,
		)

		resp, err := provider.Update(ctx, req)
		require.NoError(t, err)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "desired-bucket", props["bucketName"].StringValue())
	})

	t.Run("StandardResource/WriteOnlyResend", func(t *testing.T) {
		provider.resourceMap.Resources["aws:rds/dbInstance:DBInstance"] = metadata.CloudAPIResource{
			CfType: "AWS::RDS::DBInstance",
			Inputs: map[string]schema.PropertySpec{
				"password": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"engine":   {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			Outputs: map[string]schema.PropertySpec{
				"engine": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			WriteOnly: []string{"password"},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:rds/dbInstance:DBInstance", "name"))
		req.News = mustMarshalProperties(t, resource.PropertyMap{
			"password": resource.NewStringProperty("secret"),
			"engine":   resource.NewStringProperty("mysql"),
		})
		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"engine": resource.NewStringProperty("postgres"),
			"__inputs": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
				"password": resource.NewStringProperty("secret"),
				"engine":   resource.NewStringProperty("postgres"),
			})),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::RDS::DBInstance", "resource-id", jsonPatchEquals([]jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Engine", Value: "mysql"},
			{Operation: "add", Path: "/Password", Value: "secret"},
		})).Return(
			map[string]interface{}{
				"engine": "mysql",
			}, nil,
		)

		resp, err := provider.Update(ctx, req)
		require.NoError(t, err)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "mysql", props["engine"].StringValue())
		assert.Equal(t, "secret", props["password"].StringValue())
	})

	t.Run("StandardResource/CreateOnlyWriteOnlyCarriedButNotResent", func(t *testing.T) {
		provider.resourceMap.Resources["aws:rds/dbInstance:DBInstance"] = metadata.CloudAPIResource{
			CfType: "AWS::RDS::DBInstance",
			Inputs: map[string]schema.PropertySpec{
				"password": {TypeSpec: schema.TypeSpec{Type: "string"}},
				"engine":   {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			Outputs: map[string]schema.PropertySpec{
				"engine": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
			WriteOnly:  []string{"password"},
			CreateOnly: []string{"password"},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:rds/dbInstance:DBInstance", "name"))
		req.News = mustMarshalProperties(t, resource.PropertyMap{
			"password": resource.NewStringProperty("secret"),
			"engine":   resource.NewStringProperty("mysql"),
		})
		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"engine": resource.NewStringProperty("postgres"),
			"__inputs": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
				"password": resource.NewStringProperty("secret"),
				"engine":   resource.NewStringProperty("postgres"),
			})),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::RDS::DBInstance", "resource-id", jsonPatchEquals([]jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/Engine", Value: "mysql"},
		})).Return(
			map[string]interface{}{
				"engine": "mysql",
			}, nil,
		)

		resp, err := provider.Update(ctx, req)
		require.NoError(t, err)
		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "mysql", props["engine"].StringValue())
		assert.Equal(t, "secret", props["password"].StringValue())
	})

	t.Run("StandardResource/Error", func(t *testing.T) {
		provider.resourceMap.Resources["aws:s3/bucket:Bucket"] = metadata.CloudAPIResource{
			CfType: "AWS::S3::Bucket",
			Inputs: map[string]schema.PropertySpec{
				"bucketName": {TypeSpec: schema.TypeSpec{Type: "string"}},
			},
		}
		req.Urn = string(resource.NewURN("stack", "project", "parent", "aws:s3/bucket:Bucket", "name"))
		req.News = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName": resource.NewStringProperty("new-bucket"),
		})
		req.Olds = mustMarshalProperties(t, resource.PropertyMap{
			"bucketName": resource.NewStringProperty("old-bucket"),
			"__inputs":   resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{"bucketName": resource.NewStringProperty("old-bucket")})),
		})

		mockCCC.EXPECT().Update(ctx, gomock.Any(), "AWS::S3::Bucket", "resource-id", jsonPatchEquals([]jsonpatch.JsonPatchOperation{
			{Operation: "replace", Path: "/BucketName", Value: "new-bucket"},
		})).Return(
			nil, assert.AnError,
		)

		resp, err := provider.Update(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

}

func TestStandardResourceDiffUsesActualOutputBaseline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name: "test-provider",
		resourceMap: &metadata.CloudAPIMetadata{
			Resources: map[string]metadata.CloudAPIResource{
				"aws:rds/dbInstance:DBInstance": {
					CfType:       "AWS::RDS::DBInstance",
					TagsProperty: "tags",
					Inputs: map[string]schema.PropertySpec{
						"backupTarget": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"engine":       {TypeSpec: schema.TypeSpec{Type: "string"}},
						"password":     {TypeSpec: schema.TypeSpec{Type: "string"}},
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &schema.TypeSpec{Type: "string"},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"backupTarget": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"engine":       {TypeSpec: schema.TypeSpec{Type: "string"}},
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &schema.TypeSpec{Type: "string"},
							},
						},
					},
					WriteOnly: []string{"password"},
				},
				"aws:logs/logGroup:LogGroup": {
					CfType:       "AWS::Logs::LogGroup",
					TagsProperty: "tags",
					TagsStyle:    default_tags.TagsStyleKeyValueArray,
					Inputs: map[string]schema.PropertySpec{
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:index:Tag"},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:index:Tag"},
							},
						},
					},
				},
				"aws:iam:Role": {
					CfType: "AWS::IAM::Role",
					Inputs: map[string]schema.PropertySpec{
						"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policies": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policies": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
							},
						},
					},
				},
			},
			Types: map[string]metadata.CloudAPIType{
				"aws-native:index:Tag": {
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
						"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
				"aws-native:iam:RolePolicy": {
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"policyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policyName":     {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
			},
		},
		customResources: map[string]resources.CustomResource{},
		ccc:             client.NewMockCloudControlClient(ctrl),
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}
	urn := resource.NewURN("stack", "project", "parent", "aws:rds/dbInstance:DBInstance", "name")

	t.Run("external drift on managed input reports diff", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"engine": resource.NewStringProperty("abc")}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"engine":   resource.NewStringProperty("xyz"),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
			News: mustMarshalProperties(t, oldInputs),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_UNKNOWN, resp.Changes)
	})

	t.Run("newly visible optional computed default is ignored when unowned", func(t *testing.T) {
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("region"),
				"__inputs":     resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{})),
			}),
			News: mustMarshalProperties(t, resource.PropertyMap{}),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("user starts managing same optional computed value without cloud operation", func(t *testing.T) {
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("region"),
				"__inputs":     resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{})),
			}),
			News: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("region"),
			}),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("user starts managing different optional computed value reports diff", func(t *testing.T) {
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("region"),
				"__inputs":     resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{})),
			}),
			News: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("local"),
			}),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_UNKNOWN, resp.Changes)
	})

	t.Run("owned tag map drift reports diff", func(t *testing.T) {
		desiredTags := resource.PropertyMap{
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			}),
		}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"tags": resource.NewObjectProperty(resource.PropertyMap{
					"owner": resource.NewStringProperty("team"),
					"extra": resource.NewStringProperty("external"),
				}),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(desiredTags)),
			}),
			News: mustMarshalProperties(t, desiredTags),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_UNKNOWN, resp.Changes)
	})

	t.Run("aws managed tag in actual baseline does not report diff", func(t *testing.T) {
		desiredTags := resource.PropertyMap{
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			}),
		}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"tags": resource.NewObjectProperty(resource.PropertyMap{
					"owner":       resource.NewStringProperty("team"),
					"aws:managed": resource.NewStringProperty("external"),
				}),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(desiredTags)),
			}),
			News: mustMarshalProperties(t, desiredTags),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("key value array tag reorder in actual baseline does not report diff", func(t *testing.T) {
		arrayURN := resource.NewURN("stack", "project", "parent", "aws:logs/logGroup:LogGroup", "name")
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
		})
		actualTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
		})
		desiredTags := resource.PropertyMap{"tags": oldTags}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(arrayURN),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"tags":     actualTags,
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(desiredTags)),
			}),
			News: mustMarshalProperties(t, desiredTags),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("normalized IAM policy document actual baseline does not report diff", func(t *testing.T) {
		roleURN := resource.NewURN("stack", "project", "parent", "aws:iam:Role", "name")
		oldInputs := resource.PropertyMap{
			"assumeRolePolicyDocument": resource.NewStringProperty(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Principal": {"Service": "ec2.amazonaws.com"},
					"Action": "sts:AssumeRole"
				}]
			}`),
			"policies": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"policyName": resource.NewStringProperty("test-policy"),
					"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
						"Version": resource.NewStringProperty("2012-10-17"),
						"Statement": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewObjectProperty(resource.PropertyMap{
								"Effect":   resource.NewStringProperty("Allow"),
								"Action":   resource.NewStringProperty("*"),
								"Resource": resource.NewStringProperty("*"),
							}),
						}),
					}),
				}),
			}),
		}
		actualOutputs := resource.PropertyMap{
			"assumeRolePolicyDocument": resource.NewObjectProperty(resource.PropertyMap{
				"Version": resource.NewStringProperty("2012-10-17"),
				"Statement": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewObjectProperty(resource.PropertyMap{
						"Effect":    resource.NewStringProperty("Allow"),
						"Principal": resource.NewObjectProperty(resource.PropertyMap{"Service": resource.NewStringProperty("ec2.amazonaws.com")}),
						"Action":    resource.NewStringProperty("sts:AssumeRole"),
					}),
				}),
			}),
			"policies": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"policyName": resource.NewStringProperty("test-policy"),
					"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
						"Version": resource.NewStringProperty("2012-10-17"),
						"Statement": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewObjectProperty(resource.PropertyMap{
								"Effect":   resource.NewStringProperty("Allow"),
								"Action":   resource.NewArrayProperty([]resource.PropertyValue{resource.NewStringProperty("*")}),
								"Resource": resource.NewStringProperty("*"),
							}),
						}),
					}),
				}),
			}),
		}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(roleURN),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"assumeRolePolicyDocument": actualOutputs["assumeRolePolicyDocument"],
				"policies":                 actualOutputs["policies"],
				"__inputs":                 resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
			News: mustMarshalProperties(t, oldInputs),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("write-only input from checkpoint does not become a diff", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"password": resource.NewStringProperty("secret"),
		}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(urn),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
			News: mustMarshalProperties(t, oldInputs),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_NONE, resp.Changes)
	})

	t.Run("custom resource uses old input diff fallback", func(t *testing.T) {
		mockCustomResource := resources.NewMockCustomResource(ctrl)
		provider.customResources["custom:resource"] = mockCustomResource
		oldInputs := resource.PropertyMap{"name": resource.NewStringProperty("old")}
		resp, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(resource.NewURN("stack", "project", "parent", "custom:resource", "name")),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
			News: mustMarshalProperties(t, resource.PropertyMap{
				"name": resource.NewStringProperty("new"),
			}),
		})
		require.NoError(t, err)
		assert.Equal(t, pulumirpc.DiffResponse_DIFF_UNKNOWN, resp.Changes)
	})

	t.Run("missing standard spec returns error", func(t *testing.T) {
		_, err := provider.Diff(ctx, &pulumirpc.DiffRequest{
			Urn: string(resource.NewURN("stack", "project", "parent", "aws:missing:Resource", "name")),
			Olds: mustMarshalProperties(t, resource.PropertyMap{
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{})),
			}),
			News: mustMarshalProperties(t, resource.PropertyMap{}),
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Resource type aws:missing:Resource not found")
	})
}

func TestStandardResourceReadReturnsOwnershipFilteredInputs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCCC := client.NewMockCloudControlClient(ctrl)
	ctx, cancel := context.WithCancel(context.Background())
	provider := &cfnProvider{
		name: "test-provider",
		resourceMap: &metadata.CloudAPIMetadata{
			Resources: map[string]metadata.CloudAPIResource{
				"aws:rds/dbInstance:DBInstance": {
					CfType: "AWS::RDS::DBInstance",
					Inputs: map[string]schema.PropertySpec{
						"backupTarget": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"engine":       {TypeSpec: schema.TypeSpec{Type: "string"}},
						"password":     {TypeSpec: schema.TypeSpec{Type: "string"}},
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &schema.TypeSpec{Type: "string"},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"backupTarget": {TypeSpec: schema.TypeSpec{Type: "string"}},
						"engine":       {TypeSpec: schema.TypeSpec{Type: "string"}},
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type:                 "object",
								AdditionalProperties: &schema.TypeSpec{Type: "string"},
							},
						},
					},
					TagsProperty: "tags",
					WriteOnly:    []string{"password"},
					CreateOnly:   []string{"password"},
				},
				"aws:logs/logGroup:LogGroup": {
					CfType: "AWS::Logs::LogGroup",
					Inputs: map[string]schema.PropertySpec{
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: "#/types/aws-native:index:Tag",
								},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"tags": {
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Ref: "#/types/aws-native:index:Tag",
								},
							},
						},
					},
					TagsProperty: "tags",
					TagsStyle:    default_tags.TagsStyleKeyValueArray,
				},
				"aws:iam:Role": {
					CfType: "AWS::IAM::Role",
					Inputs: map[string]schema.PropertySpec{
						"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policies": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
							},
						},
					},
					Outputs: map[string]schema.PropertySpec{
						"assumeRolePolicyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policies": {
							TypeSpec: schema.TypeSpec{
								Type:  "array",
								Items: &schema.TypeSpec{Ref: "#/types/aws-native:iam:RolePolicy"},
							},
						},
					},
				},
			},
			Types: map[string]metadata.CloudAPIType{
				"aws-native:index:Tag": {
					Properties: map[string]schema.PropertySpec{
						"key":   {TypeSpec: schema.TypeSpec{Type: "string"}},
						"value": {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
				"aws-native:iam:RolePolicy": {
					Type: "object",
					Properties: map[string]schema.PropertySpec{
						"policyDocument": {TypeSpec: schema.TypeSpec{Ref: "pulumi.json#/Any"}},
						"policyName":     {TypeSpec: schema.TypeSpec{Type: "string"}},
					},
				},
			},
		},
		customResources: map[string]resources.CustomResource{},
		ccc:             mockCCC,
		canceler: &cancellationContext{
			context: ctx,
			cancel:  cancel,
		},
	}
	urn := resource.NewURN("stack", "project", "parent", "aws:rds/dbInstance:DBInstance", "name")

	t.Run("managed drift uses live actual value", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"engine": resource.NewStringProperty("postgres")}
		mockCCC.EXPECT().Read(ctx, "AWS::RDS::DBInstance", "resource-id").Return(
			map[string]interface{}{
				"engine": "mysql",
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(urn),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"engine":   resource.NewStringProperty("postgres"),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "mysql", props["engine"].StringValue())
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.Equal(t, "mysql", inputs["engine"].StringValue())
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		assert.Equal(t, "mysql", checkpoint["engine"].StringValue())
	})

	t.Run("managed optional computed drift uses live actual value", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"backupTarget": resource.NewStringProperty("local"),
		}
		mockCCC.EXPECT().Read(ctx, "AWS::RDS::DBInstance", "resource-id").Return(
			map[string]interface{}{
				"backupTarget": "region",
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(urn),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"backupTarget": resource.NewStringProperty("local"),
				"__inputs":     resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "region", props["backupTarget"].StringValue())
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.Equal(t, "region", inputs["backupTarget"].StringValue())
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		assert.Equal(t, "region", checkpoint["backupTarget"].StringValue())
	})

	t.Run("unowned optional computed default stays absent from inputs", func(t *testing.T) {
		oldInputs := resource.PropertyMap{"engine": resource.NewStringProperty("postgres")}
		mockCCC.EXPECT().Read(ctx, "AWS::RDS::DBInstance", "resource-id").Return(
			map[string]interface{}{
				"engine":       "postgres",
				"backupTarget": "region",
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(urn),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"engine":   resource.NewStringProperty("postgres"),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "region", props["backupTarget"].StringValue())
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.False(t, inputs.HasValue("backupTarget"))
		assert.Equal(t, "postgres", inputs["engine"].StringValue())
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		assert.False(t, checkpoint.HasValue("backupTarget"))
		assert.Equal(t, "postgres", checkpoint["engine"].StringValue())
	})

	t.Run("aws managed tag drift is not checkpointed as owned input", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"tags": resource.NewObjectProperty(resource.PropertyMap{
				"owner": resource.NewStringProperty("team"),
			}),
		}
		mockCCC.EXPECT().Read(ctx, "AWS::RDS::DBInstance", "resource-id").Return(
			map[string]interface{}{
				"tags": map[string]interface{}{
					"owner":       "team",
					"aws:managed": "external",
				},
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(urn),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"tags":     resource.NewObjectProperty(oldInputs["tags"].ObjectValue()),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		props := mustUnmarshalProperties(t, resp.Properties)
		stateTags := props["tags"].ObjectValue()
		assert.Equal(t, resource.NewStringProperty("external"), stateTags["aws:managed"])

		inputs := mustUnmarshalProperties(t, resp.Inputs)
		inputTags := inputs["tags"].ObjectValue()
		assert.False(t, inputTags.HasValue("aws:managed"))
		checkpointTags := props["__inputs"].SecretValue().Element.ObjectValue()["tags"].ObjectValue()
		assert.False(t, checkpointTags.HasValue("aws:managed"))
	})

	t.Run("key value array tag reorder is not checkpointed as owned input drift", func(t *testing.T) {
		arrayURN := resource.NewURN("stack", "project", "parent", "aws:logs/logGroup:LogGroup", "name")
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Environment"),
				"value": resource.NewStringProperty("prod"),
			}),
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("Name"),
				"value": resource.NewStringProperty("my-resource"),
			}),
		})
		refreshedTags := []interface{}{
			map[string]interface{}{"key": "Name", "value": "my-resource"},
			map[string]interface{}{"key": "Environment", "value": "prod"},
		}
		oldInputs := resource.PropertyMap{
			"tags": oldTags,
		}
		mockCCC.EXPECT().Read(ctx, "AWS::Logs::LogGroup", "resource-id").Return(
			map[string]interface{}{
				"tags": refreshedTags,
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(arrayURN),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"tags":     oldTags,
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.True(t, oldTags.DeepEquals(inputs["tags"]))
		props := mustUnmarshalProperties(t, resp.Properties)
		checkpointTags := props["__inputs"].SecretValue().Element.ObjectValue()["tags"]
		assert.True(t, oldTags.DeepEquals(checkpointTags))
	})

	t.Run("secret tag value remains secret after refresh", func(t *testing.T) {
		arrayURN := resource.NewURN("stack", "project", "parent", "aws:logs/logGroup:LogGroup", "name")
		oldTags := resource.NewArrayProperty([]resource.PropertyValue{
			resource.NewObjectProperty(resource.PropertyMap{
				"key":   resource.NewStringProperty("secretfoo"),
				"value": resource.MakeSecret(resource.NewStringProperty("secretbar")),
			}),
		})
		oldInputs := resource.PropertyMap{
			"tags": oldTags,
		}
		mockCCC.EXPECT().Read(ctx, "AWS::Logs::LogGroup", "resource-id").Return(
			map[string]interface{}{
				"tags": []interface{}{
					map[string]interface{}{"key": "secretfoo", "value": "secretbar"},
				},
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(arrayURN),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"tags":     oldTags,
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		inputs := mustUnmarshalProperties(t, resp.Inputs)
		inputTagValue := inputs["tags"].ArrayValue()[0].ObjectValue()["value"]
		require.True(t, inputTagValue.IsSecret())
		assert.Equal(t, resource.NewStringProperty("secretbar"), inputTagValue.SecretValue().Element)

		props := mustUnmarshalProperties(t, resp.Properties)
		outputTagValue := props["tags"].ArrayValue()[0].ObjectValue()["value"]
		require.True(t, outputTagValue.IsSecret())
		assert.Equal(t, resource.NewStringProperty("secretbar"), outputTagValue.SecretValue().Element)

		checkpointTagValue := props["__inputs"].SecretValue().Element.ObjectValue()["tags"].
			ArrayValue()[0].ObjectValue()["value"]
		require.True(t, checkpointTagValue.IsSecret())
		assert.Equal(t, resource.NewStringProperty("secretbar"), checkpointTagValue.SecretValue().Element)
	})

	t.Run("normalized IAM policy document is not checkpointed as owned input drift", func(t *testing.T) {
		roleURN := resource.NewURN("stack", "project", "parent", "aws:iam:Role", "name")
		oldInputs := resource.PropertyMap{
			"assumeRolePolicyDocument": resource.NewStringProperty(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Principal": {"Service": "ec2.amazonaws.com"},
					"Action": "sts:AssumeRole"
				}]
			}`),
			"policies": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"policyName": resource.NewStringProperty("test-policy"),
					"policyDocument": resource.NewObjectProperty(resource.PropertyMap{
						"Version": resource.NewStringProperty("2012-10-17"),
						"Statement": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewObjectProperty(resource.PropertyMap{
								"Effect":   resource.NewStringProperty("Allow"),
								"Action":   resource.NewStringProperty("*"),
								"Resource": resource.NewStringProperty("*"),
							}),
						}),
					}),
				}),
			}),
		}
		mockCCC.EXPECT().Read(ctx, "AWS::IAM::Role", "resource-id").Return(
			map[string]interface{}{
				"AssumeRolePolicyDocument": map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []interface{}{
						map[string]interface{}{
							"Effect":    "Allow",
							"Principal": map[string]interface{}{"Service": "ec2.amazonaws.com"},
							"Action":    "sts:AssumeRole",
						},
					},
				},
				"Policies": []interface{}{
					map[string]interface{}{
						"PolicyName": "test-policy",
						"PolicyDocument": map[string]interface{}{
							"Version": "2012-10-17",
							"Statement": []interface{}{
								map[string]interface{}{
									"Effect":   "Allow",
									"Action":   "*",
									"Resource": "*",
								},
							},
						},
					},
				},
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(roleURN),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"assumeRolePolicyDocument": oldInputs["assumeRolePolicyDocument"],
				"policies":                 oldInputs["policies"],
				"__inputs":                 resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.True(t, oldInputs["assumeRolePolicyDocument"].DeepEquals(inputs["assumeRolePolicyDocument"]))
		assert.True(t, oldInputs["policies"].DeepEquals(inputs["policies"]))
		props := mustUnmarshalProperties(t, resp.Properties)
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		assert.True(t, oldInputs["assumeRolePolicyDocument"].DeepEquals(checkpoint["assumeRolePolicyDocument"]))
		assert.True(t, oldInputs["policies"].DeepEquals(checkpoint["policies"]))
	})

	t.Run("create-only write-only value is carried forward through refresh", func(t *testing.T) {
		oldInputs := resource.PropertyMap{
			"engine":   resource.NewStringProperty("postgres"),
			"password": resource.NewStringProperty("secret"),
		}
		mockCCC.EXPECT().Read(ctx, "AWS::RDS::DBInstance", "resource-id").Return(
			map[string]interface{}{
				"engine": "postgres",
			}, true, nil,
		)

		resp, err := provider.Read(ctx, &pulumirpc.ReadRequest{
			Urn: string(urn),
			Id:  "resource-id",
			Properties: mustMarshalProperties(t, resource.PropertyMap{
				"engine":   resource.NewStringProperty("postgres"),
				"password": resource.NewStringProperty("secret"),
				"__inputs": resource.MakeSecret(resource.NewObjectProperty(oldInputs)),
			}),
		})
		require.NoError(t, err)

		props := mustUnmarshalProperties(t, resp.Properties)
		assert.Equal(t, "secret", props["password"].StringValue())
		inputs := mustUnmarshalProperties(t, resp.Inputs)
		assert.Equal(t, "secret", inputs["password"].StringValue())
		checkpoint := props["__inputs"].SecretValue().Element.ObjectValue()
		assert.Equal(t, "secret", checkpoint["password"].StringValue())
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

func jsonPatchEquals(expected []jsonpatch.JsonPatchOperation) gomock.Matcher {
	return jsonPatchMatcher{expected: expected}
}

type jsonPatchMatcher struct {
	expected []jsonpatch.JsonPatchOperation
}

func (m jsonPatchMatcher) Matches(x interface{}) bool {
	actual, ok := x.([]jsonpatch.JsonPatchOperation)
	return ok && reflect.DeepEqual(m.expected, actual)
}

func (m jsonPatchMatcher) String() string {
	return fmt.Sprintf("equals JSON patch %v", m.expected)
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

func int32Ptr(i int32) *int32 {
	return &i
}
