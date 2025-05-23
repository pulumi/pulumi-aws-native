// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// The destination configuration for failed invocations.
    /// </summary>
    public sealed class EventInvokeConfigOnFailureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the destination resource.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        public EventInvokeConfigOnFailureArgs()
        {
        }
        public static new EventInvokeConfigOnFailureArgs Empty => new EventInvokeConfigOnFailureArgs();
    }
}
