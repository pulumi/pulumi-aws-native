module github.com/pulumi/pulumi-aws-native/provider

go 1.15

require (
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/aws/aws-sdk-go v1.37.31
	github.com/goccy/go-yaml v1.8.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/hashicorp/hcl/v2 v2.3.0
	github.com/jpillora/backoff v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.2.1
	github.com/stretchr/testify v1.6.1
	github.com/zclconf/go-cty v1.3.1
)

replace github.com/aws/aws-sdk-go => ../aws-sdk-go-preview
