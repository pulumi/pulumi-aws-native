// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConvert.Inputs
{

    public sealed class JobTemplateAccelerationSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode", required: true)]
        public Input<string> Mode { get; set; } = null!;

        public JobTemplateAccelerationSettingsArgs()
        {
        }
        public static new JobTemplateAccelerationSettingsArgs Empty => new JobTemplateAccelerationSettingsArgs();
    }
}