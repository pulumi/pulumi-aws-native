// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// An object used to store task related info
    /// </summary>
    public sealed class FlowTaskPropertiesObjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The task property key.
        /// </summary>
        [Input("key", required: true)]
        public Input<Pulumi.AwsNative.AppFlow.FlowOperatorPropertiesKeys> Key { get; set; } = null!;

        /// <summary>
        /// The task property value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public FlowTaskPropertiesObjectArgs()
        {
        }
        public static new FlowTaskPropertiesObjectArgs Empty => new FlowTaskPropertiesObjectArgs();
    }
}
