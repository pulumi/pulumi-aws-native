// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Aps.Inputs
{

    /// <summary>
    /// Series label
    /// </summary>
    public sealed class WorkspaceLabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the label
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of the label
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public WorkspaceLabelArgs()
        {
        }
        public static new WorkspaceLabelArgs Empty => new WorkspaceLabelArgs();
    }
}
