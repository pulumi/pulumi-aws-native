// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleTimestreamTimestamp
    {
        public readonly string Unit;
        public readonly string Value;

        [OutputConstructor]
        private TopicRuleTimestreamTimestamp(
            string unit,

            string value)
        {
            Unit = unit;
            Value = value;
        }
    }
}