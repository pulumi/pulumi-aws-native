// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    public sealed class ConfigurationRecorderRecordingStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("useOnly", required: true)]
        public Input<string> UseOnly { get; set; } = null!;

        public ConfigurationRecorderRecordingStrategyArgs()
        {
        }
        public static new ConfigurationRecorderRecordingStrategyArgs Empty => new ConfigurationRecorderRecordingStrategyArgs();
    }
}