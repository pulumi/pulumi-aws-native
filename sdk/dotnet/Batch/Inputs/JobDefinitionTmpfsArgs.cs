// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionTmpfsArgs : global::Pulumi.ResourceArgs
    {
        [Input("containerPath", required: true)]
        public Input<string> ContainerPath { get; set; } = null!;

        [Input("mountOptions")]
        private InputList<string>? _mountOptions;
        public InputList<string> MountOptions
        {
            get => _mountOptions ?? (_mountOptions = new InputList<string>());
            set => _mountOptions = value;
        }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        public JobDefinitionTmpfsArgs()
        {
        }
        public static new JobDefinitionTmpfsArgs Empty => new JobDefinitionTmpfsArgs();
    }
}