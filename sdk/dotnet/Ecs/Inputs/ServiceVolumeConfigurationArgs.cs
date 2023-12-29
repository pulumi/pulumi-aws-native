// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class ServiceVolumeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("managedEbsVolume")]
        public Input<Inputs.ServiceManagedEbsVolumeConfigurationArgs>? ManagedEbsVolume { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ServiceVolumeConfigurationArgs()
        {
        }
        public static new ServiceVolumeConfigurationArgs Empty => new ServiceVolumeConfigurationArgs();
    }
}