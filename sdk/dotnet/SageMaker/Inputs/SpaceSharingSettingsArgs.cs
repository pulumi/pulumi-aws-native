// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class SpaceSharingSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sharingType", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.SpaceSharingSettingsSharingType> SharingType { get; set; } = null!;

        public SpaceSharingSettingsArgs()
        {
        }
        public static new SpaceSharingSettingsArgs Empty => new SpaceSharingSettingsArgs();
    }
}