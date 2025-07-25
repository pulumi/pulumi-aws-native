// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless.Outputs
{

    /// <summary>
    /// The IAM IdentityCenter configuration for trusted-identity-propagation on this application. Supported with release labels emr-7.8.0 and above.
    /// </summary>
    [OutputType]
    public sealed class ApplicationIdentityCenterConfiguration
    {
        /// <summary>
        /// The IAM IdentityCenter instance arn
        /// </summary>
        public readonly string? IdentityCenterInstanceArn;

        [OutputConstructor]
        private ApplicationIdentityCenterConfiguration(string? identityCenterInstanceArn)
        {
            IdentityCenterInstanceArn = identityCenterInstanceArn;
        }
    }
}
