module github.com/pulumi/pulumi-aws-native/provider

go 1.15

require (
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/aws/aws-sdk-go-v2 v1.9.1
	github.com/aws/aws-sdk-go-v2/config v1.8.2
	github.com/aws/aws-sdk-go-v2/credentials v1.4.2
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.0.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.10.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.18.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.11.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.1
	github.com/aws/smithy-go v1.8.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/evanphx/json-patch v0.5.2 // indirect
	github.com/goccy/go-yaml v1.8.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/hashicorp/hcl/v2 v2.3.0
	github.com/jpillora/backoff v1.0.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/lestrrat-go/jspointer v0.0.0-20181205001929-82fadba7561c // indirect
	github.com/lestrrat-go/jsref v0.0.0-20181205001954-1b590508f37d // indirect
	github.com/lestrrat-go/jsschema v0.0.0-20181205002244-5c81c58ffcc3
	github.com/lestrrat-go/jsval v0.0.0-20181205002323-20277e9befc0 // indirect
	github.com/lestrrat-go/pdebug v0.0.0-20210111095411-35b07dbf089b // indirect
	github.com/lestrrat-go/structinfo v0.0.0-20210312050401-7f8bd69d6acb // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi/pkg/v3 v3.12.0
	github.com/pulumi/pulumi/sdk/v3 v3.12.0
	github.com/stretchr/testify v1.7.0
	github.com/zclconf/go-cty v1.3.1
	google.golang.org/grpc v1.37.0
)

replace github.com/lestrrat-go/jsschema => github.com/mikhailshilkov/jsschema v0.0.0-20210924145243-fc93fd28ee1b
