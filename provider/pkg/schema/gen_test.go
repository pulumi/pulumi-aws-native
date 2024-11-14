package schema

import (
	"encoding/json"
	"fmt"
	"testing"

	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"
)

func marshalSpec(spec map[string]interface{}) *jsschema.Schema {
	var s *jsschema.Schema
	jsonSpec, _ := json.Marshal(spec)
	json.Unmarshal(jsonSpec, &s)
	return s
}

func runTest(t *testing.T, spec map[string]interface{}, docs Docs) (*schema.PackageSpec, *metadata.CloudAPIMetadata, *Reports) {
	return runTestWithRegions(t, spec, docs, []RegionInfo{
		{
			Name:        "us-west-2",
			Description: "US West (Oregon)",
		},
		{
			Name:        "us-east-1",
			Description: "US East (N. Virginia)",
		},
	})
}

func runTestWithRegions(t *testing.T, spec map[string]interface{}, docs Docs, regions []RegionInfo) (*schema.PackageSpec, *metadata.CloudAPIMetadata, *Reports) {
	s := marshalSpec(spec)
	semanticsDocument, err := GatherSemantics("../../../provider/cmd/pulumi-resource-aws-native")
	if err != nil {
		t.Fatalf("Error gathering semantics: %v", err)
	}
	packageSpec, metadata, reports, err := GatherPackage([]string{s.Extras["typeName"].(string)}, []*jsschema.Schema{s}, false, &semanticsDocument, &docs, regions)
	if err != nil {
		t.Fatalf("GatherPackage failed: %v", err)
	}
	return packageSpec, metadata, reports
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

	packageSpec, _, _ := runTest(t, spec, docs)
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

	packageSpec, _, _ := runTest(t, spec, docs)
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

	packageSpec, _, _ := runTest(t, spec, docs)
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

	packageSpec, _, _ := runTest(t, spec, docs)
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

	packageSpec, _, _ := runTest(t, spec, docs)
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

	packageSpec, _, _ := runTest(t, spec, docs)
	assert.Equal(
		t,
		"An array of server URLs.",
		packageSpec.Types["aws-native:pipes:PipeSourceSelfManagedKafkaParameters"].Properties["additionalBootstrapServers"].Description,
	)
}

func TestGatherPackage_docs_topLevelResource(t *testing.T) {
	spec := map[string]interface{}{
		"typeName":    "AWS::EC2::Volume",
		"properties":  map[string]interface{}{},
		"definitions": map[string]interface{}{},
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::EC2::Volume": {
				Description: "Specifies an Amazon Elastic Block Store (Amazon EBS) volume.\n\nWhen you use AWS CloudFormation to update an Amazon EBS volume that modifies `Iops` , `Size` , or `VolumeType` , there is a cooldown period before another operation can occur. This can cause your stack to report being in `UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long periods of time.\n\nAmazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation does not attempt to modify an Amazon EBS volume to a smaller size on rollback.\n\nSome common scenarios when you might encounter a cooldown period for Amazon EBS include:\n\n- You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.\n- You successfully update an Amazon EBS volume and the update succeeds but another change in your `update-stack` call fails. The rollback will be subject to a cooldown period.\n\nFor more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html) .\n\n*DeletionPolicy attribute*\n\nTo control how AWS CloudFormation handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .\n\n> If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.",
				Properties:  map[string]string{},
			},
		},
	}

	packageSpec, _, _ := runTest(t, spec, docs)
	assert.Equal(
		t,
		"Specifies an Amazon Elastic Block Store (Amazon EBS) volume.\n\nWhen you use AWS CloudFormation to update an Amazon EBS volume that modifies `Iops` , `Size` , or `VolumeType` , there is a cooldown period before another operation can occur. This can cause your stack to report being in `UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long periods of time.\n\nAmazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation does not attempt to modify an Amazon EBS volume to a smaller size on rollback.\n\nSome common scenarios when you might encounter a cooldown period for Amazon EBS include:\n\n- You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.\n- You successfully update an Amazon EBS volume and the update succeeds but another change in your `update-stack` call fails. The rollback will be subject to a cooldown period.\n\nFor more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html) .\n\n*DeletionPolicy attribute*\n\nTo control how AWS CloudFormation handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .\n\n> If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.",
		packageSpec.Resources["aws-native:ec2:Volume"].Description,
	)
}

