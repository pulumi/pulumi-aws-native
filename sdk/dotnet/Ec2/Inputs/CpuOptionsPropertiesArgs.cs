// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The CPU options for the instance.
    /// </summary>
    public sealed class CpuOptionsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("coreCount")]
        public Input<int>? CoreCount { get; set; }

        [Input("threadsPerCore")]
        public Input<int>? ThreadsPerCore { get; set; }

        public CpuOptionsPropertiesArgs()
        {
        }
        public static new CpuOptionsPropertiesArgs Empty => new CpuOptionsPropertiesArgs();
    }
}