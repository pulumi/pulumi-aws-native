// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormButtonArgs : global::Pulumi.ResourceArgs
    {
        [Input("children")]
        public Input<string>? Children { get; set; }

        [Input("excluded")]
        public Input<bool>? Excluded { get; set; }

        [Input("position")]
        public object? Position { get; set; }

        public FormButtonArgs()
        {
        }
        public static new FormButtonArgs Empty => new FormButtonArgs();
    }
}