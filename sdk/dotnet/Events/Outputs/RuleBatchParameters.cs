// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class RuleBatchParameters
    {
        public readonly Outputs.RuleBatchArrayProperties? ArrayProperties;
        public readonly string JobDefinition;
        public readonly string JobName;
        public readonly Outputs.RuleBatchRetryStrategy? RetryStrategy;

        [OutputConstructor]
        private RuleBatchParameters(
            Outputs.RuleBatchArrayProperties? arrayProperties,

            string jobDefinition,

            string jobName,

            Outputs.RuleBatchRetryStrategy? retryStrategy)
        {
            ArrayProperties = arrayProperties;
            JobDefinition = jobDefinition;
            JobName = jobName;
            RetryStrategy = retryStrategy;
        }
    }
}