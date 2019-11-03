// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/types"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

type PropertySpec struct {
	Documentation     string `json:"Documentation"`
	DuplicatesAllowed bool   `json:"DuplicatesAllowed"`
	ItemType          string `json:"ItemType,omitempty"`
	PrimitiveItemType string `json:"PrimitiveItemType,omitempty"`
	PrimitiveType     string `json:"PrimitiveType,omitempty"`
	Required          bool   `json:"Required"`
	Type              string `json:"Type,omitempty"`
	UpdateType        string `json:"UpdateType"`
}

type PropertyTypeSpec struct {
	Documentation string                  `json:"Documentation"`
	Properties    map[string]PropertySpec `json:"Properties"`
}

type AttributeSpec struct {
	ItemType          string `json:"ItemType,omitempty"`
	PrimitiveItemType string `json:"PrimitiveItemType,omitempty"`
	PrimitiveType     string `json:"PrimitiveType,omitempty"`
	Type              string `json:"Type,omitempty"`
}

type ResourceSpec struct {
	Attributes map[string]AttributeSpec `json:"Attributes"`
	PropertyTypeSpec
}

type CloudFormationSchema struct {
	PropertyTypes                map[string]PropertyTypeSpec `json:"PropertyTypes"`
	ResourceSpecificationVersion string                      `json:"ResourceSpecificationVersion"`
	ResourceTypes                map[string]ResourceSpec     `json:"ResourceTypes"`
}

var serviceSDKPackageNames = map[string]string{
	"AmazonMQ":               "mq",
	"Amplify":                "amplify",
	"ApiGateway":             "apigateway",
	"ApiGatewayV2":           "apigatewayv2",
	"AppMesh":                "appmesh",
	"AppStream":              "appstream",
	"AppSync":                "appsync",
	"ApplicationAutoScaling": "applicationautoscaling",
	"Athena":                 "athena",
	"AutoScaling":            "autoscaling",
	"AutoScalingPlans":       "autoscalingplans",
	"Backup":                 "backup",
	"Batch":                  "batch",
	"Budgets":                "budgets",
	"CertificateManager":     "acm",
	"Cloud9":                 "cloud9",
	"CloudFormation":         "cloudformation",
	"CloudFront":             "cloudfront",
	"CloudTrail":             "cloudtrail",
	"CloudWatch":             "cloudwatch",
	"CodeBuild":              "codebuild",
	"CodeCommit":             "codecommit",
	"CodeDeploy":             "codedeploy",
	"CodePipeline":           "codepipeline",
	"CodeStar":               "codestar",
	"Cognito":                "cognitoidentity",
	"Config":                 "configservice",
	"DAX":                    "dax",
	"DLM":                    "dlm",
	"DMS":                    "databasemigrationservice",
	"DataPipeline":           "datapipeline",
	"DirectoryService":       "directoryservice",
	"DocDB":                  "docdb",
	"DynamoDB":               "dynamodb",
	"EC2":                    "ec2",
	"ECR":                    "ecr",
	"ECS":                    "ecs",
	"EFS":                    "efs",
	"EKS":                    "eks",
	"EMR":                    "emr",
	"ElastiCache":            "elasticache",
	"ElasticBeanstalk":       "elasticbeanstalk",
	"ElasticLoadBalancing":   "elb",
	"ElasticLoadBalancingV2": "elbv2",
	"Elasticsearch":          "elasticsearchservice",
	"Events":                 "eventbridge",
	"FSx":                    "fsx",
	"GameLift":               "gamelift",
	"Glue":                   "glue",
	"Greengrass":             "greengrass",
	"GuardDuty":              "guardduty",
	"IAM":                    "iam",
	"Inspector":              "inspector",
	"IoT":                    "iot",
	"IoT1Click":              "iot1clickprojects",
	"IoTAnalytics":           "iotanalytics",
	"IoTEvents":              "iotevents",
	"IoTThingsGraph":         "iotthingsgraph",
	"KMS":                    "kms",
	"Kinesis":                "kinesis",
	"KinesisAnalytics":       "kinesisanalytics",
	"KinesisAnalyticsV2":     "kinesisanalyticsv2",
	"KinesisFirehose":        "firehose",
	"LakeFormation":          "lakeformation",
	"Lambda":                 "lambda",
	"Logs":                   "cloudwatchlogs",
	"MSK":                    "kafka",
	"MediaLive":              "medialive",
	"MediaStore":             "mediastore",
	"Neptune":                "neptune",
	"OpsWorks":               "opsworks",
	"OpsWorksCM":             "opsworkscm",
	"Pinpoint":               "pinpoint",
	"PinpointEmail":          "pinpointemail",
	"QLDB":                   "qldb",
	"RAM":                    "ram",
	"RDS":                    "rds",
	"Redshift":               "redshift",
	"RoboMaker":              "robomaker",
	"Route53":                "route53",
	"Route53Resolver":        "route53resolver",
	"S3":                     "s3",
	"SDB":                    "simpledb",
	"SES":                    "ses",
	"SNS":                    "sns",
	"SQS":                    "sqs",
	"SSM":                    "ssm",
	"SageMaker":              "sagemaker",
	"SecretsManager":         "secretsmanager",
	"SecurityHub":            "securityhub",
	"ServiceCatalog":         "servicecatalog",
	"ServiceDiscovery":       "servicediscovery",
	"StepFunctions":          "sfn",
	"Transfer":               "transfer",
	"WAF":                    "waf",
	"WAFRegional":            "wafregional",
	"WorkSpaces":             "workspaces",
}

