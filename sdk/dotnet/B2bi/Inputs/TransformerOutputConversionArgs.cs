// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class TransformerOutputConversionArgs : global::Pulumi.ResourceArgs
    {
        [Input("formatOptions")]
        public Input<Inputs.TransformerFormatOptionsPropertiesArgs>? FormatOptions { get; set; }

        [Input("toFormat", required: true)]
        public Input<Pulumi.AwsNative.B2bi.TransformerToFormat> ToFormat { get; set; } = null!;

        public TransformerOutputConversionArgs()
        {
        }
        public static new TransformerOutputConversionArgs Empty => new TransformerOutputConversionArgs();
    }
}
