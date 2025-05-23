// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies options for Spot Instances.
    ///  ``SpotOptions`` is a property of [AWS::EC2::LaunchTemplate InstanceMarketOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html).
    /// </summary>
    public sealed class LaunchTemplateSpotOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated.
        /// </summary>
        [Input("blockDurationMinutes")]
        public Input<int>? BlockDurationMinutes { get; set; }

        /// <summary>
        /// The behavior when a Spot Instance is interrupted. The default is ``terminate``.
        /// </summary>
        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// The maximum hourly price you're willing to pay for a Spot Instance. We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price. If you do specify this parameter, it must be more than USD $0.001. Specifying a value below USD $0.001 will result in an ``InvalidParameterValue`` error message when the launch template is used to launch an instance.
        ///   If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter.
        /// </summary>
        [Input("maxPrice")]
        public Input<string>? MaxPrice { get; set; }

        /// <summary>
        /// The Spot Instance request type.
        ///  If you are using Spot Instances with an Auto Scaling group, use ``one-time`` requests, as the ASlong service handles requesting new Spot Instances whenever the group is below its desired capacity.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        /// <summary>
        /// The end date of the request, in UTC format (*YYYY-MM-DD*T*HH:MM:SS*Z). Supported only for persistent requests.
        ///   +  For a persistent request, the request remains active until the ``ValidUntil`` date and time is reached. Otherwise, the request remains active until you cancel it.
        ///   +  For a one-time request, ``ValidUntil`` is not supported. The request remains active until all instances launch or you cancel the request.
        ///   
        ///  Default: 7 days from the current date
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public LaunchTemplateSpotOptionsArgs()
        {
        }
        public static new LaunchTemplateSpotOptionsArgs Empty => new LaunchTemplateSpotOptionsArgs();
    }
}
