module github.com/pulumi/pulumi-cloudformation/provider

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.5
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v2 v2.6.1
	github.com/pulumi/pulumi/sdk/v2 v2.6.1
	github.com/sanathkr/yaml v1.0.0
	github.com/satori/go.uuid v1.2.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

replace github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pkg
