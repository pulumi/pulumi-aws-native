// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TopicDefaultFormatting
    {
        public readonly Pulumi.AwsNative.QuickSight.TopicDisplayFormat? DisplayFormat;
        public readonly Outputs.TopicDisplayFormatOptions? DisplayFormatOptions;

        [OutputConstructor]
        private TopicDefaultFormatting(
            Pulumi.AwsNative.QuickSight.TopicDisplayFormat? displayFormat,

            Outputs.TopicDisplayFormatOptions? displayFormatOptions)
        {
            DisplayFormat = displayFormat;
            DisplayFormatOptions = displayFormatOptions;
        }
    }
}