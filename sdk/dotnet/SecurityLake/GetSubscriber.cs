// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake
{
    public static class GetSubscriber
    {
        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::Subscriber
        /// </summary>
        public static Task<GetSubscriberResult> InvokeAsync(GetSubscriberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriberResult>("aws-native:securitylake:getSubscriber", args ?? new GetSubscriberArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::Subscriber
        /// </summary>
        public static Output<GetSubscriberResult> Invoke(GetSubscriberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriberResult>("aws-native:securitylake:getSubscriber", args ?? new GetSubscriberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Security Lake subscriber.
        /// </summary>
        [Input("subscriberArn", required: true)]
        public string SubscriberArn { get; set; } = null!;

        public GetSubscriberArgs()
        {
        }
        public static new GetSubscriberArgs Empty => new GetSubscriberArgs();
    }

    public sealed class GetSubscriberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Security Lake subscriber.
        /// </summary>
        [Input("subscriberArn", required: true)]
        public Input<string> SubscriberArn { get; set; } = null!;

        public GetSubscriberInvokeArgs()
        {
        }
        public static new GetSubscriberInvokeArgs Empty => new GetSubscriberInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriberResult
    {
        /// <summary>
        /// You can choose to notify subscribers of new objects with an Amazon Simple Queue Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by the subscriber.
        /// 
        /// Subscribers can consume data by directly querying AWS Lake Formation tables in your Amazon S3 bucket through services like Amazon Athena. This subscription type is defined as `LAKEFORMATION` .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.SecurityLake.SubscriberAccessTypesItem> AccessTypes;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Security Lake subscriber.
        /// </summary>
        public readonly string? ResourceShareArn;
        /// <summary>
        /// The ARN name of the Amazon Security Lake subscriber.
        /// </summary>
        public readonly string? ResourceShareName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the S3 bucket.
        /// </summary>
        public readonly string? S3BucketArn;
        /// <summary>
        /// The supported AWS services from which logs and events are collected.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubscriberSource> Sources;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Security Lake subscriber.
        /// </summary>
        public readonly string? SubscriberArn;
        /// <summary>
        /// The description for your subscriber account in Security Lake.
        /// </summary>
        public readonly string? SubscriberDescription;
        /// <summary>
        /// The AWS identity used to access your data.
        /// </summary>
        public readonly Outputs.SubscriberIdentityProperties? SubscriberIdentity;
        /// <summary>
        /// The name of your Security Lake subscriber account.
        /// </summary>
        public readonly string? SubscriberName;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the role used to create the Security Lake subscriber.
        /// </summary>
        public readonly string? SubscriberRoleArn;
        /// <summary>
        /// An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetSubscriberResult(
            ImmutableArray<Pulumi.AwsNative.SecurityLake.SubscriberAccessTypesItem> accessTypes,

            string? resourceShareArn,

            string? resourceShareName,

            string? s3BucketArn,

            ImmutableArray<Outputs.SubscriberSource> sources,

            string? subscriberArn,

            string? subscriberDescription,

            Outputs.SubscriberIdentityProperties? subscriberIdentity,

            string? subscriberName,

            string? subscriberRoleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AccessTypes = accessTypes;
            ResourceShareArn = resourceShareArn;
            ResourceShareName = resourceShareName;
            S3BucketArn = s3BucketArn;
            Sources = sources;
            SubscriberArn = subscriberArn;
            SubscriberDescription = subscriberDescription;
            SubscriberIdentity = subscriberIdentity;
            SubscriberName = subscriberName;
            SubscriberRoleArn = subscriberRoleArn;
            Tags = tags;
        }
    }
}