// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterConfigurationInfo
    {
        /// <summary>
        /// ARN of the configuration to use.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The revision of the configuration to use.
        /// </summary>
        public readonly int Revision;

        [OutputConstructor]
        private ClusterConfigurationInfo(
            string arn,

            int revision)
        {
            Arn = arn;
            Revision = revision;
        }
    }
}