func TestGatherPackage_docs_augmentation_topLevelResource(t *testing.T) {
	spec := map[string]interface{}{
		// AWS::EC2::Volume has forced augmentation turned on to replace bad schema descriptions with the docs descriptions
		"typeName":    "AWS::EC2::Volume",
		"description": "bad description",
		"properties":  map[string]interface{}{},
		"definitions": map[string]interface{}{},
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::EC2::Volume": {
				Description: "Specifies an Amazon Elastic Block Store (Amazon EBS) volume.\n\nWhen you use AWS CloudFormation to update an Amazon EBS volume that modifies `Iops` , `Size` , or `VolumeType` , there is a cooldown period before another operation can occur. This can cause your stack to report being in `UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long periods of time.\n\nAmazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation does not attempt to modify an Amazon EBS volume to a smaller size on rollback.\n\nSome common scenarios when you might encounter a cooldown period for Amazon EBS include:\n\n- You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.\n- You successfully update an Amazon EBS volume and the update succeeds but another change in your `update-stack` call fails. The rollback will be subject to a cooldown period.\n\nFor more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html) .\n\n*DeletionPolicy attribute*\n\nTo control how AWS CloudFormation handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .\n\n> If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.",
				Properties:  map[string]string{},
			},
		},
	}

	packageSpec, _, _ := runTest(t, spec, docs)
	assert.Equal(
		t,
		"Specifies an Amazon Elastic Block Store (Amazon EBS) volume.\n\nWhen you use AWS CloudFormation to update an Amazon EBS volume that modifies `Iops` , `Size` , or `VolumeType` , there is a cooldown period before another operation can occur. This can cause your stack to report being in `UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long periods of time.\n\nAmazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation does not attempt to modify an Amazon EBS volume to a smaller size on rollback.\n\nSome common scenarios when you might encounter a cooldown period for Amazon EBS include:\n\n- You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.\n- You successfully update an Amazon EBS volume and the update succeeds but another change in your `update-stack` call fails. The rollback will be subject to a cooldown period.\n\nFor more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html) .\n\n*DeletionPolicy attribute*\n\nTo control how AWS CloudFormation handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .\n\n> If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.",
		packageSpec.Resources["aws-native:ec2:Volume"].Description,
	)
}

func TestGatherPackage_docs_augmentation_nestedProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"EncryptionConfiguration": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"EncryptionType": map[string]interface{}{
						"$ref":        "#/definitions/EncryptionType",
						"description": "The encryption type to use.\n If you u... Truncated",
					},
				},
			},
			"EncryptionType": map[string]interface{}{
				"type":        "string",
				"description": "The encryption type to use.",
				"enum":        []interface{}{"AES256", "KMS"},
			},
		},
		"properties": map[string]interface{}{

			"EncryptionConfiguration": map[string]interface{}{
				"$ref":        "#/definitions/EncryptionConfiguration",
				"description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
			},
		},
		"typeName":    "AWS::ECR::Repository",
		"description": "The ``AWS::ECR::Repository`` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.",
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::ECR::Repository.EncryptionConfiguration": {
				Description: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the `AES256` encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more control over the encryption of the contents of your repository, you can use server-side encryption with AWS Key Management Service key stored in AWS Key Management Service ( AWS KMS ) to encrypt your images. For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide* .",
				Properties: map[string]string{
					"EncryptionType": "The encryption type to use.\n\nIf you use the `KMS` encryption type, the contents of the repository will be encrypted using server-side encryption with AWS Key Management Service key stored in AWS KMS . When you use AWS KMS to encrypt your data, you can either use the default AWS managed AWS KMS key for Amazon ECR, or specify your own AWS KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an AWS KMS key stored in AWS Key Management Service (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .\n\nIf you use the `AES256` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .",
				},
			},
		},
	}

	packageSpec, _, _ := runTest(t, spec, docs)
	assert.Equal(
		t,
		"The encryption type to use.\n\nIf you use the `KMS` encryption type, the contents of the repository will be encrypted using server-side encryption with AWS Key Management Service key stored in AWS KMS . When you use AWS KMS to encrypt your data, you can either use the default AWS managed AWS KMS key for Amazon ECR, or specify your own AWS KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an AWS KMS key stored in AWS Key Management Service (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .\n\nIf you use the `AES256` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide* .",
		packageSpec.Types["aws-native:ecr:RepositoryEncryptionConfiguration"].Properties["encryptionType"].Description,
	)
}

func TestGatherPackage_docs_augmentation_unknownProperty(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{
			"EncryptionConfiguration": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"EncryptionType": map[string]interface{}{
						"$ref":        "#/definitions/EncryptionType",
						"description": "The encryption type to use.\n If you u... Truncated",
					},
				},
			},
			"EncryptionType": map[string]interface{}{
				"type":        "string",
				"description": "The encryption type to use.",
				"enum":        []interface{}{"AES256", "KMS"},
			},
		},
		"properties": map[string]interface{}{

			"EncryptionConfiguration": map[string]interface{}{
				"$ref":        "#/definitions/EncryptionConfiguration",
				"description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
			},
		},
		"typeName":    "AWS::ECR::Repository",
		"description": "The ``AWS::ECR::Repository`` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.",
	}
	docs := Docs{
		Types: map[string]DocsTypes{},
	}

	packageSpec, _, reports := runTest(t, spec, docs)
	assert.Equal(
		t,
		"The encryption type to use.\n If you u... Truncated",
		packageSpec.Types["aws-native:ecr:RepositoryEncryptionConfiguration"].Properties["encryptionType"].Description,
	)
	assert.Contains(t, reports.MissingDocs, "AWS::ECR::Repository")
	assert.Contains(t, reports.MissingDocs["AWS::ECR::Repository"], "EncryptionType")
	assert.Equal(t, fmt.Sprint(1), reports.MissingDocs["AWS::ECR::Repository"]["EncryptionType"])
}

