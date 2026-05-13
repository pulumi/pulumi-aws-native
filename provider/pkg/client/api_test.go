// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudControlApiClientListResourcesSerializesRequest(t *testing.T) {
	type requestBody struct {
		TypeName      string `json:"TypeName"`
		ResourceModel string `json:"ResourceModel"`
		NextToken     string `json:"NextToken"`
		MaxResults    int32  `json:"MaxResults"`
		RoleArn       string `json:"RoleArn"`
	}

	var actual requestBody
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header.Get("X-Amz-Target"), "ListResources")
		require.NoError(t, json.NewDecoder(r.Body).Decode(&actual))

		w.Header().Set("Content-Type", "application/x-amz-json-1.1")
		_, err := w.Write([]byte(`{
			"ResourceDescriptions": [{"Identifier": "id-1"}],
			"NextToken": "more"
		}`))
		require.NoError(t, err)
	}))
	t.Cleanup(server.Close)

	cctl := cloudcontrol.New(cloudcontrol.Options{
		Region:       "us-west-2",
		BaseEndpoint: aws.String(server.URL),
		Credentials:  credentials.NewStaticCredentialsProvider("access-key", "secret-key", "session-token"),
		Retryer:      aws.NopRetryer{},
	})
	roleArn := "arn:aws:iam::123456789012:role/cloudcontrol"
	client := NewCloudControlApiClient(cctl, &roleArn)
	resourceModel := `{"DomainName":"example"}`
	nextToken := "next"
	maxResults := int32(25)

	descriptions, continuation, err := client.ListResources(
		context.Background(),
		"AWS::CustomerProfiles::DomainObjectType",
		&resourceModel,
		&nextToken,
		&maxResults,
	)

	require.NoError(t, err)
	require.Len(t, descriptions, 1)
	assert.Equal(t, "id-1", *descriptions[0].Identifier)
	require.NotNil(t, continuation)
	assert.Equal(t, "more", *continuation)
	assert.Equal(t, requestBody{
		TypeName:      "AWS::CustomerProfiles::DomainObjectType",
		ResourceModel: resourceModel,
		NextToken:     nextToken,
		MaxResults:    maxResults,
		RoleArn:       roleArn,
	}, actual)
}
