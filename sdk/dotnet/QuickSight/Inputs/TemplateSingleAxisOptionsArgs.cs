// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSingleAxisOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("yAxisOptions")]
        public Input<object>? YAxisOptions { get; set; }

        public TemplateSingleAxisOptionsArgs()
        {
        }
        public static new TemplateSingleAxisOptionsArgs Empty => new TemplateSingleAxisOptionsArgs();
    }
}