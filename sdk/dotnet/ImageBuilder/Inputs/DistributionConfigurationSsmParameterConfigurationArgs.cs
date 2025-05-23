// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// The SSM parameter configuration for AMI distribution.
    /// </summary>
    public sealed class DistributionConfigurationSsmParameterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID for the AMI to update the parameter with.
        /// </summary>
        [Input("amiAccountId")]
        public Input<string>? AmiAccountId { get; set; }

        /// <summary>
        /// The data type of the SSM parameter.
        /// </summary>
        [Input("dataType")]
        public Input<Pulumi.AwsNative.ImageBuilder.DistributionConfigurationSsmParameterConfigurationDataType>? DataType { get; set; }

        /// <summary>
        /// The name of the SSM parameter.
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        public DistributionConfigurationSsmParameterConfigurationArgs()
        {
        }
        public static new DistributionConfigurationSsmParameterConfigurationArgs Empty => new DistributionConfigurationSsmParameterConfigurationArgs();
    }
}
