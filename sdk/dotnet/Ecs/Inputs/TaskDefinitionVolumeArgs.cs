// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class TaskDefinitionVolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("dockerVolumeConfiguration")]
        public Input<Inputs.TaskDefinitionDockerVolumeConfigurationArgs>? DockerVolumeConfiguration { get; set; }

        [Input("efsVolumeConfiguration")]
        public Input<Inputs.TaskDefinitionEfsVolumeConfigurationArgs>? EfsVolumeConfiguration { get; set; }

        [Input("host")]
        public Input<Inputs.TaskDefinitionHostVolumePropertiesArgs>? Host { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public TaskDefinitionVolumeArgs()
        {
        }
        public static new TaskDefinitionVolumeArgs Empty => new TaskDefinitionVolumeArgs();
    }
}