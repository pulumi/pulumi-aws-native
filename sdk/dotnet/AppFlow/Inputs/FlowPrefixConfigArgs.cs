// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowPrefixConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("pathPrefixHierarchy")]
        private InputList<Pulumi.AwsNative.AppFlow.FlowPathPrefix>? _pathPrefixHierarchy;
        public InputList<Pulumi.AwsNative.AppFlow.FlowPathPrefix> PathPrefixHierarchy
        {
            get => _pathPrefixHierarchy ?? (_pathPrefixHierarchy = new InputList<Pulumi.AwsNative.AppFlow.FlowPathPrefix>());
            set => _pathPrefixHierarchy = value;
        }

        [Input("prefixFormat")]
        public Input<Pulumi.AwsNative.AppFlow.FlowPrefixFormat>? PrefixFormat { get; set; }

        [Input("prefixType")]
        public Input<Pulumi.AwsNative.AppFlow.FlowPrefixType>? PrefixType { get; set; }

        public FlowPrefixConfigArgs()
        {
        }
        public static new FlowPrefixConfigArgs Empty => new FlowPrefixConfigArgs();
    }
}