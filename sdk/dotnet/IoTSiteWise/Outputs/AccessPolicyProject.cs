// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    /// <summary>
    /// A project resource.
    /// </summary>
    [OutputType]
    public sealed class AccessPolicyProject
    {
        /// <summary>
        /// The ID of the project.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private AccessPolicyProject(string? id)
        {
            Id = id;
        }
    }
}
