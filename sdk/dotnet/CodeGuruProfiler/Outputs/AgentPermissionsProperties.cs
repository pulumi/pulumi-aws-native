// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeGuruProfiler.Outputs
{

    /// <summary>
    /// The agent permissions attached to this profiling group.
    /// </summary>
    [OutputType]
    public sealed class AgentPermissionsProperties
    {
        /// <summary>
        /// The principals for the agent permissions.
        /// </summary>
        public readonly ImmutableArray<string> Principals;

        [OutputConstructor]
        private AgentPermissionsProperties(ImmutableArray<string> principals)
        {
            Principals = principals;
        }
    }
}
