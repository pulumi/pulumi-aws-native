// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Outputs
{

    /// <summary>
    /// A notification target.
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanNotificationTargetItem
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon SNS topic.
        /// </summary>
        public readonly string? SnsTopicArn;

        [OutputConstructor]
        private ResponsePlanNotificationTargetItem(string? snsTopicArn)
        {
            SnsTopicArn = snsTopicArn;
        }
    }
}
