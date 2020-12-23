module github.com/pulumi/pulumi-cloudformation/provider

go 1.14

require (
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/aws/aws-sdk-go v1.33.5
	github.com/goccy/go-yaml v1.8.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/hcl/v2 v2.3.0
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v2 v2.16.1
	github.com/pulumi/pulumi/sdk/v2 v2.16.1
	github.com/sanathkr/yaml v1.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/zclconf/go-cty v1.3.1
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