func parseCloudFormationResourceType(resourceType string) (string, string, error) {
	typeComponents := strings.Split(resourceType, "::")
	if len(typeComponents) != 3 {
		return "", "", errors.Errorf("expected a type of the form <Namespace>::<Service>::<Type>", resourceType)
	}
	serviceName, typeName := typeComponents[1], typeComponents[2]

	return serviceName, typeName, nil
}

func discoverServiceClient(serviceName string) (types.Type, error) {
	serviceSDKPackageName, ok := serviceSDKPackageNames[serviceName]
	if !ok {
		return nil, errors.Errorf("could not find SDK package")
	}

	serviceSDKPackageImportPath := fmt.Sprintf("github.com/aws/aws-sdk-go/service/%s", serviceSDKPackageName)

	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadTypes}, serviceSDKPackageImportPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, errors.Errorf("found multiple SDK packages")
	}
	sdkPackage := pkgs[0].Types

	// Find the client type.
	newFunc, ok := sdkPackage.Scope().Lookup("New").(*types.Func)
	if !ok {
		return nil, errors.New("could not find package-level New function")
	}
	newFuncResults := newFunc.Type().(*types.Signature).Results()
	if newFuncResults.Len() != 1 {
		return nil, errors.New("unexpected signature for package-level New function")
	}

	clientType, ok := newFuncResults.At(0).Type().(*types.Pointer)
	if !ok {
		return nil, errors.New("unexpected type for service client")
	}

	return clientType, nil
}

func discoverServiceClients(schema CloudFormationSchema) (map[string]types.Type, error) {
	resourceTypes := make([]string, 0, len(schema.ResourceTypes))
	for resourceType := range schema.ResourceTypes {
		resourceTypes = append(resourceTypes, resourceType)
	}
	sort.Strings(resourceTypes)

	serviceClients := make(map[string]types.Type)
	for _, resourceType := range resourceTypes {
		if len(schema.ResourceTypes[resourceType].Attributes) == 0 {
			continue
		}

		serviceName, _, err := parseCloudFormationResourceType(resourceType)
		if err != nil {
			return nil, errors.Wrapf(err, "could not parse resource type %v", resourceType)
		}
		if _, ok := serviceClients[serviceName]; !ok {
			clientType, err := discoverServiceClient(serviceName)
			if err != nil {
				return nil, errors.Wrapf(err, "could not discover SDK client for %v", serviceName)
			}
			serviceClients[serviceName] = clientType
			log.Printf("%v ok", serviceName)
		}
	}
	return serviceClients, nil
}

func discoverGetAPI(clientType types.Type, typeName string) (string, string, bool) {
	methodMatcher, err := regexp.Compile(fmt.Sprintf("(?i)^(Get|Describe)%s(s)?WithContext$", typeName))
	if err != nil {
		return "", "", false
	}

	// Look for a "GetWithContext" or "DescribeWithContext" method that matches the type name.
	clientMethodSet := types.NewMethodSet(clientType)
	for i := 0; i < clientMethodSet.Len(); i++ {
		fn := clientMethodSet.At(i).Obj().(*types.Func)
		if !methodMatcher.MatchString(fn.Name()) {
			continue
		}

		sig := fn.Type().(*types.Signature)

		params := sig.Params()
		if params.Len() < 2 {
			continue
		}

		results := sig.Results()
		if results.Len() != 2 {
			continue
		}

		inputTypePtr, ok := params.At(1).Type().(*types.Pointer)
		if !ok {
			continue
		}
		inputType, ok := inputTypePtr.Elem().(*types.Named)
		if !ok {
			continue
		}

		return fn.Name(), inputType.Obj().Name(), true
	}

	return "", "", false
}

var resourceGetterTemplate = template.Must(template.New("resourceGetter").Parse(`package get

import (
	"context"

	{{if not .HasGetFunc}}_ {{end}}"github.com/aws/aws-sdk-go/service/{{.SDKPackage}}"
)

func (g *Getter) get{{.ResourceName}}Attributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	{{- if .HasGetFunc}}
	resp, err := g.{{.SDKPackage}}.{{.GetFunc}}(ctx, &{{.SDKPackage}}.{{.GetInputType}}{
		// Add params here
	})
	_ = resp
	{{- else}}
	// resp, err := g.{{.SDKPackage}}.
	var err error
	{{- end}}
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		{{- range $name, $spec := .Attributes}}
		"{{$name}}": nil,
		{{- end}}
	}
	return attrs, nil
}
`))

