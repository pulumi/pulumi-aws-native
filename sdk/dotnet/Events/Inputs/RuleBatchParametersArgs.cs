// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleBatchParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("arrayProperties")]
        public Input<Inputs.RuleBatchArrayPropertiesArgs>? ArrayProperties { get; set; }

        [Input("jobDefinition", required: true)]
        public Input<string> JobDefinition { get; set; } = null!;

        [Input("jobName", required: true)]
        public Input<string> JobName { get; set; } = null!;

        [Input("retryStrategy")]
        public Input<Inputs.RuleBatchRetryStrategyArgs>? RetryStrategy { get; set; }

        public RuleBatchParametersArgs()
        {
        }
        public static new RuleBatchParametersArgs Empty => new RuleBatchParametersArgs();
    }
}