// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Outputs
{

    /// <summary>
    /// ``ThrottleSettings`` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.
    /// </summary>
    [OutputType]
    public sealed class UsagePlanThrottleSettings
    {
        public readonly int? BurstLimit;
        public readonly double? RateLimit;

        [OutputConstructor]
        private UsagePlanThrottleSettings(
            int? burstLimit,

            double? rateLimit)
        {
            BurstLimit = burstLimit;
            RateLimit = rateLimit;
        }
    }
}
