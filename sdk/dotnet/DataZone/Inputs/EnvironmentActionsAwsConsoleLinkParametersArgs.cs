// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// The parameters of the console link specified as part of the environment action
    /// </summary>
    public sealed class EnvironmentActionsAwsConsoleLinkParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URI of the console link specified as part of the environment action.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public EnvironmentActionsAwsConsoleLinkParametersArgs()
        {
        }
        public static new EnvironmentActionsAwsConsoleLinkParametersArgs Empty => new EnvironmentActionsAwsConsoleLinkParametersArgs();
    }
}