// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    /// <summary>
    /// The traffic configuration of your continuous deployment.
    /// </summary>
    public sealed class ContinuousDeploymentPolicyTrafficConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines which HTTP requests are sent to the staging distribution.
        /// </summary>
        [Input("singleHeaderConfig")]
        public Input<Inputs.ContinuousDeploymentPolicySingleHeaderConfigArgs>? SingleHeaderConfig { get; set; }

        /// <summary>
        /// Contains the percentage of traffic to send to the staging distribution.
        /// </summary>
        [Input("singleWeightConfig")]
        public Input<Inputs.ContinuousDeploymentPolicySingleWeightConfigArgs>? SingleWeightConfig { get; set; }

        /// <summary>
        /// The type of traffic configuration.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.CloudFront.ContinuousDeploymentPolicyTrafficConfigType> Type { get; set; } = null!;

        public ContinuousDeploymentPolicyTrafficConfigArgs()
        {
        }
        public static new ContinuousDeploymentPolicyTrafficConfigArgs Empty => new ContinuousDeploymentPolicyTrafficConfigArgs();
    }
}
