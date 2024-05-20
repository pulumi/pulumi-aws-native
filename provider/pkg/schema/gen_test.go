package schema

import (
	"encoding/json"
	"testing"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func marshalSpec(spec map[string]interface{}) *jsschema.Schema {
	var s *jsschema.Schema
	jsonSpec, _ := json.Marshal(spec)
	json.Unmarshal(jsonSpec, &s)
	return s
}

func runTest(t *testing.T, spec map[string]interface{}, docs Docs) *schema.PackageSpec {
	s := marshalSpec(spec)
	semanticsDocument, err := GatherSemantics("../../../provider/cmd/pulumi-resource-aws-native")
	if err != nil {
		t.Fatalf("Error gathering semantics: %v", err)
	}
	packageSpec, _, _, err := GatherPackage([]string{s.Extras["typeName"].(string)}, []*jsschema.Schema{s}, false, &semanticsDocument, &docs)
	if err != nil {
		t.Fatalf("GatherPackage failed: %v", err)
	}
	return packageSpec
}

func TestGatherPackage_docs_inputProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{},
		"properties": map[string]interface{}{
			"Description": map[string]interface{}{
				"description": "",
				"type":        "string",
			},
		},
		"typeName": "AWS::Events::Rule",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Events::Rule": {
				Properties: map[string]string{
					"Description": "The description of the rule.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		packageSpec.Resources["aws-native:events:Rule"].Properties["description"].Description,
		"The description of the rule.",
	)
}

