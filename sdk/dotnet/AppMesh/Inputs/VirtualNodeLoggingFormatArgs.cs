// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualNodeLoggingFormatArgs : global::Pulumi.ResourceArgs
    {
        [Input("json")]
        private InputList<Inputs.VirtualNodeJsonFormatRefArgs>? _json;
        public InputList<Inputs.VirtualNodeJsonFormatRefArgs> Json
        {
            get => _json ?? (_json = new InputList<Inputs.VirtualNodeJsonFormatRefArgs>());
            set => _json = value;
        }

        [Input("text")]
        public Input<string>? Text { get; set; }

        public VirtualNodeLoggingFormatArgs()
        {
        }
        public static new VirtualNodeLoggingFormatArgs Empty => new VirtualNodeLoggingFormatArgs();
    }
}