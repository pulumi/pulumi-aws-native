// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class VolumeSnaplockConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditLogVolume")]
        public Input<string>? AuditLogVolume { get; set; }

        [Input("autocommitPeriod")]
        public Input<Inputs.VolumeAutocommitPeriodArgs>? AutocommitPeriod { get; set; }

        [Input("privilegedDelete")]
        public Input<string>? PrivilegedDelete { get; set; }

        [Input("retentionPeriod")]
        public Input<Inputs.VolumeSnaplockRetentionPeriodArgs>? RetentionPeriod { get; set; }

        [Input("snaplockType", required: true)]
        public Input<string> SnaplockType { get; set; } = null!;

        [Input("volumeAppendModeEnabled")]
        public Input<string>? VolumeAppendModeEnabled { get; set; }

        public VolumeSnaplockConfigurationArgs()
        {
        }
        public static new VolumeSnaplockConfigurationArgs Empty => new VolumeSnaplockConfigurationArgs();
    }
}