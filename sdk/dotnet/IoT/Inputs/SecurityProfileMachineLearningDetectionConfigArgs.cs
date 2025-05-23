// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The configuration of an ML Detect Security Profile.
    /// </summary>
    public sealed class SecurityProfileMachineLearningDetectionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.
        /// </summary>
        [Input("confidenceLevel")]
        public Input<Pulumi.AwsNative.IoT.SecurityProfileMachineLearningDetectionConfigConfidenceLevel>? ConfidenceLevel { get; set; }

        public SecurityProfileMachineLearningDetectionConfigArgs()
        {
        }
        public static new SecurityProfileMachineLearningDetectionConfigArgs Empty => new SecurityProfileMachineLearningDetectionConfigArgs();
    }
}
