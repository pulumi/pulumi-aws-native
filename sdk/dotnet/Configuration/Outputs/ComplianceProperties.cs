// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    /// <summary>
    /// Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.
    /// </summary>
    [OutputType]
    public sealed class ComplianceProperties
    {
        /// <summary>
        /// Compliance type determined by the Config rule
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ComplianceProperties(string? type)
        {
            Type = type;
        }
    }
}
