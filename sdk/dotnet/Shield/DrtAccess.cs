// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Shield
{
    /// <summary>
    /// Config the role and list of Amazon S3 log buckets used by the Shield Response Team (SRT) to access your AWS account while assisting with attack mitigation.
    /// </summary>
    [AwsNativeResourceType("aws-native:shield:DrtAccess")]
    public partial class DrtAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the account that submitted the template.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
        /// </summary>
        [Output("logBucketList")]
        public Output<ImmutableArray<string>> LogBucketList { get; private set; } = null!;

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a DrtAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DrtAccess(string name, DrtAccessArgs args, CustomResourceOptions? options = null)
            : base("aws-native:shield:DrtAccess", name, args ?? new DrtAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DrtAccess(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:shield:DrtAccess", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DrtAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DrtAccess Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DrtAccess(name, id, options);
        }
    }

    public sealed class DrtAccessArgs : global::Pulumi.ResourceArgs
    {
        [Input("logBucketList")]
        private InputList<string>? _logBucketList;

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
        /// </summary>
        public InputList<string> LogBucketList
        {
            get => _logBucketList ?? (_logBucketList = new InputList<string>());
            set => _logBucketList = value;
        }

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public DrtAccessArgs()
        {
        }
        public static new DrtAccessArgs Empty => new DrtAccessArgs();
    }
}
