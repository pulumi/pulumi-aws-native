// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateNullValueFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("nullString", required: true)]
        public Input<string> NullString { get; set; } = null!;

        public TemplateNullValueFormatConfigurationArgs()
        {
        }
        public static new TemplateNullValueFormatConfigurationArgs Empty => new TemplateNullValueFormatConfigurationArgs();
    }
}