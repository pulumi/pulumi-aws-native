// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PinpointEmail.Inputs
{

    public sealed class ConfigurationSetReputationOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("reputationMetricsEnabled")]
        public Input<bool>? ReputationMetricsEnabled { get; set; }

        public ConfigurationSetReputationOptionsArgs()
        {
        }
        public static new ConfigurationSetReputationOptionsArgs Empty => new ConfigurationSetReputationOptionsArgs();
    }
}