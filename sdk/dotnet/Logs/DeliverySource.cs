// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    /// <summary>
    ///  A delivery source is an AWS resource that sends logs to an AWS destination. The destination can be CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.
    /// 
    /// Only some AWS services support being configured as a delivery source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from AWS services](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html).
    /// </summary>
    [AwsNativeResourceType("aws-native:logs:DeliverySource")]
    public partial class DeliverySource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that uniquely identifies this delivery source.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
        /// </summary>
        [Output("logType")]
        public Output<string?> LogType { get; private set; } = null!;

        /// <summary>
        /// The unique name of the Log source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the resource that will be sending the logs.
        /// </summary>
        [Output("resourceArn")]
        public Output<string?> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array.
        /// </summary>
        [Output("resourceArns")]
        public Output<ImmutableArray<string>> ResourceArns { get; private set; } = null!;

        /// <summary>
        /// The AWS service that is sending logs.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// The tags that have been assigned to this delivery source.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DeliverySourceTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DeliverySource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeliverySource(string name, DeliverySourceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:logs:DeliverySource", name, args ?? new DeliverySourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeliverySource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:logs:DeliverySource", name, null, MakeResourceOptions(options, id))
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
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeliverySource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeliverySource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeliverySource(name, id, options);
        }
    }

    public sealed class DeliverySourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
        /// </summary>
        [Input("logType")]
        public Input<string>? LogType { get; set; }

        /// <summary>
        /// The unique name of the Log source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the resource that will be sending the logs.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        [Input("tags")]
        private InputList<Inputs.DeliverySourceTagArgs>? _tags;

        /// <summary>
        /// The tags that have been assigned to this delivery source.
        /// </summary>
        public InputList<Inputs.DeliverySourceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DeliverySourceTagArgs>());
            set => _tags = value;
        }

        public DeliverySourceArgs()
        {
        }
        public static new DeliverySourceArgs Empty => new DeliverySourceArgs();
    }
}