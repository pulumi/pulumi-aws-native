// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// When exiting this state, perform these ``actions`` if the specified ``condition`` is ``TRUE``.
    /// </summary>
    [OutputType]
    public sealed class DetectorModelOnExit
    {
        /// <summary>
        /// Specifies the ``actions`` that are performed when the state is exited and the ``condition`` is ``TRUE``.
        /// </summary>
        public readonly ImmutableArray<Outputs.DetectorModelEvent> Events;

        [OutputConstructor]
        private DetectorModelOnExit(ImmutableArray<Outputs.DetectorModelEvent> events)
        {
            Events = events;
        }
    }
}
