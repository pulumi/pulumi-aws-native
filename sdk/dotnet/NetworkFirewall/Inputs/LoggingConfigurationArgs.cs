// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    public sealed class LoggingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("logDestinationConfigs", required: true)]
        private InputList<Inputs.LoggingConfigurationLogDestinationConfigArgs>? _logDestinationConfigs;
        public InputList<Inputs.LoggingConfigurationLogDestinationConfigArgs> LogDestinationConfigs
        {
            get => _logDestinationConfigs ?? (_logDestinationConfigs = new InputList<Inputs.LoggingConfigurationLogDestinationConfigArgs>());
            set => _logDestinationConfigs = value;
        }

        public LoggingConfigurationArgs()
        {
        }
        public static new LoggingConfigurationArgs Empty => new LoggingConfigurationArgs();
    }
}