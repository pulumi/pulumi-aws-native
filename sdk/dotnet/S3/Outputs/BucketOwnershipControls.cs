// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class BucketOwnershipControls
    {
        public readonly ImmutableArray<Outputs.BucketOwnershipControlsRule> Rules;

        [OutputConstructor]
        private BucketOwnershipControls(ImmutableArray<Outputs.BucketOwnershipControlsRule> rules)
        {
            Rules = rules;
        }
    }
}