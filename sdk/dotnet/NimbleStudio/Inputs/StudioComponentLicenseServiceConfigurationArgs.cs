// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    /// <summary>
    /// &lt;p&gt;The configuration for a license service that is associated with a studio
    ///             resource.&lt;/p&gt;
    /// </summary>
    public sealed class StudioComponentLicenseServiceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The endpoint of the license service that is accessed by the studio component
        ///             resource.&lt;/p&gt;
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        public StudioComponentLicenseServiceConfigurationArgs()
        {
        }
        public static new StudioComponentLicenseServiceConfigurationArgs Empty => new StudioComponentLicenseServiceConfigurationArgs();
    }
}