module github.com/pulumi/pulumi-cloudformation/examples

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.5
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v2 v2.6.1
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/pulumi/pulumi/sdk/v2 => github.com/pulumi/pulumi/sdk/v2 v2.6.1
)
