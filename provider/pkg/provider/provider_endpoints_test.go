package provider_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/provider"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestProviderEndpoints(t *testing.T) {
	stsGetCallerIdentityResponse := `
<GetCallerIdentityResponse xmlns="https://sts.amazonaws.com/doc/2011-06-15/">
	<GetCallerIdentityResult>
	<Arn>arn:aws:iam::123456789012:user/Alice</Arn>
	<UserId>AIDACKCEVSQ6C2EXAMPLE</UserId>
	<Account>123456789012</Account>
	</GetCallerIdentityResult>
	<ResponseMetadata>
		<RequestId>01234567-89ab-cdef-0123-456789abcdef</RequestId>
	</ResponseMetadata>
</GetCallerIdentityResponse>`

	t.Run("preview calls sts caller identity", func(t *testing.T) {
		t.Parallel()
		stsRequestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stsRequestCount++
			w.Write([]byte(stsGetCallerIdentityResponse))
		}))
		t.Cleanup(server.Close)

		provider, err := testProviderServer()
		require.NoError(t, err)
		ctx := context.Background()
		_, err = provider.Configure(ctx, &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:region":    "us-west-2",
				"aws-native:config:endpoints": `{"sts": "` + server.URL + `"}`,
				// Set fake credentials to avoid making real requests and override any local credentials.
				"aws-native:config:accessKey": "fake",
				"aws-native:config:secretKey": "fake",
				"aws-native:config:token":     "fake",
			},
		})
		require.NoError(t, err)
		assert.Equal(t, 1, stsRequestCount)

		invokeResponse, err := provider.Invoke(ctx, &pulumirpc.InvokeRequest{
			Tok: "aws-native:index:getAccountId",
		})
		require.NoError(t, err)
		assert.Equal(t, map[string]interface{}{
			"accountId": "123456789012",
		}, invokeResponse.Return.AsMap())
	})


	t.Run("configure calls set the APN/1.1 marketplace identifier in the User-Agent header", func(t *testing.T) {
		t.Parallel()
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "APN/1.1 ("+provider.PulumiAWSMarketplaceCode+")", r.Header.Get("User-Agent"))
			requestCount++
			w.Write([]byte(stsGetCallerIdentityResponse))
		}))
		t.Cleanup(server.Close)

		provider, err := testProviderServer()
		require.NoError(t, err)
		ctx := context.Background()
		_, err = provider.Configure(ctx, &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:region":    "us-west-2",
				"aws-native:config:endpoints": `{"sts": "` + server.URL + `"}`,
				// Set fake credentials to avoid making real requests and override any local credentials.
				"aws-native:config:accessKey": "fake",
				"aws-native:config:secretKey": "fake",
				"aws-native:config:token":     "fake",
			},
		})
		require.NoError(t, err)
		assert.Equal(t, 1, requestCount)
	})

	t.Run("skipping credentials doesn't break getAccountId", func(t *testing.T) {
		t.Parallel()
		stsRequestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stsRequestCount++
			w.Write([]byte(stsGetCallerIdentityResponse))
		}))
		t.Cleanup(server.Close)

		provider, err := testProviderServer()
		require.NoError(t, err)
		ctx := context.Background()
		_, err = provider.Configure(ctx, &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:region":                    "us-west-2",
				"aws-native:config:endpoints":                 `{"sts": "` + server.URL + `"}`,
				"aws-native:config:skipCredentialsValidation": "true",
				// Set fake credentials to avoid making real requests and override any local credentials.
				"aws-native:config:accessKey": "fake",
				"aws-native:config:secretKey": "fake",
				"aws-native:config:token":     "fake",
			},
		})
		require.NoError(t, err)
		assert.Equal(t, 0, stsRequestCount) // No STS request should be made due to skipping verification.

		invokeResponse, err := provider.Invoke(ctx, &pulumirpc.InvokeRequest{
			Tok: "aws-native:index:getAccountId",
		})
		require.NoError(t, err)
		assert.Equal(t, 1, stsRequestCount)
		assert.Equal(t, map[string]interface{}{
			"accountId": "123456789012",
		}, invokeResponse.Return.AsMap())
	})

	t.Run("resource create", func(t *testing.T) {
		t.Parallel()

		reqCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCount++
			if reqCount == 1 {
				w.Write([]byte(`
			{
				"ProgressEvent": {
					"Identifier": "CloudApiLogGroup",
					"OperationStatus": "IN_PROGRESS",
					"EventTime": 1707429589,
					"TypeName": "AWS::Logs::LogGroup",
					"RequestToken": "f2fcf5a1-7f17-4c7a-b67f-ab0123456789",
					"Operation": "CREATE"
				}
			}`))
				return
			}
			if reqCount == 2 {
				w.Write([]byte(`
			{
				"ProgressEvent": {
					"Identifier": "CloudApiLogGroup",
					"OperationStatus": "SUCCESS",
					"EventTime": 1707429589,
					"TypeName": "AWS::Logs::LogGroup",
					"RequestToken": "f2fcf5a1-7f17-4c7a-b67f-ab0123456789",
					"Operation": "CREATE"
				}
			}`))
				return
			}
			w.Write([]byte(`
			{
				"ResourceDescription": {
					"Identifier": "CloudApiLogGroup",
					"Properties": "{\"LogGroupName\": \"lg\"}"
				},
				"TypeName": "AWS::Logs::LogGroup"
			}`))
		}))
		t.Cleanup(server.Close)

		provider, err := testProviderServer()
		require.NoError(t, err)
		ctx := context.Background()
		_, err = provider.Configure(ctx, &pulumirpc.ConfigureRequest{
			Variables: map[string]string{
				"aws-native:config:region":                    "us-west-2",
				"aws-native:config:endpoints":                 `{"cloudcontrol": "` + server.URL + `"}`,
				"aws-native:config:skipCredentialsValidation": "true", // Skip so we don't have to mock STS.
				// Set fake credentials to avoid making real requests and override any local credentials.
				"aws-native:config:accessKey": "fake",
				"aws-native:config:secretKey": "fake",
				"aws-native:config:token":     "fake",
			},
		})
		require.NoError(t, err)

		createResponse, err := provider.Create(ctx, &pulumirpc.CreateRequest{
			Urn: "urn:pulumi:stack::project::aws-native:logs:LogGroup::lg",
			Properties: &structpb.Struct{
				Fields: map[string]*structpb.Value{},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, "CloudApiLogGroup", createResponse.Id)
		assert.Subset(t, createResponse.Properties.AsMap(), map[string]interface{}{
			"logGroupName": "lg",
		})
	})
}
