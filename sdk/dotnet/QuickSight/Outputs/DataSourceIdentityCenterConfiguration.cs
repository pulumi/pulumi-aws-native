// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The parameters for an IAM Identity Center configuration.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSourceIdentityCenterConfiguration
    {
        /// <summary>
        /// &lt;p&gt;A Boolean option that controls whether Trusted Identity Propagation should be used.&lt;/p&gt;
        /// </summary>
        public readonly bool? EnableIdentityPropagation;

        [OutputConstructor]
        private DataSourceIdentityCenterConfiguration(bool? enableIdentityPropagation)
        {
            EnableIdentityPropagation = enableIdentityPropagation;
        }
    }
}
