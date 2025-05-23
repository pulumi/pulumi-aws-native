// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts.Inputs
{

    public sealed class BucketFilterAndOperatorPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prefix identifies one or more objects to which the rule applies.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tags", required: true)]
        private InputList<Inputs.BucketFilterTagArgs>? _tags;

        /// <summary>
        /// All of these tags must exist in the object's tag set in order for the rule to apply.
        /// </summary>
        public InputList<Inputs.BucketFilterTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BucketFilterTagArgs>());
            set => _tags = value;
        }

        public BucketFilterAndOperatorPropertiesArgs()
        {
        }
        public static new BucketFilterAndOperatorPropertiesArgs Empty => new BucketFilterAndOperatorPropertiesArgs();
    }
}
