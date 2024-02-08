package provider_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
}
