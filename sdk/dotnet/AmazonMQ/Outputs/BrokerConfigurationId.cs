// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMQ.Outputs
{

    [OutputType]
    public sealed class BrokerConfigurationId
    {
        public readonly string Id;
        public readonly int Revision;

        [OutputConstructor]
        private BrokerConfigurationId(
            string id,

            int revision)
        {
            Id = id;
            Revision = revision;
        }
    }
}