func TestGatherPackage_docs_refProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"Target": map[string]interface{}{
				"additionalProperties": false,
				"properties": map[string]interface{}{
					"Id": map[string]string{
						"type": "string",
					},
				},
				"required": []string{"Id", "Arn"},
				"type":     "object",
			},
		},
		"properties": map[string]interface{}{
			"Targets": map[string]interface{}{
				"insertionOrder": false,
				"items": map[string]interface{}{
					"$ref": "#/definitions/Target",
				},
				"type":        "array",
				"uniqueItems": true,
			},
		},
		"typeName": "AWS::Events::Rule",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Events::Rule.Target": {
				Description: "", // want to make sure we don't pick this one
				Properties: map[string]string{
					"Id": "The ID of the target within the specified rule. Use this ID to reference the target when updating the rule. We recommend using a memorable and unique string.",
				},
			},
			"AWS::Events::Rule": {
				Properties: map[string]string{
					"Targets": "Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.\n\nTargets are the resources that are invoked when a rule is triggered.\n\nThe maximum number of entries per request is 10.\n\n> Each rule can have up to five (5) targets associated with it at one time. \n\nFor a list of services you can configure as targets for events, see [EventBridge targets](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html) in the *Amazon EventBridge User Guide* .\n\nCreating rules with built-in targets is supported only in the AWS Management Console . The built-in targets are:\n\n- `Amazon EBS CreateSnapshot API call`\n- `Amazon EC2 RebootInstances API call`\n- `Amazon EC2 StopInstances API call`\n- `Amazon EC2 TerminateInstances API call`\n\nFor some target types, `PutTargets` provides target-specific parameters. If the target is a Kinesis data stream, you can optionally specify which shard the event goes to by using the `KinesisParameters` argument. To invoke a command on multiple EC2 instances with one rule, you can use the `RunCommandParameters` field.\n\nTo be able to make API calls against the resources that you own, Amazon EventBridge needs the appropriate permissions:\n\n- For AWS Lambda and Amazon SNS resources, EventBridge relies on resource-based policies.\n- For EC2 instances, Kinesis Data Streams, AWS Step Functions state machines and API Gateway APIs, EventBridge relies on IAM roles that you specify in the `RoleARN` argument in `PutTargets` .\n\nFor more information, see [Authentication and Access Control](https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html) in the *Amazon EventBridge User Guide* .\n\nIf another AWS account is in the same region and has granted you permission (using `PutPermission` ), you can send events to that account. Set that account's event bus as a target of the rules in your account. To send the matched events to the other account, specify that account's event bus as the `Arn` value when you run `PutTargets` . If your account sends events to another account, your account is charged for each sent event. Each event sent to another account is charged as a custom event. The account receiving the event is not charged. For more information, see [Amazon EventBridge Pricing](https://docs.aws.amazon.com/eventbridge/pricing/) .\n\n> `Input` , `InputPath` , and `InputTransformer` are not available with `PutTarget` if the target is an event bus of a different AWS account. \n\nIf you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .\n\n> If you have an IAM role on a cross-account event bus target, a `PutTargets` call without a role on the same target (same `Id` and `Arn` ) will not remove the role. \n\nFor more information about enabling cross-account events, see [PutPermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html) .\n\n*Input* , *InputPath* , and *InputTransformer* are mutually exclusive and optional parameters of a target. When a rule is triggered due to a matched event:\n\n- If none of the following arguments are specified for a target, then the entire event is passed to the target in JSON format (unless the target is Amazon EC2 Run Command or Amazon ECS task, in which case nothing from the event is passed to the target).\n- If *Input* is specified in the form of valid JSON, then the matched event is overridden with this constant.\n- If *InputPath* is specified in the form of JSONPath (for example, `$.detail` ), then only the part of the event specified in the path is passed to the target (for example, only the detail part of the event is passed).\n- If *InputTransformer* is specified, then one or more specified JSONPaths are extracted from the event and used as values in a template that you specify as the input to the target.\n\nWhen you specify `InputPath` or `InputTransformer` , you must use JSON dot notation, not bracket notation.\n\nWhen you add targets to a rule and the associated rule triggers soon after, new or updated targets might not be immediately invoked. Allow a short period of time for changes to take effect.\n\nThis action can partially fail if too many requests are made at the same time. If that happens, `FailedEntryCount` is non-zero in the response and each entry in `FailedEntries` provides the ID of the failed target and the error code.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		"The ID of the target within the specified rule. Use this ID to reference the target when updating the rule. We recommend using a memorable and unique string.",
		packageSpec.Types["aws-native:events:RuleTarget"].Properties["id"].Description,
	)

	assert.Equal(
		t,
		"Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.\n\nTargets are the resources that are invoked when a rule is triggered.\n\nThe maximum number of entries per request is 10.\n\n> Each rule can have up to five (5) targets associated with it at one time. \n\nFor a list of services you can configure as targets for events, see [EventBridge targets](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html) in the *Amazon EventBridge User Guide* .\n\nCreating rules with built-in targets is supported only in the AWS Management Console . The built-in targets are:\n\n- `Amazon EBS CreateSnapshot API call`\n- `Amazon EC2 RebootInstances API call`\n- `Amazon EC2 StopInstances API call`\n- `Amazon EC2 TerminateInstances API call`\n\nFor some target types, `PutTargets` provides target-specific parameters. If the target is a Kinesis data stream, you can optionally specify which shard the event goes to by using the `KinesisParameters` argument. To invoke a command on multiple EC2 instances with one rule, you can use the `RunCommandParameters` field.\n\nTo be able to make API calls against the resources that you own, Amazon EventBridge needs the appropriate permissions:\n\n- For AWS Lambda and Amazon SNS resources, EventBridge relies on resource-based policies.\n- For EC2 instances, Kinesis Data Streams, AWS Step Functions state machines and API Gateway APIs, EventBridge relies on IAM roles that you specify in the `RoleARN` argument in `PutTargets` .\n\nFor more information, see [Authentication and Access Control](https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html) in the *Amazon EventBridge User Guide* .\n\nIf another AWS account is in the same region and has granted you permission (using `PutPermission` ), you can send events to that account. Set that account's event bus as a target of the rules in your account. To send the matched events to the other account, specify that account's event bus as the `Arn` value when you run `PutTargets` . If your account sends events to another account, your account is charged for each sent event. Each event sent to another account is charged as a custom event. The account receiving the event is not charged. For more information, see [Amazon EventBridge Pricing](https://docs.aws.amazon.com/eventbridge/pricing/) .\n\n> `Input` , `InputPath` , and `InputTransformer` are not available with `PutTarget` if the target is an event bus of a different AWS account. \n\nIf you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .\n\n> If you have an IAM role on a cross-account event bus target, a `PutTargets` call without a role on the same target (same `Id` and `Arn` ) will not remove the role. \n\nFor more information about enabling cross-account events, see [PutPermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html) .\n\n*Input* , *InputPath* , and *InputTransformer* are mutually exclusive and optional parameters of a target. When a rule is triggered due to a matched event:\n\n- If none of the following arguments are specified for a target, then the entire event is passed to the target in JSON format (unless the target is Amazon EC2 Run Command or Amazon ECS task, in which case nothing from the event is passed to the target).\n- If *Input* is specified in the form of valid JSON, then the matched event is overridden with this constant.\n- If *InputPath* is specified in the form of JSONPath (for example, `$.detail` ), then only the part of the event specified in the path is passed to the target (for example, only the detail part of the event is passed).\n- If *InputTransformer* is specified, then one or more specified JSONPaths are extracted from the event and used as values in a template that you specify as the input to the target.\n\nWhen you specify `InputPath` or `InputTransformer` , you must use JSON dot notation, not bracket notation.\n\nWhen you add targets to a rule and the associated rule triggers soon after, new or updated targets might not be immediately invoked. Allow a short period of time for changes to take effect.\n\nThis action can partially fail if too many requests are made at the same time. If that happens, `FailedEntryCount` is non-zero in the response and each entry in `FailedEntries` provides the ID of the failed target and the error code.",
		packageSpec.Resources["aws-native:events:Rule"].InputProperties["targets"].Description,
	)

}

