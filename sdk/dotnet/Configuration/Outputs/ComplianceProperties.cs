// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    /// <summary>
    /// Compliance details of the Config rule
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