func TestGatherPackage_regionGeneration(t *testing.T) {
	spec := map[string]interface{}{
		"typeName":    "AWS::EC2::Volume",
		"properties":  map[string]interface{}{},
		"definitions": map[string]interface{}{},
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::EC2::Volume": {
				Description: "Specifies an Amazon Elastic Block Store (Amazon EBS) volume.\n\nWhen you use AWS CloudFormation to update an Amazon EBS volume that modifies `Iops` , `Size` , or `VolumeType` , there is a cooldown period before another operation can occur. This can cause your stack to report being in `UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long periods of time.\n\nAmazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation does not attempt to modify an Amazon EBS volume to a smaller size on rollback.\n\nSome common scenarios when you might encounter a cooldown period for Amazon EBS include:\n\n- You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.\n- You successfully update an Amazon EBS volume and the update succeeds but another change in your `update-stack` call fails. The rollback will be subject to a cooldown period.\n\nFor more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html) .\n\n*DeletionPolicy attribute*\n\nTo control how AWS CloudFormation handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .\n\n> If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.",
				Properties:  map[string]string{},
			},
		},
	}
	packageSpec, _, _ := runTestWithRegions(t, spec, docs, []RegionInfo{
		{
			Name:        "us-west-2",
			Description: "US West (Oregon)",
		},
		{
			Name:        "us-east-1",
			Description: "US East (N. Virginia)",
		},
		{
			Name:        "us-gov-west-1",
			Description: "AWS GovCloud (US-West)",
		},
		{
			Name:        "cn-north-1",
			Description: "China (Beijing)",
		},
		{
			Name:        "eu-isoe-west-1",
			Description: "EU ISOE West",
		},
	})

	assert.Contains(t, packageSpec.Types, "aws-native:index/Region:Region")
	regionEnum := packageSpec.Types["aws-native:index/Region:Region"].Enum

	assert.Equal(t, []schema.EnumValueSpec{
		{Value: "cn-north-1", Name: "CnNorth1", Description: "China (Beijing)"},
		{Value: "eu-isoe-west-1", Name: "EuIsoeWest1", Description: "EU ISOE West"},
		{Value: "us-east-1", Name: "UsEast1", Description: "US East (N. Virginia)"},
		{Value: "us-gov-west-1", Name: "UsGovWest1", Description: "AWS GovCloud (US-West)"},
		{Value: "us-west-2", Name: "UsWest2", Description: "US West (Oregon)"},
	}, regionEnum)
}

func TestGatherPackage_autonaming(t *testing.T) {
	spec := map[string]interface{}{
		"definitions": map[string]interface{}{},
		"properties": map[string]interface{}{
			"RoleName": map[string]interface{}{
				"description": "A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*.\n This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The role name must be unique within the account. Role names are not distinguished by case. For example, you cannot create roles named both \"Role1\" and \"role1\".\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the role name.\n If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).\n  Naming an IAM resource can cause an unrecoverable error if you reuse the same template in multiple Regions. To prevent this, we recommend using ``Fn::Join`` and ``AWS::Region`` to create a Region-specific name, as in the following example: ``{\"Fn::Join\": [\"\", [{\"Ref\": \"AWS::Region\"}, {\"Ref\": \"MyResourceName\"}]]}``.",
				"type":        "string",
			},
		},
		"typeName":             "AWS::IAM::Role",
		"primaryIdentifier":    []string{"/properties/RoleName"},
		"createOnlyProperties": []string{"/properties/Path", "/properties/RoleName"},
		"readOnlyProperties":   []string{"/properties/Arn", "/properties/RoleId"},
	}
	docs := Docs{
		Types: map[string]DocsTypes{
			"AWS::IAM::Role": {
				Description: "Some desc",
				Properties:  map[string]string{},
			},
		},
	}

	_, meta, _ := runTest(t, spec, docs)
	auto := meta.Resources["aws-native:iam:Role"].AutoNamingSpec
	assert.Equal(t, *auto, metadata.AutoNamingSpec{
		SdkName:   "roleName",
		MinLength: 1,
		MaxLength: 64,
	})

}