func TestGatherPackage_docs_nestedProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"PipeSourceKinesisStreamParameters": map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
			"PipeSourceParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"KinesisStreamParameters": map[string]interface{}{
						"$ref": "#/definitions/PipeSourceKinesisStreamParameters",
					},
				},
			},
		},
		"properties": map[string]interface{}{
			"SourceParameters": map[string]interface{}{
				"$ref": "#/definitions/PipeSourceParameters",
			},
		},
		"typeName":    "AWS::Pipes::Pipe",
		"description": "Definition of AWS::Pipes::Pipe Resource Type",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Pipes::Pipe.PipeSourceParameters": {
				Properties: map[string]string{
					"KinesisStreamParameters": "The parameters for using a Kinesis stream as a source.",
				},
			},
			"AWS::Pipes::Pipe.PipeSourceKinesisStreamParameters": {
				Description: "",
				Properties: map[string]string{
					"MaximumBatchingWindowInSeconds": "The maximum length of a time to wait for events.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		"The parameters for using a Kinesis stream as a source.",
		packageSpec.Types["aws-native:pipes:PipeSourceParameters"].Properties["kinesisStreamParameters"].Description,
	)

}
func TestGatherPackage_docs_nestedPropertyProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"PipeSourceKinesisStreamParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"MaximumBatchingWindowInSeconds": map[string]interface{}{
						"type":    "integer",
						"maximum": 300,
						"minimum": 0,
					},
				},
			},
			"PipeSourceParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"KinesisStreamParameters": map[string]interface{}{
						"$ref": "#/definitions/PipeSourceKinesisStreamParameters",
					},
				},
			},
		},
		"properties": map[string]interface{}{
			"SourceParameters": map[string]interface{}{
				"$ref": "#/definitions/PipeSourceParameters",
			},
		},
		"typeName":    "AWS::Pipes::Pipe",
		"description": "Definition of AWS::Pipes::Pipe Resource Type",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Pipes::Pipe.PipeSourceKinesisStreamParameters": {
				Description: "The parameters for using a Kinesis stream as a source.",
				Properties: map[string]string{
					"MaximumBatchingWindowInSeconds": "The maximum length of a time to wait for events.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		"The maximum length of a time to wait for events.",
		packageSpec.Types["aws-native:pipes:PipeSourceKinesisStreamParameters"].Properties["maximumBatchingWindowInSeconds"].Description,
	)

}

