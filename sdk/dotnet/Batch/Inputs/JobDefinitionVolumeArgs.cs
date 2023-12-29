// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionVolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("efsVolumeConfiguration")]
        public Input<Inputs.JobDefinitionEfsVolumeConfigurationArgs>? EfsVolumeConfiguration { get; set; }

        [Input("host")]
        public Input<Inputs.JobDefinitionHostArgs>? Host { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public JobDefinitionVolumeArgs()
        {
        }
        public static new JobDefinitionVolumeArgs Empty => new JobDefinitionVolumeArgs();
    }
}