// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelOutputLockingSettings
    {
        public readonly Outputs.ChannelEpochLockingSettings? EpochLockingSettings;
        public readonly Outputs.ChannelPipelineLockingSettings? PipelineLockingSettings;

        [OutputConstructor]
        private ChannelOutputLockingSettings(
            Outputs.ChannelEpochLockingSettings? epochLockingSettings,

            Outputs.ChannelPipelineLockingSettings? pipelineLockingSettings)
        {
            EpochLockingSettings = epochLockingSettings;
            PipelineLockingSettings = pipelineLockingSettings;
        }
    }
}