func generateResourceGetter(outdir string, serviceClients map[string]types.Type, resourceType string, spec ResourceSpec) (string, error) {
	// If this type has no attributes, there's nothing we need to do.
	if len(spec.Attributes) == 0 {
		return "", nil
	}

	serviceName, typeName, err := parseCloudFormationResourceType(resourceType)
	if err != nil {
		return "", errors.Wrapf(err, "could not parse resource type %v", resourceType)
	}

	serviceSDKPackageName, ok := serviceSDKPackageNames[serviceName]
	if !ok {
		return "", errors.Errorf("could not find SDK package for service %v", serviceName)
	}

	getFunc, getInputType, hasGetFunc := discoverGetAPI(serviceClients[serviceName], typeName)

	resourceName := serviceName + typeName

	fileName := fmt.Sprintf("%s_%s.go", strings.ToLower(serviceName), strings.ToLower(typeName))
	filePath := path.Join(outdir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "could not create output file %v", filePath)
	}
	defer file.Close()

	getter := fmt.Sprintf("get%sAttributes", resourceName)
	return getter, resourceGetterTemplate.Execute(file, map[string]interface{}{
		"SDKPackage":   serviceSDKPackageName,
		"GetFunc":      getFunc,
		"GetInputType": getInputType,
		"HasGetFunc":   hasGetFunc,
		"ResourceName": resourceName,
		"Attributes":   spec.Attributes,
	})
}

var getterTemplate = template.Must(template.New("getter").Parse(`package get

import (
	"context"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws/client"
	{{- range $name, $client := .ServiceClients}}
	"github.com/aws/aws-sdk-go/service/{{$name}}"
	{{- end}}
)

type NotFoundError struct {
	cause error
}

func NewNotFoundError(cause error) error {
	return &NotFoundError{cause}
}

func (e *NotFoundError) Error() string {
	if e.cause == nil {
		return "Resource not found"
	}
	return e.cause.Error()
}

func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

type Getter struct {
	{{- range $name, $client := .ServiceClients}}
	{{$name}} {{$client}}
	{{- end}}
}

func NewGetter(configProvider client.ConfigProvider) *Getter {
	return &Getter{
		{{- range $name, $client := .ServiceClients}}
		{{$name}}: {{$name}}.New(configProvider),
		{{- end}}
	}
}

func (g *Getter) GetResourceAttributes(ctx context.Context, resourceType, physicalResourceID string) (map[string]interface{}, error) {
	switch resourceType {
	{{- range $type, $getter := .Getters}}
	case "{{$type}}":
		return {{if $getter}}g.{{$getter}}(ctx, physicalResourceID){{else}}nil, nil{{end}}
	{{- end}}
	default:
		return nil, errors.Errorf("unknown resource type '%v'", resourceType)
	}
}
`))

func generateGetter(outdir string, serviceClients map[string]types.Type, getters map[string]string) error {
	namedServiceClients := make(map[string]string)
	for serviceName, clientType := range serviceClients {
		serviceSDKPackageName, ok := serviceSDKPackageNames[serviceName]
		if !ok {
			return errors.Errorf("could not find SDK package for service %v", serviceName)
		}

		namedServiceClients[serviceSDKPackageName] = types.TypeString(clientType, (*types.Package).Name)
	}

	filePath := path.Join(outdir, "getter.go")
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "could not create output file %v", filePath)
	}
	defer file.Close()

	return getterTemplate.Execute(file, map[string]interface{}{
		"ServiceClients": namedServiceClients,
		"Getters":        getters,
	})
}

func main() {
	outdir := flag.String("outdir", "", "the output directory for generated sources")
	flag.Parse()

	var schema CloudFormationSchema
	if err := json.NewDecoder(os.Stdin).Decode(&schema); err != nil {
		log.Fatalf("failed to decode CloudFormation schema: %v", err)
	}

	log.Printf("Discovering service clients...")
	serviceClients, err := discoverServiceClients(schema)
	if err != nil {
		log.Fatalf("failed to discover service clients: %v", err)
	}

	log.Printf("Generating resource getters...")
	resourceTypes := make([]string, 0, len(schema.ResourceTypes))
	for resourceType := range schema.ResourceTypes {
		resourceTypes = append(resourceTypes, resourceType)
	}
	sort.Strings(resourceTypes)

	getters := make(map[string]string)
	for _, resourceType := range resourceTypes {
		getter, err := generateResourceGetter(*outdir, serviceClients, resourceType, schema.ResourceTypes[resourceType])
		if err != nil {
			log.Fatalf("failed to generate getter for %v: %v", resourceType, err)
		}
		getters[resourceType] = getter
	}

	log.Printf("Generating getter...")
	if err := generateGetter(*outdir, serviceClients, getters); err != nil {
		log.Fatalf("failed to generate root getter: %v", err)
	}
}
