// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class TransformerSampleDocumentKeysArgs : global::Pulumi.ResourceArgs
    {
        [Input("input")]
        public Input<string>? Input { get; set; }

        [Input("output")]
        public Input<string>? Output { get; set; }

        public TransformerSampleDocumentKeysArgs()
        {
        }
        public static new TransformerSampleDocumentKeysArgs Empty => new TransformerSampleDocumentKeysArgs();
    }
}
