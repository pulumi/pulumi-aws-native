// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func configToProvider(config pschema.ComplexTypeSpec) pschema.ComplexTypeSpec {
	// The Provider schema is the Config schema with an additional Language block for each Property.
	for k, v := range config.Properties {
		v.Language = map[string]pschema.RawMessage{
			"python": rawMessage(map[string]interface{}{
				"mapCase": false,
			}),
		}
		config.Properties[k] = v
	}
	return config
}

var assumeRole = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration for a Provider to assume a role.",
		Properties: map[string]pschema.PropertySpec{
			"durationSeconds": {
				Description: "Number of seconds to restrict the assume role session duration.",
				TypeSpec:    pschema.TypeSpec{Type: "integer"},
			},
			"externalId": {
				Description: "External identifier to use when assuming the role.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"policy": {
				Description: "IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"policyArns": {
				Description: "Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the role.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
			"roleArn": {
				Description: "Amazon Resource Name (ARN) of the IAM Role to assume.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sessionName": {
				Description: "Session name to use when assuming the role.",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"tags": {
				Description: "Map of assume role session tags.",
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"}},
			},
			"transitiveTagKeys": {
				Description: "A list of keys for session tags that you want to set as transitive. If you set a tag key as transitive, the corresponding key and value passes to subsequent sessions in a role chain.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var defaultTags = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration with resource tag settings to apply across all resources handled by this provider. This is designed to replace redundant per-resource `tags` configurations. Provider tags can be overridden with new values, but not excluded from specific resources. To override provider tag values, use the `tags` argument within a resource to configure new tag values for matching keys.",
		Properties: map[string]pschema.PropertySpec{
			"tags": {
				Description: "A group of tags to set across all resources.",
				TypeSpec: pschema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var endpoints = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration for for customizing service endpoints.",
		Properties: map[string]pschema.PropertySpec{
			"accessanalyzer": {
				Description: "Override the default endpoint for AWS Access Analyzer",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"acm": {
				Description: "Override the default endpoint for AWS Certificate Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"acmpca": {
				Description: "Override the default endpoint for AWS Certificate Manager Private Certificate Authority",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"amplify": {
				Description: "Override the default endpoint for AWS Amplify Console",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"appconfig": {
				Description: "Override the default endpoint for AWS AppConfig",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"applicationautoscaling": {
				Description: "Override the default endpoint for AWS Application Auto Scaling",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"applicationinsights": {
				Description: "Override the default endpoint for AWS CloudWatch Application Insights",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"appmesh": {
				Description: "Override the default endpoint for AWS App Mesh",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"apprunner": {
				Description: "Override the default endpoint for AWS App Runner",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"appstream": {
				Description: "Override the default endpoint for AWS AppStream 2.0",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"appsync": {
				Description: "Override the default endpoint for AWS AppSync",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"athena": {
				Description: "Override the default endpoint for AWS Athena",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"auditmanager": {
				Description: "Override the default endpoint for AWS Audit Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"autoscaling": {
				Description: "Override the default endpoint for AWS Auto Scaling",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"backup": {
				Description: "Override the default endpoint for AWS Backup",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"batch": {
				Description: "Override the default endpoint for AWS Batch",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"budgets": {
				Description: "Override the default endpoint for AWS Budgets",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"chime": {
				Description: "Override the default endpoint for Amazon Chime",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloud9": {
				Description: "Override the default endpoint for AWS Cloud9",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudformation": {
				Description: "Override the default endpoint for AWS CloudFormation",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudfront": {
				Description: "Override the default endpoint for AWS CloudFront",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudhsm": {
				Description: "Override the default endpoint for AWS CloudHSM",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudsearch": {
				Description: "Override the default endpoint for AWS CloudSearch",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudtrail": {
				Description: "Override the default endpoint for AWS CloudTrail",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudwatch": {
				Description: "Override the default endpoint for AWS CloudWatch",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudwatchevents": {
				Description: "Override the default endpoint for AWS CloudWatch Events",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cloudwatchlogs": {
				Description: "Override the default endpoint for AWS CloudWatch Logs",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codeartifact": {
				Description: "Override the default endpoint for AWS CodeArtifact",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codebuild": {
				Description: "Override the default endpoint for AWS CodeBuild",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codecommit": {
				Description: "Override the default endpoint for AWS CodeCommit",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codedeploy": {
				Description: "Override the default endpoint for AWS CodeDeploy",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codepipeline": {
				Description: "Override the default endpoint for AWS CodePipeline",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"codestarconnections": {
				Description: "Override the default endpoint for AWS CodeStart Connections",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cognitoidentity": {
				Description: "Override the default endpoint for Amazon Cognito",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"configservice": {
				Description: "Override the default endpoint for AWS Config",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"connect": {
				Description: "Override the default endpoint for Amazon Connect",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"cur": {
				Description: "Override the default endpoint for AWS Cost and Usage Reports",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"dataexchange": {
				Description: "Override the default endpoint for AWS Data Exchange",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"datapipeline": {
				Description: "Override the default endpoint for AWS Data Pipeline",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"datasync": {
				Description: "Override the default endpoint for AWS DataSync",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"dax": {
				Description: "Override the default endpoint for AWS DynamoDB Accelerator",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"detective": {
				Description: "Override the default endpoint for AWS Detective",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"devicefarm": {
				Description: "Override the default endpoint for AWS Device Farm",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"directconnect": {
				Description: "Override the default endpoint for AWS Direct Connect",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"dlm": {
				Description: "Override the default endpoint for AWS Data Lifecycle Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"dms": {
				Description: "Override the default endpoint for AWS Database Migration Service",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"docdb": {
				Description: "Override the default endpoint for AWS DocumentDB",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ds": {
				Description: "Override the default endpoint for AWS Directory Service",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"dynamodb": {
				Description: "Override the default endpoint for AWS DynamoDB",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ec2": {
				Description: "Override the default endpoint for AWS Elastic Compute Cloud (EC2)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ecr": {
				Description: "Override the default endpoint for AWS Elastic Container Registry (ECR)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ecrpublic": {
				Description: "Override the default endpoint for AWS Elastic Container Registry (ECR) Public",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ecs": {
				Description: "Override the default endpoint for AWS Elastic Container Service (ECS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"efs": {
				Description: "Override the default endpoint for AWS Elastic File System (EFS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"eks": {
				Description: "Override the default endpoint for AWS Elastic Kubernetes Service (EKS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"elasticache": {
				Description: "Override the default endpoint for AWS ElastiCache",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"elasticbeanstalk": {
				Description: "Override the default endpoint for AWS Elastic Beanstalk",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"elastictranscoder": {
				Description: "Override the default endpoint for AWS Elastic Transcoder",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"elb": {
				Description: "Override the default endpoint for AWS Elastic Load Balancing",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"elbv2": {
				Description: "Override the default endpoint for AWS Elastic Load Balancing V2",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"emr": {
				Description: "Override the default endpoint for AWS EMR",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"emrcontainers": {
				Description: "Override the default endpoint for AWS EMR on EKS",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"es": {
				Description: "Override the default endpoint for AWS OpenSearch Service (formerly Elasticsearch)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"firehose": {
				Description: "Override the default endpoint for AWS Kinesis Data Firehose",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"fms": {
				Description: "Override the default endpoint for AWS Firewall Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"forecast": {
				Description: "Override the default endpoint for Amazon Forecast",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"fsx": {
				Description: "Override the default endpoint for AWS FSx",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"gamelift": {
				Description: "Override the default endpoint for AWS GameLift",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"glacier": {
				Description: "Override the default endpoint for Amazon S3 Glacier",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"globalaccelerator": {
				Description: "Override the default endpoint for AWS Global Accelerator",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"glue": {
				Description: "Override the default endpoint for AWS Glue",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"greengrass": {
				Description: "Override the default endpoint for AWS IoT Greengrass",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"guardduty": {
				Description: "Override the default endpoint for AWS GuardDuty",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"iam": {
				Description: "Override the default endpoint for AWS Identity and Access Management",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"identitystore": {
				Description: "Override the default endpoint for AWS Single Sign-On (SSO) Identity Store",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"imagebuilder": {
				Description: "Override the default endpoint for AWS Image Builder",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"inspector": {
				Description: "Override the default endpoint for Amazon Inspector",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"iot": {
				Description: "Override the default endpoint for AWS IoT",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"iotanalytics": {
				Description: "Override the default endpoint for AWS IoT Analytics",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"iotevents": {
				Description: "Override the default endpoint for AWS IoT Events",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kafka": {
				Description: "Override the default endpoint for Amazon Managed Streaming for Apache Kafka (MSK)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kinesis": {
				Description: "Override the default endpoint for Amazon Kinesis",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kinesisanalytics": {
				Description: "Override the default endpoint for Amazon Kinesis Data Analytics",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kinesisanalyticsv2": {
				Description: "Override the default endpoint for Amazon Kinesis Data Analytics V2",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kinesisvideo": {
				Description: "Override the default endpoint for Amazon Kinesis Video Streams",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"kms": {
				Description: "Override the default endpoint for AWS Key Management Service",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"lakeformation": {
				Description: "Override the default endpoint for AWS Lake Formation",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"lambda": {
				Description: "Override the default endpoint for AWS Lambda",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"lexmodels": {
				Description: "Override the default endpoint for Amazon Lex",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"licensemanager": {
				Description: "Override the default endpoint for AWS License Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"lightsail": {
				Description: "Override the default endpoint for Amazon Lightsail",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"location": {
				Description: "Override the default endpoint for Amazon Location",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"macie": {
				Description: "Override the default endpoint for Amazon Macie",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"macie2": {
				Description: "Override the default endpoint for Amazon Macie V2",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"managedblockchain": {
				Description: "Override the default endpoint for Amazon Managed Blockchain",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"marketplacecatalog": {
				Description: "Override the default endpoint for AWS Marketplace Catalog",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mediaconnect": {
				Description: "Override the default endpoint for AWS MediaConnect",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mediaconvert": {
				Description: "Override the default endpoint for AWS MediaConvert",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"medialive": {
				Description: "Override the default endpoint for AWS MediaLive",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mediapackage": {
				Description: "Override the default endpoint for AWS MediaPackage",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mediastore": {
				Description: "Override the default endpoint for AWS Elemental MediaStore container",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mediastoredata": {
				Description: "Override the default endpoint for AWS Elemental MediaStore asset",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"memorydb": {
				Description: "Override the default endpoint for AWS MemoryDB for Redis",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mq": {
				Description: "Override the default endpoint for Amazon MQ",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"mwaa": {
				Description: "Override the default endpoint for Amazon Managed Workflows for Apache Airflow",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"neptune": {
				Description: "Override the default endpoint for Amazon Neptune",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"networkfirewall": {
				Description: "Override the default endpoint for AWS Network Firewall",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"networkmanager": {
				Description: "Override the default endpoint for AWS Network Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"opsworks": {
				Description: "Override the default endpoint for AWS OpsWorks",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"organizations": {
				Description: "Override the default endpoint for AWS Organizations",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"outposts": {
				Description: "Override the default endpoint for AWS Outposts",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"personalize": {
				Description: "Override the default endpoint for Amazon Personalize",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"pinpoint": {
				Description: "Override the default endpoint for Amazon Pinpoint",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"pricing": {
				Description: "Override the default endpoint for Amazon Web Services Price List Service",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"qldb": {
				Description: "Override the default endpoint for Amazon QLDB",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"quicksight": {
				Description: "Override the default endpoint for Amazon QuickSight",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ram": {
				Description: "Override the default endpoint for AWS Resource Access Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"rds": {
				Description: "Override the default endpoint for Amazon Relational Database Service",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"redshift": {
				Description: "Override the default endpoint for Amazon Redshift",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"resourcegroups": {
				Description: "Override the default endpoint for AWS Resource Groups",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"resourcegroupstaggingapi": {
				Description: "Override the default endpoint for AWS Resource Groups Tagging API",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"route53": {
				Description: "Override the default endpoint for Amazon Route 53",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"route53domains": {
				Description: "Override the default endpoint for Amazon Route 53 Domains",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"route53recoverycontrolconfig": {
				Description: "Override the default endpoint for Amazon Route 53 Recovery Control",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"route53recoveryreadiness": {
				Description: "Override the default endpoint for Amazon Route 53 Recovery Readiness",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"route53resolver": {
				Description: "Override the default endpoint for Amazon Route 53 Resolver",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"s3": {
				Description: "Override the default endpoint for Amazon Simple Storage Service (S3)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"s3control": {
				Description: "Override the default endpoint for Amazon Simple Storage Service (S3) Control",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"s3outposts": {
				Description: "Override the default endpoint for Amazon S3 on Outposts",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sagemaker": {
				Description: "Override the default endpoint for AWS SageMaker",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"schemas": {
				Description: "Override the default endpoint for Amazon EventBridge Schema Registry",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sdb": {
				Description: "Override the default endpoint for Amazon SimpleDB",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"secretsmanager": {
				Description: "Override the default endpoint for AWS Secrets Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"securityhub": {
				Description: "Override the default endpoint for AWS Security Hub",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"serverlessrepo": {
				Description: "Override the default endpoint for AWS Serverless Application Repository",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"servicecatalog": {
				Description: "Override the default endpoint for AWS Service Catalog",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"servicediscovery": {
				Description: "Override the default endpoint for AWS Cloud Map",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"servicequotas": {
				Description: "Override the default endpoint for AWS Service Quotas",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ses": {
				Description: "Override the default endpoint for Amazon Simple Email Service (SES)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"shield": {
				Description: "Override the default endpoint for AWS Shield Advanced API",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"signer": {
				Description: "Override the default endpoint for AWS Signer",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sns": {
				Description: "Override the default endpoint for Amazon Simple Notification Service (SNS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sqs": {
				Description: "Override the default endpoint for Amazon Simple Queue Service (SQS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ssm": {
				Description: "Override the default endpoint for AWS Systems Manager",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"ssoadmin": {
				Description: "Override the default endpoint for AWS Single Sign On (SSO)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"stepfunctions": {
				Description: "Override the default endpoint for AWS Step Functions",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"storagegateway": {
				Description: "Override the default endpoint for AWS Storage Gateway",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"sts": {
				Description: "Override the default endpoint for AWS Security Token Service (STS)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"swf": {
				Description: "Override the default endpoint for Amazon Simple Workflow Service (SWF)",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"synthetics": {
				Description: "Override the default endpoint for Amazon CloudWatch Synthetics",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"timestreamwrite": {
				Description: "Override the default endpoint for Amazon Timestream",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"transfer": {
				Description: "Override the default endpoint for AWS Transfer Family",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"waf": {
				Description: "Override the default endpoint for AWS WAF Classic",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"wafregional": {
				Description: "Override the default endpoint for AWS WAF Regional Classic",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"wafv2": {
				Description: "Override the default endpoint for AWS WAF V2",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"worklink": {
				Description: "Override the default endpoint for Amazon WorkLink",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"workmail": {
				Description: "Override the default endpoint for Amazon WorkMail",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"workspaces": {
				Description: "Override the default endpoint for Amazon WorkSpaces",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
			"xray": {
				Description: "Override the default endpoint for AWS X-Ray",
				TypeSpec:    pschema.TypeSpec{Type: "string"},
			},
		},
		Type: "object",
	},
}

var ignoreTags = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "The configuration with resource tag settings to ignore across all resources handled by this provider (except any individual service tag resources such as `ec2.Tag`) for situations where external systems are managing certain resource tags.",
		Properties: map[string]pschema.PropertySpec{
			"keyPrefixes": {
				Description: "List of exact resource tag keys to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning the tag in any `tags` attributes and displaying any configuration difference for the tag value. If any resource configuration still has this tag key configured in the `tags` argument, it will display a perpetual difference until the tag is removed from the argument or `ignoreChanges` is also used.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
			"keys": {
				Description: "List of resource tag key prefixes to ignore across all resources handled by this provider. This configuration prevents Pulumi from returning any tag key matching the prefixes in any `tags` attributes and displaying any configuration difference for those tag values. If any resource configuration still has a tag matching one of the prefixes configured in the `tags` argument, it will display a perpetual difference until the tag is removed from the argument or `ignoreChanges` is also used.",
				TypeSpec:    pschema.TypeSpec{Type: "array", Items: &pschema.TypeSpec{Type: "string"}},
			},
		},
		Type: "object",
	},
}

var region = pschema.ComplexTypeSpec{
	ObjectTypeSpec: pschema.ObjectTypeSpec{
		Description: "A Region represents any valid Amazon region that may be targeted with deployments.",
		Type:        "string",
	},
	Enum: []pschema.EnumValueSpec{
		{Name: "AFSouth1", Value: "af-south-1", Description: "Africa (Cape Town)"},
		{Name: "APEast1", Value: "ap-east-1", Description: "Asia Pacific (Hong Kong)"},
		{Name: "APNortheast1", Value: "ap-northeast-1", Description: "Asia Pacific (Tokyo)"},
		{Name: "APNortheast2", Value: "ap-northeast-2", Description: "Asia Pacific (Seoul)"},
		{Name: "APNortheast3", Value: "ap-northeast-3", Description: "Asia Pacific (Osaka)"},
		{Name: "APSouth1", Value: "ap-south-1", Description: "Asia Pacific (Mumbai)"},
		{Name: "APSoutheast1", Value: "ap-southeast-1", Description: "Asia Pacific (Singapore)"},
		{Name: "APSoutheast2", Value: "ap-southeast-2", Description: "Asia Pacific (Sydney)"},
		{Name: "CACentral", Value: "ca-central-1", Description: "Canada (Central)"},
		{Name: "CNNorth1", Value: "cn-north-1", Description: "China (Beijing)"},
		{Name: "CNNorthwest1", Value: "cn-northwest-1", Description: "China (Ningxia)"},
		{Name: "EUCentral1", Value: "eu-central-1", Description: "Europe (Frankfurt)"},
		{Name: "EUNorth1", Value: "eu-north-1", Description: "Europe (Stockholm)"},
		{Name: "EUWest1", Value: "eu-west-1", Description: "Europe (Ireland)"},
		{Name: "EUWest2", Value: "eu-west-2", Description: "Europe (London)"},
		{Name: "EUWest3", Value: "eu-west-3", Description: "Europe (Paris)"},
		{Name: "EUSouth1", Value: "eu-south-1", Description: "Europe (Milan)"},
		{Name: "MESouth1", Value: "me-south-1", Description: "Middle East (Bahrain)"},
		{Name: "SAEast1", Value: "sa-east-1", Description: "South America (São Paulo)"},
		{Name: "USGovEast1", Value: "us-gov-east-1", Description: "AWS GovCloud (US-East)"},
		{Name: "USGovWest1", Value: "us-gov-west-1", Description: "AWS GovCloud (US-West)"},
		{Name: "USEast1", Value: "us-east-1", Description: "US East (N. Virginia)"},
		{Name: "USEast2", Value: "us-east-2", Description: "US East (Ohio)"},
		{Name: "USWest1", Value: "us-west-1", Description: "US West (N. California)"},
		{Name: "USWest2", Value: "us-west-2", Description: "US West (Oregon)"},
	},
}

// typeOverlays augment the types defined by the schema.
var typeOverlays = map[string]pschema.ComplexTypeSpec{
	"aws-native:config:AssumeRole":         assumeRole,
	"aws-native:index:ProviderAssumeRole":  configToProvider(assumeRole),
	"aws-native:config:DefaultTags":        defaultTags,
	"aws-native:index:ProviderDefaultTags": configToProvider(defaultTags),
	"aws-native:config:Endpoints":          endpoints,
	"aws-native:index:ProviderEndpoint":    configToProvider(endpoints),
	"aws-native:config:IgnoreTags":         ignoreTags,
	"aws-native:index:ProviderIgnoreTags":  configToProvider(ignoreTags),
	"aws-native:index/Region:Region":       region,
}
