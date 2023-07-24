// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Response body contents that indicate success or failure of a login request
    /// </summary>
    [OutputType]
    public sealed class WebACLResponseInspectionBodyContains
    {
        public readonly ImmutableArray<string> FailureStrings;
        public readonly ImmutableArray<string> SuccessStrings;

        [OutputConstructor]
        private WebACLResponseInspectionBodyContains(
            ImmutableArray<string> failureStrings,

            ImmutableArray<string> successStrings)
        {
            FailureStrings = failureStrings;
            SuccessStrings = successStrings;
        }
    }
}