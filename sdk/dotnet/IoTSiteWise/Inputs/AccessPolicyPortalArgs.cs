// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// A portal resource.
    /// </summary>
    public sealed class AccessPolicyPortalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the portal.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public AccessPolicyPortalArgs()
        {
        }
        public static new AccessPolicyPortalArgs Empty => new AccessPolicyPortalArgs();
    }
}
