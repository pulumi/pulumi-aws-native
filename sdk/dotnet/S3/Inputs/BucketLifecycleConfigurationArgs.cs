// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.
    /// </summary>
    public sealed class BucketLifecycleConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.BucketRuleArgs>? _rules;

        /// <summary>
        /// A lifecycle rule for individual objects in an Amazon S3 bucket.
        /// </summary>
        public InputList<Inputs.BucketRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Indicates which default minimum object size behavior is applied to the lifecycle configuration.
        ///   This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.
        ///    +  ``all_storage_classes_128K`` - Objects smaller than 128 KB will not transition to any storage class by default.
        ///   +  ``varies_by_storage_class`` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB. 
        ///   
        ///  To customize the minimum object size for any transition you can add a filter that specifies a custom ``ObjectSizeGreaterThan`` or ``ObjectSizeLessThan`` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.
        /// </summary>
        [Input("transitionDefaultMinimumObjectSize")]
        public Input<Pulumi.AwsNative.S3.BucketLifecycleConfigurationTransitionDefaultMinimumObjectSize>? TransitionDefaultMinimumObjectSize { get; set; }

        public BucketLifecycleConfigurationArgs()
        {
        }
        public static new BucketLifecycleConfigurationArgs Empty => new BucketLifecycleConfigurationArgs();
    }
}
