// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class ApplicationQAppsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status information about whether end users can create and use Amazon Q Apps in the web experience.
        /// </summary>
        [Input("qAppsControlMode", required: true)]
        public Input<Pulumi.AwsNative.QBusiness.ApplicationQAppsControlMode> QAppsControlMode { get; set; } = null!;

        public ApplicationQAppsConfigurationArgs()
        {
        }
        public static new ApplicationQAppsConfigurationArgs Empty => new ApplicationQAppsConfigurationArgs();
    }
}