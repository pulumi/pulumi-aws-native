// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda.Inputs
{

    /// <summary>
    /// Configuration to be applied to this Object lambda Access Point. It specifies Supporting Access Point, Transformation Configurations. Customers can also set if they like to enable Cloudwatch metrics for accesses to this Object lambda Access Point. Default setting for Cloudwatch metrics is disable.
    /// </summary>
    public sealed class AccessPointObjectLambdaConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedFeatures")]
        private InputList<string>? _allowedFeatures;

        /// <summary>
        /// A container for allowed features. Valid inputs are `GetObject-Range` , `GetObject-PartNumber` , `HeadObject-Range` , and `HeadObject-PartNumber` .
        /// </summary>
        public InputList<string> AllowedFeatures
        {
            get => _allowedFeatures ?? (_allowedFeatures = new InputList<string>());
            set => _allowedFeatures = value;
        }

        /// <summary>
        /// A container for whether the CloudWatch metrics configuration is enabled.
        /// </summary>
        [Input("cloudWatchMetricsEnabled")]
        public Input<bool>? CloudWatchMetricsEnabled { get; set; }

        /// <summary>
        /// Standard access point associated with the Object Lambda Access Point.
        /// </summary>
        [Input("supportingAccessPoint", required: true)]
        public Input<string> SupportingAccessPoint { get; set; } = null!;

        [Input("transformationConfigurations", required: true)]
        private InputList<Inputs.AccessPointTransformationConfigurationArgs>? _transformationConfigurations;

        /// <summary>
        /// A container for transformation configurations for an Object Lambda Access Point.
        /// </summary>
        public InputList<Inputs.AccessPointTransformationConfigurationArgs> TransformationConfigurations
        {
            get => _transformationConfigurations ?? (_transformationConfigurations = new InputList<Inputs.AccessPointTransformationConfigurationArgs>());
            set => _transformationConfigurations = value;
        }

        public AccessPointObjectLambdaConfigurationArgs()
        {
        }
        public static new AccessPointObjectLambdaConfigurationArgs Empty => new AccessPointObjectLambdaConfigurationArgs();
    }
}
