// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes
{
    /// <summary>
    /// Definition of AWS::Pipes::Pipe Resource Type
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testPipe = new AwsNative.Pipes.Pipe("testPipe", new()
    ///     {
    ///         Name = "PipeCfnExample",
    ///         RoleArn = "arn:aws:iam::123456789123:role/Pipe-Dev-All-Targets-Dummy-Execution-Role",
    ///         Source = "arn:aws:sqs:us-east-1:123456789123:pipeDemoSource",
    ///         Enrichment = "arn:aws:execute-api:us-east-1:123456789123:53eo2i89p9/*/POST/pets",
    ///         Target = "arn:aws:states:us-east-1:123456789123:stateMachine:PipeTargetStateMachine",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testPipe = new AwsNative.Pipes.Pipe("testPipe", new()
    ///     {
    ///         Name = "PipeCfnExample",
    ///         RoleArn = "arn:aws:iam::123456789123:role/Pipe-Dev-All-Targets-Dummy-Execution-Role",
    ///         Source = "arn:aws:sqs:us-east-1:123456789123:pipeDemoSource",
    ///         Enrichment = "arn:aws:execute-api:us-east-1:123456789123:53eo2i89p9/*/POST/pets",
    ///         Target = "arn:aws:states:us-east-1:123456789123:stateMachine:PipeTargetStateMachine",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:pipes:Pipe")]
    public partial class Pipe : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the pipe.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The time the pipe was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The state the pipe is in.
        /// </summary>
        [Output("currentState")]
        public Output<Pulumi.AwsNative.Pipes.PipeState> CurrentState { get; private set; } = null!;

        /// <summary>
        /// A description of the pipe.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The state the pipe should be in.
        /// </summary>
        [Output("desiredState")]
        public Output<Pulumi.AwsNative.Pipes.PipeRequestedPipeState?> DesiredState { get; private set; } = null!;

        /// <summary>
        /// The ARN of the enrichment resource.
        /// </summary>
        [Output("enrichment")]
        public Output<string?> Enrichment { get; private set; } = null!;

        /// <summary>
        /// The parameters required to set up enrichment on your pipe.
        /// </summary>
        [Output("enrichmentParameters")]
        public Output<Outputs.PipeEnrichmentParameters?> EnrichmentParameters { get; private set; } = null!;

        /// <summary>
        /// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
        /// 
        /// To update a pipe that is using the default AWS owned key to use a customer managed key instead, or update a pipe that is using a customer managed key to use a different customer managed key, specify a customer managed key identifier.
        /// 
        /// To update a pipe that is using a customer managed key to use the default AWS owned key , specify an empty string.
        /// 
        /// For more information, see [Managing keys](https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html) in the *AWS Key Management Service Developer Guide* .
        /// </summary>
        [Output("kmsKeyIdentifier")]
        public Output<string?> KmsKeyIdentifier { get; private set; } = null!;

        /// <summary>
        /// When the pipe was last updated, in [ISO-8601 format](https://docs.aws.amazon.com/https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The logging configuration settings for the pipe.
        /// </summary>
        [Output("logConfiguration")]
        public Output<Outputs.PipeLogConfiguration?> LogConfiguration { get; private set; } = null!;

        /// <summary>
        /// The name of the pipe.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the role that allows the pipe to send data to the target.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the source resource.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// The parameters required to set up a source for your pipe.
        /// </summary>
        [Output("sourceParameters")]
        public Output<Outputs.PipeSourceParameters?> SourceParameters { get; private set; } = null!;

        /// <summary>
        /// The reason the pipe is in its current state.
        /// </summary>
        [Output("stateReason")]
        public Output<string> StateReason { get; private set; } = null!;

        /// <summary>
        /// The list of key-value pairs to associate with the pipe.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ARN of the target resource.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// The parameters required to set up a target for your pipe.
        /// 
        /// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
        /// </summary>
        [Output("targetParameters")]
        public Output<Outputs.PipeTargetParameters?> TargetParameters { get; private set; } = null!;


        /// <summary>
        /// Create a Pipe resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipe(string name, PipeArgs args, CustomResourceOptions? options = null)
            : base("aws-native:pipes:Pipe", name, args ?? new PipeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipe(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:pipes:Pipe", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "source",
                    "sourceParameters.activeMqBrokerParameters.queueName",
                    "sourceParameters.dynamoDbStreamParameters.startingPosition",
                    "sourceParameters.kinesisStreamParameters.startingPosition",
                    "sourceParameters.kinesisStreamParameters.startingPositionTimestamp",
                    "sourceParameters.managedStreamingKafkaParameters.consumerGroupId",
                    "sourceParameters.managedStreamingKafkaParameters.startingPosition",
                    "sourceParameters.managedStreamingKafkaParameters.topicName",
                    "sourceParameters.rabbitMqBrokerParameters.queueName",
                    "sourceParameters.rabbitMqBrokerParameters.virtualHost",
                    "sourceParameters.selfManagedKafkaParameters.additionalBootstrapServers[*]",
                    "sourceParameters.selfManagedKafkaParameters.consumerGroupId",
                    "sourceParameters.selfManagedKafkaParameters.startingPosition",
                    "sourceParameters.selfManagedKafkaParameters.topicName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pipe resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipe Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Pipe(name, id, options);
        }
    }

    public sealed class PipeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the pipe.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The state the pipe should be in.
        /// </summary>
        [Input("desiredState")]
        public Input<Pulumi.AwsNative.Pipes.PipeRequestedPipeState>? DesiredState { get; set; }

        /// <summary>
        /// The ARN of the enrichment resource.
        /// </summary>
        [Input("enrichment")]
        public Input<string>? Enrichment { get; set; }

        /// <summary>
        /// The parameters required to set up enrichment on your pipe.
        /// </summary>
        [Input("enrichmentParameters")]
        public Input<Inputs.PipeEnrichmentParametersArgs>? EnrichmentParameters { get; set; }

        /// <summary>
        /// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
        /// 
        /// To update a pipe that is using the default AWS owned key to use a customer managed key instead, or update a pipe that is using a customer managed key to use a different customer managed key, specify a customer managed key identifier.
        /// 
        /// To update a pipe that is using a customer managed key to use the default AWS owned key , specify an empty string.
        /// 
        /// For more information, see [Managing keys](https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html) in the *AWS Key Management Service Developer Guide* .
        /// </summary>
        [Input("kmsKeyIdentifier")]
        public Input<string>? KmsKeyIdentifier { get; set; }

        /// <summary>
        /// The logging configuration settings for the pipe.
        /// </summary>
        [Input("logConfiguration")]
        public Input<Inputs.PipeLogConfigurationArgs>? LogConfiguration { get; set; }

        /// <summary>
        /// The name of the pipe.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the role that allows the pipe to send data to the target.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The ARN of the source resource.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// The parameters required to set up a source for your pipe.
        /// </summary>
        [Input("sourceParameters")]
        public Input<Inputs.PipeSourceParametersArgs>? SourceParameters { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of key-value pairs to associate with the pipe.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ARN of the target resource.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// The parameters required to set up a target for your pipe.
        /// 
        /// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
        /// </summary>
        [Input("targetParameters")]
        public Input<Inputs.PipeTargetParametersArgs>? TargetParameters { get; set; }

        public PipeArgs()
        {
        }
        public static new PipeArgs Empty => new PipeArgs();
    }
}
