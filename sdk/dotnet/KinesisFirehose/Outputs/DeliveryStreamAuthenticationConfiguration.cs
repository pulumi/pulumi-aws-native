// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamAuthenticationConfiguration
    {
        /// <summary>
        /// The type of connectivity used to access the Amazon MSK cluster.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAuthenticationConfigurationConnectivity Connectivity;
        /// <summary>
        /// The ARN of the role used to access the Amazon MSK cluster.
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private DeliveryStreamAuthenticationConfiguration(
            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAuthenticationConfigurationConnectivity connectivity,

            string roleArn)
        {
            Connectivity = connectivity;
            RoleArn = roleArn;
        }
    }
}
