// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// The destination configuration for successful invocations.
    /// </summary>
    [OutputType]
    public sealed class EventInvokeConfigOnSuccess
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the destination resource.
        /// </summary>
        public readonly string Destination;

        [OutputConstructor]
        private EventInvokeConfigOnSuccess(string destination)
        {
            Destination = destination;
        }
    }
}
