// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Settings that determine if a Lambda function should be invoked to fulfill a specific intent.
    /// </summary>
    [OutputType]
    public sealed class BotFulfillmentCodeHookSetting
    {
        public readonly bool Enabled;
        public readonly Outputs.BotFulfillmentUpdatesSpecification? FulfillmentUpdatesSpecification;
        public readonly Outputs.BotPostFulfillmentStatusSpecification? PostFulfillmentStatusSpecification;

        [OutputConstructor]
        private BotFulfillmentCodeHookSetting(
            bool enabled,

            Outputs.BotFulfillmentUpdatesSpecification? fulfillmentUpdatesSpecification,

            Outputs.BotPostFulfillmentStatusSpecification? postFulfillmentStatusSpecification)
        {
            Enabled = enabled;
            FulfillmentUpdatesSpecification = fulfillmentUpdatesSpecification;
            PostFulfillmentStatusSpecification = postFulfillmentStatusSpecification;
        }
    }
}