func TestGatherPackage_docs_nestedListProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"EcsEnvironmentFile": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"Value": map[string]interface{}{
						"type": "string",
					},
				},
			},
			"EcsContainerOverride": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"EnvironmentFiles": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"$ref": "#/definitions/EcsEnvironmentFile",
						},
					},
				},
			},
			"EcsTaskOverride": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"ContainerOverrides": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"$ref": "#/definitions/EcsContainerOverride",
						},
					},
				},
			},
			"PipeTargetEcsTaskParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"Overrides": map[string]interface{}{
						"$ref": "#/definitions/EcsTaskOverride",
					},
				},
			},
			"PipeTargetParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"EcsTaskParameters": map[string]interface{}{
						"$ref": "#/definitions/PipeTargetEcsTaskParameters",
					},
				},
			},
		},
		"properties": map[string]interface{}{
			"TargetParameters": map[string]interface{}{
				"$ref": "#/definitions/PipeTargetParameters",
			},
		},
		"typeName":    "AWS::Pipes::Pipe",
		"description": "Definition of AWS::Pipes::Pipe Resource Type",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Pipes::Pipe.EcsContainerOverride": {
				Properties: map[string]string{
					"EnvironmentFiles": "A list of files containing the environment variables to pass to a container. You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file should contain an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .\n\nIf there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .\n\nThis parameter is only supported for tasks hosted on Fargate using the following platform versions:\n\n- Linux platform version `1.4.0` or later.\n- Windows platform version `1.0.0` or later.",
				},
			},
			"AWS::Pipes::Pipe.EcsEnvironmentFile": {
				Description: "A list of files containing the environment variables to pass to a container. You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file should contain an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .\n\nIf there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .\n\nThis parameter is only supported for tasks hosted on Fargate using the following platform versions:\n\n- Linux platform version `1.4.0` or later.\n- Windows platform version `1.0.0` or later.",
				Properties: map[string]string{
					"Type":  "The file type to use. The only supported value is `s3` .",
					"Value": "The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		"A list of files containing the environment variables to pass to a container. You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file should contain an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .\n\nIf there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .\n\nThis parameter is only supported for tasks hosted on Fargate using the following platform versions:\n\n- Linux platform version `1.4.0` or later.\n- Windows platform version `1.0.0` or later.",
		packageSpec.Types["aws-native:pipes:PipeEcsContainerOverride"].Properties["environmentFiles"].Description,
	)

}

func TestGatherPackage_docs_listProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"PipeSourceSelfManagedKafkaParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"AdditionalBootstrapServers": map[string]interface{}{
						"type": "array",
						"items": map[string]interface{}{
							"type": "string",
						},
						"maxItems": 2,
						"minItems": 0,
					},
				},
			},
			"PipeSourceParameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"SelfManagedKafkaParameters": map[string]interface{}{
						"$ref": "#/definitions/PipeSourceSelfManagedKafkaParameters",
					},
				},
			},
		},
		"properties": map[string]interface{}{
			"SourceParameters": map[string]interface{}{
				"$ref": "#/definitions/PipeSourceParameters",
			},
		},
		"typeName":    "AWS::Pipes::Pipe",
		"description": "Definition of AWS::Pipes::Pipe Resource Type",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::Pipes::Pipe.PipeSourceSelfManagedKafkaParameters": {
				Description: "The parameters for using a self-managed Apache Kafka stream as a source.\n\nA *self managed* cluster refers to any Apache Kafka cluster not hosted by AWS . This includes both clusters you manage yourself, as well as those hosted by a third-party provider, such as [Confluent Cloud](https://docs.aws.amazon.com/https://www.confluent.io/) , [CloudKarafka](https://docs.aws.amazon.com/https://www.cloudkarafka.com/) , or [Redpanda](https://docs.aws.amazon.com/https://redpanda.com/) . For more information, see [Apache Kafka streams as a source](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kafka.html) in the *Amazon EventBridge User Guide* .",
				Properties: map[string]string{
					"AdditionalBootstrapServers": "An array of server URLs.",
				},
			},
		},
	}

	packageSpec := runTest(t, spec, docs)
	assert.Equal(
		t,
		"An array of server URLs.",
		packageSpec.Types["aws-native:pipes:PipeSourceSelfManagedKafkaParameters"].Properties["additionalBootstrapServers"].Description,
	)
}
