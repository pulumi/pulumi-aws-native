// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// Job Sample
    /// </summary>
    public sealed class JobSampleArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode")]
        public Input<Pulumi.AwsNative.DataBrew.JobSampleMode>? Mode { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        public JobSampleArgs()
        {
        }
        public static new JobSampleArgs Empty => new JobSampleArgs();
    }
}